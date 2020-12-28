package model

type Order struct {
	ID       int     `json:"id" gorm:"primaryKey"`
	OrderNo  string  `json:"order_no" gorm:"unique;not null,type:varchar(30)"`
	UserName string  `json:"user_name" gorm:"type:varchar(30)"`
	Amount   float64 `json:"amount" gorm:"type:float(10,2)"`
	Status   string  `json:"status" gorm:"type:varchar(30)"`
	FileUrl  string  `json:"file_url" gorm:"type:varchar(120)"`
}

func (Order) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "demo_order"
}
