package main

import (
	"Go/internal/api"
	"Go/pkg/tgin/middleware"
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//initDefaultRouter()
	//initBlankRouter()
	//TwoServer()
	StartAndExit()
}

func initDefaultRouter()  {
	// 写入日志的文件
	f, _ := os.Create("./logs/gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果你需要同时写入日志文件和控制台上显示，使用下面代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api.SimpleDemo(router)
	api.PathDemo(router)
	api.QueryStringDemo(router)
	api.MultipartUrlencodedDemo(router)
	api.UploadFileDemo(router)
	api.UploadFilesDemo(router)
	api.GroupRouterDemo(router)
	api.BindDemo(router)			//模型绑定和验证
	api.BookableDateDemo(router)	//自定义验证器
	api.ShouldBindQuery(router)		//只绑定查询字符串
	api.ShouldBind(router)			//绑定查询字符串或 post 数据
	api.BindCheckBox(router)		//绑定 HTML 复选框
	api.XmlJsonYamlRender(router)   //XML, JSON 和 YAML 渲染
	api.SecureJSON(router)			//SecureJSON
	api.JsonP(router)				//JSONP
	api.AsciiJson(router)			//AsciiJSON
	api.StaticFileService(router)	//静态文件服务
	api.ReadFromReader(router)		//从reader读取数据
	api.Redirect(router)			//路由重定向

	router.Run(":8080")
}

func initBlankRouter()  {
	r := gin.New()
	r.Use(middleware.Logger())

	api.CustomMiddleware(r)
	r.Run(":8080")
}

/**
使用 Gin 运行多个服务
 */
func TwoServer()  {
	var (
		g errgroup.Group
	)
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      api.Router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      api.Router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

/**
正常的重启或停止
 */
func StartAndExit() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 连接服务器
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号超时　５　秒正常关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}