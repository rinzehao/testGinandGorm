package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"log"
	"net/http"
	_ "testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
	"testGinandGorm/pkg/service"
	_ "testGinandGorm/pkg/service"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: service}
}

func (handler *OrderHandler) DeleteOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := handler.orderService.DeleteOrderById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	} else {
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}
}

func (handler *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	if order, err := handler.orderService.QueryOrderById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("获取条目失败：找不到指定条目")
	} else {
		c.JSON(200, &order)
	}
}

func (handler *OrderHandler) UpdateOrder(c *gin.Context) {
	var order model.DemoOrder
	if err := c.BindJSON(&order); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("更新条目失败：JSON绑定错误")
	}
	if err := handler.orderService.UpdateByOrderNo(&order); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("更新条目失败")
	} else {
		c.JSON(200, order)
	}
}

func (handler *OrderHandler) CreateOrder(c *gin.Context) {

	var order model.DemoOrder
	if err := c.BindJSON(&order); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("创建条目失败：JSON绑定失败")
	}
	if err := handler.orderService.CreateOrder(&order); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("创建条目失败：service失败")
	}
	c.JSON(200, order)
}

func (handler *OrderHandler) GetOrders(c *gin.Context) {
	//var orders []model.DemoOrder
	if orders, err := handler.orderService.QueryOrders(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, orders)
	}
}

//根据user_name做模糊查找、根据创建时间、金额排序
func (handler *OrderHandler) GetSortedOrders(c *gin.Context) {
	var err error
	var order model.DemoOrder
	var orders []*model.DemoOrder
	if err := c.BindJSON(&order); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("模糊查找条目失败:JSO你绑定错误")
	}
	userName := order.UserName
	if orders, err = handler.orderService.QueryOrdersByName(userName); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("模糊查找条目失败：sql查询出错")
	}
	fmt.Println(&orders)
	c.JSON(200, &orders)
}

//下载DemoOrder,以excel形式导出
func (handler *OrderHandler) DownLoadExcel(c *gin.Context) {
	var err error
	var outPutFileUrl = "order.xlsx"
	file := xlsx.NewFile()
	if err := handler.orderService.DownLoadExcel(file); err != nil {
		fmt.Println(err)
	}
	err = file.Save(outPutFileUrl)
	if err != nil {
		fmt.Println("下载失败：保存表单失败")
		fmt.Println(err.Error())
	}
	fmt.Println("\n\n export success")

}

//获取文件url并保存
func (handler *OrderHandler) GetUploadUrlAndUpdate(c *gin.Context) {
	id := c.Params.ByName("id")
	url := singleFileUpload(c)
	fmt.Println("id:" + id)
	if err := handler.orderService.UpdateUrlById(id, url); err != nil {
		fmt.Println(err)
		panic("上传错误")
	}
}

//单文件上传
func singleFileUpload(c *gin.Context) string {
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
