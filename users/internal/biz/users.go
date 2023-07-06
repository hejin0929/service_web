package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/uuid"
	"gorm.io/gorm"
	pb "users/api/v1"
	"users/pkg"
)

// UsersBiz 用户相关逻辑
type UsersBiz struct {
	repo UsersRepo
	log  *log.Helper
}

type Users struct {
	gorm.Model
	UID       string `json:"uid"`
	IP        string `json:"ip"`
	Account   string `json:"account"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Equipment int    `json:"equipment"`
	Status    int    `json:"status"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Sex       int    `json:"sex"`
	Address   string `json:"address"`
	School    string `json:"school"`
	CardID    string `json:"card_id"`
	CardName  string `json:"cardName"`
}

// UsersRepo is a user repo.
type UsersRepo interface {
	CreateUser(ctx context.Context, usersType *Users) error
	GetUser(ctx context.Context, name string, value interface{}) (*Users, error)
}

type UsersUse struct {
	repo UsersRepo
	log  *log.Helper
}

// NewUsersUse new a Greeter NewUsersUse.
func NewUsersUse(repo UsersRepo, logger log.Logger) *UsersUse {
	return &UsersUse{repo: repo, log: log.NewHelper(logger)}
}

func (r *UsersUse) CreateUsers(ctx context.Context, req *pb.CreateUsersRequest) (res *pb.CreateUsersReply, err error) {

	user := &Users{}

	err = pkg.CopyTo(req, user)

	user.UID = uuid.New().String()

	if tp, ok := transport.FromServerContext(ctx); ok {
		println(fmt.Sprintf("remote_addr: %v", tp.RequestHeader().Get("RemoteAddr")))
	}
	err = r.repo.CreateUser(ctx, user)

	res = new(pb.CreateUsersReply)

	if err != nil {
		res.Success = false
		res.Message = err.Error()
		return
	}
	res.Message = "添加成功"
	res.Success = true
	return
}

func (r *UsersUse) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (res *pb.GetUsersReply, err error) {

	user, err := r.repo.GetUser(ctx, "uid", req.Uuid)

	res = new(pb.GetUsersReply)

	if user.ID == 0 {
		res.Success = false
		res.Message = "获取用户信息失败"

		return
	}

	return
}
