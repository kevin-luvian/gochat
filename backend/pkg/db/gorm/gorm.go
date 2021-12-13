package gorm

import (
	"sync"

	"gochat/env"

	"gorm.io/gorm"
)

var once sync.Once
var instance *gorm.DB

func GetDB() *gorm.DB {
	if instance == nil {
		once.Do(makeInstance)
	}
	return instance
}

func makeInstance() {
	if env.IsDevEnv() {
		instance = DevSetup()
	} else if env.IsTestEnv() {
		instance = TestSetup()
	} else {
		instance = ProdSetup()
	}
}
