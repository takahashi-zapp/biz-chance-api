package router

import (
	"github.com/gin-gonic/gin"
	"github.com/one-connect/biz-chance-api/infrastructure/api/handler"
)


func NewRouter(r *gin.Engine, handler handler.AppHandler) {
	r.POST("/users", handler.CreateUser)
	r.GET("/users", handler.CreateUser)
	r.GET("/users", handler.GetUsers)
}