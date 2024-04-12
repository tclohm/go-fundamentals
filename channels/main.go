package main

import "fmt"

// channels can be used as a signal too

type Server struct {
	users map[string]string
	userChannel chan string
	quitChannel chan struct{}
}

func NewServer() *Server {
	return &Server{
		users: make(map[string]string),
		userChannel: make(chan string),
		quitChannel: make(chan struct{}),
	}
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {

	running:
		for {
			select {
			case msg := <-s.userChannel:
				fmt.Println(msg)
			case <- s.quitChannel:
				fmt.Println("shutting down")
				break running
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
	server := NewServer()
	server.Start()

	server.quitChannel <- struct{}{}
}