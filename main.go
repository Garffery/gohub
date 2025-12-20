package main

import (
	"flag"
	"fmt"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
	"gohub/pkg/sms"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {
	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	fmt.Println("配置加载开始")
	config.InitConfig(env)
	fmt.Println("配置加载完成")

	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 初始化 DB
	fmt.Println("数据库初始化开始")
	bootstrap.SetupDB()

	// 初始化 Redis
	bootstrap.SetupRedis()
	fmt.Println("数据库初始化完成")

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	gin.SetMode(gin.ReleaseMode)

	sms.NewSMS().Send("这里填入你的手机号", sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": "123456"},
	})

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}

}
