package user

type User struct {
	Id          string
	UserName    string
	Email       string
	PhoneNumber string
}

func Current() (User, error) {
	user := User{
		Id:          "7d2bfcab-60a7-422c-b4e7-1d0f4a6df250",
		UserName:    "administrator",
		Email:       "",
		PhoneNumber: "",
	}
	return user, nil
}
