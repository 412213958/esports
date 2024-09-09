package team

import (
	model "esports/models"
	"esports/utils"
	"github.com/gin-gonic/gin"
)

// GetLegacyList godoc
//
//	@Summary		获取战队列表
//	@Description	获取所有战队的列表
//	@Tags			战队
//	@Produce		json
//	@Success		200	{object}	tools.ECode{data}
//	@Router			/api/v1/team/getLegacyList [get]
func GetLegacyList(c *gin.Context) {
	videoList, err := getLegacyList()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取战队列表成功",
		Data:    videoList,
	})
}

// GetPlayerAndTeamDataByTeam 获取选手和战队平均数据
// @Summary 获取选手信息和战队平均游戏数据
// @Description 获取选手信息和战队平均游戏数据
// @Tags 所有战队成员+具体比赛数据
// @Produce json
// @Param typeSF query int true " 值"
// @Success 200 {object} tools.ECode{data}
// @Router /api/v1/team/GetPlayerAndTeamData [get]
func GetPlayerAndTeamDataByTeam(c *gin.Context) {
	team := c.Query("team")
	raceCalendarList, setMatchList, err := getPlayerAndTeamDataByTeam(team)
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}

	data := struct {
		Players  []model.Players
		TeamInfo TeamInfo1
	}{
		Players:  raceCalendarList,
		TeamInfo: setMatchList,
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取选手和战队平均数据成功",
		Data:    data,
	})
}

// GetAllPlayerAndTeamDataByTeam 获取选手和战队平均数据
// @Summary 获取选手信息和战队平均游戏数据
// @Description 获取选手信息和战队平均游戏数据
// @Tags 所有战队成员+具体比赛数据("老P")
// @Produce json
// @Param typeSF query int true " 值"
// @Success 200 {object} tools.ECode{data}
// @Router /api/v1/team/GetAllPlayerAndTeamDataByTeam [get]
func GetAllPlayerAndTeamDataByTeam(c *gin.Context) {
	setMatchList, err := getAllPlayerAndTeamDataByTeam()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}

	data := struct {
		//Players  [][]model.Players
		TeamInfo []TeamInfo2
	}{
		//Players:  raceCalendarList,
		TeamInfo: setMatchList,
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "老P专属接口:获取选手和战队平均数据成功",
		Data:    data,
	})
}

// Dynamicodds 动态赔率
// @Summary 动态赔率
// @Description 动态赔率
// @Tags 动态赔率("老P")
// @Produce json
// @Param typeSF query int true " 值"
// @Success 200 {object} tools.ECode{data}
// @Router /api/v1/team/Dynamicodds [get]
func Dynamicodds(c *gin.Context) {
	setMatchList, err := getAllPlayerAndTeamDataByTeam()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}

	data := struct {
		//Players  [][]model.Players
		TeamInfo []TeamInfo2
	}{
		//Players:  raceCalendarList,
		TeamInfo: setMatchList,
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "老P专属接口:获取选手和战队平均数据成功",
		Data:    data,
	})
}
