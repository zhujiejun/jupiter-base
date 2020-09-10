package user

import (
	"jupiter-base/internal/app/model/db"
)

type Repository interface {
	Get(id int) (user db.User,err error)
}
