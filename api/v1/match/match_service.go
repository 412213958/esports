package match

import (
	model "esports/models"
	"esports/tools"
	"esports/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// GetRaceCalendarList godoc
//
//	@Summary		获取具体比赛列表
//	@Description	获取所有具体比赛日程的列表
//	@Tags			所有具体比赛
//	@Produce		json
//	@Success		200	{object}	tools.ECode{data}
//	@Router			/match/getRaceCalendarList [get]
func GetRaceCalendarList(c *gin.Context) {
	raceCalendarList, err := getRaceCalendarList()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取比赛日程列表成功",
		Data:    raceCalendarList,
	})
}

// GetRaceCalendarListByTitle godoc
//
//	@Summary		根据赛事获取具体比赛列表
//	@Description	根据赛事获取具体比赛列表
//	@Tags			赛事具体比赛列表
//	@Produce		json
//	@Param			title	query		string	true	"赛事标题"
//	@Success		200		{object}	tools.ECode{data}
//	@Router			/match/getRaceCalendarList [get]
func GetRaceCalendarListByTitle(c *gin.Context) {
	title := c.Query("title")
	raceCalendarList, err := getRaceCalendarListByTitle(title)
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取比赛列表成功",
		Data:    raceCalendarList,
	})
}

// GetRecommendCardList godoc
//
//	@Summary		获取列表
//	@Description	获取所有的列表
//	@Tags			推荐卡片
//	@Produce		json
//	@Success		200	{object}	tools.ECode{data}
//	@Router			/recommend/getRecommendCardList [get]
func GetRecommendCardList(c *gin.Context) {
	recommendCardList, err := getRecommendCardList()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取推荐卡片列表成功",
		Data:    recommendCardList,
	})
}

// GetContestList godoc
//
//	@Summary		获取比赛列表
//	@Description	获取所有比赛的列表
//	@Tags			比赛
//	@Produce		json
//	@Success		200	{object}	tools.ECode{data}
//	@Router			/contest/getContestList [get]
func GetContestList(c *gin.Context) {
	contestList, err := getContestList()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取比赛列表成功",
		Data:    contestList,
	})
}

// GetVideoList godoc
//
//	@Summary		根据参数获取指定类别的视频指定数量
//	@Description	根据传入的参数 type 和 num 获取指定类别的视频指定数量
//	@Tags			视频
//	@Produce		json
//	@Param			type	query	string	true	"视频类别，如 1:英雄联盟，2:王者荣耀，3:第五人格，4:无畏契约，5:DOTA2"
//	@Param			num		query	int		true	"要获取的数量"
//	@Success		200	{object}	tools.ECode{data}
//	@Router			/api/v1/match/GetVideoList [get]
func GetVideoList(c *gin.Context) {
	videoType := c.Query("type")
	pagenumstr := c.Query("pagenum")
	pagesizestr := c.Query("pagesize")
	if pagenumstr == "" || pagesizestr == "" {
		pagenumstr = "0"
		pagesizestr = "0"
	}
	pagenum, err := strconv.Atoi(pagenumstr)
	pagesize, err := strconv.Atoi(pagesizestr)
	num := (pagenum - 1) * pagesize
	if err != nil {
		utils.HandleError(c, err, 400)
		return
	}
	videoList, err := getVideoListByTypeAndNum(videoType, pagesize, num)
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取视频列表成功",
		Data:    videoList,
	})
}

// GetAllData godoc
//
//	@Summary		获取专区联合数据
//	@Description	获取每组包含RecommendCard的一行数据，Contest的两行数据，Video的六行数据
//	@Tags			综合数据
//	@Produce		json
//	@Success		200	{object}	tools.ECode{data=map[string]interface{}}
//	@Router			/api/v1/match/getAllData [get]
func GetAllData(c *gin.Context) {
	allData, err := getAllData()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取数据成功",
		Data:    allData,
	})
}

// GetRaceCalendarListWF 获取具体比赛列表
// @Summary 获取具体比赛列表
// @Description 获取所有具体比赛日程的列表
// @Tags 所有具体比赛
// @Produce json
// @Param typeSF query int true "TypeSF 值"
// @Success 200 {object} tools.ECode{data}
// @Router /api/v1/match/getRaceCalendarList [get]
func GetRaceCalendarListWF(c *gin.Context) {
	title := "2024LPL夏季赛"
	ses := "组内赛"
	typeSFStr := c.Query("typeSF")
	typeSF, err := strconv.Atoi(typeSFStr)
	if err != nil {
		utils.HandleError(c, fmt.Errorf("invalid typeSF value"), 400)
		return
	}

	raceCalendarList, stageContainerList, err := getRaceCalendarListWF(title, ses, typeSF)
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}

	data := struct {
		RaceCalendarList   []model.RaceCalendar
		StageContainerList []model.StageContainer
	}{
		RaceCalendarList:   raceCalendarList,
		StageContainerList: stageContainerList,
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取比赛日程列表成功",
		Data:    data,
	})
}

