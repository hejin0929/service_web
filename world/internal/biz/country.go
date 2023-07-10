package biz

import (
	"context"
	"gorm.io/gorm"
	"world/pkg"

	"github.com/go-kratos/kratos/v2/log"
	pb "world/api/country/v1"
)

type Country struct {
	gorm.Model
	Announcement    string `json:"announcement"`      // 公告
	Boost           string `json:"boost"`             // 加成
	Capital         string `json:"capital"`           // 国都
	CityNums        int32  `json:"city_nums"`         // 城池数量
	Declaration     string `json:"declaration"`       // 宣言
	Efficiency      string `json:"efficiency"`        // 效率
	IdleScienceNums int32  `json:"idle_science_nums"` // 闲置科技点
	Creator         string `json:"creator"`           // 国王
	MemberNums      int32  `json:"member_nums"`       // 成员数量
	Name            string `json:"name"`              // 名称
	NationalPower   string `json:"national_power"`    // 国力
	No              string `json:"no"`                // 国家号
	Rank            int32  `json:"rank"`              // 排名
	ScienceNums     int32  `json:"science_nums"`      // 科技等级
	SciencePoint    int32  `json:"science_point"`     // 积分点
	DistrictId      int    `json:"district_id"`       // id
}

// Boost 加成
type Boost struct {
	Attack   int32 `json:"attack"`   // 攻击
	Def      int32 `json:"def"`      // 防御
	Resource int32 `json:"resource"` // 资源
}

type ListsCountry struct {
	Lists []Country  `json:"lists"`
	Page  Pagination `json:"page"`
}

// Pagination 统一分页器
type Pagination struct {
	Total    int32 `json:"total"`
	Current  int32 `json:"current"`
	PageSize int32 `json:"page_size"`
}

// CountryRepo is a Greater repo.
type CountryRepo interface {
	Save(context.Context, *Country) (*Country, error)
	Update(context.Context, *Country) (*Country, error)
	FindByID(context.Context, int64) (*Country, error)
	ListByHello(context.Context, string) ([]*Country, error)
	ListAll(ctx context.Context, current int32, size int32) (*ListsCountry, error)
}

// CountryUse is a Greeter use.
type CountryUse struct {
	repo CountryRepo
	log  *log.Helper
}

// NewCountryUse new a Greeter use.
func NewCountryUse(repo CountryRepo, logger log.Logger) *CountryUse {
	return &CountryUse{repo: repo, log: log.NewHelper(logger)}
}

// CreateCountry creates a Country, and returns the new Country.
func (uc *CountryUse) CreateCountry(ctx context.Context, req *pb.CreateCountryRequest) (res *pb.CreateCountryReply, err error) {
	newCountry := new(Country)
	err = pkg.CopyTo(req, newCountry)

	newCountry.DistrictId = 1

	_, err = uc.repo.Save(ctx, newCountry)

	res = new(pb.CreateCountryReply)

	if err != nil && err.Error() == "国家已存在! 建国失败。" {
		res.Message = err.Error()
		res.Success = false
		err = nil
		return
	}

	if err == nil {
		res.Message = "建国成功"
		res.Success = true
	}
	return
}

func (uc *CountryUse) GetCountry() {

}

func (uc *CountryUse) GetCountryLists(ctx context.Context, req *pb.ListCountryRequest) (res *pb.ListCountryReply, err error) {

	data, err := uc.repo.ListAll(ctx, req.Current, req.PageSize)

	if err != nil {
		return
	}

	res = new(pb.ListCountryReply)
	res.Success = true
	res.Message = "获取成功"

	res.Data = new(pb.ListsData)

	res.Data.Page = new(pb.Pagination)

	res.Data.Page.PageSize = data.Page.PageSize
	res.Data.Page.Current = data.Page.Current
	res.Data.Page.Total = data.Page.Total

	res.Data.Lists = make([]*pb.CountryListsData, len(data.Lists))

	for i, list := range data.Lists {

		item := new(*pb.CountryListsData)

		err = pkg.CopyTo(list, item)
		res.Data.Lists[i] = *item
	}

	return
}
