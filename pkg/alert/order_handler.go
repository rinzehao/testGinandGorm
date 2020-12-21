package alert

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testGinandGorm/common"
	_ "testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
	_ "testGinandGorm/pkg/service"
)

type MyOrderHandler struct {
	orderService *MyOrderService
}

func NewOrderHandler(service *MyOrderService) *MyOrderHandler {
	return &MyOrderHandler{orderService: service}
}

func (handler *MyOrderHandler)CreateOrder(c *gin.Context) {
	var ctx model.OrderMould
	if err := c.BindJSON(&ctx); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "JSON绑定错误",
		})
		return
	}
	if err := handler.orderService.Create(&ctx); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "105",
			ErrMsg:  "创建条目失败",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx,
	})
}

