package main

import "fmt"

type Server struct {
	users map[string]string
	userChannel chan string
	quitChannel chan string
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {
	for {
		select {

		}
	}
}

func (s *Server) addUser(user string) {
	s.users[user] = user
}

func sendMessage(messageChannel chan <- string) {
	messageChannel <- "hello"
}

func readMessage(messageChannel <- chan string) {
	msg := <- messageChannel
	fmt.Println(msg)
}

func main() {

}