package model

// RaceCalendar 比赛信息表
type RaceCalendar struct {
	Id         int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Title      string `gorm:"column:title;default:NULL"`
	Session    string `gorm:"column:session;default:NULL"`
	Date       string `gorm:"column:date;default:NULL"`
	Time       string `gorm:"column:time;default:NULL"`
	Team1Name  string `gorm:"column:team1_name;default:NULL"`
	Team1Logo  string `gorm:"column:team1_logo;default:NULL"`
	Team2Name  string `gorm:"column:team2_name;default:NULL"`
	Team2Logo  string `gorm:"column:team2_logo;default:NULL"`
	Team1Score int32  `gorm:"column:team1_score;default:NULL"`
	Team2Score int32  `gorm:"column:team2_score;default:NULL"`
	Method     string `gorm:"column:method;default:NULL"`
	Schedule   string `gorm:"column:schedule;default:NULL"`
}

func (r *RaceCalendar) TableName() string {
	return "race_calendar"
}

type StageContainer struct {
	Id       int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Ranking  int32  `gorm:"column:ranking;default:NULL"`
	TeamName string `gorm:"column:team_name;default:NULL"`
	TeamLogo string `gorm:"column:team_logo;default:NULL"`
	Wins     int32  `gorm:"column:wins;default:NULL"`
	Losses   int32  `gorm:"column:losses;default:NULL"`
	Points   int32  `gorm:"column:points;default:NULL"`
	TypeSF   string `gorm:"column:typeSF;default:NULL"`
	Typestr  string `gorm:"column:typestr;default:NULL"`
}

func (s *StageContainer) TableName() string {
	return "stage_container"
}

type SetMatch struct {
	Id       int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Ranking  int32  `gorm:"column:ranking;default:NULL"`
	TeamName string `gorm:"column:team_name;default:NULL"`
	TeamLogo string `gorm:"column:team_logo;default:NULL"`
	Wins     int32  `gorm:"column:wins;default:NULL"`
	Losses   int32  `gorm:"column:losses;default:NULL"`
	Points   int32  `gorm:"column:points;default:NULL"`
	Types    string `gorm:"column:types;default:NULL"`
	Typestr  string `gorm:"column:typestr;default:NULL"`
}

func (s *SetMatch) TableName() string {
	return "set_match"
}
