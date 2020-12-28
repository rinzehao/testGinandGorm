package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"net/http"
	"strconv"
	"testGinandGorm/common"
	"testGinandGorm/common/logger"
	_ "testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
	model2 "testGinandGorm/pkg/server/model"
	"testGinandGorm/pkg/service"
	_ "testGinandGorm/pkg/service"
)

type Handler interface {
	DeleteOrderById(*gin.Context)
	CreateOrder(*gin.Context)
	UpdateOrder(c *gin.Context)
	QueryOrderById(c *gin.Context)
	QueryAllOrders(c *gin.Context)
	QueryOrders(c *gin.Context)
	DownLoadExcel(c *gin.Context)
	UploadAndUpdate(c *gin.Context)
}

//
//type Handler struct {
//	*builder.BuilderService
//}
//
//func NewHandler() *Handler {
//	return &Handler{BuilderService: builder.NewService()}
//}

const itemOrder = "order"

type OrderHandler struct {
	runtimeProfile *service.ProfileRuntime
}

func NewOrderHandler(runtime *service.ProfileRuntime) *OrderHandler {
	return &OrderHandler{runtimeProfile: runtime}
	//return &MyOrderHandler{orderService: &service}
}

func (handler *OrderHandler) DeleteOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		logger.SugarLogger.Warnf("Got Some Error At InputID, InputID:%s", id)
		return
	}
	ctx := &model2.QueryCtx{
		ItemTyp: itemOrder,
		Req:     id,
	}
	logger.SugarLogger.Debugf("Trying to Query Order By OrderID : OrderID =%s", id)
	//if err := handler.orderService.QueryOrderById(ctx); err != nil {
	if err := handler.runtimeProfile.QueryById(ctx); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "103",
			ErrMsg:  "获取条目失败：找不到指定条目, 或目标条目已删除",
		})
		logger.SugarLogger.Errorf("Fail to Query Order : Error = %s", err)
		return
	}
	ctx = &model2.QueryCtx{
		ItemTyp: itemOrder,
		Req:     ctx.Req,
	}
	logger.SugarLogger.Debugf("Trying to Delete Order By InputID : InputID =%s", id)
	//if err := handler.orderService.DeleteOrderById(ctx); err != nil {
	if err := handler.runtimeProfile.Delete(ctx); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "102",
			ErrMsg:  "order删除错误",
		})
		logger.SugarLogger.Errorf("Fail to Delete Order By InputID : InputID =%s , Error = %s", id, err)
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
	})
	logger.SugarLogger.Infof("Success!Succeed in Deleting Order By InputID : InputID =%s", id)
}

func (handler *OrderHandler) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "JSON绑定错误",
		})
		logger.SugarLogger.Errorf("Fail to Create Order : Error = %s", err)
		return
	}
	ctx := &model2.CreateCtx{
		ItemTyp: itemOrder,
		Req:     order,
	}
	logger.SugarLogger.Debug("Trying to Create Order ")
	//if err := handler.orderService.CreateOrder(ctx); err != nil {
	if err := handler.runtimeProfile.Push(ctx); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "105",
			ErrMsg:  "创建条目失败",
		})
		logger.SugarLogger.Errorf("Fail to Create Order : Error = %s", err)
		return
	}

	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.GetResult(),
	})
	logger.SugarLogger.Infof("Success!Succeed in Creating Order : OrderID =%s", order.ID)
}

func (handler *OrderHandler) UpdateOrder(c *gin.Context) {
	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "JSON绑定错误",
		})
		logger.SugarLogger.Errorf("Fail to Create Order : Error = %s", err)
		return
	}
	ctx := &model2.UpdateCtx{
		ItemTyp:  itemOrder,
		Identify: order.OrderNo,
		Req:      handler.mapTransformer(&order),
	}
	logger.SugarLogger.Debugf("Trying to Update Order By OrderNo : OrderNo =%s", order.OrderNo)
	//if err := handler.orderService.UpdateByOrderNo(ctx); err != nil {
	if err := handler.runtimeProfile.UpdateByNo(ctx); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "104",
			ErrMsg:  "更新条目失败",
		})
		logger.SugarLogger.Errorf("Fail to Update Order : Error = %s", err)
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    order,
	})
	logger.SugarLogger.Infof("Success!Succeed in Updating Order  : OrderNo =%s", order.OrderNo)
}

func (handler *OrderHandler) QueryOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		logger.SugarLogger.Warnf("Got Some Error At InputID, InputID:%s", id)
		return
	}
	ctx := &model2.QueryCtx{
		ItemTyp: itemOrder,
		Req:     id,
	}
	logger.SugarLogger.Debugf("Trying to Query Order By InputID : InputID =%s", id)
	//err := handler.orderService.QueryOrderById(ctx)
	err := handler.runtimeProfile.QueryById(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "103",
			ErrMsg:  "获取条目失败：找不到指定条目",
		})
		logger.SugarLogger.Errorf("Fail to Query Order : Error = %s", err)
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.GetResult(),
	})
	logger.SugarLogger.Infof("Success!Succeed in Querying Order By InputID : InputID =%s", id)
}

