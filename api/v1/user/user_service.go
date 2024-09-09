package user

import (
	"database/sql"
	model "esports/models"
	"esports/utils"
	"fmt"
)

type UserService struct{}

// NewUserService 创建新的用户服务实例
func NewUserService() *UserService {
	return &UserService{}
}

// GetAllUsers 获取所有用户
func (us UserService) GetAllUsers() ([]model.User, error) {
	db, err := utils.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []model.User
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserByID 根据 ID 获取用户
func (us UserService) GetUserByID(userID string) (*model.User, error) {
	db, err := utils.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user model.User
	query := "SELECT * FROM users WHERE user_id =?"
	err = db.QueryRow(query, userID).Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("用户不存在")
		}
		return nil, err
	}

	return &user, nil
}

// CreateUser 创建新用户
func (us UserService) CreateUser(user *model.User) error {
	db, err := utils.InitDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO users (username, password, email, created_at) VALUES (?,?,?,?)"
	_, err = db.Exec(query, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser 更新用户信息
func (us UserService) UpdateUser(userID string, updatedUser *model.User) error {
	db, err := utils.InitDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE users SET username =?, password =?, email =? WHERE user_id =?"
	_, err = db.Exec(query, updatedUser.Username, updatedUser.Password, updatedUser.Email, userID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser 删除用户
func (us UserService) DeleteUser(userID string) error {
	db, err := utils.InitDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM users WHERE user_id =?"
	_, err = db.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}
