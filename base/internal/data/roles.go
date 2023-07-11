package data

import (
	"base/internal/biz"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type rolesRepo struct {
	data *Data
	log  *log.Helper
}

func NewRolesRepo(data *Data, logger log.Logger) biz.RolesRepo {
	return rolesRepo{data: data, log: log.NewHelper(logger)}
}

func (r *rolesRepo) Save(ctx context.Context, role *biz.Role) (err error) {
	if !r.data.Mysql.Migrator().HasTable(role) {
		_ = r.data.Mysql.AutoMigrate(role)
	}

	var num int64

	err = r.data.Mysql.Model(role).Where("name = ? ", role.Name).Count(&num).WithContext(ctx).Error

	if num != 0 {
		return fmt.Errorf("角色名已存在")
	}
	r.data.Mysql.Create(role)
	return
}

func (r *rolesRepo) Search(ctx context.Context, key string, value interface{}) (res *biz.Role, err error) {
	res = new(biz.Role)
	err = r.data.Mysql.Where(key+" = ? ", value).Find(res).WithContext(ctx).Error
	return
}

func (r *rolesRepo) Updates(ctx context.Context, id string, value *biz.Role) error {
	err := r.data.Mysql.Where("id = ? ", id).Updates(value).Error
	return err
}
