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
------
如果你想要持久性地设置环境变量，你可以在系统的环境变量设置中进行配置，这样你就不需要每次手动设置。在 Windows 操作系统中，你可以按照以下步骤进行：

在 Windows 搜索栏中输入 "环境变量" 并选择 "编辑系统环境变量"。
在弹出的窗口中，点击 "环境变量" 按钮。
在 "用户变量" 或 "系统变量" 部分，点击 "新建" 按钮。
输入变量名为 ENVIRONMENT，变量值为 test，然后点击 "确定"。
重新打开一个命令行窗口，运行 go run 命令，程序将会使用设置好的环境变量。
这样设置之后，你就不需要每次手动输入环境变量，它将会在每个新的命令行会话中自动生效。

或者


在 Windows 命令行中，可以使用以下命令来设置环境变量：

setx ENVIRONMENT test

这将会将 ENVIRONMENT 设置为 test，并使其持久化，以便在新的命令行窗口中也能生效。

请注意，setx 命令设置的环境变量会在新的命令行窗口中立即生效，但在当前窗口中不会立即生效，你可能需要关闭当前窗口并重新打开一个新的窗口来查看更改。