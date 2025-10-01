package errorhandler

type AppHttpError struct {
	Code    int
	Status  string
	Message string
}

func (e *AppHttpError) Error() string {
	return e.Message
}
