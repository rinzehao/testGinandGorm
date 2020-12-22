package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"log"
	"net/http"
	"strconv"
	"testGinandGorm/common"
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
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		return
	}
	if err := handler.orderService.DeleteOrderById(id); err != nil {
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

func (handler *OrderHandler) QueryOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		return
	}
	order, err := handler.orderService.QueryOrderById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "103",
			ErrMsg:  "获取条目失败：找不到指定条目",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    order,
	})
}

func (handler *OrderHandler) UpdateOrder(c *gin.Context) {
	var order model.DemoOrder
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "JSON绑定错误",
		})
		return
	}
	if err := handler.orderService.UpdateByOrderNo(MapTransformer(&order), order.OrderNo); err != nil {
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

func (handler *OrderHandler) CreateOrder(c *gin.Context) {
	var order model.DemoOrder
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "JSON绑定错误",
		})
		return
	}
	if err := handler.orderService.CreateOrder(&order); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "105",
			ErrMsg:  "创建条目失败",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    order,
	})
}

func (handler *OrderHandler) QueryAllOrders(c *gin.Context) {
	var page, pageSize = 1, 100
	orders, err := handler.orderService.QueryOrders(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "106",
			ErrMsg:  "获取orderList失败",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    orders,
	})
}

//根据user_name做模糊查找、根据创建时间、金额排序
func (handler *OrderHandler) QueryOrders(c *gin.Context) {
	var order model.DemoOrder
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "模糊查找条目失败:JSON绑定错误",
		})
		return
	}
	orders, err := handler.orderService.QueryOrdersByName(order.UserName, "amount", "desc")
	if err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "107",
			ErrMsg:  "模糊查找条目失败:sql查询出错",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    orders,
	})
}

//下载DemoOrder,以excel形式导出
func (handler *OrderHandler) DownLoadExcel(c *gin.Context) {
	var sheetName = "order_List"
	var outPutFileUrl = "order.xlsx"
	if err := handler.excelHandler(sheetName, outPutFileUrl); err != nil {
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
func (handler *OrderHandler) UploadAndUpdate(c *gin.Context) {
	id := c.Params.ByName("id")
	if id, _ := strconv.Atoi(id); id == 0 || id < 0 {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		return
	}
	if _, err := handler.orderService.QueryOrderById(id); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "103",
			ErrMsg:  "获取条目失败：找不到指定条目",
		})
		return
	}
	m := map[string]interface{}{
		"file_url": singleFileUpload(c),
	}
	if err := handler.orderService.UpdateById(m, id); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "109",
			ErrMsg:  "上传错误",
		})
		return
	}
	c.JSON(http.StatusOK, &common.HttpResp{
		Success: true,
		Data:    m,
	})
}

//下载DemoOrder,以excel形式导出
func (handler *OrderHandler) excelHandler(sheetName, outPutFileUrl string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		return err
	}
	orders, err := handler.orderService.QueryOrders(1, 100)
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
	for _, order := range orders {
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
		fmt.Println("下载失败：保存表单失败")
		return err
	}
	return nil
}

//单文件上传
func singleFileUpload(c *gin.Context) string {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("ERROR: upload file failed. ", err)
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "110",
			ErrMsg:  "单文件上传错误",
		})
		return ""
	}
	dst := fmt.Sprintf("./file/" + file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Println("ERROR: save file failed. ", err)
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

func MapTransformer(order *model.DemoOrder) map[string]interface{} {
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
