package model

type TournamentAward struct {
	AwardID      int    `json:"award_id" gorm:"primaryKey;autoIncrement"`
	TournamentID int    `json:"tournament_id"`
	AwardName    string `json:"award_name" gorm:"not null"`
	TeamID       int    `json:"team_id"`
}

type Table1 struct {
	V                string `gorm:"column:v;default:NULL"`
	BetKey           string `gorm:"column:bet_key;default:NULL"`
	NewDay           string `gorm:"column:new_day;default:NULL"`
	DateTxt          string `gorm:"column:date_txt;default:NULL"`
	ViewType         int32  `gorm:"column:view_type;default:NULL"`
	TournamentId     string `gorm:"column:tournament_id;default:NULL"`
	TournamentName   string `gorm:"column:tournament_name;default:NULL"`
	TournamentEnName string `gorm:"column:tournament_en_name;default:NULL"`
	TournamentImage  string `gorm:"column:tournament_image;default:NULL"`
	RoundName        string `gorm:"column:round_name;default:NULL"`
	RoundSonName     string `gorm:"column:round_son_name;default:NULL"`
	MatchId          string `gorm:"column:match_id;default:NULL"`
	TeamImageThumbA  string `gorm:"column:team_image_thumb_a;default:NULL"`
	TeamImageThumbB  string `gorm:"column:team_image_thumb_b;default:NULL"`
	MatchDate        string `gorm:"column:match_date;default:NULL"`
	MatchDate1       string `gorm:"column:match_date1;default:NULL"`
	MatchTeamA       string `gorm:"column:match_team_a;default:NULL"`
	MatchTeamB       string `gorm:"column:match_team_b;default:NULL"`
	GameCount        string `gorm:"column:game_count;default:NULL"`
	StartId          string `gorm:"column:start_id;default:NULL"`
}

func (t *Table1) TableName() string {
	return "table1"
}

type Table2 struct {
	BetId         string `gorm:"column:bet_id;default:NULL"`
	Title         string `gorm:"column:title;default:NULL"`
	CategoryName  string `gorm:"column:category_name;default:NULL"`
	BetEndTime    string `gorm:"column:bet_end_time;default:NULL"`
	BetEndTimeTxt string `gorm:"column:bet_end_time_txt;default:NULL"`
	Status        string `gorm:"column:status;default:NULL"`
	StatusTxt     string `gorm:"column:status_txt;default:NULL"`
	TotalPrice    string `gorm:"column:total_price;default:NULL"`
	PeopleNum     string `gorm:"column:people_num;default:NULL"`
	ResultItemId  string `gorm:"column:result_item_id;default:NULL"`
	DynamicId     int32  `gorm:"column:dynamic_id;default:NULL"`
}

func (t *Table2) TableName() string {
	return "table2"
}

type Table3 struct {
	BetItemId     string `gorm:"column:bet_item_id;default:NULL"`
	InitPrice     string `gorm:"column:init_price;default:NULL"`
	Price         string `gorm:"column:price;default:NULL"`
	MemberMaxBet  string `gorm:"column:member_max_bet;default:NULL"`
	ItemName      string `gorm:"column:item_name;default:NULL"`
	ItemNameEn    string `gorm:"column:item_name_en;default:NULL"`
	ItemNameTw    string `gorm:"column:item_name_tw;default:NULL"`
	WinRate       string `gorm:"column:win_rate;default:NULL"`
	Odds          string `gorm:"column:odds;default:NULL"`
	TeamAWin      string `gorm:"column:team_a_win;default:NULL"`
	TeamBWin      string `gorm:"column:team_b_win;default:NULL"`
	MatchStatus   string `gorm:"column:match_status;default:NULL"`
	MatchBetCount string `gorm:"column:match_bet_count;default:NULL"`
}

func (t *Table3) TableName() string {
	return "table3"
}
