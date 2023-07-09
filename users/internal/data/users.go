package data

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"users/internal/biz"
)

type UsersRepo struct {
	data *Data
	log  *log.Helper
}

type LoginMessage struct {
	Token    string `json:"token"`
	District string `json:"district"`
}

func NewUsersRepo(data *Data, logger log.Logger) *UsersRepo {
	return &UsersRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UsersRepo) CreateUser(ctx context.Context, usersType *biz.Users) (err error) {
	if !r.data.Mysql.Migrator().HasTable("users") {
		_ = r.data.Mysql.AutoMigrate(usersType)
	}
	db := r.data.Mysql.Table("users").WithContext(ctx)

	var nums int64

	db.Where("account = ?", usersType.Account).Count(&nums)

	if nums > 0 {
		err = errors.New("账户已存在")
		return
	}
	err = db.Create(usersType).Error
	return
}

func (r *UsersRepo) GetUser(ctx context.Context, name string, value interface{}) (user *biz.Users, err error) {
	user = new(biz.Users)
	r.data.Mysql.Model(biz.Users{}).Where(name+" = ?", value).Find(&user).WithContext(ctx)
	return
}

func (r *UsersRepo) patchUser() (user *biz.Users, err error) {
	return
}

func (r *UsersRepo) deleteUser() (err error) {
	return
}

func (r *UsersRepo) SaveToken(ctx context.Context, id string, token string) (err error) {
	message := new(LoginMessage)
	message.Token = token
	bytes, err := json.Marshal(message)
	err = r.data.RedisCli.Set(ctx, id, string(bytes), 0).Err()
	return
}

func (r *UsersRepo) GetToken(ctx context.Context, id string) (token string, err error) {
	message := new(LoginMessage)
	var bytes string
	bytes, err = r.data.RedisCli.Get(ctx, id).Result()
	err = json.Unmarshal([]byte(bytes), &message)
	token = message.Token

	return
}
