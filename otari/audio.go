package otari

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strings"
)

// SpeechParams are the parameters for POST /v1/audio/speech (text-to-speech).
// Extra carries any additional fields (response_format, speed, instructions,
// user, ...) forwarded verbatim.
type SpeechParams struct {
	Model string         `json:"model"`
	Input string         `json:"input"`
	Voice string         `json:"voice"`
	Extra map[string]any `json:"-"`
}

// speechBody builds the JSON request body for POST /v1/audio/speech.
func speechBody(params SpeechParams) map[string]any {
	body := structToMap(params)
	for k, v := range params.Extra {
		body[k] = v
	}
	return body
}

// Speech synthesizes speech (text-to-speech) via POST /v1/audio/speech,
// returning the raw audio bytes.
//
// The gateway returns binary audio (audio/mpeg by default) with no JSON
// response model, so this posts over the raw transport (reusing the per-mode
// auth default header and unified error mapping) rather than the generated JSON
// core, and returns the response body bytes unchanged.
func (c *Client) Speech(
	ctx context.Context,
	params SpeechParams,
) ([]byte, error) {
	if params.Model == "" {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("model is required"))
	}
	if params.Input == "" {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("input is required"))
	}
	if params.Voice == "" {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("voice is required"))
	}

	raw, err := json.Marshal(speechBody(params))
	if err != nil {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("marshal request: %w", err))
	}

	_, body, err := c.rawPOST(ctx, "/audio/speech", raw, contentTypeJSON, "")
	if err != nil {
		return nil, err
	}
	return body, nil
}

// TranscriptionParams are the parameters for POST /v1/audio/transcriptions.
// File is the raw audio bytes uploaded as multipart form data; Filename names
// the upload part (some providers infer the audio format from its extension).
// Extra carries any additional fields (language, prompt, response_format,
// temperature, user, ...) sent as multipart form fields.
type TranscriptionParams struct {
	Model    string
	File     []byte
	Filename string
	Extra    map[string]any
}

// TranscriptionResult is the result of a transcription request. Exactly one
// field is populated, chosen by the gateway response's Content-Type: JSON holds
// the decoded object for the json/verbose_json formats (otherwise nil); Text
// holds the response body decoded as a string for the text/srt/vtt formats
// (otherwise empty).
type TranscriptionResult struct {
	JSON map[string]any
	Text string
}

// Transcription transcribes audio to text via POST /v1/audio/transcriptions.
//
// File is uploaded as the multipart "file" field; Model and any Extra fields
// are sent as form fields. The generated core types its file field as a string
// and cannot perform a multipart upload, so this posts over the raw transport
// (reusing the per-mode auth default header and unified error mapping).
//
// For JSON response formats the result's JSON field holds the decoded object;
// for text/srt/vtt formats the Text field holds the raw text.
func (c *Client) Transcription(
	ctx context.Context,
	params TranscriptionParams,
) (*TranscriptionResult, error) {
	if params.Model == "" {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("model is required"))
	}
	if len(params.File) == 0 {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("file is required"))
	}

	filename := params.Filename
	if filename == "" {
		filename = "audio"
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("build multipart body: %w", err))
	}
	if _, err := part.Write(params.File); err != nil {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("write multipart file: %w", err))
	}
	if err := writer.WriteField("model", params.Model); err != nil {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("write multipart field: %w", err))
	}
	for key, value := range params.Extra {
		if err := writer.WriteField(key, fmt.Sprintf("%v", value)); err != nil {
			return nil, newInvalidRequestError(providerName, fmt.Errorf("write multipart field %q: %w", key, err))
		}
	}
	if err := writer.Close(); err != nil {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("finalize multipart body: %w", err))
	}

	resp, body, err := c.rawPOST(ctx, "/audio/transcriptions", buf.Bytes(), writer.FormDataContentType(), "")
	if err != nil {
		return nil, err
	}

	result := &TranscriptionResult{}
	if strings.Contains(resp.Header.Get(contentTypeHeader), "application/json") {
		if err := json.Unmarshal(body, &result.JSON); err != nil {
			return nil, newProviderError(providerName, fmt.Errorf("decode response: %w", err))
		}
		return result, nil
	}
	result.Text = string(body)
	return result, nil
}
