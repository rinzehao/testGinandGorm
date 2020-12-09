package rpcHandler

import (
	context2 "context"
	"golang.org/x/net/context"
	"log"
	"strconv"
	"testGinandGorm/pkg/handler/grpc/pb"
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

func (handler *OrderHandler) QueryOrderById(c context.Context, in *pb.QueryRequest) (reply *pb.QueryReply, err error) {
	id := in.Id
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		return nil, nil
	}
	order, err := handler.orderService.QueryOrderById(id)
	if err != nil {
		return nil, err
	}
	log.Print(order)
	reply.Id = int32(order.ID)
	reply.UserName = order.UserName
	reply.OrderNo = order.OrderNo
	reply.Amount = float32(order.Amount)
	reply.Status = order.Status
	reply.FileUrl = order.FileUrl
	return reply, nil
}
