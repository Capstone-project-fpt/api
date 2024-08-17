package initialize

import (
	"fmt"
	"time"

	"github.com/api/global"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errStr string) {
	if err != nil {
		global.Logger.Error(errStr, zap.Error(err))
		panic(err)
	}
}

func InitDB() {
	dbConfig := global.Config.DB
	dsn := "host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=%s"
	var dbConnection = fmt.Sprintf(dsn, dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName, dbConfig.Port, dbConfig.SslMode, dbConfig.Timezone)
	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorPanic(err, "Init database connection failed")
	global.Logger.Info("Init database connection success")
	global.MDb = db

	SetPool()
}

func SetPool() {
	dbConfig := global.Config.DB
	db, err := global.MDb.DB()
	checkErrorPanic(err, "Set database pool failed")
	db.SetConnMaxIdleTime(time.Duration(dbConfig.MaxIdleConns))
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime))
}
