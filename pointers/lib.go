package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn time.Time
}

type Member struct {
	name Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total int
	lended int
}

type Library struct {
	members map[Name]Member
	books map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMembersAudit(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}

	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("member did not checkout this book")
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books: make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Webapps in Go"] = BookEntry{
		total: 4,
		lended: 0,
	}

	library.books["See Go, go!"] = BookEntry{
		total: 3,
		lended: 0,
	}

	library.books["Learning go"] = BookEntry{
		total: 2,
		lended: 0,
	}

	library.books["Go go power!"] = BookEntry{
		total: 1,
		lended: 0,
	}

	library.members["Taylor"] = Member{"Taylor", make(map[Title]LendAudit)}
	library.members["Parker"] = Member{"Parker", make(map[Title]LendAudit)}
	library.members["Marta"] = Member{"Marta", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMembersAudit(&library)

	member := library.members["Taylor"]
	checkout := checkoutBook(&library, "Webapps in Go", &member)
	fmt.Println("\ncheckout")

	if checkout {
		printLibraryBooks(&library)
		printMembersAudit(&library)
	}

	returned := returnBook(&library, "Webapps in Go", &member)
	fmt.Println("\nreturned")
	if returned {
		printLibraryBooks(&library)
		printMembersAudit(&library)
	}
}