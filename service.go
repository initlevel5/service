package service

import (
    "github.com/initlevel5/service/v1/client"
    "github.com/initlevel5/service/v1/logger"
    "github.com/initlevel5/service/v1/server"
)

type Service interface {
    Init(...Option)
    Run() error
    Options() Options
    Client() client.Client
    Server() server.Server
}

func NewService(opts ...Option) Service {
    return newService(opts...)
}

type service struct {
    opts Options
}

func newService(opts ...Option) Service {
    service := new(service)
    options := newOptions(opts...)

    err := options.Server.Init()

    if err != nil {
        logger.Fatal(err)
    }

    service.opts = options

    return service
}

func (s *service) Init(opts ...Option) {
    for _, o := range opts {
        o(&s.opts)
    }
}

func (s *service) Start() error {
    return nil
}

func (s *service) Stop() error {
    return nil
}

func (s *service) Run() error {
    return nil
}

func (s *service) Options() Options {
    return s.opts
}

func (s *service) Client() client.Client {
    return s.opts.Client
}

func (s *service) Server() server.Server {
    return s.opts.Server
}
