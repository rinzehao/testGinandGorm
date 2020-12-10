package rpcHandler

import (
	context2 "context"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"testGinandGorm/pkg/handler/grpc/pb"
	"testGinandGorm/pkg/model"
	"testGinandGorm/pkg/service"
)

type Server struct{}

type OrderHandler struct {
	orderService *service.OrderService
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *Server) SayHelloAgain(ctx context2.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "HelloAgain " + request.Name}, nil
}

func NewHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: service}
}

func (handler *OrderHandler) Test(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (handler *OrderHandler) QueryOrderById(c context.Context, in *pb.ID) (reply *pb.OrderModel, err error) {
	id := in.Id
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		return nil, nil
	}
	order, err := handler.orderService.QueryOrderById(id)
	if err != nil {
		status.Error(codes.InvalidArgument, "订单不存在")
		return nil,err
	}
	log.Print(order)
	reply =new(pb.OrderModel)
	reply.Id = int32(order.ID)
	reply.UserName = order.UserName
	reply.OrderNo = order.OrderNo
	reply.Amount = float32(order.Amount)
	reply.Status = order.Status
	reply.FileUrl = order.FileUrl
	return reply, nil
}

func (handler *OrderHandler) CreateOrder(ctx context2.Context, orderInput *pb.OrderModel) (id *pb.ID, err error) {
	var order model.DemoOrder
	id =new(pb.ID)
	order.ID= int(orderInput.Id)
	order.OrderNo =orderInput.OrderNo
	order.Status =orderInput.Status
	order.FileUrl =orderInput.FileUrl
	order.UserName=orderInput.UserName
	if err := handler.orderService.CreateOrder(&order); err != nil {
		status.Error(codes.InvalidArgument, "插入失败")
		id.Id="/"
		return id,err
	}
	id.Id=strconv.Itoa(order.ID)
	return id, err
}

func (handler *OrderHandler) DeleteOrder(ctx context2.Context, in *pb.ID) (*pb.ID, error) {
	id := in.Id
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		return nil, nil
	}
	in.Id="/"
	if err := handler.orderService.DeleteOrderById(id) ; err != nil {
		status.Error(codes.InvalidArgument, "删除失败")
		return in,err
	}
	in.Id=id

	return in, nil
}

func (handler *OrderHandler) UpdateOrder(ctx context2.Context, orderModel *pb.OrderModel) (*pb.OrderModel, error) {
	m := map[string]interface{}{
		"Id":        orderModel.Id,
		"order_No":  orderModel.OrderNo,
		"user_name": orderModel.UserName,
		"amount":    orderModel.Amount,
		"status":    orderModel.Status,
		"file_url":  orderModel.FileUrl,
	}
	if err := handler.orderService.UpdateById(m,strconv.Itoa(int(orderModel.Id))); err != nil {
		status.Error(codes.InvalidArgument, "更新成功")
		return orderModel,err
	}
	return orderModel,nil
}