// GetRaceCalendarListByABCD 获取具体比赛列表
// @Summary 获取具体比赛列表
// @Description 获取所有具体比赛日程的列表
// @Tags 所有具体比赛
// @Produce json
// @Param typeSF query int true "TypeSF 值"
// @Success 200 {object} tools.ECode{data}
// @Router /api/v1/match/getRaceCalendarList [get]
func GetRaceCalendarListByABCD(c *gin.Context) {
	title := "2024LPL夏季赛"
	ses := "组内赛"
	types := c.Query("types")

	raceCalendarList, setMatchList, err := getRaceCalendarListByABCD(title, ses, types)
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}

	data := struct {
		RaceCalendarList   []model.RaceCalendar
		StageContainerList []model.SetMatch
	}{
		RaceCalendarList:   raceCalendarList,
		StageContainerList: setMatchList,
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取比赛列表成功",
		Data:    data,
	})
}

// GetAllRaceCalendarList 合并后的获取比赛列表接口
// @Summary 获取具体比赛列表
// @Description 获取所有具体比赛日程的列表
// @Tags 所有具体比赛
// @Produce json
// @Param typeSF query int true "TypeSF 值"
// @Param types query string false "types 值"
// @Success 200 {object} tools.ECode{data}
// @Router /api/v1/match/getRaceCalendarList [get]
func GetAllRaceCalendarList(c *gin.Context) {
	title := "2024LPL夏季赛"
	ses := "组内赛"

	raceCalendarListWin, stageContainerListWin, err1 := getRaceCalendarListWF(title, ses, 1)
	if err1 != nil {
		utils.HandleError(c, err1, 500)
		return
	}

	raceCalendarListLose, stageContainerListLose, err2 := getRaceCalendarListWF(title, ses, -1)
	if err2 != nil {
		utils.HandleError(c, err2, 500)
		return
	}

	raceCalendarList_A, setMatchList_A, err3 := getRaceCalendarListByABCD(title, ses, "A")
	if err3 != nil {
		utils.HandleError(c, err3, 500)
		return
	}

	raceCalendarList_B, setMatchList_B, err4 := getRaceCalendarListByABCD(title, ses, "B")
	if err4 != nil {
		utils.HandleError(c, err4, 500)
		return
	}

	raceCalendarList_C, setMatchList_C, err5 := getRaceCalendarListByABCD(title, ses, "C")
	if err5 != nil {
		utils.HandleError(c, err5, 500)
		return
	}

	raceCalendarList_D, setMatchList_D, err6 := getRaceCalendarListByABCD(title, ses, "D")
	if err6 != nil {
		utils.HandleError(c, err6, 500)
		return
	}

	data := struct {
		RaceCalendarListWLABCD [][]model.RaceCalendar
		StageContainerListWL   [][]model.StageContainer
		SetMatchListABCD       [][]model.SetMatch
	}{
		RaceCalendarListWLABCD: [][]model.RaceCalendar{raceCalendarListWin, raceCalendarListLose, raceCalendarList_A, raceCalendarList_B, raceCalendarList_C, raceCalendarList_D},
		StageContainerListWL:   [][]model.StageContainer{stageContainerListWin, stageContainerListLose},
		SetMatchListABCD:       [][]model.SetMatch{setMatchList_A, setMatchList_B, setMatchList_C, setMatchList_D},
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取比赛列表成功",
		Data:    data,
	})
}

// GetMatchesList godoc
//
//	@Summary		获取赛事列表
//	@Description	获取所有赛事的列表
//	@Tags			赛事
//	@Produce		json
//	@Success		200	{object}	tools.ECode{data}
//	@Router			/api/v1/match/GetMatchesList [get]
func GetMatchesList(c *gin.Context) {
	contestList, err := getMatchesList()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取赛事列表成功",
		Data:    contestList,
	})
}

// GetMatches godoc
//
// @Summary 获取赛事图数据列表
// @Description 获取赛事图数据的列表
// @Tags 赛事图数据
// @Produce json
// @Success 200 {object} tools.ECode{data}
// @Router /api/v1/match/GetMatches [get]
func GetMatches(c *gin.Context) {
	contestList, err := getMatches()
	if err != nil {
		log.Println("获取赛事列表时发生错误:", err)
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取赛事列表成功",
		Data:    contestList,
	})
}

// GetMatches1 godoc
//
// @Summary 获取赛事图1数据列表
// @Description 获取赛事图1数据的列表
// @Tags 赛事图1数据
// @Produce json
// @Success 200 {object} tools.ECode{data}
// @Router /api/v1/match/GetMatches1 [get]
func GetMatches1(c *gin.Context) {
	formatname1 := "冒泡赛"
	formatname2 := "双败赛"
	contestList1, err1 := getMatches1(formatname1)
	contestList2, err2 := getMatches1(formatname2)

	if err1 != nil {
		log.Println("获取冒泡赛赛事列表时发生错误:", err1)
		utils.HandleError(c, err1, 500)
		return
	}

	if err2 != nil {
		log.Println("获取双败赛赛事列表时发生错误:", err2)
		utils.HandleError(c, err2, 500)
		return
	}

	c.JSON(200, tools.ECode{
		Code:    0,
		Message: "获取赛事1列表成功",
		Data: gin.H{
			"maopao":    contestList1,
			"shuangbai": contestList2,
		},
	})
}
