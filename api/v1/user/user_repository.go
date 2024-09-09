package user

type UserRepository struct{}

// NewUserRepository 创建新的用户存储库实例
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
