package main

import (
	"fmt"
	"log"
	"net/http"
	// "runtime"
)

//if your goroutine cannot make progress until it gets the result from another, 
// oftentimes it is simpler to just do the work yourself rather than to delegate it
func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr: addr,
		Handler: handler,
	}

	go func() {
		<-stop // what for stop signal
		s.Shutdown(context.Background())
	}

	return s.ListenAndServe()
}

func serveApp(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Gopher")
	})
	return serve("0.0.0.0:8080", mux, stop)
}

func serveDebug(stop <-chan stuct{}) error {
	return serve("127.0.0.1:8001", http.DefaultServeMux, stop)
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- serveDebug(stop)
	}()
	go func() {
		done <- serveApp(stop)
	}()

	var stopped bool
	for i := 0 ; i < cap(done) ; i++ {
		if err := <-done ; err != nil {
			fmt.Println("error: %v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
// If your function starts a goroutine you must provide the caller with a way to explicitly stop that goroutine. 
// It is often easier to leave decision to execute a function asynchronously to the caller of that function. 