package database

import (
	"fmt"
	"github.com/zakariawahyu/go-fiber-gorm-basic/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	DbName   string
	Username string
	Password string
}

func BuildConfig() string {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		DbName:   "golang_fiber_gorm_basic",
		Username: "root",
		Password: "masukdb",
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
	)
}

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(BuildConfig()), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	defer DB.AutoMigrate(
		&models.User{},
		&models.Locker{},
	)
}
