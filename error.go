package ayoconnect

func (e ErrorResponse) Error() string {
	return e.Message
}
