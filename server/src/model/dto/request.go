package dto

import "time"

type GetHistoryRequest struct {
	Provider  string    `form:"provider"`
	Token     string    `form:"token"`
	Operation string    `form:"operation"`
	APR       *float32  `form:"apr"`
	FromDate  time.Time `form:"from_date" time_format:"2006-01-02"`
	ToDate    time.Time `form:"to_date" time_format:"2006-01-02"`
}

type GetNewestHistoryRequest struct {
	Provider  string `form:"provider"`
	Token     string `form:"token"`
	Operation string `form:"operation"`
}

type GetNewestLiquidityPoolHistoryRequest struct {
	Provider string `form:"provider"`
	TokenA   string `form:"token_a"`
	TokenB   string `form:"token_b"`
}
