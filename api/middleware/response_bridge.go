package middleware

type ResponseBridge struct {
	Error      error `json:"error"`
	StatusCode int   `json:"statusCode"`
	Data       any   `json:"data"`
}
