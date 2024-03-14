package database

import (
	"fmt"
	"net/url"
	"time"

	"github.com/nazudis/disbursement/src/config"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db = newDB()

func GetDB() *gorm.DB {
	return db
}

func newDB() (DB *gorm.DB) {
	var (
		host = viper.GetString(config.DBHost)
		port = viper.GetString(config.DBPort)
		user = viper.GetString(config.DBUser)
		pass = viper.GetString(config.DBPass)
		name = viper.GetString(config.DBName)
		tz   = viper.GetString(config.DBTimeZone)
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, pass, name, port, url.QueryEscape(tz))

	gormConfig := &gorm.Config{SkipDefaultTransaction: true}
	if viper.Get(config.Mode) == config.ModeDev {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Error)
	}

	DB, err := gorm.Open(postgres.Open(dsn), gormConfig)

	if err != nil {
		panic(err)
	}

	DB.Set("gorm:auto_preload", true)
	DB.Session(&gorm.Session{
		AllowGlobalUpdate:    true,
		FullSaveAssociations: false,
	})

	sDB, _ := DB.DB()
	sDB.SetMaxOpenConns(50)
	sDB.SetMaxIdleConns(1)
	sDB.SetConnMaxIdleTime(1 * time.Minute)

	return
}
