package otari

import "fmt"

// completionBody builds the JSON request body for POST /v1/chat/completions
// from the ergonomic CompletionParams. The params struct already carries the
// correct JSON tags for the OpenAI-compatible wire shape; this assembles a map
// so the stream flag and any Extra fields can be layered in, and so zero-value
// fields are omitted via the struct's own omitempty tags.
//
// Fields not modeled by the gateway are forwarded as-is; the gateway silently
// drops params it does not model (e.g. seed), matching its documented behavior.
func completionBody(params CompletionParams, stream bool) map[string]any {
	body := structToMap(params)
	for k, v := range params.Extra {
		body[k] = v
	}
	if stream {
		body["stream"] = true
	} else {
		delete(body, "stream")
	}
	return body
}

// MessageParams are the parameters for an Anthropic-shaped /messages request.
// max_tokens is required by the gateway. Extra carries any additional
// /messages fields (system, tools, thinking, ...) forwarded verbatim.
type MessageParams struct {
	Model         string           `json:"model"`
	Messages      []map[string]any `json:"messages"`
	MaxTokens     int              `json:"max_tokens"`
	System        any              `json:"system,omitempty"`
	Temperature   *float64         `json:"temperature,omitempty"`
	TopP          *float64         `json:"top_p,omitempty"`
	TopK          *int             `json:"top_k,omitempty"`
	StopSequences []string         `json:"stop_sequences,omitempty"`
	Tools         []map[string]any `json:"tools,omitempty"`
	ToolChoice    map[string]any   `json:"tool_choice,omitempty"`
	Thinking      map[string]any   `json:"thinking,omitempty"`
	Metadata      map[string]any   `json:"metadata,omitempty"`
	Extra         map[string]any   `json:"-"`
}

// messageBody builds the JSON request body for POST /v1/messages.
func messageBody(params MessageParams, stream bool) map[string]any {
	body := structToMap(params)
	for k, v := range params.Extra {
		body[k] = v
	}
	if stream {
		body["stream"] = true
	}
	return body
}

// ResponseParams are the parameters for an OpenAI-style /responses request.
// Input accepts any JSON-serializable value (string, message list, ...).
// Extra carries additional /responses fields forwarded verbatim.
type ResponseParams struct {
	Model string         `json:"model"`
	Input any            `json:"input"`
	Extra map[string]any `json:"-"`
}

// responseBody builds the JSON request body for POST /v1/responses.
func responseBody(params ResponseParams, stream bool) map[string]any {
	body := structToMap(params)
	for k, v := range params.Extra {
		body[k] = v
	}
	if stream {
		body["stream"] = true
	}
	return body
}

// validateCompletionParams validates chat completion parameters before sending.
func validateCompletionParams(params CompletionParams) error {
	if params.Model == "" {
		return newInvalidRequestError(providerName, fmt.Errorf("model is required"))
	}
	if len(params.Messages) == 0 {
		return newInvalidRequestError(providerName, fmt.Errorf("at least one message is required"))
	}
	for _, msg := range params.Messages {
		switch msg.Role {
		case RoleAssistant, RoleSystem, RoleTool, RoleUser:
		default:
			return newInvalidRequestError(providerName, fmt.Errorf("unknown message role: %q", msg.Role))
		}
	}
	return nil
}
