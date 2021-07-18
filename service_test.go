package service

import (
	"log"
	"os"
	"testing"
)

func TestService(t *testing.T) {
	logger := log.New(os.Stdout, "test-service: ", log.LstdFlags)

	svc := NewService(logger)

	svc.Init()

	if err := svc.Run(); err != nil {
		t.Fatal(err)
	}
}
