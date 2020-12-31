package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"testGinandGorm/common"
	"testGinandGorm/common/log"
	_ "testGinandGorm/pkg/dao"
	"testGinandGorm/pkg/model"
	model2 "testGinandGorm/pkg/server/model"
	"testGinandGorm/pkg/service"
	_ "testGinandGorm/pkg/service"
)

type HandlerInterface interface {
	DeleteOrderById(*gin.Context)
	CreateOrder(*gin.Context)
	UpdateOrder(c *gin.Context)
	QueryOrderById(c *gin.Context)
	QueryAllOrders(c *gin.Context)
	QueryOrders(c *gin.Context)
	DownLoadExcel(c *gin.Context)
	UploadAndUpdate(c *gin.Context)
}

const itemOrder = "order"

type Handler struct {
	profileManager *service.ProfileManager
}

func NewHandler(manager *service.ProfileManager) *Handler {
	return &Handler{profileManager: manager}
}

func (handler *Handler) DeleteOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		log.Logger.Warn("Got Some Error At InputID", zap.String("InputID", strconv.Itoa(id)))
		return
	}
	ctx := &model2.QueryCtx{
		ItemTyp: itemOrder,
		Req:     id,
	}
	if err := handler.profileManager.QueryProfileById(ctx); err != nil {
		c.JSON(http.StatusGone, &common.HttpResp{
			ErrCode: "102",
			ErrMsg:  "获取条目失败：找不到指定条目, 或目标条目已删除",
		})
		log.Logger.Error("Fail to Query Order ", zap.Error(err))
		return
	}
	ctx = &model2.QueryCtx{
		ItemTyp: itemOrder,
		Req:     ctx.Req,
	}
	if err := handler.profileManager.DeleteProfile(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, &common.HttpResp{
			ErrCode: "103",
			ErrMsg:  "order删除错误",
		})
		log.Logger.Error("Fail to Delete Order By InputID ", zap.String("InputID", id), zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
	})
}

func (handler *Handler) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "JSON绑定错误",
		})
		log.Logger.Error("Fail to Create Order ", zap.Error(err))
		return
	}
	ctx := &model2.CreateCtx{
		ItemTyp: itemOrder,
		Req:     order,
	}
	if err := handler.profileManager.PushProfile(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, &common.HttpResp{
			ErrCode: "104",
			ErrMsg:  "创建条目失败",
		})
		log.Logger.Error("Fail to Create Order ", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.GetResult(),
	})
}

func (handler *Handler) UpdateOrder(c *gin.Context) {
	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "JSON绑定错误",
		})
		log.Logger.Error("Fail to Create Order ", zap.Error(err))
		return
	}
	ctx := &model2.UpdateCtx{
		ItemTyp:  itemOrder,
		Identify: order.OrderNo,
		Req:      handler.mapTransformer(&order),
	}
	if err := handler.profileManager.UpdateProfileByNo(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, &common.HttpResp{
			ErrCode: "105",
			ErrMsg:  "更新条目失败",
		})
		log.Logger.Error("Fail to Update Order ", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    order,
	})
}

func (handler *Handler) QueryOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		log.Logger.Warn("Got Some Error At InputID", zap.String("InputID", strconv.Itoa(id)))
		return
	}
	ctx := &model2.QueryCtx{
		ItemTyp: itemOrder,
		Req:     id,
	}
	err := handler.profileManager.QueryProfileById(ctx)
	if err != nil {
		c.JSON(http.StatusGone, &common.HttpResp{
			ErrCode: "102",
			ErrMsg:  "获取条目失败：找不到指定条目",
		})
		log.Logger.Error("Fail to Query Order ", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.GetResult(),
	})
}

func (handler *Handler) QueryAllOrders(c *gin.Context) {
	var page, pageSize = 1, 100
	ctx := &model2.QueryCtxs{
		ItemTyp: itemOrder,
		ReqPage: page,
		ReqSize: pageSize,
	}
	err := handler.profileManager.QueryProfiles(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &common.HttpResp{
			ErrCode: "106",
			ErrMsg:  "获取orderList失败",
		})
		log.Logger.Error("Fail to Query Order ", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.GetResult(),
	})
}

