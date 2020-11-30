package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tealeg/xlsx"
	"log"
	"net/http"
	"strconv"
	"testGinandGorm/db"
	"testGinandGorm/pkg/model"
)

func DeleteOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var order model.Demo_order
	d := db.Db.Where("id = ?", id).Delete(&order)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateOrder(c *gin.Context) {

	var order model.Demo_order
	id := c.Params.ByName("id")

	if err := db.Db.Where("id = ?", id).First(&order).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&order)

	db.Db.Save(&order)
	c.JSON(200, order)

}

func CreateOrder(c *gin.Context) {

	var order model.Demo_order
	c.BindJSON(&order)

	db.Db.Create(&order)
	c.JSON(200, order)
}

func GetOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var order model.Demo_order
	if err := db.Db.Where("id = ?", id).First(&order).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, order)
	}
}

func GetOrderList(c *gin.Context) {
	var orderList []model.Demo_order
	if err := db.Db.Find(&orderList).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, orderList)
	}
}

//根据user_name做模糊查找、根据创建时间、金额排序
func GetSortedOrderList(c *gin.Context) {
	var order model.Demo_order
	var orderList []model.Demo_order
	c.BindJSON(&order)
	likeName := order.User_name
	fmt.Scan(likeName)
	db.Db = db.Db.Raw("select * from demo_order where user_name like ? ORDER BY amount DESC", "%"+likeName+"%").Scan(&orderList)
	c.JSON(200, orderList)
	fmt.Println(orderList)
}

//获取文件url并保存
func GetUploadUrl(c *gin.Context) {
	//var order model.Demo_order
	id := c.Params.ByName("id")
	fmt.Println("id:" + id)

	//开启事务
	tx := db.Db.Begin()
	str := TestUpload(c)
	sql := "UPDATE demo_order SET file_url=?  WHERE id=?"
	result := tx.Exec(sql, str, id)
	if result == nil {
		tx.Rollback()
	}
	fmt.Println(str)
	tx.Commit()
}

//下载demo_order,以excel形式导出
func DownLoadExcel(c *gin.Context) {
	var outOutFile = "order.xlsx"
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("order_list")
	if err != nil {
		fmt.Println(err.Error())
	}
	list := GetDownLoadList()
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
	for _, order := range list {
		row := sheet.AddRow()
		orderId := row.AddCell()
		orderId.Value = strconv.Itoa(order.ID)
		fmt.Println(orderId.Value)
		orderNo := row.AddCell()
		orderNo.Value = order.Order_No
		orderName := row.AddCell()
		orderName.Value = order.User_name
		orderAmount := row.AddCell()
		orderAmount.Value = strconv.FormatFloat(float64(order.Amount), 'f', 3, 64)
		orderStatus := row.AddCell()
		orderStatus.Value = order.Status
		orderFile := row.AddCell()
		orderFile.Value = order.File_url
	}
	err = file.Save(outOutFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("\n\n export success")

}

func GetDownLoadList() []model.Demo_order {
	var orderList []model.Demo_order
	db.Db = db.Db.Raw("select * from demo_order").Scan(&orderList)
	return orderList
}

//单文件上传
func TestUpload(c *gin.Context) string {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("ERROR: upload file failed. ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("ERROR: upload file failed. %s", err),
		})
	}
	log.Println(file.Filename)
	dst := fmt.Sprintf("./common/" + file.Filename)
	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Println("ERROR: save file failed. ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("ERROR: save file failed. %s", err),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":      "file upload succ.",
		"filepath": dst,
	})
	return dst
}
