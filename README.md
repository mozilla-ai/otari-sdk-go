<p align="center">
  <img src="assets/otari-logo.svg" width="320" alt="otari logo"/>
</p>

<div align="center">

# Otari Go Client SDK

![Go 1.25+](https://img.shields.io/badge/go-1.25%2B-blue.svg)
<a href="https://discord.gg/4gf3zXrQUc">
    <img src="https://img.shields.io/static/v1?label=Chat%20on&message=Discord&color=blue&logo=Discord&style=flat-square" alt="Discord">
</a>

**Go client for [otari](https://github.com/mozilla-ai/otari), the open-source core that powers [otari.ai](https://otari.ai).**
Communicate with any LLM provider through otari using a single, typed interface.

[Python SDK](https://github.com/mozilla-ai/otari-sdk-python) | [TypeScript SDK](https://github.com/mozilla-ai/otari-sdk-ts) | [Documentation](https://mozilla-ai.github.io/otari/) | [Platform (Beta)](https://otari.ai/)

</div>

> New to otari? The [otari repo](https://github.com/mozilla-ai/otari) explains what it is and why you’d use it.

## Quickstart

```bash
go get github.com/mozilla-ai/otari-sdk-go/otari
```

Generate an API token at [otari.ai/organization-settings/api-tokens](https://otari.ai/organization-settings/api-tokens), then add a provider key (e.g. OpenAI) at [otari.ai/organization-settings/provider-keys](https://otari.ai/organization-settings/provider-keys) so the gateway can route requests to that provider. Then use the client:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/mozilla-ai/otari-sdk-go/otari"
)

func main() {
    client, err := otari.New(
        otari.WithAPIKey("tk_your_api_token"),
        otari.WithPlatformMode(),
    )
    if err != nil {
        log.Fatal(err)
    }

    resp, err := client.Completion(context.Background(), otari.CompletionParams{
        Model:    "openai:gpt-4o-mini",
        Messages: []otari.Message{{Role: otari.RoleUser, Content: "Hello!"}},
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp.Choices[0].Message.ContentString())
}
```

That's it: the client defaults to the hosted gateway at `https://api.otari.ai`, so you don't need to set a base URL. Change the model string to switch between LLM providers through the gateway.

## Installation

### Requirements

- Go 1.25 or newer
- An [otari.ai](https://otari.ai/) account, or a running self-hosted [otari](https://mozilla-ai.github.io/otari/gateway/overview/) instance

### Install

```bash
go get github.com/mozilla-ai/otari-sdk-go/otari
```

### Setting up credentials

For the hosted platform, set the platform-token environment variable:

```bash
export OTARI_AI_TOKEN="tk_your_api_token"
# GATEWAY_PLATFORM_TOKEN is still accepted as a legacy alias.
```

For a self-hosted gateway, set the base URL and gateway API key:

```bash
export GATEWAY_API_BASE="http://localhost:8000"
export GATEWAY_API_KEY="your-key-here"
```

Alternatively, pass credentials directly when creating the client (see [Authentication](#authentication)).

## Authentication

The client supports two authentication modes, matching the Python and TypeScript SDKs.

**Platform mode (recommended)** uses a platform token as standard `Authorization: Bearer` auth. On the hosted platform, generate an API token at [otari.ai/organization-settings/api-tokens](https://otari.ai/organization-settings/api-tokens) and add a provider key (e.g. OpenAI) at [otari.ai/organization-settings/provider-keys](https://otari.ai/organization-settings/provider-keys) so the gateway can route requests to that provider. The base URL defaults to the hosted gateway at `https://api.otari.ai`, so it can be omitted:

```go
client, err := otari.New(
    otari.WithAPIKey("tk_your_api_token"),
    otari.WithPlatformMode(),
)
```

The token can also be supplied via the `OTARI_AI_TOKEN` environment variable (legacy alias `GATEWAY_PLATFORM_TOKEN`), in which case `otari.New(otari.WithPlatformMode())` picks it up automatically.

**Self-hosted mode** sends the API key via the custom `Otari-Key: Bearer …` header and requires a base URL. Run the gateway yourself following the setup in the [otari repo](https://github.com/mozilla-ai/otari), make sure it has provider keys configured, then point the SDK at it:

```go
client, err := otari.New(
    otari.WithBaseURL("http://localhost:8000"), // or wherever you host the gateway
    otari.WithOtariKey("your-gateway-api-key"),
)
```

The base URL and key can also come from the `GATEWAY_API_BASE` and `GATEWAY_API_KEY` environment variables. When no explicit credentials are provided, `otari.New()` resolves the mode from the environment: a platform token (`OTARI_AI_TOKEN` / `GATEWAY_PLATFORM_TOKEN`) selects platform mode against the hosted gateway, otherwise `GATEWAY_API_BASE` + `GATEWAY_API_KEY` select self-hosted mode.

## Usage

### Chat completions

```go
resp, err := client.Completion(ctx, otari.CompletionParams{
    Model:    "openai:gpt-4o-mini",
    Messages: []otari.Message{{Role: otari.RoleUser, Content: "Hello!"}},
})
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp.Choices[0].Message.ContentString())
```

### Streaming

Streaming methods return a channel of typed chunks and a buffered error channel. Drain the chunk channel, then check the error channel once it closes.

```go
chunks, errs := client.CompletionStream(ctx, otari.CompletionParams{
    Model:    "openai:gpt-4o-mini",
    Messages: []otari.Message{{Role: otari.RoleUser, Content: "Tell me a story."}},
})

for chunk := range chunks {
    if len(chunk.Choices) > 0 {
        fmt.Print(chunk.Choices[0].Delta.Content)
    }
}

if err := <-errs; err != nil {
    log.Fatal(err)
}
```

### Responses API

The OpenAI-style Responses API is exposed via `Response` (and `ResponseStream`). The gateway returns an opaque payload, so the decoded JSON is delivered as a `map[string]any`. `Input` accepts any JSON-serializable value (a string, a message list, and so on); `Extra` forwards additional `/responses` fields verbatim.

```go
resp, err := client.Response(ctx, otari.ResponseParams{
    Model: "openai:gpt-4o-mini",
    Input: "Write a haiku about Go.",
})
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp["output_text"])
```

Streaming delivers each event as a raw decoded JSON map:

```go
events, errs := client.ResponseStream(ctx, otari.ResponseParams{
    Model: "openai:gpt-4o-mini",
    Input: "Write a haiku about Go.",
})

for event := range events {
    fmt.Println(event["type"])
}

if err := <-errs; err != nil {
    log.Fatal(err)
}
```

### Messages API

The Anthropic-shaped Messages API is exposed via `Message` (and `MessageStream`). `MaxTokens` is required by the gateway. The response is opaque, so the decoded JSON is delivered as a `map[string]any`; `Extra` forwards additional `/messages` fields (`system`, `tools`, `thinking`, and so on) verbatim.

```go
resp, err := client.Message(ctx, otari.MessageParams{
    Model:     "anthropic:claude-3-5-sonnet",
    MaxTokens: 1024,
    Messages: []map[string]any{
        {"role": "user", "content": "Hello!"},
    },
})
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp["content"])
```

Streaming delivers each event as a raw decoded JSON map:

```go
events, errs := client.MessageStream(ctx, otari.MessageParams{
    Model:     "anthropic:claude-3-5-sonnet",
    MaxTokens: 1024,
    Messages: []map[string]any{
        {"role": "user", "content": "Tell me a story."},
    },
})

for event := range events {
    fmt.Println(event["type"])
}

if err := <-errs; err != nil {
    log.Fatal(err)
}
```

### Embeddings

```go
result, err := client.Embedding(ctx, otari.EmbeddingParams{
    Model: "openai:text-embedding-3-small",
    Input: "Hello world",
})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result.Data[0].Embedding)
```

### Listing models

```go
models, err := client.ListModels(ctx)
if err != nil {
    log.Fatal(err)
}
for _, m := range models.Data {
    fmt.Println(m.ID)
}
```

### Moderation

```go
result, err := client.Moderation(ctx, otari.ModerationParams{
    Model: "openai:omni-moderation-latest",
    Input: "some content to check",
})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Flagged: %v\n", result.Results[0].Flagged)
```

### Reranking

```go
result, err := client.Rerank(ctx, otari.RerankParams{
    Model:     "cohere:rerank-v3.5",
    Query:     "What is machine learning?",
    Documents: []string{"ML is a subset of AI", "The weather is nice"},
})
if err != nil {
    log.Fatal(err)
}
for _, r := range result.Results {
    fmt.Printf("doc[%d] score=%.3f\n", r.Index, r.RelevanceScore)
}
```

### Batch operations

Create a batch, poll it with `RetrieveBatch`, cancel it with `CancelBatch`, enumerate jobs with `ListBatches`, and fetch results with `RetrieveBatchResults`. Each lookup takes the batch ID and the upstream provider.

```go
// Create a batch.
batch, err := client.CreateBatch(ctx, otari.CreateBatchParams{
    Model: "openai:gpt-4o-mini",
    Requests: []otari.BatchRequestItem{
        {CustomID: "req-1", Body: map[string]any{"messages": []any{
            map[string]any{"role": "user", "content": "Hello"},
        }}},
    },
    CompletionWindow: "24h",
})
if err != nil {
    log.Fatal(err)
}

// Check status later.
status, err := client.RetrieveBatch(ctx, batch.ID, batch.Provider)
if err != nil {
    log.Fatal(err)
}
fmt.Println(status.Status)

// Retrieve results once complete.
results, err := client.RetrieveBatchResults(ctx, batch.ID, batch.Provider)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%d results\n", len(results.Results))
```

### Error handling

All errors are typed. Use `errors.Is` for sentinel checks and `errors.As` for typed details:

```go
import "errors"

resp, err := client.Completion(ctx, params)
if err != nil {
    switch {
    case errors.Is(err, otari.ErrAuthentication):
        log.Fatal("Invalid credentials")
    case errors.Is(err, otari.ErrRateLimit):
        log.Fatal("Rate limited, try again later")
    case errors.Is(err, otari.ErrInsufficientFunds):
        log.Fatal("Budget exceeded")
    default:
        log.Fatal(err)
    }
}
```

| HTTP Status | Error Type | Sentinel | Description |
|------------|-----------|----------|-------------|
| 400 (capability) | `*UnsupportedCapabilityError` | `ErrUnsupported` | Provider does not support the requested capability |
| 401, 403 | `*AuthenticationError` | `ErrAuthentication` | Invalid or missing credentials |
| 402 | `*InsufficientFundsError` | `ErrInsufficientFunds` | Budget or credits exhausted |
| 404 | `*ModelNotFoundError` | `ErrModelNotFound` | Model not found, or no provider key configured for the requested provider; the error carries the gateway's detail |
| 409 | `*BatchNotCompleteError` | `ErrBatchNotComplete` | Batch not yet finished (includes `BatchID`, `Status`) |
| 429 | `*RateLimitError` | `ErrRateLimit` | Rate limit exceeded |
| 502 | `*UpstreamProviderError` | `ErrUpstreamProvider` | Upstream provider unreachable |
| 504 | `*TimeoutError` | `ErrTimeout` | Gateway timed out waiting for provider |

## Project Structure

```
otari-sdk-go/
├── otari/                  # Main package (hand-written ergonomic shell)
│   ├── client.go           # Client struct, constructor, chat/messages/responses methods
│   ├── config.go           # Functional options (WithBaseURL, WithAPIKey, WithPlatformMode, ...)
│   ├── streaming.go        # Hand-written SSE decoder (channel-based streaming)
│   ├── errors.go           # Typed error hierarchy with sentinel errors
│   ├── types.go            # Request/response types
│   ├── batch.go            # Batch API methods
│   ├── moderation.go       # Content moderation
│   ├── rerank.go           # Document reranking
│   ├── control_plane.go    # Control-plane API (keys/users/budgets/pricing/usage)
│   ├── conversion.go       # Type conversions
│   ├── request.go          # HTTP request helpers
│   ├── transport.go        # HTTP transport helpers
│   ├── otari.go            # Package documentation
│   ├── client/             # Generated OpenAPI core (do not hand-edit; regenerated from the spec)
│   └── *_test.go           # Unit, integration, and endpoint-coverage tests
├── examples/
│   └── quickstart/         # Quick smoke test
├── sdk-endpoints.txt       # Endpoint-coverage manifest (drift gate)
├── go.mod
├── go.sum
└── README.md
```

## Development

```bash
# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Build
go build ./...

# Vet
go vet ./...
```

## Documentation

- **[Full Documentation](https://mozilla-ai.github.io/otari/)** - Complete guides and API reference
- **[Supported Providers](https://mozilla-ai.github.io/otari/providers/)** - List of all supported LLM providers
- **[Gateway Documentation](https://mozilla-ai.github.io/otari/gateway/overview/)** - Gateway setup and deployment
- **[Python SDK](https://github.com/mozilla-ai/otari-sdk-python)** - The Python SDK
- **[TypeScript SDK](https://github.com/mozilla-ai/otari-sdk-ts)** - The TypeScript SDK for Node.js applications
- **[otari Platform (Beta)](https://otari.ai/)** - Hosted control plane for key management, usage tracking, and cost visibility

## Contributing

We welcome contributions from developers of all skill levels! Please see the [Contributing Guide](https://github.com/mozilla-ai/otari/blob/main/CONTRIBUTING.md) or open an issue to discuss changes.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
