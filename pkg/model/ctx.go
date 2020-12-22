package model

type OrderMould struct {
	ID       int     `json:"id" gorm:"primaryKey"`
	OrderNo  string  `json:"order_no" gorm:"unique;not null,type:varchar(30)"`
	UserName string  `json:"user_name" gorm:"type:varchar(30)"`
	Amount   float64 `json:"amount" gorm:"type:float(10,2)"`
	Status   string  `json:"status" gorm:"type:varchar(30)"`
	FileUrl  string  `json:"file_url" gorm:"type:varchar(120)"`
}

func (o *OrderMould) ID_() int {
	return o.ID
}

func (o *OrderMould) UserName_() string {
	return o.OrderNo
}

func (o *OrderMould) OrderNo_() string {
	return o.OrderNo
}

func (o *OrderMould) Amount_() float64 {
	return o.Amount
}

func (o *OrderMould) Status_() string {
	return o.Status
}

func (o *OrderMould) FileUrl_() string {
	return o.FileUrl
}

