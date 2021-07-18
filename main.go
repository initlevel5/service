package main

import "github.com/initlevel5/service/v1"

func main() {
	s := service.NewService()

	s.Init()

	s.Run()
}
