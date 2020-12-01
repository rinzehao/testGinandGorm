package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
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

func (handler *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := handler.orderService.DeleteOrder(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	} else {
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}
}

func (handler *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var order model.DemoOrder
	if err := handler.orderService.GetOrder(id, &order); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("获取条目失败：找不到指定条目")
	} else {
		c.JSON(200, order)
	}
}

func (handler *OrderHandler) UpdateOrder(c *gin.Context) {
	var order model.DemoOrder
	if err := c.BindJSON(&order); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("更新条目失败：JSON绑定错误")
	}
	if err := handler.orderService.UpdateOrder(&order); err != nil {
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

func (handler *OrderHandler) GetOrderList(c *gin.Context) {
	var orderList []model.DemoOrder
	if err, orderList := handler.orderService.GetOrderList(orderList); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, orderList)
	}
}

//根据user_name做模糊查找、根据创建时间、金额排序
func (handler *OrderHandler) GetSortedOrderList(c *gin.Context) {
	var err error
	var order model.DemoOrder
	var orderList []model.DemoOrder
	if err = c.BindJSON(&order); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("模糊查找条目失败:JSO你绑定错误")
	}
	if err, orderList = handler.orderService.GetSortedOrderList(&order, orderList); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("模糊查找条目失败：sql查询出错")
	}
	c.JSON(200, orderList)
	fmt.Println(orderList)
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
func (handler *OrderHandler) GetUploadUrl(c *gin.Context) {
	id := c.Params.ByName("id")
	str := service.TestUpload(c)
	fmt.Println("id:" + id)
	if err := handler.orderService.GetUploadUrl(id, str); err != nil {
		fmt.Println(err)
		panic("上传错误")
	}

}
