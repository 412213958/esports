package team

import (
	model "esports/models"
	"esports/utils"
)

func getLegacyList() ([]model.Teamss, error) {
	var Teamss []model.Teamss
	if err := utils.DB.Find(&Teamss).Error; err != nil {
		return nil, err
	}
	return Teamss, nil
}

//type TeamData struct {
//	Data []TeamDataS
//}

type TeamDataS struct {
	Title string `gorm:"column:title;default:NULL"`
	Data  string `gorm:"column:data;default:NULL"`
}

type TeamInfo1 struct {
	Id       int32       `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Squadron string      `gorm:"column:squadron;default:NULL"`
	WinRate  string      `gorm:"column:win_rate;default:NULL"`
	WinDesc  string      `gorm:"column:win_desc;default:NULL"`
	Kda      string      `gorm:"column:kda;default:NULL"`
	KdaDesc  string      `gorm:"column:kda_desc;default:NULL"`
	Data     []TeamDataS `gorm:"embedded"`
}

func getPlayerAndTeamDataByTeam(team string) ([]model.Players, TeamInfo1, error) {
	var players []model.Players
	var teamInfo model.TeamInfo
	var teamInfo1 TeamInfo1

	err := utils.DB.Model(&model.Players{}).
		Where("team =?", team).
		Find(&players).Error
	if err != nil {
		return nil, teamInfo1, err
	}

	err = utils.DB.Where("squadron =?", team).
		Find(&teamInfo).Error
	if err != nil {
		return nil, teamInfo1, err
	}

	teamInfo1.Id = teamInfo.Id
	teamInfo1.Squadron = teamInfo.Squadron
	teamInfo1.WinRate = teamInfo.WinRate
	teamInfo1.Kda = teamInfo.Kda
	teamInfo1.WinDesc = teamInfo.WinDesc
	teamInfo1.KdaDesc = teamInfo.KdaDesc

	dataS := []TeamDataS{
		{Title: "分均输出", Data: teamInfo.OutputPerMinute},
		{Title: "分均经济", Data: teamInfo.EconomyPerMinute},
		{Title: "分均补刀", Data: teamInfo.CsPerMinute},
		{Title: "小龙控制率", Data: teamInfo.DragonControlRate},
		{Title: "场均击杀", Data: teamInfo.AverageKills},
		{Title: "场均助攻", Data: teamInfo.AverageAssists},
		{Title: "场均死亡", Data: teamInfo.AverageDeaths},
		{Title: "大龙控制率", Data: teamInfo.BaronControlRate},
	}

	teamInfo1.Data = dataS

	return players, teamInfo1, nil
}

type TeamInfo2 struct {
	Id                int32           `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Squadron          string          `gorm:"column:squadron;default:NULL"`
	TeamLogo          string          `gorm:"column:team_logo;default:NULL"`
	WinRate           string          `gorm:"column:win_rate;default:NULL"`
	WinDesc           string          `gorm:"column:win_desc;default:NULL"`
	Kda               string          `gorm:"column:kda;default:NULL"`
	KdaDesc           string          `gorm:"column:kda_desc;default:NULL"`
	OutputPerMinute   string          `gorm:"column:output_per_minute;default:NULL"`
	EconomyPerMinute  string          `gorm:"column:economy_per_minute;default:NULL"`
	CsPerMinute       string          `gorm:"column:cs_per_minute;default:NULL"`
	DragonControlRate string          `gorm:"column:dragon_control_rate;default:NULL"`
	AverageKills      string          `gorm:"column:average_kills;default:NULL"`
	AverageAssists    string          `gorm:"column:average_assists;default:NULL"`
	AverageDeaths     string          `gorm:"column:average_deaths;default:NULL"`
	BaronControlRate  string          `gorm:"column:baron_control_rate;default:NULL"`
	Members           []model.Players `gorm:"column:members;default:NULL"`
}

func getAllPlayerAndTeamDataByTeam() ([]TeamInfo2, error) {

	var Teamss []model.Teamss
	if err := utils.DB.Find(&Teamss).Error; err != nil {
		return nil, err
	}

	var teamInfos []model.TeamInfo
	var teamInfos2 []TeamInfo2
	err := utils.DB.Find(&teamInfos).Error
	if err != nil {
		return nil, err
	}

	var matchingPlayerss [][]model.Players
	for _, teamInfo := range teamInfos {
		var matchingPlayers []model.Players
		err := utils.DB.Where("team =?", teamInfo.Squadron).Find(&matchingPlayers).Error
		if err != nil {
			return nil, err
		}
		matchingPlayerss = append(matchingPlayerss, matchingPlayers)
	}

	for index, teamInfo := range teamInfos {
		var teamInfo2 TeamInfo2
		teamInfo2.Id = teamInfo.Id
		teamInfo2.Squadron = teamInfo.Squadron
		teamInfo2.TeamLogo = Teamss[index].TeamLogo
		teamInfo2.WinRate = teamInfo.WinRate
		teamInfo2.Kda = teamInfo.Kda
		teamInfo2.WinDesc = teamInfo.WinDesc
		teamInfo2.KdaDesc = teamInfo.KdaDesc
		teamInfo2.OutputPerMinute = teamInfo.OutputPerMinute
		teamInfo2.EconomyPerMinute = teamInfo.EconomyPerMinute
		teamInfo2.CsPerMinute = teamInfo.CsPerMinute
		teamInfo2.DragonControlRate = teamInfo.DragonControlRate
		teamInfo2.AverageKills = teamInfo.AverageKills
		teamInfo2.AverageAssists = teamInfo.AverageAssists
		teamInfo2.AverageDeaths = teamInfo.AverageDeaths
		teamInfo2.BaronControlRate = teamInfo.BaronControlRate
		teamInfo2.Members = matchingPlayerss[index]
		teamInfos2 = append(teamInfos2, teamInfo2)
	}

	return teamInfos2, nil
}
