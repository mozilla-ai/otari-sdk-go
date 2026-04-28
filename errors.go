package otari

import (
	stderrors "errors"
	"fmt"
)

// Error codes used in OtariError.Code field.
const (
	codeAuthError         = "auth_error"
	codeBatchNotComplete  = "batch_not_complete"
	codeContentFilter     = "content_filter"
	codeContextLength     = "context_length_exceeded"
	codeInsufficientFunds = "insufficient_funds"
	codeInvalidRequest    = "invalid_request"
	codeMissingAPIKey     = "missing_api_key"
	codeModelNotFound     = "model_not_found"
	codeProviderError     = "provider_error"
	codeRateLimit         = "rate_limit"
	codeTimeout           = "otari_timeout"
	codeUnsupported       = "unsupported"
	codeUpstreamProvider  = "upstream_provider"
)

// Sentinel errors for type checking with errors.Is().
var (
	// ErrAuthentication is matched by errors.Is on authentication failures (HTTP 401, 403).
	ErrAuthentication = stderrors.New("authentication failed")

	// ErrBatchNotComplete is matched by errors.Is when retrieving batch
	// results on a batch that has not yet finished processing (HTTP 409).
	ErrBatchNotComplete = stderrors.New("batch not yet complete")

	// ErrContentFilter is matched by errors.Is when content is blocked by safety filters.
	ErrContentFilter = stderrors.New("content filtered")

	// ErrContextLength is matched by errors.Is when the context exceeds the model's limit.
	ErrContextLength = stderrors.New("context length exceeded")

	// ErrInsufficientFunds is matched by errors.Is on payment-required errors (HTTP 402).
	ErrInsufficientFunds = stderrors.New("insufficient funds")

	// ErrInvalidRequest is matched by errors.Is when the request is malformed.
	ErrInvalidRequest = stderrors.New("invalid request")

	// ErrMissingAPIKey is matched by errors.Is when no API key is provided.
	ErrMissingAPIKey = stderrors.New("missing API key")

	// ErrModelNotFound is matched by errors.Is when the requested model doesn't exist (HTTP 404).
	ErrModelNotFound = stderrors.New("model not found")

	// ErrProvider is matched by errors.Is for general provider-side errors.
	ErrProvider = stderrors.New("provider error")

	// ErrRateLimit is matched by errors.Is when the API rate limit is exceeded (HTTP 429).
	ErrRateLimit = stderrors.New("rate limit exceeded")

	// ErrTimeout is matched by errors.Is on otari timeout errors (HTTP 504).
	ErrTimeout = stderrors.New("otari timeout")

	// ErrUnsupported is matched by errors.Is when an operation is not supported by the provider.
	ErrUnsupported = stderrors.New("operation not supported by provider")

	// ErrUpstreamProvider is matched by errors.Is on upstream provider errors (HTTP 502).
	ErrUpstreamProvider = stderrors.New("upstream provider error")
)

// OtariError is the base error type for all otari SDK errors.
// It wraps the original error and includes provider context.
type OtariError struct {
	// Code is a short error code (e.g., "rate_limit", "auth_error").
	Code string

	// Provider is the name of the provider that returned the error.
	Provider string

	// Err is the underlying error (original provider error).
	Err error

	// sentinel is the sentinel error for errors.Is() matching.
	sentinel error
}

// newOtariError constructs an OtariError. This exists so the unexported
// sentinel field can be set by typed error constructors.
func newOtariError(code, provider string, err, sentinel error) OtariError {
	return OtariError{
		Code:     code,
		Provider: provider,
		Err:      err,
		sentinel: sentinel,
	}
}

// Error implements the error interface.
func (e *OtariError) Error() string {
	msg := ""
	if e.Err != nil {
		msg = e.Err.Error()
	}
	if e.Provider != "" {
		return fmt.Sprintf("[%s] %s: %s", e.Provider, e.Code, msg)
	}
	return fmt.Sprintf("%s: %s", e.Code, msg)
}

// Is allows checking error types with errors.Is().
func (e *OtariError) Is(target error) bool {
	return e.sentinel != nil && target == e.sentinel
}

// Unwrap returns the underlying error for errors.Is() and errors.As().
func (e *OtariError) Unwrap() error {
	return e.Err
}

// AuthenticationError is returned when authentication fails (HTTP 401, 403).
type AuthenticationError struct {
	OtariError
}

// BatchNotCompleteError is returned when RetrieveBatchResults is called on
// a batch that has not finished processing yet (HTTP 409).
type BatchNotCompleteError struct {
	OtariError
	BatchID string
	Status  string
}

// ContentFilterError is returned when content is blocked by safety filters.
type ContentFilterError struct {
	OtariError
}

// ContextLengthError is returned when the context exceeds the model's limit.
type ContextLengthError struct {
	OtariError
}

