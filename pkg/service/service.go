package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
	orderNo :=order.OrderNo
	m :=MapTransform(order)
	if err := service.orderDao.UpdateByNo(orderNo,m); err != nil {
		return err
	}
	return nil
}

func (service *OrderService) CreateOrder(order *model.DemoOrder) error {
	orderNo :=order.OrderNo
	if order,err := service.orderDao.QueryOrderByNo(orderNo); err == gorm.ErrRecordNotFound {
		service.orderDao.CreateOrder(order)
	} else {
		return err
	}
	return nil
}

func (service *OrderService) QueryOrders() (orders []*model.DemoOrder, err error) {
	//页查询 page为页数 pagesize为单页展示条目数量 默认page=1 pagesize=100
	//当page的页数小于等于零的时候  offset不生效
	page, pageSize := 1, 100
	if orders, err = service.orderDao.QueryOrders(page, pageSize); err != nil {
		return nil, err
	}
	return orders, err
}

//根据user_name做模糊查找、根据创建时间、金额排序
func (service *OrderService) QueryOrdersByName(userName string) (orders []*model.DemoOrder, err error) {
	orderBy := "amount"
	desc := "DESC"
	if orders, err = service.orderDao.QueryOrdersByName(userName, orderBy, desc); err != nil {
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
func (service *OrderService) UpdateUrlById(id string, url string) error {
	if len(url) == 0 {
		return nil
	} else {
		m :=map[string]interface{}{
			"file_url" : url,
		}
		if err := service.orderDao.UpdateById(m, id); err != nil {
			return err
		}
	}
	return nil
}

func MapTransform(order *model.DemoOrder) map[string]interface{} {
	m :=map[string]interface{}{
		"Id" :order.ID,
		"order_No":order.OrderNo,
		"user_name" :order.UserName,
		"amount" :order.Amount,
		"status" :order.Status,
		"file_url":order.FileUrl,
	}
	return m
}
