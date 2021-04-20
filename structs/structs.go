package main

import "fmt"

// User is a user typee
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

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.Firstname, u.Lastname, u.Email)
	return desc
}

func describeGroup(g Group) string {
	if len(g.users) > 2 {
		g.spaceAvailable = false
	}
	desc := fmt.Sprintf("This user group has %d. Newest user is %s %s. Accepting New Users: %t", len(g.users), g.newestUser.Firstname, g.newestUser.Lastname, g.spaceAvailable)
	return desc
}

func main() {
	u := User{ID: 1, Firstname: "Taylor", Lastname: "Chase", Email: "tchase@example.com",}
	u2 := User{ID: 2, Firstname: "Parker", Lastname: "Lohman", Email: "ptlohman@example.com",}
	u3 := User{ID: 3, Firstname: "Marta", Lastname: "Blank", Email: "toast@example.com"}

	g := Group{
		role: "admin",
		users: []User{u, u2, u3},
		newestUser: u2,
		spaceAvailable: true,
	}

	fmt.Println(describeUser(u))
	fmt.Println(describeGroup(g))
	fmt.Println(u.Firstname)
}