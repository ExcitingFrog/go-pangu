package routers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go-pangu/conf"
	service "go-pangu/controller"
	"go-pangu/middleware"
	"go-pangu/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(sig ...os.Signal) {
	router := SetupRouter()

	if len(sig) == 0 {
		sig = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}

	signalChan := make(chan os.Signal, 1)

	go func() {
		router.Run(fmt.Sprintf(":%v", conf.GetEnv("HTTP_PORT")))
	}()
	signal.Notify(signalChan, sig...)
}

func SetupRouter() *gin.Engine {
	//防止前端跨域请求失败
	router := gin.Default()
	config := cors.DefaultConfig()
	config.ExposeHeaders = []string{"Authorization"}
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	router.Any("/ws", websocket.HandleWebsocket)
	router.GET("/ping", service.PingHandler)
	router.GET("/create", service.CreateUsersHandlerWithContext)
	router.POST("/sms", service.SMSHandler)
	router.POST("/users", service.CreateUsersHandler)
	router.POST("/alipay", service.AliPayHandler)
	router.POST("/alipay_notify", service.AliPayNotifyHandler)
	router.POST("/cities", service.CityListHandler)
	router.POST("/influx_save", service.SaveInfluxDBHandler)
	router.POST("/influx_show", service.ShowInfuxDBHandler)
	authorized := router.Group("/")
	authorized.Use(middleware.Auth("user"))
	{
		authorized.GET("/auth_ping", service.AuthPingHandler)
	}
	users := router.Group("/users")
	{
		users.POST("/sign_up", service.SignUpHandler)
		users.POST("/sign_in", service.SignInHandler)
	}
	users.Use(middleware.Auth("user"))
	{
		users.POST("/change_password", service.ChangePasswordHandler)
	}
	return router
}
