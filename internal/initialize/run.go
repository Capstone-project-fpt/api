package initialize

import (
	"fmt"

	"github.com/api/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()
	InitRedis()
	InitGoth()
	InitI18n()
	r := InitRouter()

	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)

	if global.Config.Server.Mode != "release" {
		fmt.Printf("Swagger API Docs: http://localhost:%v/swagger/index.html\n", global.Config.Server.Port)
	}

	r.Run("127.0.0.1" + serverAddr)
}
