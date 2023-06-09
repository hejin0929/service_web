package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"template/internal/biz"
)

type templateRepo struct {
	data *Data
	log  *log.Helper
}

// NewTemplateRepo .
func NewTemplateRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
