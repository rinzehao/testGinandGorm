package model

type DemoOrder struct {
	ID       int     `json:"id" gorm:"primaryKey"`
 	OrderNo  string  `json:"order_no" gorm:"unique;not null,type:varchar(30)"`
	UserName string  `json:"user_name" gorm:"type:varchar(30)"`
	Amount   float64 `json:"amount" gorm:"type:float(10,2)"`
	Status   string  `json:"status" gorm:"type:varchar(30)"`
	FileUrl  string  `json:"file_url" gorm:"type:varchar(120)"`
}

