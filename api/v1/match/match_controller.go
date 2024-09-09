package match

import (
	"errors"
	model "esports/models"
	"esports/utils"
	"fmt"
	"sort"
)

func getRaceCalendarList() ([]*model.RaceCalendar, error) {
	data := make([]*model.RaceCalendar, 0)
	err := utils.DB.Group("title, session, date, time, team1_name, team1_logo, team2_name, team2_logo, team1_score, team2_score, method, schedule").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getRaceCalendarListByTitle(title string) ([]model.RaceCalendar, error) {
	var raceCalendarList []model.RaceCalendar

	err := utils.DB.Where("title = ?", title).Find(&raceCalendarList).Error
	if err != nil {
		return nil, err
	}

	return raceCalendarList, nil
}

func getRecommendCardList() ([]model.RecommendCard, error) {
	var recommendCards []model.RecommendCard
	if err := utils.DB.Find(&recommendCards).Error; err != nil {
		return nil, err
	}
	return recommendCards, nil
}

func getContestList() ([]model.Contest, error) {
	var contests []model.Contest
	if err := utils.DB.Find(&contests).Error; err != nil {
		return nil, err
	}
	return contests, nil
}

func getVideoListByTypeAndNum(videoType string, pagesize int, num int) ([]model.Video, error) {
	var videoList []model.Video
	err := errors.New("")
	if pagesize == 0 && num == 0 {
		err = utils.DB.Where("type =? ", videoType).
			Find(&videoList).Error
	} else {
		err = utils.DB.Where("type =? ", videoType).
			Limit(pagesize).
			Offset(num).
			Find(&videoList).Error
	}
	return videoList, err
}

func getAllData() ([]map[string]interface{}, error) {

	var recommendCards []model.RecommendCard
	var contests []model.Contest
	var videos []model.Video
	Backgroundcolor := []string{"linear-gradient( to top, RGB(28, 128, 238) 0%, RGB(28, 128, 238) 50%, RGBA(28, 128, 238, 0.9) 100% )",
		"linear-gradient( to top, RGB(177, 151, 84) 50%, RGB(177, 151, 84) 50%, RGBA(177, 151, 84, 0.9) 100% )",
		"linear-gradient( to top, RGB(0, 0, 0) 50%, RGB(0, 0, 0) 50%, RGBA(0, 0, 0, 0.8) 100% )",
		"linear-gradient( to top, RGB(0, 0, 0) 50%, RGB(0, 0, 0) 50%, RGBA(0, 0, 0, 0.8) 100% )",
		"linear-gradient( to top, RGB(71, 21, 13) 50%, RGB(71, 21, 13) 50%, RGBA(71, 21, 13, 0.9) 100% )",
	}

	// Fetch all RecommendCards
	if err := utils.DB.Find(&recommendCards).Error; err != nil {
		return nil, err
	}

	// Fetch all Contests
	if err := utils.DB.Find(&contests).Error; err != nil {
		return nil, err
	}

	// Fetch all Videos
	if err := utils.DB.Find(&videos).Error; err != nil {
		return nil, err
	}

	var allData []map[string]interface{}
	contestIndex := 0
	videoIndex := 0

	for _, recommendCard := range recommendCards {
		dataGroup := map[string]interface{}{
			"recommendCard": recommendCard,
			"contests":      []model.Contest{},
			"videos":        []model.Video{},
		}

		// Add two Contests to the group
		for i := 0; i < 2 && contestIndex < len(contests); i++ {
			dataGroup["contests"] = append(dataGroup["contests"].([]model.Contest), contests[contestIndex])
			contestIndex++
		}

		// Add six Videos to the group
		for i := 0; i < 6 && videoIndex < len(videos); i++ {
			dataGroup["videos"] = append(dataGroup["videos"].([]model.Video), videos[videoIndex])
			videoIndex++
		}

		dataGroup["Backgroundcolor"] = Backgroundcolor[0]
		allData = append(allData, dataGroup)

		// 更新 Backgroundcolor 列表
		Backgroundcolor = Backgroundcolor[1:]
	}

	return allData, nil
}

func getRaceCalendarListWF(title string, ses string, typeSF int) ([]model.RaceCalendar, []model.StageContainer, error) {
	var raceCalendarList []model.RaceCalendar
	var stageContainerList []model.StageContainer

	err := utils.DB.Model(&model.RaceCalendar{}).
		Where("title = ? AND session LIKE ?", title, ses+"%").
		Joins("LEFT JOIN stage_container AS s1 ON race_calendar.team1_name = s1.team_name").
		Joins("LEFT JOIN stage_container AS s2 ON race_calendar.team2_name = s2.team_name").
		Where("s1.typeSF = ?", typeSF).
		Group("title, session, race_calendar.team1_name, race_calendar.team2_name"). // 对非 id 字段进行分组去重
		Find(&raceCalendarList).Error
	if err != nil {
		return nil, nil, err
	}

	err = utils.DB.Where("typeSF = ?", typeSF).Find(&stageContainerList).Error
	if err != nil {
		return nil, nil, err
	}

	for i := range stageContainerList {
		switch stageContainerList[i].TypeSF {
		case "1":
			stageContainerList[i].TypeSF = "登峰组"
			stageContainerList[i].Typestr = "Win"
		case "-1":
			stageContainerList[i].TypeSF = "涅槃组"
			stageContainerList[i].Typestr = "Lose"
		default:
			return nil, nil, errors.New("无效的 typeSF 值")
		}
	}

	for i := range raceCalendarList {
		if typeSF == 1 {
			raceCalendarList[i].Method = "Win"
			raceCalendarList[i].Schedule = "登峰组赛程"
		} else if typeSF == -1 {
			raceCalendarList[i].Method = "Lose"
			raceCalendarList[i].Schedule = "涅槃组赛程"
		}

	}

	return raceCalendarList, stageContainerList, nil
}

func getRaceCalendarListByABCD(title string, ses string, typeSF string) ([]model.RaceCalendar, []model.SetMatch, error) {
	var raceCalendarList []model.RaceCalendar
	var setMatchList []model.SetMatch

	err := utils.DB.Model(&model.RaceCalendar{}).
		Where("title = ? AND session LIKE ?", title, ses+"%").
		Joins("LEFT JOIN set_match AS s1 ON race_calendar.team1_name = s1.team_name").
		Joins("LEFT JOIN set_match AS s2 ON race_calendar.team2_name = s2.team_name").
		Where("s1.types = ?", typeSF).
		Find(&raceCalendarList).Error
	if err != nil {
		return nil, nil, err
	}

	err = utils.DB.Where("types = ?", typeSF).Find(&setMatchList).Error
	if err != nil {
		return nil, nil, err
	}

	for i := range raceCalendarList {
		if typeSF == "A" {
			raceCalendarList[i].Method = "SSS" + typeSF
		} else if typeSF == "B" {
			raceCalendarList[i].Method = "SSS" + typeSF
		} else if typeSF == "C" {
			raceCalendarList[i].Method = "SSS" + typeSF
		} else if typeSF == "D" {
			raceCalendarList[i].Method = "SSS" + typeSF
		}
		raceCalendarList[i].Schedule = typeSF + "组赛程"
	}

	for i := range setMatchList {
		setMatchList[i].Typestr = "SSS" + setMatchList[i].Types
	}

	return raceCalendarList, setMatchList, nil
}

//	func getMatchesList() ([]model.Matches, error) {
//		var matches []model.Matches
//		if err := utils.DB.Find(&matches).Error; err != nil {
//			return nil, err
//		}
//
//		for i := range matches {
//			switch matches[i].Type {
//			case "1":
//				matches[i].Type = "英雄联盟"
//			case "2":
//				matches[i].Type = "王者荣耀"
//			case "3":
//				matches[i].Type = "第五人格"
//			case "4":
//				matches[i].Type = "无畏契约"
//			case "5":
//				matches[i].Type = "DOTA2"
//			case "6":
//				matches[i].Type = "CSGO"
//			case "7":
//				matches[i].Type = "绝地求生"
//			case "8":
//				matches[i].Type = "和平精英"
//			case "9":
//				matches[i].Type = "综合赛事"
//			case "10":
//				matches[i].Type = "守望先锋"
//			default:
//				return nil, errors.New("无效的类型值")
//			}
//		}
//
//		return matches, nil
//	}
func getMatchesList() (map[string][]model.Matches, error) {
	var matches []model.Matches
	if err := utils.DB.Find(&matches).Error; err != nil {
		return nil, err
	}

	result := make(map[string][]model.Matches)

	for _, match := range matches {
		switch match.Type {
		case "1":
			result["英雄联盟"] = append(result["英雄联盟"], match)
		case "2":
			result["王者荣耀"] = append(result["王者荣耀"], match)
		case "3":
			result["第五人格"] = append(result["第五人格"], match)
		case "4":
			result["无畏契约"] = append(result["无畏契约"], match)
		case "5":
			result["DOTA2"] = append(result["DOTA2"], match)
		case "6":
			result["CSGO"] = append(result["CSGO"], match)
		case "7":
			result["绝地求生"] = append(result["绝地求生"], match)
		case "8":
			result["和平精英"] = append(result["和平精英"], match)
		case "9":
			result["综合赛事"] = append(result["综合赛事"], match)
		case "10":
			result["守望先锋"] = append(result["守望先锋"], match)
		default:
			return nil, errors.New("无效的类型值")
		}
	}

	return result, nil
}

func getMatches() (map[string]interface{}, error) {
	var mainMatches []model.MainMatches
	err := utils.DB.Find(&mainMatches).Error
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	for _, mainMatch := range mainMatches {
		var subMatches []model.SubMatches
		err = utils.DB.Where("main_match_id =?", mainMatch.Id).Find(&subMatches).Error
		if err != nil {
			return nil, err
		}

		mainMatchData := make(map[string]interface{})
		mainMatchData["home_team_name"] = mainMatch.HomeTeamName
		mainMatchData["away_team_name"] = mainMatch.AwayTeamName
		mainMatchData["home_team_logo"] = mainMatch.HomeTeamLogo
		mainMatchData["away_team_logo"] = mainMatch.AwayTeamLogo
		mainMatchData["home_score"] = mainMatch.HomeScore
		mainMatchData["away_score"] = mainMatch.AwayScore
		mainMatchData["win_team_id"] = mainMatch.WinTeamId
		mainMatchData["game_stage"] = mainMatch.GameStage
		mainMatchData["start_time"] = mainMatch.StartTime
		mainMatchData["end_time"] = mainMatch.EndTime

		children := []map[string]interface{}{}
		for _, subMatch := range subMatches {
			var grandSubMatches []model.GrandSubMatches
			err = utils.DB.Where("sub_match_id =?", subMatch.Id).Find(&grandSubMatches).Error
			if err != nil {
				return nil, err
			}

			subMatchData := make(map[string]interface{})
			subMatchData["home_team_name"] = subMatch.HomeTeamName
			subMatchData["away_team_name"] = subMatch.AwayTeamName
			subMatchData["home_team_logo"] = subMatch.HomeTeamLogo
			subMatchData["away_team_logo"] = subMatch.AwayTeamLogo
			subMatchData["home_score"] = subMatch.HomeScore
			subMatchData["away_score"] = subMatch.AwayScore
			subMatchData["win_team_id"] = subMatch.WinTeamId
			subMatchData["game_stage"] = subMatch.GameStage
			subMatchData["start_time"] = subMatch.StartTime
			subMatchData["end_time"] = subMatch.EndTime

			subChildren := []map[string]interface{}{}
			for _, grandSubMatch := range grandSubMatches {
				grandSubMatchData := make(map[string]interface{})
				grandSubMatchData["home_team_name"] = grandSubMatch.HomeTeamName
				grandSubMatchData["away_team_name"] = grandSubMatch.AwayTeamName
				grandSubMatchData["home_team_logo"] = grandSubMatch.HomeTeamLogo
				grandSubMatchData["away_team_logo"] = grandSubMatch.AwayTeamLogo
				grandSubMatchData["home_score"] = grandSubMatch.HomeScore
				grandSubMatchData["away_score"] = grandSubMatch.AwayScore
				grandSubMatchData["win_team_id"] = grandSubMatch.WinTeamId
				grandSubMatchData["game_stage"] = grandSubMatch.GameStage
				grandSubMatchData["start_time"] = grandSubMatch.StartTime
				grandSubMatchData["end_time"] = grandSubMatch.EndTime
				subChildren = append(subChildren, grandSubMatchData)
			}

			subMatchData["children"] = subChildren
			children = append(children, subMatchData)
		}

		mainMatchData["children"] = children
		result[fmt.Sprintf("%d", mainMatch.Id)] = mainMatchData
	}

	//jsonData, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	fmt.Println("jsonData:====", result)

	return result, nil
}

func getMatches1(formatname string) ([][]model.MatchInformation, error) {
	var matchInformations []model.MatchInformation
	if err := utils.DB.Where("format_name =?", formatname).Find(&matchInformations).Error; err != nil {
		return nil, err
	}

	// 使用 map 按照处理后的 match_part_name 分类
	matchMap := make(map[string][]model.MatchInformation)
	for _, match := range matchInformations {
		if match.MatchPart1Name != "" {
			match.MatchPart1Name = match.MatchPart1Name + "-" + match.MatchPartName
		}
		matchMap[match.MatchPartName] = append(matchMap[match.MatchPartName], match)
	}

	// 将分类后的结果转换为二维切片
	var result [][]model.MatchInformation
	for _, matches := range matchMap {
		result = append(result, matches)
	}

	// 按照每组第一个元素的 ID 排序
	sort.Slice(result, func(i, j int) bool {
		return result[i][0].Id < result[j][0].Id
	})

	return result, nil
}
