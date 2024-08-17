package initialize

import (
	"fmt"

	"github.com/api/global"
	_ "github.com/api/docs"
)

// @title Capstone API
func Run() {
	LoadConfig()
	InitLogger()
	InitDB()
	InitRedis()
	r := InitRouter()

	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)

	r.Run(serverAddr)
}
