package service

import (
	"fmt"
	"time"

	"github.com/api/global"
	"github.com/hellofresh/health-go/v5"
	healthPg "github.com/hellofresh/health-go/v5/checks/postgres"
	healthRedis "github.com/hellofresh/health-go/v5/checks/redis"
)

type IHealthCheckService interface {
	HealthCheck() *health.Health
}

type healthCheckService struct {}

func NewHealthCheckService() IHealthCheckService {
	return &healthCheckService{}
}

func (hcs *healthCheckService) HealthCheck() *health.Health {
	dbConfig := global.Config.DB
	dsn := "postgres://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=%s"
	var dbConnection = fmt.Sprintf(dsn, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.SslMode, dbConfig.Timezone)

	h, _ := health.New(health.WithComponent(health.Component{
		Name: "api",
		Version: "1.0.0",
	}), health.WithChecks(health.Config{
		Name: "db",
		Timeout: 10 * time.Second,
		SkipOnErr: false,
		Check: healthPg.New(healthPg.Config{
			DSN: dbConnection,
		}),
	}))

	h.Register(health.Config{
		Name: "redis",
		Timeout: 10 * time.Second,
		SkipOnErr: false,
		Check: healthRedis.New(healthRedis.Config{
			DSN: global.Config.Redis.Addr,
		}),
	})

	return h
}