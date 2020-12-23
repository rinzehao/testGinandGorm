package model

type OrderMade struct {
	Order     interface{}
	OrderID   string
	OrderNo   string
	UpdateMap map[string]interface{}
	Page      int
	PageSize  int
	UserName  string
	OrderBy   string
	Desc      string
	Group     []*DemoOrder
}

func (o *OrderMade) GetOrder() interface{} {
	return &o.Order
}

func (o *OrderMade) GetOrderID() string {
	return o.OrderID
}

func (o *OrderMade) GetOrderNo() string {
	return o.OrderNo
}

func (o *OrderMade) GetUpdateMap() map[string]interface{} {
	return o.UpdateMap
}

func (o *OrderMade) GetPage() int {
	return o.Page
}

func (o *OrderMade) GetPageSize() int {
	return o.PageSize
}

func (o *OrderMade) GetUserName() string {
	return o.UserName
}

func (o *OrderMade) GetOrderBy() string {
	return o.OrderBy
}

func (o *OrderMade) GetDesc() string {
	return o.Desc
}

func (o *OrderMade) GetGroup() []*DemoOrder {
	return o.Group
}
