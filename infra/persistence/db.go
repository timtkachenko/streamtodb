package persistence

import (
	"github.com/google/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"sync"
	"time"
)

type ConnectionConfig struct {
	retryCount   int
	retryTimeOut time.Duration
}

var (
	onceDB   = &sync.Once{}
	db       *gorm.DB
	dbConfig = &ConnectionConfig{5, time.Second * 3}
)

func ConnectDb() *gorm.DB {
	onceDB.Do(dbConnect)
	return db
}

// attempt to connect to the DB server with retry
func dbConnect() {
	tryCount := 1
	for {
		logger.Info("Connecting to DB ...", viper.GetString("POSTGRES_DSN"))
		var err error
		db, err = gorm.Open("postgres", viper.Get("POSTGRES_DSN"))
		if err != nil {
			logger.Info("dbConnect failed: " + err.Error())
			db = nil
		} else if db != nil {
			db.LogMode(viper.GetBool("DEBUG"))
			db.DB().SetMaxOpenConns(5)
			db.DB().SetConnMaxLifetime(10 * time.Minute)
			logger.Info("DB Connected")
			return
		}
		if tryCount++; tryCount > dbConfig.retryCount {
			logger.Fatal("unable to connect to DB")
		}
		<-time.After(dbConfig.retryTimeOut)
	}
}
