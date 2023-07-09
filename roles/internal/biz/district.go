package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "roles/api/district/v1"
	"time"
)

//var (
//	// ErrUserNotFound is user not found.
//	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
//)

// Greeter is a Greeter model.
//type Greeter struct {
//	Hello string
//}

// DistrictRepo is a Greater repo.
type DistrictRepo interface {
	Save(ctx context.Context, id string) error
	//Save(context.Context, *Greeter) (*Greeter, error)
	//Update(context.Context, *Greeter) (*Greeter, error)
	//FindByID(context.Context, int64) (*Greeter, error)
	//ListByHello(context.Context, string) ([]*Greeter, error)
	//ListAll(context.Context) ([]*Greeter, error)
}

// DistrictUse is a Greeter use.
type DistrictUse struct {
	repo DistrictRepo
	log  *log.Helper
}

// NewDistrictUse new a Greeter use.
func NewDistrictUse(repo DistrictRepo, logger log.Logger) *DistrictUse {
	return &DistrictUse{repo: repo, log: log.NewHelper(logger)}
}

// GetDistrict creates a district, and returns the new districtLists.
func (d *DistrictUse) GetDistrict(ctx context.Context, req *pb.GetDistrictRequest) (res *pb.GetDistrictReply, err error) {

	res = new(pb.GetDistrictReply)

	res.Message = "获取成功"

	res.Success = true
	// 获取当前时间
	now := time.Now()
	// 设置中国时区
	location, err := time.LoadLocation("Asia/Shanghai")
	res.Lists = append(res.Lists, &pb.DistrictData{
		Name:      "测试1区",
		Status:    "正常",
		OpenTime:  now.In(location).Format("2006-01-02 15:04:05"),
		CloseTime: "",
		Id:        "1",
		Type:      "普通区",
	})

	return
}

func (d *DistrictUse) SetDistrict(ctx context.Context, req *pb.UpdateDistrictRequest) (res *pb.UpdateDistrictReply, err error) {

	res = new(pb.UpdateDistrictReply)

	err = d.repo.Save(ctx, req.Id)

	if err == nil {
		res.Success = true
		res.Message = "进入成功"
	}

	return
}
