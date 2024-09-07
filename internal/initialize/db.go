package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/api/global"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	var gormConfig gorm.Config = gorm.Config{
		SkipDefaultTransaction: false,
		PrepareStmt:            true,
	}
	if global.Config.Server.Mode == "dev" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dbConnection), &gormConfig)
	checkErrorPanic(err, "Init database connection failed")

	rawDb, err := db.DB()
	checkErrorPanic(err, "Init database connection failed")
	SetPool(rawDb)

	global.Logger.Info("Init database connection success")
	global.Db = db
}

func SetPool(db *sql.DB) {
	dbConfig := global.Config.DB
	db.SetConnMaxIdleTime(time.Duration(dbConfig.MaxIdleConns))
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime))
}
