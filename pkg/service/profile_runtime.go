package service

import (
	"testGinandGorm/common/logger"
)

type ProfileItem interface {
	Schema() string
	Delete(DeleteContext) error
	Query(QueryContext) error
	UpdateByNo(UpdateContext) error
	Create(CreateContext) error
	QueryOrders(QueryObjectsContext) error
	QueryOrdersByName(QueryByNameContext) error
	UpdateById(UpdateContext) error
}

type ProfileRuntime struct {
	invoker map[string]ProfileItem
}

func NewProfileRuntime(items ...ProfileItem) *ProfileRuntime {
	if len(items) == 0 {
		logger.SugarLogger.Panic("profile item need item command")
	}
	pr := &ProfileRuntime{invoker: make(map[string]ProfileItem)}
	for _, v := range items {
		pr.invoker[v.Schema()] = v
	}
	return pr
}

func (p *ProfileRuntime) Push(ctx CreateContext) error {
	rt, ok := p.invoker[ctx.Schema()]
	if !ok {
		logger.SugarLogger.Panicf("not found profile schema:%s", ctx.Schema())
		//panic(fmt.Sprintf("not found profile schema:%s", ctx.Schema()))
	}
	return rt.Create(ctx)
}

func (p *ProfileRuntime) QueryOrders(ctx QueryObjectsContext) error {
	rt, ok := p.invoker[ctx.Schema()]
	if !ok {
		logger.SugarLogger.Panicf("not found profile schema:%s", ctx.Schema())
		//panic(fmt.Sprintf("not found profile schema:%s", ctx.Schema()))
	}
	return rt.QueryOrders(ctx)
}

func (p *ProfileRuntime) UpdateByNo(ctx UpdateContext) error {
	rt, ok := p.invoker[ctx.Schema()]
	if !ok {
		logger.SugarLogger.Panicf("not found profile schema:%s", ctx.Schema())
	}
	return rt.UpdateByNo(ctx)
}

func (p *ProfileRuntime) UpdateById(ctx UpdateContext) error {
	rt, ok := p.invoker[ctx.Schema()]
	if !ok {
		logger.SugarLogger.Panicf("not found profile schema:%s", ctx.Schema())
	}
	return rt.UpdateById(ctx)
}

func (p *ProfileRuntime) Delete(ctx DeleteContext) error {
	rt, ok := p.invoker[ctx.Schema()]
	if !ok {
		logger.SugarLogger.Panicf("not found profile schema:%s", ctx.Schema())

	}
	return rt.Delete(ctx)
}

func (p *ProfileRuntime) QueryById(ctx QueryContext) error {
	rt, ok := p.invoker[ctx.Schema()]
	if !ok {
		logger.SugarLogger.Panicf("not found profile schema:%s", ctx.Schema())
	}
	return rt.Query(ctx)
}

func (p *ProfileRuntime) QueryByName(ctx QueryByNameContext) error {
	rt, ok := p.invoker[ctx.Schema()]
	if !ok {
		logger.SugarLogger.Panicf("not found profile schema:%s", ctx.Schema())
	}
	return rt.QueryOrdersByName(ctx)
}
