package initialize

import (
	"strconv"
	"time"

	"go-ecommerce-be/globals"
	"go-ecommerce-be/internal/po"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	// mysql
	mysqlConfig := globals.Config.Mysql
	connection := mysqlConfig.User + ":" + mysqlConfig.Password + "@tcp(" + mysqlConfig.Host + ":" + strconv.Itoa(mysqlConfig.Port) + ")/" + mysqlConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	CheckErrorPanic(err, "Failed to connect to database")

	globals.MysqlDB = db

	SetPool()
}

func SetPool() {
	// pool
	mysqlPool, err := globals.MysqlDB.DB()
	mysqlConfig := globals.Config.Mysql

	if err != nil {
		globals.Logger.Error("Failed to set pool", zap.Error(err))
		panic(err)
	}

	mysqlPool.SetConnMaxIdleTime(time.Duration(mysqlConfig.MaxIdle))
	mysqlPool.SetMaxOpenConns(mysqlConfig.MaxOpen)
	mysqlPool.SetConnMaxLifetime(time.Duration(mysqlConfig.TimeOut))
}

func MigrateTable() {
	// migrate table
	err := globals.MysqlDB.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		globals.Logger.Error("Failed to migrate table", zap.Error(err))
		panic(err)
	}
}

func CheckErrorPanic(err error, errorString string) {
	if err != nil {
		globals.Logger.Error(errorString, zap.Error(err))
		panic(err)
	}
}
