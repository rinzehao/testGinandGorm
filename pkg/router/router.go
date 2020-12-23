package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/pkg/handler"
	_ "testGinandGorm/pkg/service"
)

//func BindRoute(orderHandler *handler.OrderHandler) {
//	r := gin.Default()
//	r.POST("/order/create", orderHandler.CreateOrder)              //1）创建 demo_order
//	r.PUT("/order/update", orderHandler.UpdateOrder)                //2)更新demo_order （amount、stuatus、file_url）
//	r.GET("/order/query/:id", orderHandler.QueryOrderById)                     //3)获取demo_order的详情
//	r.GET("/order/list", orderHandler.QueryAllOrders)                        //4）获取demo_order列表
//	r.GET("/order/query_username", orderHandler.QueryOrders)            //5)获取模糊查询结果
//	r.GET("/order_download", orderHandler.DownLoadExcel)            //6)下载xlsx表格
//	r.POST("/order_upload/:id", orderHandler.UploadAndUpdate) //7)上传文件
//	r.DELETE("/order/del/:id", orderHandler.DeleteOrderById)    //8)删除demo_order
//	r.Run()
//}

func BindRoute(orderHandler handler.OrderHandler) {
	r := gin.Default()
	r.POST("/order/create", orderHandler.CreateOrder)         //1）创建 demo_order
	r.PUT("/order/update", orderHandler.UpdateOrder)          //2)更新demo_order （amount、stuatus、file_url）
	r.GET("/order/query/:id", orderHandler.QueryOrderById)    //3)获取demo_order的详情
	r.GET("/order/list", orderHandler.QueryAllOrders)         //4）获取demo_order列表
	r.GET("/order/query_username", orderHandler.QueryOrders)  //5)获取模糊查询结果
	r.GET("/order_download", orderHandler.DownLoadExcel)      //6)下载xlsx表格
	r.POST("/order_upload/:id", orderHandler.UploadAndUpdate) //7)上传文件
	r.DELETE("/order/del/:id", orderHandler.DeleteOrderById)  //8)删除demo_order
	r.Run()
}
