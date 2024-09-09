package tournament

import (
	model "esports/models"
	"esports/utils"
	"github.com/gin-gonic/gin"
)

// GetTournaments godoc
//
//	@Summary  获取赛事战和队列表
//	@Description  获取所有赛事的列表
//	@Tags  赛事
//	@Produce  json
//	@Success  200  {object}  tools.ECode{data}
//	@Router  /api/v1/tournament/GetTournaments [get]
func GetTournaments(c *gin.Context) {
	tournamentList1, err := getTournaments("热门赛事")
	tournamentList2, err := getTournaments("热门圈子")
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	data := struct {
		Tournaments1 []model.Tournaments
		Tournaments2 []model.Tournaments
	}{
		Tournaments1: tournamentList1,
		Tournaments2: tournamentList2,
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取赛事列表成功",
		Data:    data,
	})
}

// GetWeiboData godoc
//
//	@Summary  获取微博数据列表
//	@Description  获取所有微博数据的列表
//	@Tags  微博数据
//	@Produce  json
//	@Success  200  {object}  tools.ECode{data}
//	@Router  /api/v1/tournament/GetWeiboData [get]
func GetWeiboData(c *gin.Context) {
	weiboDataList, err := getWeiboData()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取微博数据列表成功",
		Data:    weiboDataList,
	})
}

// GetArticlesData 获取文章数据列表
// @Summary 获取文章数据列表
// @Description 获取所有文章数据的列表
// @Tags 文章数据
// @Produce json
// @Success 200 {object} ECode{data}
// @Router /api/v1/tournament/GetArticlesData [get]
func GetArticlesData(c *gin.Context) {
	var articles []model.Articles
	articles, err := getArticlesData()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取文章数据列表成功",
		Data:    articles,
	})
}

// GetArticles1Data 获取文章 1 数据列表
// @Summary 获取文章 1 数据列表
// @Description 获取所有文章 1 数据的列表
// @Tags 文章 1 数据
// @Produce json
// @Success 200 {object} ECode{data}
// @Router /api/v1/tournament/GetArticles1Data [get]
func GetArticles1Data(c *gin.Context) {

	var articles1 map[string][]model.Articles1
	articles1, err := getArticles1Data()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取文章 1 数据列表成功",
		Data:    articles1,
	})
}

// GetCommentsData 获取评论数据列表
// @Summary 获取评论数据列表
// @Description 获取所有评论数据的列表
// @Tags 评论数据
// @Produce json
// @Success 200 {object} ECode{data}
// @Router /api/v1/tournament/GetCommentsData [get]
func GetCommentsData(c *gin.Context) {

	var comments []model.Comments
	comments, err := getComments()
	if err != nil {
		utils.HandleError(c, err, 500)
		return
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取评论数据列表成功",
		Data:    comments,
	})
}

// GetYCData  获取预测数据列表
// @Summary 获取预测数据列表
// @Description 获取所有预测数据的列表
// @Tags 预测论数据
// @Produce json
// @Success 200 {object} ECode{data}
// @Router /api/v1/tournament/GetYCData [get]
func GetYCData(c *gin.Context) {
	var table1Data []model.Table1
	result1 := utils.DB.Find(&table1Data)
	if result1.Error != nil {
		utils.HandleError(c, result1.Error, 500)
		return
	}

	var table2Data []model.Table2
	result2 := utils.DB.Find(&table2Data)
	if result2.Error != nil {
		utils.HandleError(c, result2.Error, 500)
		return
	}

	var table3Data []model.Table3
	result3 := utils.DB.Find(&table3Data)
	if result3.Error != nil {
		utils.HandleError(c, result3.Error, 500)
		return
	}

	combinedData := make([]struct {
		Table1Data model.Table1
		Table2Data model.Table2
		Table3Data []model.Table3
	}, len(table1Data))

	for i := range table1Data {
		combinedData[i].Table1Data = table1Data[i]
		if i < len(table2Data) {
			combinedData[i].Table2Data = table2Data[i]
		}
		if i*2 < len(table3Data) {
			combinedData[i].Table3Data = table3Data[i*2 : i*2+2]
		}
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取组合数据成功",
		Data:    combinedData,
	})
}
