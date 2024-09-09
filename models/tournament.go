package model

import (
	"database/sql"
	"time"
)

type Tournament struct {
	TournamentID   int       `json:"tournament_id" gorm:"primaryKey;autoIncrement"`
	TournamentName string    `json:"tournament_name" gorm:"not null"`
	StartDate      time.Time `json:"start_date" gorm:"not null"`
	EndDate        time.Time `json:"end_date" gorm:"not null"`
	Description    string    `json:"description" gorm:"text"`
	CreatedAt      time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Articles struct {
	Id          int32          `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Title       string         `gorm:"column:title;default:NULL"`
	Description sql.NullString `gorm:"column:description"`
	ImageLink   string         `gorm:"column:image_link;default:NULL"`
	Lp          string         `gorm:"column:lp;default:NULL"`
}

func (a *Articles) TableName() string {
	return "articles"
}

type Articles1 struct {
	Id          int32          `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	ImageLink   string         `gorm:"column:image_link;default:NULL"`
	Title       string         `gorm:"column:title;default:NULL"`
	Description sql.NullString `gorm:"column:description"`
	Author      string         `gorm:"column:author;default:NULL"`
	TypeS       string         `gorm:"column:type_s;default:NULL"`
	Lp          string         `gorm:"column:lp;default:NULL"`
}

func (a *Articles1) TableName() string {
	return "articles_1"
}

type Comments struct {
	Id           int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	UserAvatar   string `gorm:"column:user_avatar;default:NULL"`
	UserName     string `gorm:"column:user_name;default:NULL"`
	PostTime     string `gorm:"column:post_time;default:NULL"`
	CommentTitle string `gorm:"column:comment_title;default:NULL"`
	CommentReply string `gorm:"column:comment_reply;default:NULL"`
	Lp           string `gorm:"column:lp;default:NULL"`
}

func (c *Comments) TableName() string {
	return "comments"
}
