package biz

import "gorm.io/gorm"

type CityData struct {
	gorm.Model
	Name         string  `json:"name,omitempty"`
	Type         string  `json:"type,omitempty"`
	X            int32   `json:"x,omitempty"`
	Y            int32   `json:"y,omitempty"`
	Boost        string  `json:"boost,omitempty"`
	CountryId    string  `json:"country_id"`
	MasterName   int32   `json:"master_name,omitempty"`
	Production   string  `json:"production,omitempty"`
	Revenue      string  `json:"revenue,omitempty"`
	DefendMax    int32   `json:"defend_max,omitempty"`
	DefendTroops []int32 `json:"defend_troops,omitempty"`
	ManorMax     int32   `json:"manor_max,omitempty"`
	Manors       []int32 `json:"manors,omitempty"`
	Announcement string  `json:"announcement,omitempty"`
	CityWall     int32   `json:"city_wall,omitempty"`
	CityWay      int32   `json:"city_way,omitempty"`
}
