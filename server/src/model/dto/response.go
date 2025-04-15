package dto

type HistoryResponse struct {
	Provider  string  `json:"provider"`
	Token     string  `json:"token"`
	Operation string  `json:"operation"`
	APR       float32 `json:"apr"`
	CreatedAt string  `json:"created_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
