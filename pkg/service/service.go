package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tealeg/xlsx"
	"strconv"
	"testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
)

type OrderService struct {
	orderDao *dao.OrderDao
}

func NewService(dao *dao.OrderDao) *OrderService {
	return &OrderService{orderDao: dao}
}

func (service *OrderService) DeleteOrderById(id string) error {
	var err error
	if err = service.orderDao.DeleteById(id); err != nil {
		return err
	}
	return err
}

func (service *OrderService) QueryOrderById(id string) (order *model.DemoOrder, err error) {
	if order, err = service.orderDao.QueryOrderById(id); err != nil {
		return nil, err
	}
	return order, err
}

func (service *OrderService) UpdateByOrderNo(order *model.DemoOrder) error {
	var err error
	paramName :="order_No"
	m :=StructOrderToMap(order)
	if err = service.orderDao.UpdateByParam(m,paramName,order); err != nil {
		return err
	}
	return err
}

func (service *OrderService) CreateOrderByOrderNo(order *model.DemoOrder) error {
	m:=StructOrderToMap(order)
	queryParam :="order_No"
	if isExit, err := service.orderDao.QueryOrderIsExist(m ,queryParam,order); isExit == false {
		service.orderDao.CreateOrder(order)
	} else {
		return err
	}
	return nil
}

func (service *OrderService) QueryOrders() (orders []*model.DemoOrder, err error) {
	//页查询 page为页数 pagesize为单页展示条目数量 默认page=1 pagesize=100
	page, pageSize := 1, 100
	if orders, err = service.orderDao.QueryOrders(page, pageSize); err != nil {
		return nil, err
	}
	return orders, err
}

//根据user_name做模糊查找、根据创建时间、金额排序
func (service *OrderService) QuerySortedOrdersByUserName(userName string) (orders []*model.DemoOrder, err error) {

	orderBy := "amount"
	desc := "DESC"
	if orders, err = service.orderDao.QuerySortedOrdersByUserName(userName, orderBy, desc); err != nil {
		return nil, err
	}
	return orders, err
}

//下载DemoOrder,以excel形式导出
func (service *OrderService) DownLoadExcel(file *xlsx.File) error {
	sheet, err := file.AddSheet("order_list")
	var orders []*model.DemoOrder
	if err != nil {
		return err
	}
	if orders, err = service.QueryOrders(); err != nil {
		return err
	}
	//定义表头
	headeRow := sheet.AddRow()
	idCell := headeRow.AddCell()
	idCell.Value = "ID"
	noCell := headeRow.AddCell()
	noCell.Value = "Order_No"
	nameCell := headeRow.AddCell()
	nameCell.Value = "user_name"
	amountCell := headeRow.AddCell()
	amountCell.Value = "Amount"
	statusCell := headeRow.AddCell()
	statusCell.Value = "Status"
	fileCell := headeRow.AddCell()
	fileCell.Value = "File_Url"
	//写入表单
	for _, order := range orders {
		row := sheet.AddRow()
		orderId := row.AddCell()
		orderId.Value = strconv.Itoa(order.ID)
		fmt.Println(orderId.Value)
		orderNo := row.AddCell()
		orderNo.Value = order.OrderNo
		orderName := row.AddCell()
		orderName.Value = order.UserName
		orderAmount := row.AddCell()
		orderAmount.Value = strconv.FormatFloat(float64(order.Amount), 'f', 3, 64)
		orderStatus := row.AddCell()
		orderStatus.Value = order.Status
		orderFile := row.AddCell()
		orderFile.Value = order.FileUrl
	}
	return err
}

//获取文件url并保存
func (service *OrderService) GetUploadUrlAndSave(id string, url string) error {
	fmt.Println("id:" + id)
	var err error
	if len(url) == 0 {
		return err
	} else {
		m :=map[string]string{
			"file_url" : url,
		}
		if err = service.orderDao.TransactionUpdateById(m, id); err != nil {
			return err
		}
	}
	return err
}

func StructOrderToMap(order *model.DemoOrder) map[string]string {
	m :=map[string]string{
		"Id" :strconv.Itoa(order.ID),
		"order_No":order.OrderNo,
		"user_name" :order.UserName,
		"amount" :strconv.FormatFloat(order.Amount, 'E', -1, 64),
		"status" :order.Status,
		"file_url":order.FileUrl,
	}
	return m
}
