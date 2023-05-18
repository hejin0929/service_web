package biz

import "github.com/go-kratos/kratos/v2/log"

// UsersBiz 用户相关逻辑
type UsersBiz struct {
	repo UsersRepo
	log  *log.Helper
}

type createUsersType struct {
}

// UsersRepo is a user repo.
type UsersRepo interface {
}
