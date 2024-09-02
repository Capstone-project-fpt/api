package initialize

import (
	"database/sql"
	"fmt"
	"time"

	database "github.com/api/database/sqlc"
	"github.com/api/global"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func checkErrorPanic(err error, errStr string) {
	if err != nil {
		global.Logger.Error(errStr, zap.Error(err))
		panic(err)
	}
}

func InitDB() {
	dbConfig := global.Config.DB
	dsn := "postgres://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=%s"
	var dbConnection = fmt.Sprintf(dsn, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.SslMode, dbConfig.Timezone)
	db, err := sql.Open("postgres", dbConnection)
	checkErrorPanic(err, "Init database connection failed")
	SetPool(db)

	if err = db.Ping(); err != nil {
		checkErrorPanic(err, "Failed to ping database")
	}

	global.Logger.Info("Init database connection success")
	global.Db = database.New(db)
	global.RawDb = db
}

func SetPool(db *sql.DB) {
	dbConfig := global.Config.DB
	db.SetConnMaxIdleTime(time.Duration(dbConfig.MaxIdleConns))
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime))
}
