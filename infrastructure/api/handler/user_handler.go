package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/one-connect/biz-chance-api/domain/model"
	"github.com/one-connect/biz-chance-api/interfaces/controller"
	"net/http"
)

type userHandler struct {
	userController controller.UserController
}

type UserHandler interface {
	CreateUser(c *gin.Context)
	GetUsers(c *gin.Context)
}

func NewUserHandler(uc controller.UserController) UserHandler {
	return &userHandler{userController: uc}
}

func (uh *userHandler) CreateUser(c *gin.Context) {

	// リクエスト用のEntityを作成
	req := &model.User{}

	// bind
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	// validate
	//if err := c.Validate(req); err != nil {
	//	c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	//	return
	//}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := uh.userController.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "success")
}

func (uh *userHandler) GetUsers(c *gin.Context) {

	req := &model.User{}

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	u, err := uh.userController.GetUsers()
	if err != nil {
		// システム内のエラー
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, u)
}