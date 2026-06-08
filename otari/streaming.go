package otari

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// sseDecoder is a Server-Sent-Events decoder over a raw HTTP response body.
//
// The generated client buffers full responses and cannot stream, so streaming
// is hand-written here. This decoder mirrors the official openai-go
// packages/ssestream eventStreamDecoder: it scans line by line, accumulates
// "data:" lines into the current event's payload, records the "event:" name,
// ignores ":"-comments and unknown fields, and dispatches an event on a blank
// line. The gateway emits single-line events, but mirroring openai-go's
// framing keeps the decoder robust to multi-line data and named events.
type sseDecoder struct {
	rc  io.ReadCloser
	scn *bufio.Scanner
	evt sseEvent
	err error
}

// sseEvent is one decoded SSE event: an optional event name and the raw data
// payload (newline-joined when the event spanned multiple "data:" lines).
type sseEvent struct {
	Type string
	Data []byte
}

// newSSEDecoder builds a decoder over res.Body. The buffer is enlarged the same
// way openai-go does so large single-line JSON events are not truncated.
func newSSEDecoder(res *http.Response) *sseDecoder {
	scn := bufio.NewScanner(res.Body)
	scn.Buffer(nil, bufio.MaxScanTokenSize<<9)
	return &sseDecoder{rc: res.Body, scn: scn}
}

// Next advances to the next event, returning false at end-of-stream or on error.
// Mirrors openai-go's eventStreamDecoder.Next framing.
func (d *sseDecoder) Next() bool {
	if d.err != nil {
		return false
	}

	event := ""
	data := bytes.NewBuffer(nil)

	for d.scn.Scan() {
		txt := d.scn.Bytes()

		// A blank line dispatches the accumulated event.
		if len(txt) == 0 {
			d.evt = sseEvent{Type: event, Data: data.Bytes()}
			return true
		}

		// Split "field: value"; consume a single optional leading space.
		name, value, _ := bytes.Cut(txt, []byte(":"))
		if len(value) > 0 && value[0] == ' ' {
			value = value[1:]
		}

		switch string(name) {
		case "":
			// ": comment" line — ignore.
			continue
		case "event":
			event = string(value)
		case "data":
			if _, err := data.Write(value); err != nil {
				d.err = err
				break
			}
			if _, err := data.WriteRune('\n'); err != nil {
				d.err = err
				break
			}
		}
	}

	if d.scn.Err() != nil {
		d.err = d.scn.Err()
	}
	return false
}

// Event returns the most recently decoded event.
func (d *sseDecoder) Event() sseEvent { return d.evt }

// Close closes the underlying response body.
func (d *sseDecoder) Close() error { return d.rc.Close() }

// Err returns the terminal decode error, if any.
func (d *sseDecoder) Err() error { return d.err }

// eventData returns the trimmed payload of an event, dropping the trailing
// newline the decoder appends. Returns ("", true) for the [DONE] sentinel.
func eventData(e sseEvent) (payload []byte, done bool) {
	p := bytes.TrimRight(e.Data, "\n")
	if bytes.HasPrefix(p, []byte("[DONE]")) {
		return nil, true
	}
	return p, false
}

// streamSSE opens a raw streaming POST through the generated client's
// configured transport (carrying the per-mode auth default header) and invokes
// onEvent for each decoded SSE data payload, stopping at the [DONE] sentinel.
// A failed (>=400) response is mapped through the same unified error mapper.
func streamSSE(
	ctx context.Context,
	c *Client,
	path string,
	body any,
	onEvent func(payload []byte) error,
) error {
	resp, errBody, err := c.doRaw(ctx, "POST", path, body, true)
	if err != nil {
		return err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		_ = resp.Body.Close()
		return mapHTTPError(resp.StatusCode, resp.Header, errBody, "")
	}

	dec := newSSEDecoder(resp)
	defer func() { _ = dec.Close() }()

	for dec.Next() {
		payload, done := eventData(dec.Event())
		if done {
			return nil
		}
		if len(payload) == 0 {
			continue
		}
		if err := onEvent(payload); err != nil {
			return err
		}
	}
	if derr := dec.Err(); derr != nil {
		return newProviderError(providerName, derr)
	}
	return nil
}

// jsonUnmarshal decodes an SSE payload into v, wrapping decode failures as a
// provider error.
func jsonUnmarshal(payload []byte, v any) error {
	if err := json.Unmarshal(payload, v); err != nil {
		return newProviderError(providerName, err)
	}
	return nil
}
