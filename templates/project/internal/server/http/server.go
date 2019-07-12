package http

import (
	"fmt"
	"goadmin/internal/api/http/v1"
	"goadmin/internal/common/middleware/cors"
	"goadmin/internal/conf"
	"goadmin/pkg/utils"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	ginglog "github.com/szuecs/gin-glog"

	"goadmin/pkg/jsonutil"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/g/os/glog"
)

// 跨域设置
var authCors = cors.Config{
	Origins:         "*",
	Methods:         "GET, PUT, POST, DELETE",
	RequestHeaders:  "Origin, X-Token, Content-Type, X-Client-Id",
	ExposedHeaders:  "",
	MaxAge:          50 * time.Second,
	Credentials:     true,
	ValidateHeaders: false,
}

// http server配置
func New(router *gin.Engine) (srv *http.Server) {
	router.Use(cors.Middleware(authCors))
	router.Use(ginglog.Logger(3 * time.Second))
	router.Use(MyRecovery())

	// 对频繁的请求限流
	RequestLimiter(router)

	// init routes
	v1.InitRouter(router)

	// http server
	srv = &http.Server{
		Addr:         conf.ServerConf.Addr,
		ReadTimeout:  conf.ServerConf.ReadTimeout.Duration,
		WriteTimeout: conf.ServerConf.WriteTimeout.Duration,
		Handler:      router,
	}

	log.Printf("[GIN-INFO] server is runing listen %s\n", conf.ServerConf.Addr)

	// 启动http server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Fatalf("server listen error: %s\n", err)
		}
	}()

	return
}

// 针对每个IP的请求频率进行限流，每秒最多1次请求
func RequestLimiter(router gin.IRouter) {
	limiter := tollbooth.NewLimiter(1, nil)
	router.GET("/", tollbooth_gin.LimitHandler(limiter), func(c *gin.Context) {
		c.String(200, "服务繁忙，请稍后再试...")
	})
}

// 错误全局handle处理
func MyRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				glog.Infof("%v", string(debug.Stack()))
				resbody := map[string]interface{}{
					"code": 500,
					"msg":  fmt.Sprintf("%v", err),
				}

				c.String(http.StatusOK, jsonutil.MarshalToString(resbody))
				glog.Infof("MyRecovery response=%v", utils.PrettyStruct(resbody))
			}
		}()
		c.Next()
	}
}
