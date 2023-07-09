package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	pb "users/api/auth/v1"
	"users/pkg"
)

type AuthUse struct {
	repo UsersRepo
	log  *log.Helper
	ID   uint
}

// NewAuthUse new a Greeter NewUsersUse.
func NewAuthUse(repo UsersRepo, logger log.Logger) *AuthUse {
	return &AuthUse{repo: repo, log: log.NewHelper(logger)}
}

func (r *AuthUse) LoginUsers(ctx context.Context, req *pb.LoginUsersRequest) (res *pb.LoginUsersReply, err error) {

	user, err := r.repo.GetUser(ctx, "account", req.Account)

	res = new(pb.LoginUsersReply)
	res.Data = new(pb.LoginData)

	if user.ID == 0 {
		res.Success = false
		res.Message = "暂无该用户信息"
		err = errors.New(500, "USER_NOT_FOUND", "account nonentity")
		return
	}

	if user.Password != req.Password {
		res.Success = false
		res.Message = "用户密码错误"
		return
	}

	res.Success = true
	res.Message = "登陆成功"
	res.Data.Uid = user.UID
	res.Data.Token, err = pkg.GenerateToken(strconv.Itoa(int(user.ID)), user.Account)
	err = r.repo.SaveToken(ctx, strconv.Itoa(int(user.ID)), res.Data.Token)
	return
}

func (r *AuthUse) AuthLogin(ctx context.Context, req *pb.AuthLoginRequest) (res *pb.AuthLoginReply, err error) {
	userID, err := pkg.UseUserID(ctx)
	resToken, err := r.repo.GetToken(ctx, strconv.Itoa(userID))
	if resToken != req.Token {
		return nil, fmt.Errorf("auth verification failed")
	}
	res = new(pb.AuthLoginReply)
	res.Success = true
	res.Message = "验证通过"
	return
}
