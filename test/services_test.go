package test

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"io/ioutil"
//	"net/http/httptest"
//	"testGinandGorm/db"
//	dao2 "testGinandGorm/pkg/dao"
//	"testGinandGorm/pkg/service"
//	"testing"
//)
//
//var r = gin.Default()
//
//
//type Demo_order struct {
//	ID       int     `form:"id" json:"id" binding:"required"`
//	OrderNo  string  `form:"order_no" json:"order_no" binding:"required"`
//	UserName string  `form:"user_name" json:"user_name" binding:"required"`
//	Amount   float64 `form:"amount" json:"amount" binding:"required"`
//	Status   string  `form:"status" json:"status" binding:"required"`
//	FileUrl  string  `form:"file_url" json:"file_url" binding:"required"`
//}
//
//func init(){
//	r.POST("/order/create", service.CreateOrder)       //1）创建 demo_order
//	r.PUT("/order/update/:id", service.UpdateOrder)    //2)更新demo_order （amount、stuatus、file_url）
//	r.GET("/order/:id", service.GetOrder)              //3)获取demo_order的详情
//	r.GET("/order", service.GetOrderList)              //4）获取demo_order列表
//	r.GET("/orderSearch", service.GetSortedOrderList)  //5)获取模糊查询结果
//	r.GET("/orderDownload", service.DownLoadExcel)     //6)下载xlsx表格
//	r.POST("/orderUpload/:id", service.GetUploadUrl)   //7)上传文件
//	r.DELETE("/order/delete/:id", service.DeleteOrder) //8)删除demo_order
//	r.LoadHTMLGlob("127.0.0.1:8080/") //定义模板文件路径
//
//}
//
//func Get(uri string, router *gin.Engine) []byte {
//	// 构造get请求
//	req := httptest.NewRequest("GET", uri, nil)
//	// 初始化响应
//	w := httptest.NewRecorder()
//
//	// 调用相应的handler接口
//	router.ServeHTTP(w, req)
//
//	// 提取响应
//	result := w.Result()
//	defer result.Body.Close()
//
//	// 读取响应body
//	body,_ := ioutil.ReadAll(result.Body)
//	return body
//}
//
//func PostForm(uri string, param map[string]string, router *gin.Engine) []byte {
//	// 构造post请求，表单数据以querystring的形式加在uri之后
//	req := httptest.NewRequest("POST", uri+ParseToStr(param), nil)
//
//	// 初始化响应
//	w := httptest.NewRecorder()
//
//	// 调用相应handler接口
//	router.ServeHTTP(w, req)
//
//	// 提取响应
//	result := w.Result()
//	defer result.Body.Close()
//
//	// 读取响应body
//	body, _ := ioutil.ReadAll(result.Body)
//	return body
//}
//
//func PostJson(uri string, param map[string]interface{}, router *gin.Engine) []byte {
//	// 将参数转化为json比特流
//	jsonByte,_ := json.Marshal(param)
//
//	// 构造post请求，json数据以请求body的形式传递
//	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
//
//
//	// 初始化响应
//	w := httptest.NewRecorder()
//
//	// 调用相应的handler接口
//	router.ServeHTTP(w, req)
//
//	// 提取响应
//	result := w.Result()
//	defer result.Body.Close()
//
//	// 读取响应body
//	body,_ := ioutil.ReadAll(result.Body)
//	return body
//}
//
//func TestOnGetOrder(t *testing.T) {
//	// 初始化请求地址
//	uri := "order/3"
//
//	// 发起Get请求
//	body := Get(uri, r)
//	fmt.Printf("response:%v\n", string(body))
//	fmt.Println(body)
//
//	// 判断响应是否与预期一致
//	//if string(body) != "success" {
//	//	t.Errorf("响应字符串不符，body:%v\n",string(body))
//	//}
//}
//
//func TestCreateOrder(t *testing.T)  {
//	id:="1"
//	db:=db.DbInit()
//	dao:=dao2.NewDao(db)
//	s:=service.NewService(dao)
//	err:=s.DeleteOrder(id)
//}
//
//
//// ParseToStr 将map中的键值对输出成querystring形式
//func ParseToStr(mp map[string]string) string {
//	values := ""
//	for key, val := range mp {
//		values += "&" + key + "=" + val
//	}
//	temp := values[1:]
//	values = "?" + temp
//	return values
//}
