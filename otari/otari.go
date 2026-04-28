// Package otari provides the Go SDK for the otari gateway.
//
// The otari gateway is a proxy server that sits between applications and LLM
// providers, adding budget management, API key management, usage analytics,
// and multi-tenancy. This SDK provides a typed Go client for interacting with
// the gateway.
//
// # Quick Start
//
//	import "github.com/mozilla-ai/otari-sdk-go"
//
//	client, err := otari.New(
//	    otari.WithBaseURL("http://localhost:8000"),
//	    otari.WithAPIKey("tk_your_token"),
//	    otari.WithPlatformMode(),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	resp, err := client.Completion(ctx, otari.CompletionParams{
//	    Model:    "openai:gpt-4o-mini",
//	    Messages: []otari.Message{{Role: otari.RoleUser, Content: "Hello!"}},
//	})
//
// # Authentication Modes
//
// The SDK supports two authentication modes:
//
//   - Platform mode: uses a platform token as standard Bearer authentication.
//     Enable with [WithPlatformMode] and provide the token via [WithAPIKey]
//     or the GATEWAY_PLATFORM_TOKEN environment variable.
//
//   - Non-platform mode: sends an API key via the Otari-Key header. Provide
//     the key via [WithOtariKey] or the GATEWAY_API_KEY environment variable.
//
// # Error Handling
//
// All errors returned by the SDK are typed. Use [errors.Is] to check for
// sentinel errors (e.g. [ErrRateLimit], [ErrAuthentication]) and [errors.As]
// to extract typed error details (e.g. [*BatchNotCompleteError]).
package otari
