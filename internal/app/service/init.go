package service

import (
	"jupiter-base/internal/app/model"
	"jupiter-base/internal/app/service/user"
	"jupiter-base/internal/app/service/user/impl"
)

var (
	UserRepository user.Repository
)
//Init instantiate the service
func Init()  {
	UserRepository = impl.NewMysqlImpl(model.MysqlHandler)
}