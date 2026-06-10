package otari

import (
	"context"
	stderrors "errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// These tests exercise the generated-core shell: request shaping + typed decode for
// the generated-core-backed methods (including the new Message and Response
// endpoints), the unified status -> typed-error mapping in both auth modes, and
// the hand-written SSE streaming shim against mocked text/event-stream bytes.
//
// Streamed chat cannot be exercised end-to-end here (no provider key on the
// gateway in this sandbox); the SSE shim is verified with mocked bytes only.

// --- Message (NEW /messages endpoint) -------------------------------------

func TestMessageTypedDecodeAndHeaders(t *testing.T) {
	t.Parallel()

	var (
		mu       sync.Mutex
		gotPath  string
		gotBody  map[string]any
		gotOtari string
	)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		gotPath = r.URL.Path
		gotOtari = r.Header.Get(apiKeyHeaderName)
		raw, _ := io.ReadAll(r.Body)
		_ = jsonUnmarshal(raw, &gotBody)
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"id": "msg-1", "type": "message", "role": "assistant",
			"model": "anthropic:claude-3-5-sonnet",
			"content": [{"type":"text","text":"Hi"}],
			"usage": {"input_tokens":1,"output_tokens":1}
		}`))
	}))
	t.Cleanup(srv.Close)

	client, err := New(WithBaseURL(srv.URL), WithOtariKey("vk"))
	require.NoError(t, err)

	out, err := client.Message(context.Background(), MessageParams{
		Model:     "anthropic:claude-3-5-sonnet",
		Messages:  []map[string]any{{"role": "user", "content": "Hi"}},
		MaxTokens: 64,
	})
	require.NoError(t, err)
	require.Equal(t, "msg-1", out["id"])

	mu.Lock()
	defer mu.Unlock()
	require.Equal(t, "/v1/messages", gotPath)
	require.EqualValues(t, 64, gotBody["max_tokens"])
	require.Equal(t, "anthropic:claude-3-5-sonnet", gotBody["model"])
	require.Equal(t, bearerPrefix+"vk", gotOtari)
}

func TestMessageRequiresMaxTokens(t *testing.T) {
	t.Parallel()

	client, err := New(WithBaseURL("http://localhost:9999"), WithOtariKey("vk"))
	require.NoError(t, err)

	_, err = client.Message(context.Background(), MessageParams{
		Model:    "anthropic:claude",
		Messages: []map[string]any{{"role": "user", "content": "hi"}},
	})
	require.Error(t, err)
	require.ErrorIs(t, err, ErrInvalidRequest)
	require.Contains(t, err.Error(), "max_tokens is required")
}

// --- Response (/responses endpoint) ---------------------------------------

func TestResponseTypedDecode(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/v1/responses", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"resp-1","object":"response","output":[]}`))
	}))
	t.Cleanup(srv.Close)

	client, err := New(WithBaseURL(srv.URL), WithOtariKey("vk"))
	require.NoError(t, err)

	out, err := client.Response(context.Background(), ResponseParams{
		Model: "openai:gpt-4o-mini",
		Input: "hello",
	})
	require.NoError(t, err)
	require.Equal(t, "resp-1", out["id"])
}

// --- Unified error mapping (generated-core path) --------------------------

func TestUnifiedErrorMapping(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		status  int
		body    string
		wantErr error
	}{
		{"401 auth", http.StatusUnauthorized, `{"detail":"boom"}`, ErrAuthentication},
		{"403 auth", http.StatusForbidden, `{"detail":"boom"}`, ErrAuthentication},
		{"402 funds", http.StatusPaymentRequired, `{"detail":"boom"}`, ErrInsufficientFunds},
		{"404 model", http.StatusNotFound, `{"detail":"boom"}`, ErrModelNotFound},
		{"409 batch", http.StatusConflict, `{"detail":"boom"}`, ErrBatchNotComplete},
		{"429 rate", http.StatusTooManyRequests, `{"detail":"boom"}`, ErrRateLimit},
		{"502 upstream", http.StatusBadGateway, `{"detail":"boom"}`, ErrUpstreamProvider},
		{"503 upstream", http.StatusServiceUnavailable, `{"detail":"boom"}`, ErrUpstreamProvider},
		{"504 timeout", http.StatusGatewayTimeout, `{"detail":"boom"}`, ErrTimeout},
		{"418 generic", http.StatusTeapot, `{"detail":"boom"}`, ErrProvider},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.body))
			}))
			t.Cleanup(srv.Close)

			client, err := New(WithBaseURL(srv.URL), WithOtariKey("vk"))
			require.NoError(t, err)

			_, err = client.Completion(context.Background(), mockCompletionParams())
			require.Error(t, err)
			require.ErrorIs(t, err, tc.wantErr)
			require.Contains(t, err.Error(), "boom")
		})
	}
}

