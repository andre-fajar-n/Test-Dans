package postgre

import (
	"dans/entity"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewConnection(cfg *viper.Viper) *gorm.DB {
	fmt.Println("Start Connecting to DB...")

	cfgDB := cfg.Sub("postgre")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable ",
		cfgDB.GetString("host"),
		cfgDB.GetString("user"),
		cfgDB.GetString("password"),
		cfgDB.GetString("db"),
		cfgDB.GetInt64("port"),
	)

	cfgGorm := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), cfgGorm)
	if err != nil {
		log.Fatalln("error connect to db")
	}

	fmt.Println("Success to connect DB")

	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.User{},
	)
	if err != nil {
		log.Fatalln("error auto migrate")
	}
}
