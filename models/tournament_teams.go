package model

import "database/sql"

type WeiboData struct {
	Id        int32          `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	User      string         `gorm:"column:user;default:NULL"`
	Content   sql.NullString `gorm:"column:content"`
	Time      string         `gorm:"column:time;default:NULL"`
	ImageSrc  string         `gorm:"column:image_src;default:NULL"`
	AvatarSrc string         `gorm:"column:avatar_src;default:NULL"`
	Lp        string         `gorm:"column:lp;default:NULL"`
	Lp1       string         `gorm:"column:lp1;default:NULL"`
}

func (w *WeiboData) TableName() string {
	return "weibo_data"
}

type Tournaments struct {
	Id       int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Title    string `gorm:"column:title;default:NULL"`
	ImageSrc string `gorm:"column:image_src;default:NULL"`
	TypeS    string `gorm:"column:type_s;default:NULL"`
	Lp       string `gorm:"column:lp;default:NULL"`
}

func (t *Tournaments) TableName() string {
	return "tournaments"
}