func (handler *OrderHandler) QueryAllOrders(c *gin.Context) {
	var page, pageSize = 1, 100
	ctx := &model2.QueryCtxs{
		ItemTyp: itemOrder,
		ReqPage: page,
		ReqSize: pageSize,
	}
	logger.SugarLogger.Debug("Trying to Query All Orders ")
	//err := handler.orderService.QueryOrders(ctx)
	err := handler.runtimeProfile.QueryOrders(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "106",
			ErrMsg:  "获取orderList失败",
		})
		logger.SugarLogger.Errorf("Fail to Query Order : Error = %s", err)
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.GetResult(),
	})
	logger.SugarLogger.Infof("Success!Succeed in Querying All Orders : Counts Of Orders =%s", len(ctx.GetResult()))
}

//根据user_name做模糊查找、根据创建时间、金额排序
func (handler *OrderHandler) QueryOrders(c *gin.Context) {
	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "模糊查找条目失败:JSON绑定错误",
		})
		logger.SugarLogger.Errorf("Fail to Create Order : Error = %s", err)
		return
	}
	ctx := &model2.QueryByNameCtx{
		ItemTyp:     itemOrder,
		Req:         order.UserName,
		OrderOption: "amount",
		DescOrder:   true,
	}

	logger.SugarLogger.Debugf("Trying to Query Order By UserName : UserName =%s", order.UserName)
	//err := handler.orderService.QueryOrdersByName(ctx)
	err := handler.runtimeProfile.QueryByName(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "107",
			ErrMsg:  "模糊查找条目失败:sql查询出错",
		})
		logger.SugarLogger.Errorf("Fail to Query Orders By UserName : UserName = %s , Error = %s", ctx.Req, err)
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.GetResult(),
	})
	logger.SugarLogger.Infof("Success!Succeed in Querying Orders By OrderName : OrderName = %s , Counts Of Orders =%s", order.UserName, len(ctx.GetResult()))
}

//下载Order,以excel形式导出
func (handler *OrderHandler) DownLoadExcel(c *gin.Context) {
	var sheetName = "order_List"
	var outPutFileUrl = "order.xlsx"
	logger.SugarLogger.Debugf("Trying to Export Data As Excel File ")
	if err := handler.excelHandler(sheetName, outPutFileUrl); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "108",
			ErrMsg:  "下载失败：保存表单失败",
		})
		logger.SugarLogger.Errorf("Fail to DownLoad Fail : Error = %s", err)
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
	})
	logger.SugarLogger.Infof("Success!Succeed in DownLoading Data File!")

}

//获取文件url并保存
func (handler *OrderHandler) UploadAndUpdate(c *gin.Context) {
	id := c.Params.ByName("id")
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		logger.SugarLogger.Warnf("Got Some Error At InputID, InputID:%s", id)
		return
	}
	queryCtx := &model2.QueryCtx{
		ItemTyp: itemOrder,
		Req:     id,
	}
	logger.SugarLogger.Debugf("Trying to Query Order By OrderID : OrderID =%s", id)
	//if err := handler.orderService.QueryOrderById(queryCtx); err != nil {
	if err := handler.runtimeProfile.QueryById(queryCtx); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "103",
			ErrMsg:  "获取条目失败：找不到指定条目",
		})
		logger.SugarLogger.Errorf("Fail to Query Order : Error = %s", err)
		return
	}
	m := map[string]interface{}{
		"file_url": handler.singleFileUpload(c),
	}
	updateCtx := &model2.UpdateCtx{
		ItemTyp:  itemOrder,
		Identify: id,
		Req:      m,
	}
	logger.SugarLogger.Debugf("Trying to Update Order By OrderID : OrderID =%s", id)
	//if err := handler.orderService.UpdateById(updateCtx); err != nil {
	if err := handler.runtimeProfile.UpdateById(updateCtx); err != nil {
		logger.SugarLogger.Errorf("Fail to Update Order : Error = %s", err)
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "109",
			ErrMsg:  "上传错误",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    updateCtx.GetResult(),
	})
	logger.SugarLogger.Infof("Success!Succeed in Uploading And Saving  : OrderID = %s", id)
}

//下载Order,以excel形式导出
func (handler *OrderHandler) excelHandler(sheetName, outPutFileUrl string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		logger.SugarLogger.Errorf("Fail to Add Table Sheet : Error = %s", err)
		return err
	}
	var page, pageSize = 1, 100

	ctx := &model2.QueryCtxs{
		ItemTyp: itemOrder,
		ReqPage: page,
		ReqSize: pageSize,
	}
	logger.SugarLogger.Debug("Trying to Query All Orders ")
	//err = handler.orderService.QueryOrders(ctx)
	err = handler.runtimeProfile.QueryOrders(ctx)
	if err != nil {
		logger.SugarLogger.Errorf("Fail to Query Orders : Error = %s", err)
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
	array := ctx.GetResult()
	var orders []*model.Order
	for _, v := range array {
		orders = append(orders, v.(*model.Order))
	}
	for _, order := range orders {
		row := sheet.AddRow()
		orderId := row.AddCell()
		orderId.Value = strconv.Itoa(order.ID)
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
		logger.SugarLogger.Errorf("Fail to Save Form  : Error = %s", err)
		return err
	}
	return nil
}

func (handler *OrderHandler) mapTransformer(order *model.Order) map[string]interface{} {
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
func (handler *OrderHandler) singleFileUpload(c *gin.Context) string {
	file, err := c.FormFile("file")
	if err != nil {
		logger.SugarLogger.Debugf("Fail to Upload File : Error = %s", err)
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "110",
			ErrMsg:  "单文件上传错误",
		})
		return ""
	}
	dst := fmt.Sprintf("./file/" + file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		logger.SugarLogger.Errorf("Fail to Save File : Error = %s", err)
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
