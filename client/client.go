package client

type Client interface {
    Init(...Option) error
    Options() Options
}

var (
    DefaultClient Client = newJRpcClient()
)

func Init(opts ...Option) {}
