package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"testGinandGorm/pkg/server/grpc/pb"
)

const (
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

	res4, err4 := c2.QueryOrderById(context.Background(), &pb.ID{Id: "239"})
	if err4 != nil {
		log.Fatalln(err4)
	}
	log.Println(res4)

	//todo  test Create
	res6, err6 := c2.CreateOrder(context.Background(), &pb.OrderModel{
		Id:       2,
		OrderNo:  "2",
		UserName: "2",
		Amount:   22.22,
		Status:   "完成",
		FileUrl:  "../././nsbsk",
	})
	if err6 != nil {
		log.Fatalln(err6)
	}
	log.Println(res6)

	//todo test Update
	res8, err8 := c2.UpdateOrder(context.Background(), &pb.OrderModel{
		Id:       2,
		OrderNo:  "2",
		UserName: "2",
		Amount:   15.00,
		Status:   "未完成",
		FileUrl:  "南山必胜客",
	})
	if err8 != nil {
		log.Fatalln(err6)
	}
	log.Println(res8)

	//todo test Delete
	res7, err7 := c2.DeleteOrder(context.Background(), &pb.ID{Id: "226"})
	if err7 != nil {
		log.Fatalln(err6)
	}
	log.Println(res7)
}
