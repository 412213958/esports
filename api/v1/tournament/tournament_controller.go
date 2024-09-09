package tournament

import (
	model "esports/models"
	"esports/utils"
)

func getTournaments(types string) ([]model.Tournaments, error) {
	var tournaments []model.Tournaments

	err := utils.DB.Table("tournaments").Where("type_s = ?", types).Find(&tournaments).Error
	if err != nil {
		return nil, err
	}
	return tournaments, nil
}

func getWeiboData() ([]model.WeiboData, error) {
	var weiboData []model.WeiboData
	err := utils.DB.Table("weibo_data").Find(&weiboData).Error
	if err != nil {
		return nil, err
	}
	return weiboData, nil
}

func getArticlesData() ([]model.Articles, error) {
	var articles []model.Articles
	err := utils.DB.Table("articles").Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func getArticles1Data() (map[string][]model.Articles1, error) {
	var articles1 []model.Articles1
	err := utils.DB.Table("articles_1").Find(&articles1).Error
	if err != nil {
		return nil, err
	}

	// 分类逻辑
	result := make(map[string][]model.Articles1)
	for _, article := range articles1 {
		result[article.TypeS] = append(result[article.TypeS], article)
	}
	return result, nil
}

func getComments() ([]model.Comments, error) {
	var comments []model.Comments
	err := utils.DB.Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}
