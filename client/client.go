package client

type Client interface {
    Init(...Option)
    Options() Options
}

type Option func(*Options)

func Init(opts ...Option) {}
