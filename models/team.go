package model

import (
	"time"
)

type Team struct {
	TeamID    int       `json:"team_id" gorm:"primaryKey;autoIncrement"`
	TeamName  string    `json:"team_name" gorm:"not null"`
	LogoURL   string    `json:"logo_url" gorm:""`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Teamss struct {
	Id       int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	TeamName string `gorm:"column:team_name;default:NULL"`
	TeamLogo string `gorm:"column:team_logo;default:NULL"`
}

func (t *Teamss) TableName() string {
	return "teamss"
}

type Players struct {
	Id       int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name     string `gorm:"column:name;default:NULL"`
	ImageUrl string `gorm:"column:image_url;default:NULL"`
	Team     string `gorm:"column:team;default:NULL"`
}

func (p *Players) TableName() string {
	return "players"
}

type TeamInfo struct {
	Id                int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Squadron          string `gorm:"column:squadron;default:NULL"`
	WinRate           string `gorm:"column:win_rate;default:NULL"`
	WinDesc           string `gorm:"column:win_desc;default:NULL"`
	Kda               string `gorm:"column:kda;default:NULL"`
	KdaDesc           string `gorm:"column:kda_desc;default:NULL"`
	OutputPerMinute   string `gorm:"column:output_per_minute;default:NULL"`
	EconomyPerMinute  string `gorm:"column:economy_per_minute;default:NULL"`
	CsPerMinute       string `gorm:"column:cs_per_minute;default:NULL"`
	DragonControlRate string `gorm:"column:dragon_control_rate;default:NULL"`
	AverageKills      string `gorm:"column:average_kills;default:NULL"`
	AverageAssists    string `gorm:"column:average_assists;default:NULL"`
	AverageDeaths     string `gorm:"column:average_deaths;default:NULL"`
	BaronControlRate  string `gorm:"column:baron_control_rate;default:NULL"`
}

func (t *TeamInfo) TableName() string {
	return "team_info"
}
