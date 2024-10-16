package middleware

type ResponseBridge struct {
	Error error `json:"error"`
	Data  any   `json:"data"`
}

func NewResponseBridge(err error, data any) *ResponseBridge {
	return &ResponseBridge{
		Error: err,
		Data:  data,
	}
}
