package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testGinandGorm/pkg/service"
	_ "testGinandGorm/pkg/service"
)

func BindRoute() {
	r := gin.Default()
	r.POST("/order/create", service.CreateOrder)       //1）创建 demo_order
	r.PUT("/order/update/:id", service.UpdateOrder)    //2)更新demo_order （amount、stuatus、file_url）
	r.GET("/order/:id", service.GetOrder)              //3)获取demo_order的详情
	r.GET("/order", service.GetOrderList)              //4）获取demo_order列表
	r.GET("/orderSearch", service.GetSortedOrderList)  //5)获取模糊查询结果
	r.GET("/orderDownload", service.DownLoadExcel)     //6)下载xlsx表格
	r.POST("/orderUpload/:id", service.GetUploadUrl)   //7)上传文件
	r.DELETE("/order/delete/:id", service.DeleteOrder) //8)删除demo_order
	r.Run()
}
