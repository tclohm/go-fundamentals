package main

import (
	"fmt"
)

type GenericList[T comparable] struct {
	data []T
}

func Initialize[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (l *GenericList[T]) Push(value T) {
	l.data = append(l.data, value)
}

func (l *GenericList[T]) Get(index int) (T, string) {
	if index > len(l.data) - 1 {
		var result T
		err := fmt.Errorf("Error: Index: %d is outside of the list", index)
		return result, err.Error()
	}
	return l.data[index], ""
}

func (l *GenericList[T]) Remove(index int) {
	if index > len(l.data) - 1 {
		err := fmt.Errorf("Error: Index: %d is outside of the list", index)
		fmt.Println(err.Error())
		return
	}
	removed := l.data[index]
	l.data = append(l.data[:index], l.data[index+1:]...)
	fmt.Println("removed", removed)
	fmt.Println("new list", l.data)
}

func main() {
	gl := Initialize[string]()
	gl.Push("Paul")
	gl.Push("Christian")
	gl.Push("Caroline")
	gl.Push("Andrew")
	fmt.Println(gl.Get(4))
	fmt.Println(gl.Get(3))
	fmt.Println(gl.Get(2))
	fmt.Println(gl.Get(1))
	fmt.Println(gl.Get(0))
	gl.Remove(2)
	gl.Remove(0)
	gl.Remove(3)
	gl.Remove(1)
}