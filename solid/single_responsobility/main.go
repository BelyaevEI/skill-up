package main

type User struct {
	Username string
	Email    string
}

type UserService struct{}

func (us *UserService) SendWelcomeEmail(u *User) {
	// Код для отправки email
}

type UserRepository struct{}

func (ur *UserRepository) SaveToDB(u *User) {
	// Код для сохранения в базу данных
}

func main() {
	user := &User{
		Username: "Jonh",
		Email:    "john77@gmail.com",
	}

	repo := UserRepository{}
	service := UserService{}

	//Разделение зависимости
	repo.SaveToDB(user)
	service.SendWelcomeEmail(user)

}
