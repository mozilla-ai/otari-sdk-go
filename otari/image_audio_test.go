package otari

import (
	"context"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImageGeneration(t *testing.T) {
	t.Parallel()

	var (
		mu        sync.Mutex
		gotPath   string
		gotBody   map[string]any
		gotHeader http.Header
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		mu.Lock()
		gotPath = r.URL.Path
		gotHeader = r.Header.Clone()
		_ = json.Unmarshal(body, &gotBody)
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"created": 123, "data": [{"url": "https://img.example/1.png"}]}`))
	}))
	t.Cleanup(srv.Close)

	client, err := New(WithBaseURL(srv.URL), WithOtariKey("gw_key"))
	require.NoError(t, err)

	out, err := client.ImageGeneration(context.Background(), ImageGenerationParams{
		Model:  "openai:dall-e-3",
		Prompt: "a red bicycle",
		Extra:  map[string]any{"size": "1024x1024"},
	})
	require.NoError(t, err)
	require.Equal(t, float64(123), out["created"])
	data, ok := out["data"].([]any)
	require.True(t, ok)
	require.Len(t, data, 1)

	mu.Lock()
	defer mu.Unlock()
	require.Equal(t, "/v1/images/generations", gotPath)
	require.Equal(t, "openai:dall-e-3", gotBody["model"])
	require.Equal(t, "a red bicycle", gotBody["prompt"])
	require.Equal(t, "1024x1024", gotBody["size"])
	require.Equal(t, bearerPrefix+"gw_key", gotHeader.Get(apiKeyHeaderName))
}

func TestImageGenerationValidation(t *testing.T) {
	t.Parallel()

	client, err := New(WithBaseURL("http://example.invalid"), WithOtariKey("gw_key"))
	require.NoError(t, err)

	_, err = client.ImageGeneration(context.Background(), ImageGenerationParams{Prompt: "x"})
	require.ErrorIs(t, err, ErrInvalidRequest)

	_, err = client.ImageGeneration(context.Background(), ImageGenerationParams{Model: "m"})
	require.ErrorIs(t, err, ErrInvalidRequest)
}

func TestSpeechReturnsBytesAndSendsAuth(t *testing.T) {
	t.Parallel()

	audio := []byte{0x49, 0x44, 0x33, 0x04, 0x00} // fake ID3/mp3 header bytes

	var (
		mu        sync.Mutex
		gotPath   string
		gotBody   map[string]any
		gotHeader http.Header
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		mu.Lock()
		gotPath = r.URL.Path
		gotHeader = r.Header.Clone()
		_ = json.Unmarshal(body, &gotBody)
		mu.Unlock()

		w.Header().Set("Content-Type", "audio/mpeg")
		_, _ = w.Write(audio)
	}))
	t.Cleanup(srv.Close)

	client, err := New(
		WithBaseURL(srv.URL),
		WithAPIKey("tk_platform"),
		WithPlatformMode(),
	)
	require.NoError(t, err)

	got, err := client.Speech(context.Background(), SpeechParams{
		Model: "openai:tts-1",
		Input: "hello world",
		Voice: "alloy",
		Extra: map[string]any{"response_format": "mp3"},
	})
	require.NoError(t, err)
	require.Equal(t, audio, got)

	mu.Lock()
	defer mu.Unlock()
	require.Equal(t, "/v1/audio/speech", gotPath)
	require.Equal(t, "openai:tts-1", gotBody["model"])
	require.Equal(t, "hello world", gotBody["input"])
	require.Equal(t, "alloy", gotBody["voice"])
	require.Equal(t, "mp3", gotBody["response_format"])
	require.Equal(t, bearerPrefix+"tk_platform", gotHeader.Get(authorizationHeader))
}

func TestSpeechMapsErrors(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Retry-After", "7")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"detail": "rate limit exceeded"}`))
	}))
	t.Cleanup(srv.Close)

	client, err := New(WithBaseURL(srv.URL), WithOtariKey("gw_key"))
	require.NoError(t, err)

	_, err = client.Speech(context.Background(), SpeechParams{
		Model: "openai:tts-1",
		Input: "hello",
		Voice: "alloy",
	})
	require.ErrorIs(t, err, ErrRateLimit)

	var rl *RateLimitError
	require.ErrorAs(t, err, &rl)
	require.Equal(t, 7, rl.RetryAfter)
}

func TestTranscriptionSendsMultipartAndParsesJSON(t *testing.T) {
	t.Parallel()

	var (
		mu          sync.Mutex
		gotPath     string
		gotModel    string
		gotLang     string
		gotFile     []byte
		gotFilename string
		gotHeader   http.Header
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mediaType, params, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
		require.Equal(t, "multipart/form-data", mediaType)
		reader := multipart.NewReader(r.Body, params["boundary"])

		mu.Lock()
		gotPath = r.URL.Path
		gotHeader = r.Header.Clone()
		for {
			part, perr := reader.NextPart()
			if perr != nil {
				break
			}
			data, _ := io.ReadAll(part)
			switch part.FormName() {
			case "file":
				gotFile = data
				gotFilename = part.FileName()
			case "model":
				gotModel = string(data)
			case "language":
				gotLang = string(data)
			}
		}
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"text": "transcribed speech"}`))
	}))
	t.Cleanup(srv.Close)

	client, err := New(WithBaseURL(srv.URL), WithOtariKey("gw_key"))
	require.NoError(t, err)

	result, err := client.Transcription(context.Background(), TranscriptionParams{
		Model:    "openai:whisper-1",
		File:     []byte("RIFFfake-audio-bytes"),
		Filename: "speech.wav",
		Extra:    map[string]any{"language": "en"},
	})
	require.NoError(t, err)
	require.NotNil(t, result.JSON)
	require.Equal(t, "transcribed speech", result.JSON["text"])
	require.Empty(t, result.Text)

	mu.Lock()
	defer mu.Unlock()
	require.Equal(t, "/v1/audio/transcriptions", gotPath)
	require.Equal(t, "openai:whisper-1", gotModel)
	require.Equal(t, "en", gotLang)
	require.Equal(t, []byte("RIFFfake-audio-bytes"), gotFile)
	require.Equal(t, "speech.wav", gotFilename)
	require.Equal(t, bearerPrefix+"gw_key", gotHeader.Get(apiKeyHeaderName))
}

func TestTranscriptionReturnsTextForNonJSON(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("1\n00:00:00,000 --> 00:00:01,000\nhello\n"))
	}))
	t.Cleanup(srv.Close)

	client, err := New(WithBaseURL(srv.URL), WithOtariKey("gw_key"))
	require.NoError(t, err)

	result, err := client.Transcription(context.Background(), TranscriptionParams{
		Model: "openai:whisper-1",
		File:  []byte("audio"),
		Extra: map[string]any{"response_format": "srt"},
	})
	require.NoError(t, err)
	require.Nil(t, result.JSON)
	require.Contains(t, result.Text, "hello")
}

func TestTranscriptionValidation(t *testing.T) {
	t.Parallel()

	client, err := New(WithBaseURL("http://example.invalid"), WithOtariKey("gw_key"))
	require.NoError(t, err)

	_, err = client.Transcription(context.Background(), TranscriptionParams{File: []byte("x")})
	require.ErrorIs(t, err, ErrInvalidRequest)

	_, err = client.Transcription(context.Background(), TranscriptionParams{Model: "m"})
	require.ErrorIs(t, err, ErrInvalidRequest)
}
