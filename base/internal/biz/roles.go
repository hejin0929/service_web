package biz

import (
	pb "base/api/roles/v1"
	"base/pkg"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strconv"
)

type RolesRepo interface {
}

// Role sql db roles table
type Role struct {
	gorm.Model
	AchievementsId    int    `json:"achievements_id"`    // 成就
	ActionDesignation string `json:"action_designation"` // 激活称号，玩家激活称号
	Assets            []int  `json:"assets"`             // 封地，玩家封地，资产
	Avatar            string `json:"avatar"`             // 头像
	CityNums          int    `json:"city_nums"`          // 城池数量
	Consumption       int    `json:"consumption"`        // 会员值，用户消费了多少充值，比如在会员6 - 会员7之间
	Contribution      int    `json:"contribution"`       // 贡献
	Country           string `json:"country"`            // 国家，玩家所在国家
	Exp               int    `json:"exp"`                // 经验值
	Grade             int    `json:"grade"`              // 等级
	Intro             string `json:"intro"`              // 简介
	Member            int    `json:"member"`             // 会员
	MilitaryExploit   int    `json:"military_exploit"`   // 军功，通过战场获得
	Name              string `json:"name"`               // 角色名
	Points            int    `json:"points"`             // 消费积分
	Status            string `json:"status"`             // 状态
	TeamName          string `json:"team_name"`          // 战队名称
	Uid               string `json:"uid"`                // 角色ID
	UserId            int    `json:"userId"`             // 玩家ID
	DistrictId        int    `json:"district_id"`        // 区服ID
	Sex               string `json:"sex"`                // 性别
}

type RolesUse struct {
	repo RolesRepo
	log  *log.Helper
}

func NewRolesUse(repo RolesRepo, logger log.Logger) *RolesUse {
	return &RolesUse{repo: repo, log: log.NewHelper(logger)}
}

func (r *RolesUse) CreateRoles(ctx context.Context, req *pb.CreateRolesRequest) (res *pb.CreateRolesReply, err error) {
	userRole := new(Role)
	userRole.UserId, err = pkg.UseUserID(ctx)
	userRole.DistrictId, err = strconv.Atoi(req.DistrictId)
	userRole.Name = req.Name
	userRole.Avatar = req.Avatar
	userRole.Country = req.Country
	userRole.Sex = req.Sex

	return
}
