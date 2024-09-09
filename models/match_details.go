package model

type MatchDetail struct {
	MatchDetailID int `json:"match_detail_id" gorm:"primaryKey;autoIncrement"`
	MatchID       int `json:"match_id"`
	PlayerID      int `json:"player_id"`
	Kills         int `json:"kills"`
	Deaths        int `json:"deaths"`
	Assists       int `json:"assists"`
}

type RecommendCard struct {
	Id           int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	SeasonLogo   string `gorm:"column:season_logo;default:NULL"`
	SeasonTitle  string `gorm:"column:season_title;default:NULL"`
	SeasonTime   string `gorm:"column:season_time;default:NULL"`
	SeasonStatus string `gorm:"column:season_status;default:NULL"`
	GameLink     string `gorm:"column:game_link;default:NULL"`
}

func (r *RecommendCard) TableName() string {
	return "recommend_card"
}

type Contest struct {
	Id           int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	ContestTitle string `gorm:"column:contest_title;default:NULL"`
	Team1Name    string `gorm:"column:team1_name;default:NULL"`
	Team1Icon    string `gorm:"column:team1_icon;default:NULL"`
	Team2Name    string `gorm:"column:team2_name;default:NULL"`
	Team2Icon    string `gorm:"column:team2_icon;default:NULL"`
	ShadowImg    string `gorm:"column:shadow_img;default:NULL"`
}

func (c *Contest) TableName() string {
	return "contest"
}

type Video struct {
	Id            int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	VideoTitle    string `gorm:"column:video_title;default:NULL"`
	VideoCover    string `gorm:"column:video_cover;default:NULL"`
	VideoLink     string `gorm:"column:video_link;default:NULL"`
	VideoCount    string `gorm:"column:video_count;default:NULL"`
	VideoLike     string `gorm:"column:video_like;default:NULL"`
	VideoDuration string `gorm:"column:video_duration;NOT NULL"`
	Type          string `gorm:"column:type;default:NULL;comment:'1:英雄联盟，2:王者荣耀，3:第五人格，4:无畏契约，5:DOTA2'"`
}

func (v *Video) TableName() string {
	return "video"
}

type MainMatches struct {
	Id           int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	HomeTeamName string `gorm:"column:home_team_name;default:NULL"`
	AwayTeamName string `gorm:"column:away_team_name;default:NULL"`
	HomeTeamLogo string `gorm:"column:home_team_logo;default:NULL"`
	AwayTeamLogo string `gorm:"column:away_team_logo;default:NULL"`
	HomeScore    int32  `gorm:"column:home_score;default:NULL"`
	AwayScore    int32  `gorm:"column:away_score;default:NULL"`
	WinTeamId    int32  `gorm:"column:win_team_id;default:NULL"`
	GameStage    string `gorm:"column:game_stage;default:NULL"`
	StartTime    int32  `gorm:"column:start_time;default:NULL"`
	EndTime      int32  `gorm:"column:end_time;default:NULL"`
}

func (m *MainMatches) TableName() string {
	return "main_matches"
}

type GrandSubMatches struct {
	Id           int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	SubMatchId   int32  `gorm:"column:sub_match_id;default:NULL"`
	HomeTeamName string `gorm:"column:home_team_name;default:NULL"`
	AwayTeamName string `gorm:"column:away_team_name;default:NULL"`
	HomeTeamLogo string `gorm:"column:home_team_logo;default:NULL"`
	AwayTeamLogo string `gorm:"column:away_team_logo;default:NULL"`
	HomeScore    int32  `gorm:"column:home_score;default:NULL"`
	AwayScore    int32  `gorm:"column:away_score;default:NULL"`
	WinTeamId    int32  `gorm:"column:win_team_id;default:NULL"`
	GameStage    string `gorm:"column:game_stage;default:NULL"`
	StartTime    int32  `gorm:"column:start_time;default:NULL"`
	EndTime      int32  `gorm:"column:end_time;default:NULL"`
}

func (g *GrandSubMatches) TableName() string {
	return "grand_sub_matches"
}

type SubMatches struct {
	Id           int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	MainMatchId  int32  `gorm:"column:main_match_id;default:NULL"`
	HomeTeamName string `gorm:"column:home_team_name;default:NULL"`
	AwayTeamName string `gorm:"column:away_team_name;default:NULL"`
	HomeTeamLogo string `gorm:"column:home_team_logo;default:NULL"`
	AwayTeamLogo string `gorm:"column:away_team_logo;default:NULL"`
	HomeScore    int32  `gorm:"column:home_score;default:NULL"`
	AwayScore    int32  `gorm:"column:away_score;default:NULL"`
	WinTeamId    int32  `gorm:"column:win_team_id;default:NULL"`
	GameStage    string `gorm:"column:game_stage;default:NULL"`
	StartTime    int32  `gorm:"column:start_time;default:NULL"`
	EndTime      int32  `gorm:"column:end_time;default:NULL"`
}

func (s *SubMatches) TableName() string {
	return "sub_matches"
}
