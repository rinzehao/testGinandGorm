package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testGinandGorm/pkg/server/grpc/pb"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	host = "127.0.0.1:3031"
)

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	crudClient := pb.NewOrderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	for {
		fmt.Println("OrderDemo CRUD")
		fmt.Print(" 1 : create \n 2 : read \n 3 : update \n 4: delete \n 5: 退出 \n Enter choice :  ")

		choice := bufio.NewReader(os.Stdin)
		text, _ := choice.ReadString('\n')

		switch text {
		case "1\n":
			fmt.Print("\nOrder_ID: ")
			idR := bufio.NewReader(os.Stdin)
			orderID, _ := idR.ReadString('\n')
			orderID = strings.Trim(orderID, "\n")
			orderId, _ := strconv.Atoi(orderID)

			fmt.Print("\nOrder_NO: ")
			noR := bufio.NewReader(os.Stdin)
			no, _ := noR.ReadString('\n')
			no = strings.Trim(no, "\n")

			fmt.Print("\nuser_name: ")
			nameR := bufio.NewReader(os.Stdin)
			name, _ := nameR.ReadString('\n')
			name = strings.Trim(name, "\n")

			fmt.Print("\namount: ")
			amountR := bufio.NewReader(os.Stdin)
			amount, _ := amountR.ReadString('\n')
			amount = strings.Trim(amount, "\n")
			orderAmount, _ := strconv.ParseFloat(amount, 32)

			fmt.Print("\nstatus: ")
			statusR := bufio.NewReader(os.Stdin)
			status, _ := statusR.ReadString('\n')
			status = strings.Trim(status, "\n")

			fmt.Print("\nfile_Url: ")
			urlR := bufio.NewReader(os.Stdin)
			fileUrl, _ := urlR.ReadString('\n')
			fileUrl = strings.Trim(fileUrl, "\n")

			order, err := crudClient.CreateOrder(ctx, &pb.OrderModel{
				Id:       int32(orderId),
				OrderNo:  no,
				UserName: name,
				Amount:   float32(orderAmount),
				Status:   status,
				FileUrl:  fileUrl,
			})
			if err != nil {
				log.Fatalf("创建订单失败: %v", err)
			}
			fmt.Println("\n插入成功，ID ： ", order.Id)

		case "2\n":
			fmt.Print("\n请输入ID  : ")
			idR := bufio.NewReader(os.Stdin)
			orderID, _ := idR.ReadString('\n')
			orderID = strings.Trim(orderID, "\n")

			order, err := crudClient.QueryOrderById(ctx, &pb.ID{Id: orderID})
			if err != nil {
				log.Fatalf("Error in getting product: %v", err)
			}
			fmt.Println("\n 查询成功!")
			fmt.Println(
				"order_No:", order.OrderNo, "\n",
				" user_Name: ", order.UserName, "\n",
				"order_amount : ", order.Amount, "\n",
				"order_status :", order.Status, "\n",
				"order_fileUrl :", order.FileUrl)

		case "3\n":

			fmt.Print("\n请输入ID  : ")
			qidR := bufio.NewReader(os.Stdin)
			queryID, _ := qidR.ReadString('\n')
			queryID = strings.Trim(queryID, "\n")

			order, err := crudClient.QueryOrderById(ctx, &pb.ID{Id: queryID})
			if err != nil {
				log.Fatalf("Error in getting product: %v", err)
			}
			fmt.Println("\n 存在订单")
			fmt.Println(
				"order_No:", order.OrderNo, "\n",
				" user_Name: ", order.UserName, "\n",
				"order_amount : ", order.Amount, "\n",
				"order_status :", order.Status, "\n",
				"order_fileUrl :", order.FileUrl)

			fmt.Println("\n ————修改订单————")

			fmt.Print("\nOrder_ID: ")
			idR := bufio.NewReader(os.Stdin)
			orderID, _ := idR.ReadString('\n')
			orderID = strings.Trim(orderID, "\n")
			orderId, _ := strconv.Atoi(orderID)

			fmt.Print("\nOrder_NO: ")
			noR := bufio.NewReader(os.Stdin)
			no, _ := noR.ReadString('\n')
			no = strings.Trim(no, "\n")

			fmt.Print("\nuser_name: ")
			nameR := bufio.NewReader(os.Stdin)
			name, _ := nameR.ReadString('\n')
			name = strings.Trim(name, "\n")

			fmt.Print("\namount: ")
			amountR := bufio.NewReader(os.Stdin)
			amount, _ := amountR.ReadString('\n')
			amount = strings.Trim(amount, "\n")
			orderAmount, _ := strconv.ParseFloat(amount, 32)

			fmt.Print("\nstatus: ")
			statusR := bufio.NewReader(os.Stdin)
			status, _ := statusR.ReadString('\n')
			status = strings.Trim(status, "\n")

			fmt.Print("\nfile_Url: ")
			urlR := bufio.NewReader(os.Stdin)
			fileUrl, _ := urlR.ReadString('\n')
			fileUrl = strings.Trim(fileUrl, "\n")

			updateOrder, err := crudClient.UpdateOrder(ctx, &pb.OrderModel{
				Id:       int32(orderId),
				OrderNo:  no,
				UserName: name,
				Amount:   float32(orderAmount),
				Status:   status,
				FileUrl:  fileUrl,
			})
			if err != nil {
				log.Fatalf("更新订单失败: %v", err)
			}
			fmt.Println("\n订单：", updateOrder.Id, " 更新成功")

		case "4\n":

			fmt.Print("\n————删除订单，请输入ID————")
			fmt.Print("\nID: ")
			idR := bufio.NewReader(os.Stdin)
			id, _ := idR.ReadString('\n')
			id = strings.Trim(id, "\n")
			product, _ := crudClient.DeleteOrder(ctx, &pb.ID{Id: id})
			if product != nil {
				log.Printf("\n订单 ID ： %s 删除失败 ", product.Id)
			} else {
				log.Printf("删除成功！")
			}
		case "5\n":
			goto End
		default:
			fmt.Println("\n操 作 异 常!")
		}
	}
 End:
}
