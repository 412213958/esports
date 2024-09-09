package model

import "time"

type Users struct {
	Id        int32     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Money     string    `gorm:"column:money;default:0.00"`
	DeletedAt string    `gorm:"column:deleted_at;default:NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:NULL"`
}

func (u *Users) TableName() string {
	return "users"
}

type Predictions struct {
	Id        int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	UserId    int32  `gorm:"column:user_id;default:NULL"`
	Team      int32  `gorm:"column:team;default:NULL"`
	BetAmount string `gorm:"column:bet_amount;default:NULL"`
}

func (p *Predictions) TableName() string {
	return "predictions"
}
