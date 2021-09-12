package errors

func EmptyRow(reason string, message string) *Error {
	return Newf(110001, reason, message)
}

