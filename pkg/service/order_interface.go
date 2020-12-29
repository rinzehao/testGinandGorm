package service

type CreateContext interface {
	Schema() string
	Param() interface{}
	GetResult() interface{}
	SetResult(interface{})
}

type UpdateContext interface {
	Schema() string
	GetIdentify() string
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

type QueryObjectsContext interface {
	Schema() string
	Page() int
	PageSize() int
	GetResult() []interface{}
	SetResult([]interface{})
}

type QueryByNameContext interface {
	Schema() string
	Param() interface{}
	Order() string                  //指出查询结果排序顺序，非订单
	Desc() bool                     //true为按照DESC顺序排序，false为按照ASC顺序排序
	GetResult() []interface{}
	SetResult([]interface{})
}

