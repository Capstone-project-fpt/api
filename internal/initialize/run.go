package initialize

import (
	"fmt"

	"github.com/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()
	InitRedis()
	r := InitRouter()

	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)

	r.Run(serverAddr)
}
