package model

type Demo_order struct {
	ID        int     `json:"id"`
	Order_No  string  `json:"order_no"`
	User_name string  `json:"user_name"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	File_url  string  `json:"file_url"`
}

const (
	SR_File_Max_Bytes = 1024 * 1024 * 4
)
