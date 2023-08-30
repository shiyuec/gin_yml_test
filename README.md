在使用 Gin（一个 Go 语言的 Web 框架）来区分测试和生产环境的情况下，你可以通过不同的配置文件和环境变量来实现。以下是一个基本的代码示例，演示如何在 Gin 中区分测试和生产环境。

首先，你需要安装 Gin 框架和一个用于加载配置的库，比如 viper。你可以使用以下命令安装它们：

```bash
go mod init your_project
go get -u github.com/gin-gonic/gin
go get -u github.com/spf13/viper
```

然后，创建一个目录结构如下：

```bash
your_project/
├── config/
│   ├── config.yaml
│   ├── config.test.yaml
├── main.go
```

在 `config/config.yaml` 中，存储生产环境的配置：

```yaml
app:
  environment: production
  port: 8080
```

在 `config/config.test.yaml` 中，存储测试环境的配置：

```yaml
app:
  environment: test
  port: 3000
```

接下来，在 `main.go` 中编写代码来加载不同环境下的配置文件，并根据环境设置 Gin 服务器：

```go
package main

import (
	"log"
	"os"

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
```

在上述代码中，我们根据环境变量 `"ENVIRONMENT"` 的值来选择加载不同的配置文件。如果环境变量是 `"test"`，则加载测试环境的配置文件；否则，加载生产环境的配置文件。

确保在实际部署时设置正确的环境变量，比如在生产环境中设置为 `"production"`，在测试环境中设置为 `"test"`。

这个示例只是一个基本的框架，你可以根据实际需求进一步扩展和优化代码。例如，可以添加数据库连接、日志记录、中间件等功能来完善你的应用程序。