//根据user_name做模糊查找、根据创建时间、金额排序
func (handler *Handler) QueryOrders(c *gin.Context) {
	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "模糊查找条目失败:JSON绑定错误",
		})
		log.Logger.Error("Fail to Create Order ", zap.Error(err))
		return
	}
	ctx := &model2.QueryByNameCtx{
		ItemTyp:     itemOrder,
		Req:         order.UserName,
		OrderOption: "amount",
		DescOrder:   true,
	}
	err := handler.profileManager.QueryProfileByName(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &common.HttpResp{
			ErrCode: "107",
			ErrMsg:  "模糊查找条目失败:sql查询出错",
		})
		log.Logger.Error("Fail to Query Orders By UserName ", zap.String("userName", ctx.Req.(string)), zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    ctx.GetResult(),
	})
}

//下载Order,以excel形式导出
func (handler *Handler) DownLoadExcel(c *gin.Context) {
	var sheetName = "order_List"
	var outPutFileUrl = "order.xlsx"
	if err := handler.excelHandler(sheetName, outPutFileUrl); err != nil {
		c.JSON(http.StatusInternalServerError, &common.HttpResp{
			ErrCode: "108",
			ErrMsg:  "下载失败：保存表单失败",
		})
		log.Logger.Error("Fail to DownLoad Fail ", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
	})
}

//获取文件url并保存
func (handler *Handler) UploadAndUpdate(c *gin.Context) {
	id := c.Params.ByName("id")
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		log.Logger.Warn("Got Some Error At InputID", zap.String("InputID", strconv.Itoa(id)))
		return
	}
	queryCtx := &model2.QueryCtx{
		ItemTyp: itemOrder,
		Req:     id,
	}
	if err := handler.profileManager.QueryProfileById(queryCtx); err != nil {
		c.JSON(http.StatusGone, &common.HttpResp{
			ErrCode: "103",
			ErrMsg:  "获取条目失败：找不到指定条目",
		})
		log.Logger.Error("Fail to Query Order ", zap.Error(err))
		return
	}
	url, err := handler.singleFileUpload(c)
	if err != nil {
		log.Logger.Error("Fail to Get UploadUrl ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, &common.HttpResp{
			ErrCode: "109",
			ErrMsg:  "上传错误",
		})
		return
	}
	m := map[string]interface{}{
		"file_url": url,
	}
	updateCtx := &model2.UpdateCtx{
		ItemTyp:  itemOrder,
		Identify: id,
		Req:      m,
	}
	if err := handler.profileManager.UpdateProfileById(updateCtx); err != nil {
		log.Logger.Error("Fail to Update Order ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, &common.HttpResp{
			ErrCode: "109",
			ErrMsg:  "上传错误",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    updateCtx.GetResult(),
	})
}

//下载Order,以excel形式导出
func (handler *Handler) excelHandler(sheetName, outPutFileUrl string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		log.Logger.Error("Fail to Add Table Sheet ", zap.Error(err))
		return err
	}
	var page, pageSize = 1, 100

	ctx := &model2.QueryCtxs{
		ItemTyp: itemOrder,
		ReqPage: page,
		ReqSize: pageSize,
	}
	err = handler.profileManager.QueryProfiles(ctx)
	if err != nil {
		log.Logger.Error("Fail to Query Orders ", zap.Error(err))
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
		log.Logger.Error("Fail to Save Form  ", zap.Error(err))
		return err
	}
	return nil
}

func (handler *Handler) mapTransformer(order *model.Order) map[string]interface{} {
	m := map[string]interface{}{
		"id":        order.ID,
		"order_no":  order.OrderNo,
		"user_name": order.UserName,
		"amount":    order.Amount,
		"status":    order.Status,
		"file_url":  order.FileUrl,
	}
	return m
}

//单文件上传
func (handler *Handler) singleFileUpload(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "110",
			ErrMsg:  "单文件上传错误",
		})
		return "", err
	}
	dst := fmt.Sprintf("./file/" + file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Logger.Error("Fail to Save File ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, &common.HttpResp{
			ErrCode: "111",
			ErrMsg:  "单文件保存失败",
		})
		return "", err
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    dst,
	})
	return dst, nil
}
