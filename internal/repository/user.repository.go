package repository

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (u *userRepository) GetUserByEmail(email string) bool {
	return true
}