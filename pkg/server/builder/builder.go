package builder

import (
	"testGinandGorm/common/mySQL"
	"testGinandGorm/common/redis"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/db"
	"testGinandGorm/pkg/service"
	"testGinandGorm/pkg/service/profile"
)

type BuilderService struct {
	ProfileRuntime *service.ProfileRuntime
}

func NewService() *BuilderService {
	s := &BuilderService{}
	buildMysql(s)
	return s
}

func buildMysql(s *BuilderService) {
	sqlDB := mySQL.DbInit()
	orderDB := db.NewOrderDB(sqlDB)
	orderCache := redis.NewRedisCache(redis.DEFAULT)
	orderDao := dao.NewOrderDao(orderDB, &orderCache)
	orderService := profile.NewOrderService(orderDao)
	s.ProfileRuntime = service.NewProfileRuntime(orderService)
}
