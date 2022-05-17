package main

import "fmt"

type Coordinate struct {
	X, Y int
}

func (coord *Coordinate) shiftBy(x, y int) {
	coord.X += x
	coord.Y += y
}

func (coord Coordinate) Dist(other Coordinate) Coordinate {
	return Coordinate{coord.X - other.X, coord.Y - other.Y}
}

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func (lot *ParkingLot) occupy(spaceNumber int) {
	lot.spaces[spaceNumber - 1].occupied = true
}

func (lot *ParkingLot) vacate(spaceNumber int) {
	lot.spaces[spaceNumber - 1].occupied = false
}

func main(){
	coord := Coordinate{5, 5}
	coord.shiftBy(1, 1)

	first := Coordinate{2, 2}
	second := Coordinate{1, 5}

	distance := first.Dist(second)
	fmt.Println(distance)

	lot := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println(lot)
	lot.occupy(1)
	fmt.Println(lot)
	lot.vacate(1)
	fmt.Println(lot)
}