package model

type OrderInterface interface {
	QueryOrder() interface{}
	QueryOrderID() string
	QueryOrderNo() string
	QueryUpdateMap() map[string]interface{}
	QueryPage() int
	QueryPageSize() int
	QueryUserName() string
	QueryOrderBy() string
	QueryDesc() string
	QueryGroup() []*DemoOrder
}
