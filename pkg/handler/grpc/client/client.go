package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"testGinandGorm/pkg/handler/grpc/pb"
)

const (
	//Address = "0.0.0.0:9090"
	Address  = "127.0.0.1:3030"
	Address2 = "127.0.0.1:3031"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	//初始化客户端
	c := pb.NewGreeterClient(conn)
	// 调用 SayHello 方法
	res, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "Hello World"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Message)

	res2, err2 := c.SayHello(context.Background(), &pb.HelloRequest{Name: "你好 世界"})
	if err2 != nil {
		log.Fatalln(err2)
	}
	log.Println(res2.Message)

	res5, err5 := c.SayHelloAgain(context.Background(), &pb.HelloRequest{Name: "你好 世界ooo "})
	if err5 != nil {
		log.Fatalln(err5)
	}
	log.Println(res5.Message)

	//调用OrderService的方法

	conn2, err2 := grpc.Dial(Address2, grpc.WithInsecure())
	if err2 != nil {
		log.Fatalln(err)
	}
	defer conn2.Close()
	c2 := pb.NewOrderClient(conn2)

	res3, err3 := c2.Test(context.Background(), &pb.HelloRequest{Name: "239"})
	if err3 != nil {
		log.Fatalln(err3)
	}
	log.Println(res3)

	res4, err4 := c2.QueryOrderById(context.Background(), &pb.QueryRequest{Id: "239"})
	if err4 != nil {
		log.Fatalln(err4)
	}
	log.Println(res4)
}
