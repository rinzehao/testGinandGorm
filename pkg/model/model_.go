package model

type DemoOrder struct {
	ID       int     `json:"id"`
	OrderNo  string  `json:"order_no"`
	UserName string  `json:"user_name"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	FileUrl  string  `json:"file_url"`
}
