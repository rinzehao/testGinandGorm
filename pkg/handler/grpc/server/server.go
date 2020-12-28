package main

import (
	_ "github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testGinandGorm/common/mySQL"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/handler/grpc/pb"
	"testGinandGorm/pkg/handler/grpc/server/rpc-handler"
	"testGinandGorm/pkg/service"
	"testGinandGorm/pkg/service/profile"
)

const (
	Address = "127.0.0.1:3031"
)

func main() {
	Db := mySQL.DbInit()
	orderDao := dao.NewOrderDao(Db)
	orderService := service.NewService(orderDao)
	go crudService(orderService)
	select {}
}

func crudService(orderService *profile.OrderService) {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s2 := grpc.NewServer()
	// 服务注册
	pb.RegisterOrderServer(s2, rpc_handler.NewHandler(orderService))
	reflection.Register(s2)
	if err := s2.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
