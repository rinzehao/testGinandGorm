package main

import (
	_ "github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testGinandGorm/common/db"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/handler/grpc/pb"
	"testGinandGorm/pkg/handler/grpc/server/rpcHandler"
	"testGinandGorm/pkg/service"
)

const (
	//Address = "0.0.0.0:9090"
	Address  = "127.0.0.1:3030"
	Address2 = "127.0.0.1:3031"
)

func main() {
	Db := db.DbInit()
	orderDao := dao.NewOrderDao(Db)
	orderService := service.NewService(orderDao)
	go test1()
	go test2(orderService)
	select {}
}

func test1() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &rpcHandler.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func test2(orderService *service.OrderService) {
	listen, err := net.Listen("tcp", Address2)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s2 := grpc.NewServer()
	// 服务注册
	pb.RegisterOrderServer(s2, rpcHandler.NewHandler(orderService))
	reflection.Register(s2)
	if err := s2.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
