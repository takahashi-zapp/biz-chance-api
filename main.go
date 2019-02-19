package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/one-connect/biz-chance-api/conf"
	"github.com/one-connect/biz-chance-api/infrastructure/persistence/datastore"
	"log"

	//"gopkg.in/go-playground/validator.v8"
)

type User struct {
	gorm.Model
	Name string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	Email string `gorm:"size:255"`
}

func main() {
	// conf読み取り
	conf.ReadConf()

	// DB取得
	conn := datastore.NewMySqlDB()
	// db stop
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	conn.AutoMigrate(&User{})

	// interactor
	// r := registry.NewInteractor(conn)

	// 依存解決
	// h := r.NewAppHandler()

	// gin
	//g := gin.New()

	//// middleware
	//g.Use(middleware.Logger())
	//g.Use(middleware.Recover())
	//
	//// validate
	//g.Validator = &validater.CustomValidator{Validator: validator.New()}

	// router
	//router.NewRouter(g, h)
	//g.Run(":8000")
}