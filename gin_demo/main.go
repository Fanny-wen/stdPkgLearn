package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 1.创建路由
var r = gin.Default()

func init() {
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	ApiParameter(r)
	UrlParameter(r)
	FormParameter(r)
}

func main() {
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	_ = r.Run("localhost:8000")
}

// ApiParameter API参数
func ApiParameter(r *gin.Engine) {
	// 通过Context的Param方法来获取API参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
}

// UrlParameter URL参数
func UrlParameter(r *gin.Engine) {
	//URL参数可以通过DefaultQuery()或Query()方法获取
	//DefaultQuery()若参数不村则，返回默认值，Query()若不存在，返回空串
	r.GET("/welcome", func(c *gin.Context) {
		fmt.Printf("%T, %#v\n", c.Request.URL, c.Request.URL)
		fmt.Printf("%v, %[1]T\n", c.Request.URL.Query())
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
}

// FormParameter 表单传参
func FormParameter(r *gin.Engine) {
	// 表单参数可以通过PostForm(), PostFormArray(), PostFormMap()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数
	r.POST("/student", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		name := c.PostForm("name")
		pwd := c.PostForm("password")
		hobby, _ := c.GetPostFormArray("hobby")
		parent := c.PostFormMap("parent")
		fmt.Printf("name: %v, pwd: %s, hobby: %v, parent: %v, type: %s\n", name, pwd, hobby, parent, types)
		c.JSON(http.StatusCreated, gin.H{"name": name, "pwd": pwd, "hobby": hobby, "parent": parent})
	})
}
