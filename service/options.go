package service

import (
    "github.com/initlevel5/service/client"
    "github.com/initlevel5/service/server"
)

type Option func(o *Options)

type Options struct {
    Client client.Client
    Server server.Server
}

func NewOptions(opts ...Option) Options {
    options := Options{
        Client: client.DefaultClient,
        Server: server.DefaultServer,
    }

    for _, o := range opts {
        o(&options)
    }

    return options
}

func Client(c client.Client) Option {
    return func(o *Options) {
        o.Client = c
    }
}

func Server(s server.Server) Option {
    return func(o *Options) {
        o.Server = s
    }
}
