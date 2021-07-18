package client

type jrpcClient struct {
	opts Options
}

func newJRpcClient(opts ...Option) Client {
	options := newOptions(opts...)

	return &jrpcClient{
		opts: options,
	}
}

func (c *jrpcClient) Init(opts ...Option) error {
	return nil
}

func (c *jrpcClient) Options() Options {
	return c.opts
}
