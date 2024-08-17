package initialize

import (
	"fmt"
	"time"

	"github.com/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errStr string) {
	if err != nil {
		global.Logger.Error(errStr, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	mysqlConfig := global.Config.MySQL
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var dbConnection = fmt.Sprintf(dsn, mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DbName)
	db, err := gorm.Open(mysql.Open(dbConnection), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorPanic(err, "Init database connection failed")
	global.Logger.Info("Init database connection success")
	global.MDb = db

	SetPool()
}

func SetPool() {
	mysqlConfig := global.Config.MySQL
	db, err := global.MDb.DB()
	checkErrorPanic(err, "Set database pool failed")
	db.SetConnMaxIdleTime(time.Duration(mysqlConfig.MaxIdleConns))
	db.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(mysqlConfig.ConnMaxLifetime))
}
