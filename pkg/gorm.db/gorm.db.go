package gormdb

import (
	"fmt"
	"sync"

	"github.com/aldyN25/todolist/app/configs"
	"github.com/aldyN25/todolist/app/models"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}
var db *gorm.DB

func GetInstance() (*gorm.DB, error) {
	configs := configs.GetInstance()
	// if db == nil {
	dsn := fmt.Sprintf("%v:%v@(%v:%v)/%v?parseTime=true",
		configs.Dbconfig.Username,
		configs.Dbconfig.Password,
		configs.Dbconfig.Host,
		configs.Dbconfig.Port,
		configs.Dbconfig.Dbname,
	)

	lock.Lock()
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	lock.Unlock()
	if err != nil {
		return nil, err
	}
	// Migrate Here
	errors := db.AutoMigrate(&models.Activities{})
	if errors != nil {
		return nil, errors
	}
	errors = db.AutoMigrate(&models.Todos{})
	if errors != nil {
		return nil, errors
	}
	// fmt.Println("[DATABASE] : ", db)
	return db, nil
}
