package main

import (
	"fmt"
)

type User struct {
	ID	int
	Firstname, Lastname, Email string
}

// Group is a set of users
type Group struct {
	role string
	users []User
	newestUser User
	spaceAvailable bool
}

func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.Firstname, u.Lastname, u.Email)
	return desc
}

// method more for changing state

func main() {
	user := User{ID: 1, Firstname: "Yaya", Lastname: "Ding Dong", Email: "yaya@dingdong.edu"}

	desc := user.describe()

	fmt.Println(desc)
}