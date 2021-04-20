package main

import "fmt"

type Mom struct {
	name: string
	last: string
	Kids: []Kid
	Friends: []Friend
}

type Dad struct {
	name: string
	last: string
	Kids: []Kid
	Friends: []Friend
}

func (d Dad) CallFrequently(kids []Kids) {
	fmt.Println("Clean your room")
}

func (m Mom) CallFrequently(kids []Kids) {
	fmt.Println("Did you brush your teeth")
}

func (d Dad) CheckIfOvenIsOff() bool {}
func (m Mom) CheckIfOvenIsOff() bool {}

// describes whatever struct is passed in
func DoDescribing(d Describer) string {
	return d.describe()
}


type Worrier interface {
	CallFrequently([]Kids)
	CheckIfOvenIsOff()
}

// becomes more generic
type Describer interface {
	describe() string
}

// empty interface, thing of the any type
interface {}

var people map[string]interface{}