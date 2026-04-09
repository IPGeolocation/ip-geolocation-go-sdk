package ipgeolocation

import "fmt"

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type TransportError struct {
	Message string
	Err     error
}

func (e *TransportError) Error() string {
	return e.Message
}

func (e *TransportError) Unwrap() error {
	return e.Err
}

type RequestTimeoutError struct {
	TransportError
}

type SerializationError struct {
	Message string
	Err     error
}

func (e *SerializationError) Error() string {
	return e.Message
}

func (e *SerializationError) Unwrap() error {
	return e.Err
}

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return e.Message
}

type BadRequestError struct{ APIError }
type UnauthorizedError struct{ APIError }
type ForbiddenError struct{ APIError }
type NotFoundError struct{ APIError }
type MethodNotAllowedError struct{ APIError }
type ContentTooLargeError struct{ APIError }
type UnsupportedMediaTypeError struct{ APIError }
type LockedError struct{ APIError }
type TooManyRequestsError struct{ APIError }
type CustomStatus499Error struct{ APIError }
type InternalServerError struct{ APIError }

func apiErrorMessage(statusCode int) string {
	switch statusCode {
	case 400:
		return "Bad request"
	case 401:
		return "Unauthorized"
	case 403:
		return "Forbidden"
	case 404:
		return "Not found"
	case 405:
		return "Method not allowed"
	case 413:
		return "Content too large"
	case 415:
		return "Unsupported media type"
	case 423:
		return "Locked"
	case 429:
		return "Too many requests"
	case 499:
		return "Client closed request"
	case 500:
		return "Internal server error"
	default:
		return fmt.Sprintf("API request failed with status code %d", statusCode)
	}
}
