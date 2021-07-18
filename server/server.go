package server

type Server interface {
    Init(...Option)
    Run() error
    Options() Options
}

type Option func(*Options)

func Init(opts ...Option) {}

func Run() error {
    return nil
}

func Start() error {
    return nil
}

func Stop() error {
    return nil
}
