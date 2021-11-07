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
	PageNotFindDemo(r)
	NORouterDemo(r)
	ApiParameter(r)
	UrlParameter(r)
	FormParameter(r)
	RedirectDemo(r)
	// Grouping routes
	upload := r.Group("/upload")
	{
		upload.POST("", UploadDemo)
		upload.POST("/multiple", UploadMultipleDemo)
	}
}

func main() {
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	_ = r.Run("localhost:8000")
}

// NORouterDemo 重新向到页面
func NORouterDemo(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/404")
	})
}

// PageNotFindDemo  404页面
func PageNotFindDemo(r *gin.Engine) {
	r.LoadHTMLGlob("./*")
	r.GET("/404", func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "404",
		})
	})
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
	//URL参数可以通过DefaultQuery()或Query(), QueryArray(), QueryMap()方法获取
	//DefaultQuery()若参数不存在，返回默认值，Query()若不存在，返回空串
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

// UploadDemo 上传文件
func UploadDemo(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	if file != nil {
		fmt.Printf("filename: %s\n", file.Filename)
		err := c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			fmt.Printf("save uploaded file failed! err: %v\n", err)
			panic("save uploaded file failed!")
		}
		c.JSON(http.StatusOK, gin.H{
			"file":    file.Filename,
			"size":    file.Size,
			"message": "success",
			"status":  200,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
			"status":  400,
		})
	}
}

// UploadMultipleDemo 多文件
func UploadMultipleDemo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err,
				"status":  400,
			})
		}
	}()
	var data []interface{}
	form, err := c.MultipartForm()
	if form == nil && err != nil {
		panic(err)
	}
	files := form.File["file"]
	value := form.Value // Value 属性保存除 文件外的其他body数据
	fmt.Printf("form's value: %v\n", value)

	for _, file := range files {
		fmt.Printf("filename: %s\n", file.Filename)
		err := c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			fmt.Printf("save uploaded file failed! err: %v\n", err)
			panic("save uploaded file failed!")
		}
		info := map[string]interface{}{
			"filename": file.Filename,
			"size":     file.Size,
		}
		data = append(data, info)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// RedirectDemo 重定向
func RedirectDemo(r *gin.Engine) {
	r.GET("/redirect", func(c *gin.Context) {
		//c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
		c.Redirect(http.StatusMovedPermanently, "/404")
	})
}