// InsufficientFundsError is returned when the account has insufficient funds (HTTP 402).
type InsufficientFundsError struct {
	OtariError
}

// InvalidRequestError is returned when the request is malformed.
type InvalidRequestError struct {
	OtariError
}

// MissingAPIKeyError is returned when no API key is provided.
type MissingAPIKeyError struct {
	OtariError
	EnvVar string // The environment variable that should contain the key
}

// ModelNotFoundError is returned when the requested model doesn't exist (HTTP 404).
type ModelNotFoundError struct {
	OtariError
}

// ProviderError is returned for general provider-side errors.
type ProviderError struct {
	OtariError
	StatusCode int
}

// RateLimitError is returned when the API rate limit is exceeded (HTTP 429).
type RateLimitError struct {
	OtariError
	RetryAfter int // Seconds until retry is allowed, if known
}

// TimeoutError is returned when the otari gateway times out (HTTP 504).
type TimeoutError struct {
	OtariError
}

// UnsupportedCapabilityError indicates a provider does not support a given
// operation (e.g. moderation). Use errors.As to detect:
//
//	var unsup *otari.UnsupportedCapabilityError
//	if errors.As(err, &unsup) {
//	    // unsup.Operation, unsup.Provider ...
//	}
//
// Also matches errors.Is(err, ErrUnsupported).
type UnsupportedCapabilityError struct {
	OtariError
	// Operation names the unsupported capability (e.g. "moderation",
	// "multimodal_moderation").
	Operation string
}

// UpstreamProviderError is returned when the upstream provider is
// unreachable (HTTP 502).
type UpstreamProviderError struct {
	OtariError
}

// newAuthenticationError creates a new AuthenticationError.
func newAuthenticationError(provider string, err error) *AuthenticationError {
	return &AuthenticationError{
		OtariError: newOtariError(codeAuthError, provider, err, ErrAuthentication),
	}
}

// newBatchNotCompleteError constructs a BatchNotCompleteError for the given
// batch and upstream-reported status. status may be empty when the server
// did not include a structured status field in the response.
func newBatchNotCompleteError(batchID, status string, cause error) *BatchNotCompleteError {
	if cause == nil {
		switch {
		case batchID != "" && status != "":
			cause = fmt.Errorf("batch %q is not yet complete (status: %s)", batchID, status)
		case batchID != "":
			cause = fmt.Errorf("batch %q is not yet complete", batchID)
		default:
			cause = stderrors.New("batch is not yet complete")
		}
	}
	return &BatchNotCompleteError{
		OtariError: newOtariError(codeBatchNotComplete, providerName, cause, ErrBatchNotComplete),
		BatchID:    batchID,
		Status:     status,
	}
}

// newContentFilterError creates a new ContentFilterError.
func newContentFilterError(provider string, err error) *ContentFilterError {
	return &ContentFilterError{
		OtariError: newOtariError(codeContentFilter, provider, err, ErrContentFilter),
	}
}

// newContextLengthError creates a new ContextLengthError.
func newContextLengthError(provider string, err error) *ContextLengthError {
	return &ContextLengthError{
		OtariError: newOtariError(codeContextLength, provider, err, ErrContextLength),
	}
}

// newInsufficientFundsError creates a new InsufficientFundsError.
func newInsufficientFundsError(provider string, err error) *InsufficientFundsError {
	return &InsufficientFundsError{
		OtariError: newOtariError(codeInsufficientFunds, provider, err, ErrInsufficientFunds),
	}
}

// newInvalidRequestError creates a new InvalidRequestError.
func newInvalidRequestError(provider string, err error) *InvalidRequestError {
	return &InvalidRequestError{
		OtariError: newOtariError(codeInvalidRequest, provider, err, ErrInvalidRequest),
	}
}

// newModelNotFoundError creates a new ModelNotFoundError.
func newModelNotFoundError(provider string, err error) *ModelNotFoundError {
	return &ModelNotFoundError{
		OtariError: newOtariError(codeModelNotFound, provider, err, ErrModelNotFound),
	}
}

// newProviderError creates a new ProviderError.
func newProviderError(provider string, err error) *ProviderError {
	return &ProviderError{
		OtariError: newOtariError(codeProviderError, provider, err, ErrProvider),
	}
}

// newRateLimitError creates a new RateLimitError.
func newRateLimitError(provider string, err error) *RateLimitError {
	return &RateLimitError{
		OtariError: newOtariError(codeRateLimit, provider, err, ErrRateLimit),
	}
}

// newUnsupportedCapabilityError creates a new UnsupportedCapabilityError for
// the given provider and operation (e.g. "moderation").
func newUnsupportedCapabilityError(provider, operation string, err error) *UnsupportedCapabilityError {
	return &UnsupportedCapabilityError{
		OtariError: newOtariError(codeUnsupported, provider, err, ErrUnsupported),
		Operation:  operation,
	}
}
