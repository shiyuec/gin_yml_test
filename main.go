package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 设置配置文件的路径和名称
	configFile := "config/config.yaml"
	env := os.Getenv("ENVIRONMENT")
	if env == "test" {
		configFile = "config/config.test.yaml"
	}

	// 使用 viper 加载配置文件
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// 初始化 Gin 路由
	router := gin.Default()

	// 获取配置中的端口号
	port := viper.GetInt("app.port")

	// 设置路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
			"env":     viper.GetString("app.environment"),
		})
	})

	// 启动 Gin 服务器
	router.Run(":" + strconv.Itoa(port))
}
