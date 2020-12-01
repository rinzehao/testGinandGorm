package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/handler"
	_ "testGinandGorm/pkg/service"
)

func BindRoute(orderHandler *handler.OrderHandler) {
	r := gin.Default()
	r.POST("/order/create", orderHandler.CreateOrder)       //1）创建 demo_order
	r.PUT("/orderUpdate", orderHandler.UpdateOrder)         //2)更新demo_order （amount、stuatus、file_url）
	r.GET("/order/:id", orderHandler.GetOrder)              //3)获取demo_order的详情
	r.GET("/order", orderHandler.GetOrderList)              //4）获取demo_order列表
	r.GET("/orderSearch", orderHandler.GetSortedOrderList)  //5)获取模糊查询结果
	r.GET("/orderDownload", orderHandler.DownLoadExcel)     //6)下载xlsx表格
	r.POST("/orderUpload/:id", orderHandler.GetUploadUrl)   //7)上传文件
	r.DELETE("/order/delete/:id", orderHandler.DeleteOrder) //8)删除demo_order
	r.Run()
}
