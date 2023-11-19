package ports

type UserService interface {
	Login(email, password string) (string, error)
	SignUp(email, password string) (string, error)
}

type UserRepository interface {
	Login(email, password string) (string, error)
	SignUp(email, password string) (string, error)
}
