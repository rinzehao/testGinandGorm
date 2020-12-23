package model

type OrderInterface interface {
	GetOrder() interface{}
	GetOrderID() string
	GetOrderNo() string
	GetUpdateMap() map[string]interface{}
	GetPage() int
	GetPageSize() int
	GetUserName() string
	GetOrderBy() string
	GetDesc() string
	GetGroup() []*DemoOrder
}
