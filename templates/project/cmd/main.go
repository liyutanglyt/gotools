package main

import (
	"context"
	"flag"
	"goadmin/internal/common/DB"
	"goadmin/internal/conf"
	"goadmin/internal/model"
	"goadmin/internal/server/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gogf/gf/g/os/glog"

	_ "expvar"

	httpv "net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	flag.Parse()

	// configs init
	conf.Init()
	// db start
	DB.Start()
	// create tables
	model.Init()

	// start http server
	srv := http.New(gin.New())
	httpv.ListenAndServe(":1234", nil)

	// 以下代码用于平滑重启，重启时，旧的请求不会中断，新的请求会被新启动的程序接管
	// 注意重启不要使用kill -9 pid,否则无法捕获到信号，请使用kill -2 pid
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	glog.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		glog.Fatal("Server Shutdown: ", err)
	}
}
