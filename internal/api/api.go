package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/**
简单实例
 */
func SimpleDemo(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

/**
path 中的参数
 */
func PathDemo(router *gin.Engine) {
	// 这个处理器可以匹配 /user/john ， 但是它不会匹配 /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 但是，这个可以匹配 /user/john 和 /user/john/send
	// 如果没有其他的路由匹配 /user/john ， 它将重定向到 /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
}

/**
查询字符串参数
 */
func QueryStringDemo(router *gin.Engine) {
	// 查询字符串参数使用现有的底层 request 对象解析。
	// 请求响应匹配的 URL： /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		// 这个是 c.Request.URL.Query().Get("lastname") 的快捷方式。
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
}

/**
Multipart/Urlencoded 表单
 */
func MultipartUrlencodedDemo(router *gin.Engine) {
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
}

/**
上传文件
通过 postman 上传文件：https://blog.csdn.net/maowendi/article/details/80537304
file: 文件
 */
func UploadFileDemo(router *gin.Engine) {
	// 为 multipart 表单设置一个较低的内存限制（默认是 32 MiB）
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload-file", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 上传文件到指定的 dst 。
		//c.SaveUploadedFile(file, "/Users/tt/Downloads/"+file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
}

/**
上传多个文件
通过postman上传 https://blog.csdn.net/hl449006540/article/details/85015782
upload[]: file1
upload[]: file2
 */
func UploadFilesDemo(router *gin.Engine) {
	// 为 multipart 表单设置一个较低的内存限制（默认是 32 MiB）
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload-files", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件到指定的 dst.
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
}

/**
组路由
 */
func GroupRouterDemo(router *gin.Engine) {
	// 简单组： v1
	v1 := router.Group("/v1")
	{
		v1.GET("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "I am v1 test",
			})
		})
	}

	// 简单组： v2
	v2 := router.Group("/v2")
	{
		v2.GET("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "I am v2 test",
			})
		})
	}
}

/**
XML, JSON 和 YAML 渲染
 */
func XmlJsonYamlRender(r *gin.Engine) {
	// gin.H 是一个 map[string]interface{} 的快捷方式
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// 你也可以使用一个结构
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 在 JSON 中会变成 "user"
		// 将会输出： {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
}

/**
SecureJSON
 */
func SecureJSON(r *gin.Engine) {
	// 你也可以使用自己的安装 json 前缀
	//r.SecureJsonPrefix("while(tt)',\n")

	r.GET("/secure-json", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将会输出  :   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})
}

/**
JSONP 在不同的域中使用 JSONP 从一个服务器请求数据。如果请求参数中存在 callback，添加 callback 到 response body
JSONP是一种非正式传输协议，该协议的一个要点就是允许用户传递一个callback参数给服务端，
然后服务端返回数据时会将这个callback参数作为函数名来包裹住JSON数据，这样客户端就可以随意定制自己的函数来自动处理返回数据了
 */
func JsonP(r *gin.Engine) {
	r.GET("/json-p", func(c *gin.Context) {		// 访问 /json-p?callback=ttFunc
		data := map[string]interface{}{
			"foo": "bar",
		}
		//callback 是 ttFunc
		// 将会输出  :   ttFunc({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})
}

/**
使用 AsciiJSON 生成仅有 ASCII 字符的 JSON，非 ASCII 字符将会被转义
*/
func AsciiJson(r *gin.Engine) {
	r.GET("/ascii-json", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 将会输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})
}

/**
静态文件服务
 */
func StaticFileService(router *gin.Engine) {
	router.StaticFS("/more_static", http.Dir("assets"))
	router.StaticFile("/favicon.ico", "./assets/haha/test.go")
}

/**
从 reader 读取数据
 */
func ReadFromReader(router *gin.Engine) {
	router.GET("/read-from-reader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
}

/**
重定向
 */
func Redirect(router *gin.Engine) {
	router.GET("/redirect", redirect)
	router.GET("/redirect-with-handle", func(context *gin.Context) {
		context.Request.URL.Path = "/redirect-with-handle-test"
		router.HandleContext(context)
	})
	router.GET("/redirect-with-handle-test", func(context *gin.Context) {
		context.JSON(200, gin.H{"hello": "world"})
	})
}

func redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
}

/**
使用 Gin 运行多个服务
 */
func Router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})
	return e
}

func Router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})
	return e
}
