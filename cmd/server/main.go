package main

import (
	"github.com/api/internal/initialize"
)

// @title Capstone Project FPT API
// @description Capstone Project FPT API
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
// @contact.email minhbeardev@gmail.com
// @Schemes http https
// @query.collection.format multi
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	initialize.Run()
}
