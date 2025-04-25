package dto

type HistoryResponse struct {
	Provider  string  `json:"provider"`
	Token     string  `json:"token"`
	Operation string  `json:"operation"`
	Link      string  `json:"link"`
	APR       float32 `json:"apr"`
	CreatedAt string  `json:"created_at"`
}

type HistoryLinkResponse struct {
	Provider  string `json:"provider"`
	Token     string `json:"token"`
	Operation string `json:"operation"`
	Link      string `json:"link"`
}

type LiquidityPoolHistoryResponse struct {
	Provider  string  `json:"provider"`
	TokenA    string  `json:"token_a"`
	TokenB    string  `json:"token_b"`
	Link      string  `json:"link"`
	APR       float32 `json:"apr"`
	CreatedAt string  `json:"created_at"`
}

type LiquidityPoolHistoryLinkResponse struct {
	Provider string `json:"provider"`
	TokenA   string `json:"token_a"`
	TokenB   string `json:"token_b"`
	Link     string `json:"link"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
