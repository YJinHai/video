package datasource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var dbLock sync.Mutex
var masterInstance *gorm.DB
var slaveInstance *gorm.DB

// 得到唯一的主库实例
func InstanceDbMaster() *gorm.DB {
	if masterInstance != nil {
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()

	if masterInstance != nil {
		return masterInstance
	}
	return NewDbMaster()
}


func NewDbMaster() *gorm.DB {
	sourcename := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
		true,
		"Local",
		)
	fmt.Print(sourcename)
	fmt.Print("db.setting",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))

	Instance, err := gorm.Open("mysql", sourcename)
	//defer slaveInstance.Close()
	if err != nil {
		log.Fatal("dbhelper gorm.Open error ", err)
		return nil
	}

	masterInstance = Instance
	return masterInstance
}
