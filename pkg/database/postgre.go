package database

import (
	"fmt"
	"log"
	"sync"

	usermodel "github.com/adityarifqyfauzan/go-chat/internal/authentication/domain/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once sync.Once
	conn *gorm.DB
)

func InitPostgreDB(conf *viper.Viper) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s port=%s",
			conf.GetString("db.postgre.host"),
			conf.GetString("db.postgre.name"),
			conf.GetString("db.postgre.user"),
			conf.GetString("db.postgre.pass"),
			conf.GetString("db.postgre.ssl"),
			conf.GetString("db.postgre.port"),
		)

		var err error
		conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("unable to connecting database: %v", err)
		}

	})
	conn.AutoMigrate(&usermodel.User{})
	return conn
}
