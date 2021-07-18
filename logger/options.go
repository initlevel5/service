package logger

type Option func(*Options)

type Options struct{}

func NewOptions(opts ...Option) Options {
    options := Options{}

    for _, o := range opts {
        o(&options)
    }

    return options
}
