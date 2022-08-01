package user

import (
	"MY-WEB/server/db"
	"MY-WEB/server/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var welcome = ".\\webpage\\welcome.html"
var register = ".\\webpage\\register.html"
var getuser = ".\\webpage\\getuser.html"
var getlogin = ".\\webpage\\getlogin.html"

var UserMap = make(map[string]string) //将用户IP与id相联系

func helloworld(c *gin.Context) {
	c.String(http.StatusOK, "hello user")
}

func sendjson(c *gin.Context) {
	data := gin.H{
		"name": "apowerfulmei",
		"psw":  "a good man",
		"id":   20,
	}
	c.JSON(http.StatusOK, data)
}

func welcome_page(c *gin.Context) {
	//加载欢迎html页面
	models.LoadHTML(c, welcome)
}
func register_page(c *gin.Context) {
	//加载信息注册html页面
	models.LoadHTML(c, register)
}
func login_page(c *gin.Context) {
	//加载登录html页面
	var DB db.MyDB
	DB.Linkdb()
	_ID := c.PostForm("id")
	_psw := c.PostForm("psw")
	fmt.Println(_psw)
	msg := DB.GetData(_ID)
	fmt.Println("aa", msg.Psw)
	if msg.Psw == _psw {
		models.LoadHTML(c, getlogin)
		//显示登录日志
		logs := c.ClientIP() + " 用户" + msg.Name + "登录了"
		models.Write_userlog(_ID, logs)
		log.Println(logs)
		UserMap[c.ClientIP()] = msg.Userid
	} else {
		c.String(http.StatusOK, "password or ID wrong")
	}
}
func post_page(c *gin.Context) {
	//用户提交对话框
	models.LoadHTML(c, getlogin)
	_Word := c.PostForm("words")
	_SID := UserMap[c.ClientIP()]
	_DID := c.PostForm("id")
	//写入日志
	logs := _SID + "->" + _DID + ":" + _Word
	fmt.Printf(logs)
	models.Write_userlog(_SID, logs)
}
func getm_page(c *gin.Context) {
	//获取用户请求并发送其相关信息
	_SID := UserMap[c.ClientIP()]
	var DB db.MyDB
	DB.Linkdb()
	msg := DB.GetData(_SID)
	picpath := "./usersdata/" + _SID + "/pic/" + msg.Pic
	c.JSON(200, gin.H{
		"name": msg.Name,
		"id":   msg.Userid,
		"pic":  models.ReadPic(picpath),
	})
}
func getuser_page(c *gin.Context) {
	//注册
	var DB db.MyDB
	DB.Linkdb()
	models.LoadHTML(c, getuser)
	//获取用户传输的信息
	_name := c.PostForm("name")
	_psw := c.PostForm("psw")
	_pic := "1.jpg"
	ID := models.FormID(db.TotalUser)
	//存储在map中
	UserMap[c.ClientIP()] = ID
	//存储到数据库中
	DB.InsertData(db.FormUser(_name, _psw, ID, _pic))
	db.TotalUser += 1
	//在user/users/目录下为用户创建数据文件夹
	models.Init_user_dir(ID)
	//为用户随机分配一个头像
	src_name := "./usersdata/default_pic/" + strconv.Itoa(models.Form_rannum(6)) + ".jpg"
	dst_name := "./usersdata/" + ID + "/pic/" + _pic
	models.Tran_file(src_name, dst_name)

}
func User(e *gin.Engine) {
	//初始界面
	e.GET("/user", welcome_page)

}

func RegUser(e *gin.Engine) {
	//显示用户注册界面
	e.GET("/user_reg", register_page)
	e.POST("/user_reg", getuser_page)
}

func LoginUser(e *gin.Engine) {
	//用户登录
	e.POST("/user_log", login_page)
	e.POST("/user_post", post_page)
	e.GET("/get_msg", getm_page)
}
