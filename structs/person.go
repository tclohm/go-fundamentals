package main

// perfer fewer, larger packages

type Person struct {
	Name string
	Age int
}


// use interfaces to describe the behavior your functions or methods require
// avoid the use of global state
func AverageAge(people []Person) int {
	if len(people) == 0 {
		return 0
	}

	var count, sum int
	for _, p := range people {
		sum += p.Age
		count += 1
	}

	return sum / count
}