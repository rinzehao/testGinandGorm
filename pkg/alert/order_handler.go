package alert

import (
	"container/list"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"golang.org/x/net/xsrftoken"
	"net/http"
	"reflect"
	"strconv"
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
	var order model.DemoOrder
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "JSON绑定错误",
		})
		return
	}
	ctx := &model.OrderMade{
		Order: &model.DemoOrder{
			ID: order.ID,
			OrderNo: order.OrderNo,
			UserName: order.UserName,
			Amount: order.Amount,
			Status: order.Status,
			FileUrl: order.FileUrl,
			},
	}
	if err := handler.orderService.Create(ctx); err != nil {
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
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "JSON绑定错误",
		})
		return
	}
	ctx := &model.OrderMade{
		OrderNo: order.OrderNo,
		UpdateMap: mapTransformer(&order),
	}
	if err := handler.orderService.UpdateNo(ctx); err != nil {
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
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "101",
			ErrMsg:  "id输入错误",
		})
		return
	}
	ctx := &model.OrderMade{
		OrderID: id,
	}
	err := handler.orderService.QueryID(ctx)
	if err != nil {
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
		Page: page,
		PageSize: pageSize,
	}
    err := handler.orderService.QueryMulti(ctx)
	if err != nil {
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
		c.JSON(http.StatusBadRequest, &common.HttpResp{
			ErrCode: "100",
			ErrMsg:  "模糊查找条目失败:JSON绑定错误",
		})
		return
	}
	ctx := &model.OrderMade{
		UserName: order.UserName,
		OrderBy: "amount",
		Desc: "desc",
	}
	err := handler.orderService.QueryName(ctx)
	if err != nil {
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

//下载DemoOrder,以excel形式导出
func (handler *MyOrderHandler) excelHandler(sheetName, outPutFileUrl string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		return err
	}
	var page, pageSize = 1, 100
	ctx := &model.OrderMade{
		Page: page,
		PageSize: pageSize,
	}
	err = handler.orderService.QueryMulti(ctx)
	if err != nil {
		return err
	}

	origin := []string{"123", "345", "abc"}
	if true {
		switch reflect.TypeOf(origin).Kind(){
		case reflect.Slice, reflect.Array:
			v := reflect.ValueOf(origin)
			for i :=0 ;i < v.Len();i++ {
				fmt.Println(v.Index(i))
			}
		case reflect.String:
			v := reflect.ValueOf(origin)
			fmt.Println(v.String(),"是String类型的")
		case reflect.Int:
			v := reflect.ValueOf(origin)
			t := v.Int()
			fmt.Println(t,"int type")
		}
	}
	orders := list.New()
	if true {
		switch reflect.TypeOf(ctx.Group).Kind(){
		case reflect.Slice, reflect.Array:
			v := reflect.ValueOf(origin)
			for i :=0 ;i < v.Len();i++ {
				orders.PushBack(v.Index(i))
			}
		}
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
	for order := orders.Front(); order != nil; order = order.Next() {
		fmt.Print(i.Value, " ")
	}

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


func mapTransformer(order *model.DemoOrder) map[string]interface{} {
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

