package data

import (
	"gorm.io/gorm"
	"users/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	user Users
}

type Users struct {
	gorm.Model
	UUID      string `json:"uuid"`
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
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
