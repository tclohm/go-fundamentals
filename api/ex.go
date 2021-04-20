package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

const base = "https://swapi.dev/api/"

// Planet
type Planet struct {
	Name string `json:"name"`
	Population string `json:"population"`
	Terrain string `json:"terrain"`
}

// Person
type Person struct {
	Name string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld Planet
}

// All People
type AllPeople struct {
	People []Person `json:"results"`
}

// Method
func (p *Person) getHomeWorld() {
	res, err := http.Get(p.HomeworldURL)
	if err != nill {
		log.Print("Error fetching homeworld", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != {
		log.Print("Error reading response body", err)
	}

	json.Unmarshal(bytes, &p.Homeworld)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(base + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request SW peeps")
	}

	fmt.Println(res)

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request body")
	}

	var people AllPeople

	fmt.Println(string(bytes))

	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}


	fmt.Println(people)

	for _, pers := range people.People {
		pers.getHomeWorld()
		fmt.Println(pers)
	}
}

func main() {
	http.HandleFunc("/people", getPeople)
	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}