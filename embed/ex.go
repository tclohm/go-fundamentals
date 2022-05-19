package main

import "fmt"

type Whisperer interface {
	Whisper() string
}

type Yeller interface {
	Yell() string
}

type Talker interface {
	Whisperer
	Yeller
}

func talk(t Talker) {
	fmt.Println(t.Yell())
	fmt.Println(t.Whisper())
}

type Account struct {
	accountID int
	balance int
	name string
}

type ManagerAccount struct {
	Account
}

func (a *Account) GetBalance() int {
	return a.balance
}

func (a Account) String() string {
	return fmt.Sprintf("Standard (%v) $%v \"%v\"", a.accountID, a.balance, a.name)
}

func (m ManagerAccount) String() string {
	return fmt.Sprintf("Manager (%v) $%v \"%v\"", m.accountID, m.balance, m.name)
}

func main() {
	mgr := ManagerAccount{Account{2, 30, "Bianca"}}
	fmt.Printf("%v\n", mgr)
	fmt.Printf("%v\n", mgr.GetBalance())
	fmt.Printf("%v\n", mgr.accountID)
}