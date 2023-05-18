package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"users/internal/biz"
)

type usersRepo struct {
	data *Data
	log  *log.Helper
}

func NewUsersRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *usersRepo) createUser() (err error) {
	return
}

func (r *usersRepo) getUser() (user Users, err error) {
	return
}

func (r *usersRepo) patchUser() (user Users, err error) {
	return
}

func (r *usersRepo) deleteUser() (err error) {
	return
}
