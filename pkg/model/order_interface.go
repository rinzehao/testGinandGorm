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


type CreateContext interface{
	Schema() string
	Param() interface{}
	GetResult() interface{}
	SetResult(interface{})
}

type UpdateContext interface{
	Schema() string
	GetIdentify () string
	Param() interface{}
	GetResult() interface{}
	SetResult(interface{})
}

type QueryContext interface {
	Schema() string
	Param() interface{}
	GetResult() interface{}
	SetResult(interface{})
}

type DeleteContext interface {
	Schema() string
	Param() interface{}
	GetResult() interface{}
	SetResult(interface{})
}

type QueryOrdersContext interface {
	Schema() string
	Page() interface{}
	PageSize() interface{}
	GetResult() []interface{}
	SetResult([]interface{})
}

