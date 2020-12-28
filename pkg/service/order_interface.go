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
	Order() string
	Desc() bool
	GetResult() []interface{}
	SetResult([]interface{})
}
