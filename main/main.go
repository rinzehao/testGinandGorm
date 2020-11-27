package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Demo_order struct {
	ID        uint    `json:"id"`
	Order_No  string  `json:"order_no"`
	user_name string  `json:"user_name"`
	amount    float64 `json:"amount"`
	status    string  `json:"status"`
	file_url  string  `json:"file_url"`
}

func main() {

	db, _ = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test_gorm?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Demo_order{})
	r := gin.Default()
	r.POST("/order/create", CreateOrder)    //1）创建 demo_order
	r.PUT("/order/update/:id", UpdateOrder) //2)更新demo_order （amount、stuatus、file_url）
	r.GET("/order/:id", GetOrder)           //3)获取demo_order的详情
	r.GET("/order/", GetOrderList)          //4）获取demo_order列表
	r.DELETE("/order/delete/:id", DeleteOrder)
	r.Run(":8080")

}
func DeleteOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Demo_order
	d := db.Where("id = ?", id).Delete(&order)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateOrder(c *gin.Context) {

	var order Demo_order
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&order).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&order)

	db.Save(&order)
	c.JSON(200, order)

}

func CreateOrder(c *gin.Context) {

	var order Demo_order
	c.BindJSON(&order)

	db.Create(&order)
	c.JSON(200, order)
}

func GetOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Demo_order
	if err := db.Where("id = ?", id).First(&order).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, order)
	}
}

func GetOrderList(c *gin.Context) {
	var orderList []Demo_order
	if err := db.Find(&orderList).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, orderList)
	}

}
