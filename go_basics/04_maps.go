package main

import "fmt"

func main() {
	// var userEmails map[int]string = make(map[int]string)

	// userEmails[1] = "user1@gmail.com"
	// userEmails[2] = "user2@gmail.com"

	// fmt.Println(userEmails)

	userEmails := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	fmt.Println(userEmails)

	firstEmail := userEmails[1]

	thirdEmail, ok := userEmails[3]

	fmt.Println(firstEmail)

	fmt.Println(thirdEmail, ok)

	if _, ok := userEmails[4]; ok {
		fmt.Println("email exists")
	} else {
		fmt.Println("email doesn't exists")
	}

	// delete
	delete(userEmails, 2)
	fmt.Println(userEmails)
}