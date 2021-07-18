package server

type jrpcServer struct {
	opts Options
}

func newJRpcServer(opts ...Option) Server {
	options := newOptions(opts...)

	return &jrpcServer{
		opts: options,
	}
}

func (s *jrpcServer) Init(opts ...Option) error {
	return nil
}

func (s *jrpcServer) Run() error {
	return nil
}

func (s *jrpcServer) Options() Options {
	return s.opts
}
