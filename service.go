package service

import (
    "github.com/initlevel5/microservice/client"
    "github.com/initlevel5/microservice/server"
)

type Service interface {
    Init(...Option)
    Run() error
    Options() Options
    Client() client.Client
    Server() server.Server
}

type Option func(*Options)

func NewService(opts ...Option) Service {
    return newService(opts...)
}
