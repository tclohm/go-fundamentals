package main

import (
	"fmt"
	"net/http"
	"log"
	//"html/template"
)

func home(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("Home!")
}

// comment
type Todo struct {
	Title string
	Content string
}

// comment
type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

var todos []Todo

func getTodos(writer http.ResponseWriter, req *http.Request) {
	pageVariables := PageVariables {
		PageTitle: "Get Todos",
		PageTodos: todos,
	}
	t, err := template.ParseFiles("todos.html")
	//fmt.Fprint(writer, "TADA!")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("template parsing error")
	}

	err = t.Execute(writer, pageVariables)
}

func addTodo(writer http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("request parsing error: ", err)
	}

	todo := Todo{
		Title: req.FormValue("title"),
		Content: req.FormValue("content"),
	}

	todos = append(todos, todo)
	log.Print(todos)
	http.Redirect(w, req, "/todos/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo/", addTodo)
	fmt.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}