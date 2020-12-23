package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"testGinandGorm/common"
	"testGinandGorm/common/logger"
	_ "testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
	"testGinandGorm/pkg/service"
	_ "testGinandGorm/pkg/service"
)

type OrderHandler interface {
	DeleteOrderById(*gin.Context)
	CreateOrder(*gin.Context)
	UpdateOrder(c *gin.Context)
	QueryOrderById(c *gin.Context)
	QueryAllOrders(c *gin.Context)
	QueryOrders(c *gin.Context)
	DownLoadExcel(c *gin.Context)
	UploadAndUpdate(c *gin.Context)
}

type MyOrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(service service.OrderService) *MyOrderHandler {
	return &MyOrderHandler{orderService: service}
}

func (handler *MyOrderHandler) DeleteOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		logger.Warn("#Input ID Has error")
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		return
	}
	ctx := &model.OrderMade{OrderID: id}
	if err := handler.orderService.DeleteOrderById(ctx); err != nil {
		logger.Error("#Delete Order fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "102",
			ErrMsg:  "order删除错误",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
	})
}

func (handler *MyOrderHandler) CreateOrder(c *gin.Context) {
	var order model.DemoOrder
	if err := c.BindJSON(&order); err != nil {
		logger.Error("#JSON Binding fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "JSON绑定错误",
		})
		return
	}
	ctx := &model.OrderMade{
		OrderNo: order.OrderNo,
		Order: &model.DemoOrder{
			ID:       order.ID,
			OrderNo:  order.OrderNo,
			UserName: order.UserName,
			Amount:   order.Amount,
			Status:   order.Status,
			FileUrl:  order.FileUrl,
		},
	}
	if err := handler.orderService.CreateOrder(ctx); err != nil {
		logger.Error("#Order Create fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "105",
			ErrMsg:  "创建条目失败",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.Order,
	})
}

func (handler *MyOrderHandler) UpdateOrder(c *gin.Context) {
	var order model.DemoOrder
	if err := c.BindJSON(&order); err != nil {
		logger.Error("#JSON Binding fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "JSON绑定错误",
		})
		return
	}
	ctx := &model.OrderMade{
		OrderNo:   order.OrderNo,
		UpdateMap: handler.mapTransformer(&order),
	}
	if err := handler.orderService.UpdateByOrderNo(ctx); err != nil {
		logger.Error("#Order Update fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "104",
			ErrMsg:  "更新条目失败",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    order,
	})
}

func (handler *MyOrderHandler) QueryOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		logger.Warn("#Input ID Has error")
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		return
	}
	ctx := &model.OrderMade{
		OrderID: id,
	}
	err := handler.orderService.QueryOrderById(ctx)
	if err != nil {
		logger.Error("#Query Order fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "103",
			ErrMsg:  "获取条目失败：找不到指定条目",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.Order,
	})
}

func (handler *MyOrderHandler) QueryAllOrders(c *gin.Context) {
	var page, pageSize = 1, 100
	ctx := &model.OrderMade{
		Page:     page,
		PageSize: pageSize,
	}
	err := handler.orderService.QueryOrders(ctx)
	if err != nil {
		logger.Error("#Query Orders fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "106",
			ErrMsg:  "获取orderList失败",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.Group,
	})
}

//根据user_name做模糊查找、根据创建时间、金额排序
func (handler *MyOrderHandler) QueryOrders(c *gin.Context) {
	var order model.DemoOrder
	if err := c.BindJSON(&order); err != nil {
		logger.Error("#JSON Binding fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "模糊查找条目失败:JSON绑定错误",
		})
		return
	}
	ctx := &model.OrderMade{
		UserName: order.UserName,
		OrderBy:  "amount",
		Desc:     "desc",
	}
	err := handler.orderService.QueryOrdersByName(ctx)
	if err != nil {
		logger.Error("#Query Orders By UserName fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "107",
			ErrMsg:  "模糊查找条目失败:sql查询出错",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.Group,
	})
}

//下载DemoOrder,以excel形式导出
func (handler *MyOrderHandler) DownLoadExcel(c *gin.Context) {
	var sheetName = "order_List"
	var outPutFileUrl = "order.xlsx"
	if err := handler.excelHandler(sheetName, outPutFileUrl); err != nil {
		logger.Error("#DownLoad fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "108",
			ErrMsg:  "下载失败：保存表单失败",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
	})
}

//获取文件url并保存
func (handler *MyOrderHandler) UploadAndUpdate(c *gin.Context) {
	id := c.Params.ByName("id")
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		return
	}
	ctx := &model.OrderMade{
		OrderID: id,
	}
	if err := handler.orderService.QueryOrderById(ctx); err != nil {
		logger.Error("#Query Order By OrderID fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "103",
			ErrMsg:  "获取条目失败：找不到指定条目",
		})
		return
	}
	m := map[string]interface{}{
		"file_url": handler.singleFileUpload(c),
	}
	ctx = &model.OrderMade{
		OrderID:   id,
		UpdateMap: m,
	}
	if err := handler.orderService.UpdateById(ctx); err != nil {
		logger.Error("#Update Order fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "109",
			ErrMsg:  "上传错误",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.UpdateMap,
	})
}

//下载DemoOrder,以excel形式导出
func (handler *MyOrderHandler) excelHandler(sheetName, outPutFileUrl string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		return err
	}
	var page, pageSize = 1, 100
	ctx := &model.OrderMade{
		Page:     page,
		PageSize: pageSize,
	}
	err = handler.orderService.QueryOrders(ctx)
	if err != nil {
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
	for _, order := range ctx.Group {
		row := sheet.AddRow()
		orderId := row.AddCell()
		orderId.Value = strconv.Itoa(order.ID)
		fmt.Println(orderId.Value)
		orderNo := row.AddCell()
		orderNo.Value = order.OrderNo
		orderName := row.AddCell()
		orderName.Value = order.UserName
		orderAmount := row.AddCell()
		orderAmount.Value = strconv.FormatFloat(order.Amount, 'f', 3, 64)
		orderStatus := row.AddCell()
		orderStatus.Value = order.Status
		orderFile := row.AddCell()
		orderFile.Value = order.FileUrl
	}
	err = file.Save(outPutFileUrl)
	if err != nil {
		logger.Error("#DownLoad Fail:failed to save form fail", zap.Error(err))
		return err
	}
	return nil
}

func (handler *MyOrderHandler) mapTransformer(order *model.DemoOrder) map[string]interface{} {
	m := map[string]interface{}{
		"Id":        order.ID,
		"order_No":  order.OrderNo,
		"user_name": order.UserName,
		"amount":    order.Amount,
		"status":    order.Status,
		"file_url":  order.FileUrl,
	}
	return m
}

//单文件上传
func (handler *MyOrderHandler) singleFileUpload(c *gin.Context) string {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Error("#Upload File fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "110",
			ErrMsg:  "单文件上传错误",
		})
		return ""
	}
	dst := fmt.Sprintf("./file/" + file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		logger.Error("#Save File fail", zap.Error(err))
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "111",
			ErrMsg:  "单文件保存失败",
		})
		return ""
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    dst,
	})
	return dst
}
