# AGENTS.md

Guidance for agentic coding tools working in this repository.
Scope: entire repo.

`CLAUDE.md` is a symlink to this file. Always edit `AGENTS.md` directly; never modify `CLAUDE.md`.

## Project Snapshot
- Project: the **Go client SDK** for the otari gateway / platform.
- Module: `github.com/mozilla-ai/otari-sdk-go`; the ergonomic client is the `otari` package
  (`import "github.com/mozilla-ai/otari-sdk-go/otari"`).
- Go: 1.25. Test/assert: `github.com/stretchr/testify`.

## Architecture (Big Picture)
This SDK is a **thin, ergonomic shell over an OpenAPI-generated typed core**. Read these together
before changing request behavior.

- **Generated core** (`otari/client/`): produced by OpenAPI Generator from the gateway's OpenAPI
  spec (`api_*.go`, `model_*.go`, `client.go`, `configuration.go`). It is **generated, not
  hand-edited.** Regeneration happens upstream in the gateway repo
  (`.github/workflows/gateway-sdk-codegen.yml`), which opens a `sdk-codegen/client-core` PR here.
  The generated tree is `gofmt`'d as part of codegen, so the whole module is held to `gofmt`.
- **Hand-written shell, package `otari` (`otari/*.go`)**:
  - `client.go`: `Client`, constructor `New(opts ...Option)`, and the inference methods
    (`Completion`/`CompletionStream`, `Message`/`MessageStream`, `Response`/`ResponseStream`,
    `Embedding`, `ListModels`).
  - `config.go`: functional options (`WithAPIKey`, `WithBaseURL`, `WithHTTPClient`,
    `WithTimeout`, `WithOtariKey`, `WithPlatformMode`).
  - `streaming.go`: hand-written SSE decoder (mirrors openai-go); the generated core can't
    stream. Streaming methods return `(<-chan T, <-chan error)` and honor `context` cancellation.
  - `errors.go`: typed error hierarchy with **sentinel errors** (`ErrAuthentication`,
    `ErrRateLimit`, `ErrInsufficientFunds`, `ErrUpstreamProvider`, `ErrTimeout`,
    `ErrBatchNotComplete`, `ErrUnsupported`, ...) plus structs carrying fields (e.g.
    `BatchNotCompleteError{BatchID, Status}`); HTTP status → typed error in one place.
  - `batch.go`, `moderation.go`, `rerank.go`, `control_plane.go`: endpoint wrappers
    (`control_plane.go` exposes the management API: keys/users/budgets/pricing/usage).
  - `types.go`, `conversion.go`, `request.go`, `transport.go`: request/response types,
    constants, and HTTP helpers.

### Two auth modes (must both keep working)
- **Platform** (`WithPlatformMode()` + token, env `OTARI_AI_TOKEN`): `Authorization: Bearer`,
  base URL defaults to the hosted gateway.
- **Self-hosted** (`WithOtariKey()`, env `GATEWAY_API_KEY` / `GATEWAY_API_BASE`): `Otari-Key`
  header. Auth headers are set as defaults on the generated client's configuration.
Typed error mapping applies in both modes.

### Error-handling idiom
Callers use `errors.Is(err, otari.ErrRateLimit)` for category checks and `errors.As(err, &target)`
to read fields (e.g. `*otari.BatchNotCompleteError`). Preserve sentinel wrapping when you add or
change error paths.

### Endpoint-coverage drift gate
`otari/endpoint_coverage_test.go` fetches the canonical gateway spec
(`https://raw.githubusercontent.com/mozilla-ai/otari/main/docs/public/openapi.json`) and asserts
every endpoint is accounted for in `sdk-endpoints.txt` (`[covered]` / `[excluded]` with a reason).
Update that manifest when you add or intentionally skip an endpoint. The gate skips offline
(`OTARI_SKIP_NETWORK_TESTS=1`).

## Build / Test Commands
- Build: `go build ./...`
- Test (all): `go test ./...`
- Single test: `go test -run TestNew ./otari`
- Vet: `go vet ./...`
- Format check (CI fails on unformatted files): `gofmt -l .`
- Drift gate: `go test ./otari/ -run 'TestManifestParses|TestSpecEndpointsAreAccountedFor|TestManifestHasNoStaleEntries' -v`

There is **no `golangci-lint`** config; CI runs `gofmt`, `go vet`, `go build`, `go test`, and the
drift gate. Integration tests (e.g. `control_plane_integration_test.go`) require a real gateway
and skip when one is not available.

## Repository Conventions
- Package `otari` is the public surface; `otari/client` is generated and imported internally.
- All hand-written code must be `gofmt`-clean and pass `go vet`.
- Tests are table-driven with `t.Run` subtests, use `httptest.NewServer` to mock the gateway,
  and assert with `testify/require`. `option_c_test.go` is the reference for new-endpoint and
  error-mapping coverage.

## Change Validation Checklist
- Touched request handling, auth, or errors → `go test ./otari/...` and confirm `errors.Is`/`As`
  still classify errors in both auth modes.
- Touched streaming → exercise the channel-based stream tests; verify context cancellation closes
  channels.
- Added/removed an endpoint wrapper → update `sdk-endpoints.txt` and run the drift gate.
- Always run `gofmt -l .` (expect empty), `go vet ./...`, and `go test ./...` before a PR.

## Writing style

- Avoid em dashes and double hyphens (`--`) used as separators in prose
  (README, docs, doc comments, commit messages, PR descriptions). Use commas,
  semicolons, colons, parentheses, or periods, or rephrase. This does not apply
  to code (for example CLI flags like `--all`) or en-dash numeric ranges like `3–4`.

## Notes for Agents
- Never hand-edit `otari/client/`; it is regenerated from the gateway spec.
- Prefer minimal, targeted edits; match existing error-handling and option-pattern idioms.
