package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	pb "users/api/v1"
)

// UsersBiz 用户相关逻辑
type UsersBiz struct {
	repo UsersRepo
	log  *log.Helper
}

type CreateUsersType struct {
	Account   string `json:"account"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
	Sex       int    `json:"sex"`
	Email     string `json:"email"`
	School    string `json:"school"`
	Address   string `json:"address"`
	UUID      string `json:"uuid"`
	IP        string `json:"ip"`
	Equipment string `json:"equipment"`
}

// UsersRepo is a user repo.
type UsersRepo interface {
	CreateUser(usersType *CreateUsersType) error
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

	user := CreateUsersType{}

	bytes, err := json.Marshal(req)

	err = json.Unmarshal(bytes, &user)

	if tp, ok := transport.FromServerContext(ctx); ok {
		println(fmt.Sprintf("remote_addr: %v", tp.RequestHeader().Get("RemoteAddr")))
	}

	return
}
