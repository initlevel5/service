package client

type Option func(*Options)

type Options struct{}

func newOptions(opts ...Option) Options {
    options := Options{}

    for _, o := range opts {
        o(&options)
    }

    return options
}
