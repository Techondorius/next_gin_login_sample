package model

type User struct {
	ID       int
	UserID   string
	Password string
}

var USERS = []User{
	{
		ID:       1,
		UserID:   "asdl0606",
		Password: "password",
	},
}

func FindUserByUserID(userID string) *User {
	return &USERS[0]
}
