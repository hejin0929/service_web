package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"users/internal/biz"
)

type UsersRepo struct {
	data *Data
	log  *log.Helper
}

func NewUsersRepo(data *Data, logger log.Logger) *UsersRepo {
	return &UsersRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UsersRepo) CreateUser(usersType *biz.CreateUsersType) (err error) {
	return
}

//func (r *UsersRepo) getUser() (user Users, err error) {
//	return
//}
//
//func (r *UsersRepo) patchUser() (user Users, err error) {
//	return
//}
//
//func (r *UsersRepo) deleteUser() (err error) {
//	return
//}
