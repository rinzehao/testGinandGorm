package main

import (
	_ "github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testGinandGorm/common/mysql"
	"testGinandGorm/common/redis"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/dao/mysql"
	"testGinandGorm/pkg/server/grpc/pb"
	"testGinandGorm/pkg/server/grpc/server/rpc-handler"
	"testGinandGorm/pkg/service"
	"testGinandGorm/pkg/service/profile"
	"testGinandGorm/pkg/service/profile/profile-item"
)

const (
	Address = "127.0.0.1:3031"
)

func main() {
	Db := mysql.DbInit()
	orderDB := mysql.NewOrderDB(Db)
	orderCache := redis.NewRedisCache(redis.DEFAULT)
	orderDao := dao.NewOrderDao(orderDB, &orderCache)
	orderService := profile_item.NewOrderService(orderDao)
	runtime := profile.NewProfileRuntime(orderService)
	go crudService(runtime)
	select {}
}

func crudService(runtime *service.ProfileRuntime) {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s2 := grpc.NewServer()
	// 服务注册
	pb.RegisterOrderServer(s2, rpc_handler.NewOrderHandler(runtime))
	reflection.Register(s2)
	if err := s2.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}