package model

type TeamMember struct {
	MemberID int    `json:"member_id" gorm:"primaryKey;autoIncrement"`
	TeamID   int    `json:"team_id"`
	UserID   int    `json:"user_id"`
	Position string `json:"position"`
}
