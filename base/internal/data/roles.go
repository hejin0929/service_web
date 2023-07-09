package data

import (
	"base/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type rolesRepo struct {
	data *Data
	log  *log.Helper
}

func NewRolesRepo(data *Data, logger log.Logger) biz.RolesRepo {
	return rolesRepo{data: data, log: log.NewHelper(logger)}
}

func (r *rolesRepo) Save(ctx context.Context, role *biz.Role) {
	if !r.data.Mysql.Migrator().HasTable("roles") {
		_ = r.data.Mysql.AutoMigrate(role)
	}

}
