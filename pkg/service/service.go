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
		fmt.Println(404)
		return err
	}
	return err
}

func (service *OrderService) GetOrder(id string, order *model.Demo_order) error {
	var err error
	if err = service.orderDao.GetOrder(id, order); err != nil {
		fmt.Println(err)
		fmt.Println("获取条目失败：找不到指定条目")
	}
	return err
}

func (service *OrderService) UpdateOrder(id string, order *model.Demo_order) error {
	var err error
	if err = service.orderDao.FindOrder(id); err != nil {
		fmt.Println(err)
		fmt.Println("更新条目失败：查询不到相应条目")
	} else if err = service.orderDao.Update(order); err != nil {
		fmt.Println(err)
		fmt.Println("更新条目失败：保存错误")
	}
	return err
}

func (service *OrderService) CreateOrder(order *model.Demo_order) error {
	var err error
	var id = strconv.Itoa(order.ID)
	if err = service.orderDao.GetOrder(id, order); err != nil {
		service.orderDao.Create(order)
	}
	return err
}

func (service *OrderService) GetOrderList(list []model.Demo_order) (error, []model.Demo_order) {
	var err error
	if err = service.orderDao.GetOrderList(&list); err != nil {
		fmt.Println(err)
	}
	return err, list
}

//根据user_name做模糊查找、根据创建时间、金额排序
func (service *OrderService) GetSortedOrderList(order *model.Demo_order,
	orderList []model.Demo_order) (error, []model.Demo_order) {

	var err error
	likeName := order.UserName
	fmt.Scan(likeName)
	if err = service.orderDao.GetSortedOrderList(likeName, &orderList); err != nil {
		fmt.Println(err)
		fmt.Println("模糊查找条目失败：sql查询出错")
	}
	fmt.Println(orderList)
	return err, orderList
}

//下载demo_order,以excel形式导出
func (service *OrderService) DownLoadExcel(file *xlsx.File) error {
	sheet, err := file.AddSheet("order_list")
	var list []model.Demo_order
	if err != nil {
		fmt.Println("下载失败：添加表单失败")
		fmt.Println(err.Error())
	}
	if err = service.orderDao.GetOrderList(&list); err != nil {
		fmt.Println(err)
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
	//var order model.Demo_order
	fmt.Println("id:" + id)
	//开启事务
	err, tx := service.orderDao.GetSessionBegin()
	if err != nil {
		fmt.Println(err)
	}
	if str == " " {
		fmt.Println("更新url失败：获取新url为空")
		tx.Rollback()
	}
	if err := service.orderDao.UpdateUrl(str, id, tx); err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
	fmt.Println(str)
	if err := tx.Commit().Error; err != nil {
		fmt.Println(err)
		fmt.Println("更新url失败：事务提交出错")
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
