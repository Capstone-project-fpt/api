package initialize

import (
	"fmt"

	"github.com/api/global"
	"github.com/api/internal/worker"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-playground/validator/v10"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()
	InitRedis()
	InitAsynq()
	InitGoth()
	InitI18n()
	r := InitRouter()

	serverAddr := fmt.Sprintf("%s:%v", global.Config.Server.ServerIP, global.Config.Server.Port)

	if global.Config.Server.Mode != "release" {
		fmt.Printf("Swagger API Docs: http://%v/swagger/index.html\n", serverAddr)
	}

	go func() {
		if err := worker.InitWorker(); err != nil {
			panic(err)
		}
	}()

	global.AwsSession, _ = session.NewSession(&aws.Config{
		Region: &global.Config.AWS.Region,
		Credentials: credentials.NewStaticCredentials(
			global.Config.AWS.AccessKey,
			global.Config.AWS.SecretKey,
			"",
		),
	})

	global.S3Client = s3.New(global.AwsSession)
	global.Validator = validator.New(validator.WithRequiredStructEnabled())

	r.Run(serverAddr)
}
