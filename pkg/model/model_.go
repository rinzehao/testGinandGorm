package model

type Demo_order struct {
	ID       int     `json:"id"`
	OrderNo  string  `json:"order_no"`
	UserName string  `json:"user_name"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	FileUrl  string  `json:"file_url"`
}

const (
	SR_File_Max_Bytes = 1024 * 1024 * 4
)
