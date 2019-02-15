package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/one-connect/biz-chance-api/conf"
	"github.com/one-connect/biz-chance-api/infrastructure/api/router"
	"github.com/one-connect/biz-chance-api/infrastructure/persistence/datastore"
	"github.com/one-connect/biz-chance-api/registry"
	"gopkg.in/go-playground/validator.v8"
)

func main() {
	// conf読み取り
	conf.ReadConf()

	// DB取得
	conn := datastore.NewMySqlDB()

	// interactor
	r := registry.NewInteractor(conn)

	// 依存解決
	h := r.NewAppHandler()

	// gin
	g := gin.New()

	// middleware
	g.Use(middleware.Logger())
	g.Use(middleware.Recover())

	// validate
	g.Validator = &validater.CustomValidator{Validator: validator.New()}

	// router
	router.NewRouter(g, h)
	g.Run(":8000")

	// DB stop
	defer func() {
		if err := conn.Close(); err != nil {
			e.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()

	// server start
	e.Logger.Fatal(e.Start(":" + conf.C.Server.Address))
}