func TestErrorRetryAfterAndCorrelationID(t *testing.T) {
	t.Parallel()

	t.Run("429 carries retry-after", func(t *testing.T) {
		t.Parallel()
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Retry-After", "30")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"detail":"slow down"}`))
		}))
		t.Cleanup(srv.Close)

		client, err := New(WithBaseURL(srv.URL), WithOtariKey("vk"))
		require.NoError(t, err)

		_, err = client.Completion(context.Background(), mockCompletionParams())
		require.Error(t, err)
		var rl *RateLimitError
		require.True(t, stderrors.As(err, &rl))
		require.Equal(t, 30, rl.RetryAfter)
	})

	t.Run("correlation id appears in message", func(t *testing.T) {
		t.Parallel()
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("X-Correlation-Id", "abc-123")
			w.WriteHeader(http.StatusPaymentRequired)
			_, _ = w.Write([]byte(`{"detail":"no funds"}`))
		}))
		t.Cleanup(srv.Close)

		client, err := New(WithBaseURL(srv.URL), WithOtariKey("vk"))
		require.NoError(t, err)

		_, err = client.Completion(context.Background(), mockCompletionParams())
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInsufficientFunds)
		require.Contains(t, err.Error(), "abc-123")
	})
}

// --- SSE streaming shim (mocked bytes only) -------------------------------

func TestChatStreamingTypedChunksAndDone(t *testing.T) {
	t.Parallel()

	var (
		mu      sync.Mutex
		gotAcc  string
		gotAuth string
	)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		gotAcc = r.Header.Get("Accept")
		gotAuth = r.Header.Get(authorizationHeader)
		mu.Unlock()
		w.Header().Set("Content-Type", "text/event-stream")
		// Two typed chunks, an SSE comment, an unrelated event line, then [DONE].
		_, _ = fmt.Fprint(w, ": keep-alive comment\n\n")
		_, _ = fmt.Fprintf(w, "data: %s\n\n",
			`{"id":"c","object":"chat.completion.chunk","created":1,"model":"m","choices":[{"index":0,"delta":{"role":"assistant","content":"He"}}]}`)
		_, _ = fmt.Fprintf(w, "event: ping\ndata: %s\n\n",
			`{"id":"c","object":"chat.completion.chunk","created":1,"model":"m","choices":[{"index":0,"delta":{"content":"llo"}}]}`)
		_, _ = fmt.Fprint(w, "data: [DONE]\n\n")
		// Anything after [DONE] must be ignored.
		_, _ = fmt.Fprint(w, "data: {\"should\":\"not appear\"}\n\n")
	}))
	t.Cleanup(srv.Close)

	client, err := New(WithBaseURL(srv.URL), WithAPIKey("tk"), WithPlatformMode())
	require.NoError(t, err)

	chunks, errs := client.CompletionStream(context.Background(), mockCompletionParams())

	var got string
	n := 0
	for chunk := range chunks {
		n++
		require.Equal(t, objectChatCompletionChunk, chunk.Object)
		if len(chunk.Choices) > 0 {
			got += chunk.Choices[0].Delta.Content
		}
	}
	require.NoError(t, <-errs)
	require.Equal(t, 2, n, "two chunks before [DONE]")
	require.Equal(t, "Hello", got)

	mu.Lock()
	defer mu.Unlock()
	require.Equal(t, "text/event-stream", gotAcc)
	require.Equal(t, bearerPrefix+"tk", gotAuth)
}

func TestChatStreamingErrorMapsToTypedError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"detail":"rate limited"}`))
	}))
	t.Cleanup(srv.Close)

	client, err := New(WithBaseURL(srv.URL), WithOtariKey("vk"))
	require.NoError(t, err)

	chunks, errs := client.CompletionStream(context.Background(), mockCompletionParams())
	for range chunks {
	}
	err = <-errs
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRateLimit)
}

func TestMessageStreamingRawEvents(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/v1/messages", r.URL.Path)
		w.Header().Set("Content-Type", "text/event-stream")
		_, _ = fmt.Fprintf(w, "data: %s\n\n", `{"type":"message_start"}`)
		_, _ = fmt.Fprintf(w, "data: %s\n\n", `{"type":"content_block_delta","delta":{"text":"hi"}}`)
		_, _ = fmt.Fprint(w, "data: [DONE]\n\n")
	}))
	t.Cleanup(srv.Close)

	client, err := New(WithBaseURL(srv.URL), WithOtariKey("vk"))
	require.NoError(t, err)

	events, errs := client.MessageStream(context.Background(), MessageParams{
		Model:     "anthropic:claude-3-5-sonnet",
		Messages:  []map[string]any{{"role": "user", "content": "Hi"}},
		MaxTokens: 32,
	})

	var types []string
	for evt := range events {
		types = append(types, fmt.Sprint(evt["type"]))
	}
	require.NoError(t, <-errs)
	require.Equal(t, []string{"message_start", "content_block_delta"}, types)
}
