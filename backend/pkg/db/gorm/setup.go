package gorm

import (
	"fmt"
	"gochat/config"
	"gochat/helper/util"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// refer to: https://github.com/go-sql-driver/mysql
func setup(configType ...config.ConfigType) *gorm.DB {
	conf := config.GetConfig(configType...)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cant establish db connection " + err.Error())
	}
	return db
}

func DevSetup() *gorm.DB {
	return setup()
}

func ProdSetup() *gorm.DB {
	return setup(config.Prod)
}

func TestSetup() *gorm.DB {
	conf := gorm.Config{}
	filename := util.MakeBinDir() + "/test.db"
	db, err := gorm.Open(sqlite.Open(filename), &conf)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
