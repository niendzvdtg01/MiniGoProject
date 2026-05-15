package model

type User struct {
	UserName string
	Password string
}

func (u User) Public() User {
	return User{
		UserName: u.UserName,
		Password: u.Password,
	}
}
