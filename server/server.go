package server

type Server interface {
    Init(...Option) error
    Run() error
    Options() Options
}

var (
    DefaultServer Server = newJRpcServer()
)

func Init(opts ...Option) error {
    return nil
}

func Run() error {
    return nil
}

func Start() error {
    return nil
}

func Stop() error {
    return nil
}
