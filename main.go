package main

import (
	"os"
	"try-go-react-todo-back/router"
)

func main() {

	s := router.NewServer()

	if err := s.Start(":80"); err != nil {
		os.Exit(1)
	}
}
