package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "testGinandGorm/pkg/service"
)

func BindRoute(handler *Handler) error {
	r := gin.Default()
	r.POST("/order/create", handler.CreateOrder)         //1）创建 order
	r.PUT("/order/update", handler.UpdateOrder)          //2)更新order （amount、stuatus、file_url）
	r.GET("/order/query/:id", handler.QueryOrderById)    //3)获取order的详情
	r.GET("/order/list", handler.QueryAllOrders)         //4）获取order列表
	r.GET("/order/query_username", handler.QueryOrders)  //5)获取模糊查询结果
	r.GET("/order_download", handler.DownLoadExcel)      //6)下载xlsx表格
	r.POST("/order_upload/:id", handler.UploadAndUpdate) //7)上传文件
	r.DELETE("/order/del/:id", handler.DeleteOrderById)  //8)删除order
	if err := r.Run(); err != nil {
		return err
	}
	return nil
}
