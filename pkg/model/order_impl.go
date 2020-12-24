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

func (o *OrderMade) QueryOrder() interface{} {
	return &o.Order
}

func (o *OrderMade) QueryOrderID() string {
	return o.OrderID
}

func (o *OrderMade) QueryOrderNo() string {
	return o.OrderNo
}

func (o *OrderMade) QueryUpdateMap() map[string]interface{} {
	return o.UpdateMap
}

func (o *OrderMade) QueryPage() int {
	return o.Page
}

func (o *OrderMade) QueryPageSize() int {
	return o.PageSize
}

func (o *OrderMade) QueryUserName() string {
	return o.UserName
}

func (o *OrderMade) QueryOrderBy() string {
	return o.OrderBy
}

func (o *OrderMade) QueryDesc() string {
	return o.Desc
}

func (o *OrderMade) QueryGroup() []*DemoOrder {
	return o.Group
}
