<p align="center">
  <img src="assets/otari-logo.svg" width="320" alt="otari logo"/>
</p>

<div align="center">

# otari (Go)

![Go 1.25+](https://img.shields.io/badge/go-1.25%2B-blue.svg)
<a href="https://discord.gg/4gf3zXrQUc">
    <img src="https://img.shields.io/static/v1?label=Chat%20on&message=Discord&color=blue&logo=Discord&style=flat-square" alt="Discord">
</a>

**Go client for [otari-gateway](https://github.com/mozilla-ai/otari).**
Communicate with any LLM provider through the gateway using a single, typed interface.

[Python SDK](https://github.com/mozilla-ai/otari-sdk-python) | [TypeScript SDK](https://github.com/mozilla-ai/otari-sdk-ts) | [Documentation](https://mozilla-ai.github.io/otari/) | [Platform (Beta)](https://otari.ai/)

</div>

## Quickstart

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

That's it — the client defaults to the hosted gateway at `https://api.otari.ai`, so you don't need to set a base URL. Change the model string to switch between LLM providers through the gateway.

Prefer to keep secrets out of code? Export `OTARI_AI_TOKEN` and call `otari.New(otari.WithPlatformMode())` — the token is picked up from the environment.

## Self-hosting the gateway

Prefer to run the gateway yourself instead of using the hosted otari.ai? Follow the setup in the [otari gateway repo](https://github.com/mozilla-ai/otari), then point the SDK at it:

```go
client, err := otari.New(
    otari.WithBaseURL("http://localhost:8000"), // or wherever you host the gateway
    otari.WithOtariKey("your-gateway-api-key"),
)
```

The SDK sends the key via the custom `Otari-Key: Bearer …` header. Env: `GATEWAY_API_BASE` + `GATEWAY_API_KEY`.

Make sure your gateway has provider keys configured (e.g. OpenAI) so it can route requests upstream — see the [otari gateway repo](https://github.com/mozilla-ai/otari) for setup.

## Installation

### Requirements

- Go 1.25 or newer
- An [otari.ai](https://otari.ai/) account, or a running self-hosted [otari-gateway](https://mozilla-ai.github.io/otari/gateway/overview/) instance

### Install

```bash
go get github.com/mozilla-ai/otari-sdk-go/otari
```

### Setting Up Credentials

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

Alternatively, pass credentials directly when creating the client (see [Usage](#usage) examples).

## otari-gateway

This Go SDK is a client for [otari-gateway](https://github.com/mozilla-ai/otari), an **optional** FastAPI-based proxy server that adds enterprise-grade features on top of the core library:

- **Budget Management** - Enforce spending limits with automatic daily, weekly, or monthly resets
- **API Key Management** - Issue, revoke, and monitor virtual API keys without exposing provider credentials
- **Usage Analytics** - Track every request with full token counts, costs, and metadata
- **Multi-tenant Support** - Manage access and budgets across users and teams

The gateway sits between your applications and LLM providers, exposing an OpenAI-compatible API that works with any supported provider.

### Quick Start

```bash
docker run \
  -e GATEWAY_MASTER_KEY="your-secure-master-key" \
  -e OPENAI_API_KEY="your-api-key" \
  -p 8000:8000 \
  ghcr.io/mozilla-ai/otari/gateway:latest
```

> **Note:** You can use a specific release version instead of `latest` (e.g., `1.2.0`). See [available versions](https://github.com/orgs/mozilla-ai/packages/container/package/otari%2Fgateway).

### Managed Platform (Beta)

Prefer a hosted experience? The [otari platform](https://otari.ai/) provides a managed control plane for keys, usage tracking, and cost visibility across providers, while still building on the same `otari` interfaces.

## Usage

### Authentication Modes

The client supports two authentication modes, matching the Python and TypeScript SDKs:

#### Platform Mode (Recommended)

Uses a Bearer token in the standard Authorization header. On the hosted platform, generate an API token at [otari.ai/organization-settings/api-tokens](https://otari.ai/organization-settings/api-tokens) and add a provider key (e.g. OpenAI) at [otari.ai/organization-settings/provider-keys](https://otari.ai/organization-settings/provider-keys) so the gateway can route requests to that provider. The base URL defaults to the hosted gateway at `https://api.otari.ai`, so it can be omitted:

```go
client, err := otari.New(
    otari.WithAPIKey("tk_your_api_token"),
    otari.WithPlatformMode(),
)
```

#### Non-Platform Mode (Self-hosted)

Sends the API key via a custom `Otari-Key` header. A base URL is required:

```go
client, err := otari.New(
    otari.WithBaseURL("http://localhost:8000"),
    otari.WithOtariKey("your-api-key"),
)
```

#### Auto-Detection from Environment Variables

When no explicit credentials are provided, the client reads from environment variables:

```go
// Platform mode: OTARI_AI_TOKEN (or legacy GATEWAY_PLATFORM_TOKEN) → hosted gateway.
// Self-hosted: GATEWAY_API_BASE + GATEWAY_API_KEY.
client, err := otari.New()
```

### Chat Completions

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

### Listing Models

```go
models, err := client.ListModels(ctx)
if err != nil {
    log.Fatal(err)
}
for _, m := range models.Data {
    fmt.Println(m.ID)
}
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

### Content Moderation

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

### Batch Operations

```go
// Create a batch
batch, err := client.CreateBatch(ctx, otari.CreateBatchParams{
    Model: "openai:gpt-4o-mini",
    Requests: []otari.BatchRequestItem{
        {CustomID: "req-1", Body: map[string]any{"messages": []any{
            map[string]any{"role": "user", "content": "Hello"},
        }}},
    },
    CompletionWindow: "24h",
})

// Retrieve results when complete
results, err := client.RetrieveBatchResults(ctx, batch.ID, batch.Provider)
```

### Error Handling

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
├── otari/                  # Main package
│   ├── client.go           # Client struct, constructor, core methods
│   ├── config.go           # Functional options (WithBaseURL, WithAPIKey, etc.)
│   ├── conversion.go       # OpenAI SDK type conversions
│   ├── errors.go           # Error hierarchy with sentinel errors
│   ├── types.go            # Request/response types
│   ├── batch.go            # Batch API methods
│   ├── moderation.go       # Content moderation
│   ├── rerank.go           # Document reranking
│   ├── transport.go        # HTTP transport helpers
│   ├── otari.go            # Package documentation
│   └── client_test.go      # Tests
├── examples/
│   └── quickstart/         # Quick smoke test
├── go.mod
├── go.sum
└── README.md
```

## Why choose `otari`?

- **Simple, unified interface** - Single client for all providers through the gateway, switch models with just a string change
- **Developer friendly** - Full type safety with Go's type system and clear, actionable error messages
- **Leverages the OpenAI SDK** - Built on the official [openai-go](https://github.com/openai/openai-go) SDK for maximum compatibility
- **Stays framework-agnostic** so it can be used across different projects and use cases
- **Battle-tested** - Extracted from [any-llm-go](https://github.com/mozilla-ai/any-llm-go)'s production gateway provider

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
