package data

import (
	"base/internal/biz"
	"base/pkg"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

type districtRepo struct {
	data *Data
	log  *log.Helper
}

type LoginMessage struct {
	Token    string `json:"token"`
	District string `json:"district"`
}

// NewDistrictRepo .
func NewDistrictRepo(data *Data, logger log.Logger) biz.DistrictRepo {
	return &districtRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *districtRepo) Save(ctx context.Context, id string) (err error) {

	user, err := pkg.UseUserID(ctx)

	userid := strconv.Itoa(user)

	str, err := r.data.RedisCli.Get(ctx, userid).Result()

	message := new(LoginMessage)

	err = json.Unmarshal([]byte(str), &message)
	message.District = id

	bytes, err := json.Marshal(message)
	r.data.RedisCli.Set(ctx, userid, string(bytes), 0)
	return
}

//func (r *districtRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
//	return g, nil
//}
//
//func (r *districtRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
//	return nil, nil
//}
//
//func (r *districtRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
//	return nil, nil
//}
//
//func (r *districtRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
//	return nil, nil
//}
