package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/one-connect/biz-chance-api/domain/repository"
	"github.com/one-connect/biz-chance-api/domain/service"
	"github.com/one-connect/biz-chance-api/infrastructure/api/handler"
	"github.com/one-connect/biz-chance-api/infrastructure/persistence/datastore"
	"github.com/one-connect/biz-chance-api/interfaces/controller"
	"github.com/one-connect/biz-chance-api/interfaces/presenter"
)

type interactor struct {
	conn *gorm.DB
}

type Iteractor interface {
	NewAppHandler() handler.AppHandler
}

func NewInteractor(conn *gorm.DB) Iteractor {
	return &interactor{conn}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return i.NewUserHandler()
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserController())
}

func (i *interactor) NewUserController() controller.UserController {
	return controller.NewUserController(i.NewUserService())
}

func (i *interactor) NewUserService() service.UserService {
	return service.NewUserService(i.NewUserRepository(), i.NewUserPresenter())
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return datastore.NewUserRepository(i.conn)
}

func (i *interactor) NewUserPresenter() presenter.UserPresenter {
	return presenter.NewUserPresenter()
}