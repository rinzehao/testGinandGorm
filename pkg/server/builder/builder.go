package builder

import (
	"testGinandGorm/common/mySQL_db"
	"testGinandGorm/common/redis"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/db"
	"testGinandGorm/pkg/service"
)

type BuilderService struct {
	orderService service.Service
}

func NewBuilderService() *BuilderService {
	s := &BuilderService{}
	buildMysql(s)
	return s
}

func buildMysql(s *BuilderService) {
	sqlDB := mySQL_db.DbInit()
	defer sqlDB.Close()
	orderDB := db.NewMyOrderDB(sqlDB)
	orderCache := redis.NewRedisCache(1e10 * 6 * 20)
	orderDao := dao.NewMyOrderDao(orderDB, &orderCache)
	orderService := service.NewOrderService(orderDao)
	//ordeHandler := handler.NewOrderHandler(*orderService)
	s.orderService = orderService
}
