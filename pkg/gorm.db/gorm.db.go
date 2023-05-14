package gormdb

import (
	"fmt"
	"sync"

	"github.com/aldyN25/todolist/app/configs"
	// "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
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
	// return db, nil
	// }
	// fmt.Println("[DATABASE] : ", db)
	return db, nil
}

// type Server struct {
// 	DB     *gorm.DB
// 	Router *mux.Router
// }

// func Initialize(Dbdriver string) (*gorm.DB, error) {
// 	var server Server
// 	configs := configs.GetInstance()

// 	var err error

// 	if Dbdriver == "mysql" {
// 		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 			configs.Dbconfig.Username,
// 			configs.Dbconfig.Password,
// 			configs.Dbconfig.Host,
// 			configs.Dbconfig.Port,
// 			configs.Dbconfig.Dbname,
// 		)
// 		server.DB, err = gorm.Open(Dbdriver, DBURL)
// 		if err != nil {
// 			fmt.Printf("Cannot connect to %s database", Dbdriver)
// 			log.Fatal("This is the error:", err)
// 		} else {
// 			fmt.Printf("We are connected to the %s database", Dbdriver)
// 		}
// 	}

// 	return server.DB, nil
// }
