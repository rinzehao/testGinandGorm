package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tealeg/xlsx"
	"log"
	"net/http"
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

func (service *OrderService) DeleteOrder(id string) error {
	var err error
	if err = service.orderDao.Delete(id); err != nil {
		return err
	}
	return err
}

func (service *OrderService) GetOrder(id string, order *model.DemoOrder) error {
	var err error
	if err = service.orderDao.GetOrder(id, order); err != nil {
		return err
	}
	return err
}

func (service *OrderService) UpdateOrder(order *model.DemoOrder) error {
	var err error
	var orderNo = order.OrderNo
	if err = service.orderDao.FindOrder(orderNo); err != nil {
		return err
	} else if err = service.orderDao.Update(order); err != nil {
		return err
	}
	return err
}

func (service *OrderService) CreateOrder(order *model.DemoOrder) error {
	var err error
	orderNo := order.OrderNo
	if err = service.orderDao.FindOrder(orderNo); err == nil {
		service.orderDao.Create(order)
	}
	return err
}

func (service *OrderService) GetOrderList(list []model.DemoOrder) (error, []model.DemoOrder) {
	var err error
	if err = service.orderDao.GetOrderList(&list); err != nil {
		return err, nil
	}
	return err, list
}

//根据user_name做模糊查找、根据创建时间、金额排序
func (service *OrderService) GetSortedOrderList(order *model.DemoOrder,
	orderList []model.DemoOrder) (error, []model.DemoOrder) {

	var err error
	likeName := order.UserName
	fmt.Scan(likeName)
	if err = service.orderDao.GetSortedOrderList(likeName, &orderList); err != nil {
		return err, nil
	}
	fmt.Println(orderList)
	return err, orderList
}

//下载DemoOrder,以excel形式导出
func (service *OrderService) DownLoadExcel(file *xlsx.File) error {
	sheet, err := file.AddSheet("order_list")
	var list []model.DemoOrder
	if err != nil {
		return err
	}
	if err = service.orderDao.GetOrderList(&list); err != nil {
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
	for _, order := range list {
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
func (service *OrderService) GetUploadUrl(id string, str string) error {
	fmt.Println("id:" + id)
	var err error
	//开启事务
	if str == " " {
		return err
	} else if err := service.orderDao.GetTransactionBegin(str, id); err != nil {
		return err
	}

	return err
}

//单文件上传
func TestUpload(c *gin.Context) string {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("ERROR: upload file failed. ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("ERROR: upload file failed. %s", err),
		})
	}
	log.Println(file.Filename)
	dst := fmt.Sprintf("./common/" + file.Filename)
	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Println("ERROR: save file failed. ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("ERROR: save file failed. %s", err),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":      "file upload succ.",
		"filepath": dst,
	})
	return dst
}
