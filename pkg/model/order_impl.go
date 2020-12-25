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



type CreateCtx struct{
	ItemTyp string
	Req interface{}
	Result interface{}
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

func (c *CreateCtx) SetResult(d interface{})  {
	c.Result = d
}


type UpdateCtx struct{
	ItemTyp string
	Identify string
	Req interface{}
	Result interface{}
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

func (c *UpdateCtx) SetResult(d interface{})  {
	c.Result = d
}

func (c *UpdateCtx) GetIdentify() string {
	return c.Identify
}

type QueryCtx struct{
	ItemTyp string
	Req interface{}
	Result interface{}
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

func (c *QueryCtx) SetResult(d interface{})  {
	c.Result = d
}



type QueryCtxs struct{
	ItemTyp string
	ReqPage interface{}
	ReqSize interface{}
	Result []interface{}
}

func (c *QueryCtxs) Schema() string {
	return c.ItemTyp
}

func (c *QueryCtxs) Page() interface{} {
	return c.ReqPage
}

func (c *QueryCtxs) PageSize() interface{} {
	return c.ReqPage
}

func (c *QueryCtxs) GetResult() []interface{} {
	return c.Result
}

func (c *QueryCtxs) SetResult(d []interface{})  {
	c.Result = d
}

type DeleteCtx struct{
	ItemTyp string
	Req interface{}
	Result interface{}
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

func (c *DeleteCtx) SetResult(d interface{})  {
	c.Result = d
}