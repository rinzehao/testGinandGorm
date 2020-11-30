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
	if err := db.Db.Where("id = ?", id).Delete(&order).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateOrder(c *gin.Context) {

	var order model.Demo_order
	id := c.Params.ByName("id")

	if err := db.Db.Where("id = ?", id).First(&order).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("更新条目失败：查询不到相应条目")
	}
	if err := c.BindJSON(&order); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("更新条目失败：JSON绑定错误")
	}

	if err := db.Db.Save(&order).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("更新条目失败：保存错误")
	}
	c.JSON(200, order)

}

func CreateOrder(c *gin.Context) {

	var order model.Demo_order
	if err := c.BindJSON(&order); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("创建条目失败：JSON绑定失败")
	}

	db.Db.Create(&order)
	c.JSON(200, order)
}

func GetOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var order model.Demo_order
	db.Db.LogMode(true)
	if err := db.Db.Where("id = ?", id).First(&order).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("获取条目失败：找不到指定条目")

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

	if err := c.BindJSON(&order); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("模糊查找条目失败:JSO你绑定错误")
	}
	likeName := order.UserName
	fmt.Scan(likeName)
	if err := db.Db.Raw("select * from demo_order where user_name like ? ORDER BY amount DESC",
		"%"+likeName+"%").Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("模糊查找条目失败：sql查询出错")
	} else {
		db.Db = db.Db.Raw("select * from demo_order where user_name like ? ORDER BY amount DESC",
			"%"+likeName+"%").Scan(&orderList)
	}
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
	if str == " " {
		fmt.Println("更新url失败：获取新url为空")
		tx.Rollback()
	}
	sql := "UPDATE demo_order SET file_url=?  WHERE id=?"
	result := tx.Exec(sql, str, id)
	if result == nil {
		fmt.Println("更新url失败：更新操作失败")
		tx.Rollback()
	}
	fmt.Println(str)
	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		fmt.Println("更新url失败：事务提交出错")
	}

}

//下载demo_order,以excel形式导出
func DownLoadExcel(c *gin.Context) {
	db.Db.LogMode(true)
	var outOutFile = "order.xlsx"
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("order_list")
	if err != nil {
		fmt.Println("下载失败：添加表单失败")
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
		orderNo.Value = order.OrderNo
		orderName := row.AddCell()
		orderName.Value = order.UserName
		orderAmount := row.AddCell()
		orderAmount.Value = strconv.FormatFloat(float64(order.Amount), 'f', 3, 64)
		orderStatus := row.AddCell()
		orderStatus.Value = order.Status
		orderFile := row.AddCell()
		orderFile.Value = order.FileUrl
	}
	err = file.Save(outOutFile)
	if err != nil {
		fmt.Println("下载失败：保存表单失败")
		fmt.Println(err.Error())
	}
	fmt.Println("\n\n export success")

}

func GetDownLoadList() []model.Demo_order {
	var orderList []model.Demo_order

	if err := db.Db.Raw("select * from demo_order").Error; err != nil {
		fmt.Println("获取数据库表单失败：查询出错")
		fmt.Println(err.Error())
	}

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
