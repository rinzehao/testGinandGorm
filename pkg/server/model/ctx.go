package model

type CreateCtx struct {
	ItemTyp string
	Req     interface{}
	Result  interface{}
}

func (c *CreateCtx) Schema() string {
	return c.ItemTyp
}

func (c *CreateCtx) Param() interface{} {
	return c.Req
}

func (c *CreateCtx) GetResult() interface{} {
	return c.Result
}

func (c *CreateCtx) SetResult(d interface{}) {
	c.Result = d
}

type UpdateCtx struct {
	ItemTyp  string
	Identify string
	Req      interface{}
	Result   interface{}
}

func (c *UpdateCtx) Schema() string {
	return c.ItemTyp
}

func (c *UpdateCtx) Param() interface{} {
	return c.Req
}

func (c *UpdateCtx) GetResult() interface{} {
	return c.Result
}

func (c *UpdateCtx) SetResult(d interface{}) {
	c.Result = d
}

func (c *UpdateCtx) GetIdentify() string {
	return c.Identify
}

type QueryCtx struct {
	ItemTyp string
	Req     interface{}
	Result  interface{}
}

func (c *QueryCtx) Schema() string {
	return c.ItemTyp
}

func (c *QueryCtx) Param() interface{} {
	return c.Req
}

func (c *QueryCtx) GetResult() interface{} {
	return c.Result
}

func (c *QueryCtx) SetResult(d interface{}) {
	c.Result = d
}

type QueryCtxs struct {
	ItemTyp string
	ReqPage int
	ReqSize int
	Result  []interface{}
}

func (c *QueryCtxs) Schema() string {
	return c.ItemTyp
}

func (c *QueryCtxs) Page() int {
	return c.ReqPage
}

func (c *QueryCtxs) PageSize() int {
	return c.ReqSize
}

func (c *QueryCtxs) GetResult() []interface{} {
	return c.Result
}

func (c *QueryCtxs) SetResult(d []interface{}) {
	c.Result = d
}

type DeleteCtx struct {
	ItemTyp string
	Req     interface{}
	Result  interface{}
}

func (c *DeleteCtx) Schema() string {
	return c.ItemTyp
}

func (c *DeleteCtx) Param() interface{} {
	return c.Req
}

func (c *DeleteCtx) GetResult() interface{} {
	return c.Result
}

func (c *DeleteCtx) SetResult(d interface{}) {
	c.Result = d
}

type QueryByNameCtx struct {
	ItemTyp     string
	Req         interface{}
	Result      []interface{}
	OrderOption string
	DescOrder   bool
}

func (c *QueryByNameCtx) Schema() string {
	return c.ItemTyp
}

func (c *QueryByNameCtx) Param() interface{} {
	return c.Req
}

func (c *QueryByNameCtx) Order() string {
	return c.OrderOption
}

func (c *QueryByNameCtx) Desc() bool {
	return c.DescOrder
}

func (c *QueryByNameCtx) SetResult(d []interface{}) {
	c.Result = d
}

func (c *QueryByNameCtx) GetResult() []interface{} {
	return c.Result
}
