package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/2 20:07
*@Version: V1.0
 */
var db *gorm.DB
var lock sync.Mutex

const (
	driveName = "mysql"
	dsn       = "root:root@(127.0.0.1:3306)/motion?charset=utf8mb4&parseTime=True&loc=Local"
)

func GetDB() *gorm.DB {
	if db == nil {
		lock.Lock()
		if db == nil {
			InitDB()
		}
		lock.Unlock()
	}
	return db
}

func InitDB() {
	var err error
	if db, err = gorm.Open(driveName, dsn); err != nil {
		println("数据库连接错误", err.Error())
		return
	}
}
