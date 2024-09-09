package middleware

import (
	model "esports/models"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

// User 定义用户结构体
type User struct {
	gorm.Model
	ID    int
	Money float64
}

type Prediction struct {
	gorm.Model
	UserID    int
	Team      int
	BetAmount float64
}

// Result 定义比赛结果枚举
type Result int

const (
	Team1Win Result = iota
	Team2Win
	Draw
)

// 从数据库获取用户数据
func getUsersFromDB(db *gorm.DB) ([]*User, error) {
	var users []*User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 从数据库获取预测数据
func getPredictionsFromDB(db *gorm.DB) ([]*model.Predictions, error) {
	var predictions []*model.Predictions
	err := db.Find(&predictions).Error
	if err != nil {
		return nil, err
	}
	return predictions, nil
}

// SettlePredictionMiddleware 结算预测中间件函数
func SettlePredictionMiddleware(db *gorm.DB) func() {
	return func() {
		users, err := getUsersFromDB(db)
		if err != nil {
			fmt.Println("获取用户数据出错:", err)
			return
		}

		predictions, err := getPredictionsFromDB(db)
		if err != nil {
			fmt.Println("获取预测数据出错:", err)
			return
		}

		for _, prediction := range predictions {
			var winAmount float64
			if prediction.Team == 1 {
				betAmount, err := strconv.ParseFloat(prediction.BetAmount, 64)
				if err != nil {
					fmt.Println("转换出错:", err)
					continue
				}
				winAmount = betAmount * 2
			} else {
				winAmount = 0
			}

			for _, user := range users {
				if int32(user.ID) == prediction.UserId {
					user.Money += winAmount
				}
			}
		}

		// 更新用户金钱到数据库
		for _, user := range users {
			err := db.Model(&User{}).Where("id =?", user.ID).Update("money", user.Money).Error
			if err != nil {
				fmt.Println("更新用户金钱出错:", err)
				return
			}
		}
	}
}
