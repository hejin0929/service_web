package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"world/internal/biz"
)

type countryRepo struct {
	data *Data
	log  *log.Helper
}

// NewCountryRepo .
func NewCountryRepo(data *Data, logger log.Logger) biz.CountryRepo {
	return &countryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *countryRepo) Save(ctx context.Context, country *biz.Country) (res *biz.Country, err error) {
	if !r.data.Mysql.Migrator().HasTable(country) {
		_ = r.data.Mysql.AutoMigrate(country)
	}
	var num int64
	r.data.Mysql.Model(country).Where("name = ? ", country.Name).Count(&num)
	if num != 0 {
		return nil, fmt.Errorf("国家已存在! 建国失败。")
	}
	err = r.data.Mysql.Create(country).WithContext(ctx).Error
	res = country
	return
}

func (r *countryRepo) Update(ctx context.Context, g *biz.Country) (*biz.Country, error) {
	return g, nil
}

func (r *countryRepo) FindByID(context.Context, int64) (*biz.Country, error) {
	return nil, nil
}

func (r *countryRepo) ListByHello(context.Context, string) ([]*biz.Country, error) {
	return nil, nil
}

func (r *countryRepo) ListAll(ctx context.Context, current int32, size int32) (res *biz.ListsCountry, err error) {
	countries := make([]biz.Country, size)
	err = r.data.Mysql.Limit(int(current * size)).Offset(int((current - 1) * size)).Find(&countries).WithContext(ctx).Error

	res = new(biz.ListsCountry)

	res.Lists = countries

	res.Page.PageSize = size

	res.Page.Current = current

	var total int64
	err = r.data.Mysql.Model(biz.Country{}).Count(&total).Error
	res.Page.Total = int32(total)
	return
}
