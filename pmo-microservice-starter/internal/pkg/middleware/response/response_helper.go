package response

import "time"

// SuccessResponse membuat response sukses
func SuccessResponse[T any](data T, message string) Response[T] {
	return Response[T]{
		Success:   true,
		Message:   message,
		Data:      &data,
		Timestamp: time.Now(),
	}
}

// SuccessResponseWithMeta untuk list/pagination
func SuccessResponseWithMeta[T any](data T, message string, meta *Meta) Response[T] {
	return Response[T]{
		Success:   true,
		Message:   message,
		Data:      &data,
		Meta:      meta,
		Timestamp: time.Now(),
	}
}

// ErrorResponse membuat response error sederhana
func ErrorResponse[T any](message string) Response[T] {
	return Response[T]{
		Success:   false,
		Message:   message,
		Timestamp: time.Now(),
	}
}
func ValidationError[T any](errors []Error) Response[T] {
	return Response[T]{
		Success:   false,
		Message:   "Validation failed",
		Errors:    errors,
		Timestamp: time.Now(),
	}
}
