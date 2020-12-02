package httpclient

import "context"

type Options struct {
	timeout int
	params  map[string]string
	headers map[string]string
}

type Option func(*Options)

func WithTimeout(timeout int) Option {
	return func(s *Options) {
		s.timeout = timeout
	}
}

func WithParams(params map[string]string) Option {
	return func(s *Options) {
		if s.params == nil {
			s.params = make(map[string]string)
		}

		for key, val := range params {
			s.params[key] = val
		}
	}
}

func WithHeaders(headers map[string]string) Option {
	return func(s *Options) {
		if s.headers == nil {
			s.headers = make(map[string]string)
		}

		for key, val := range headers {
			s.headers[key] = val
		}
	}
}

func optionHandle(ctx context.Context, opts ...Option) (*Options, error) {
	if ctx == nil {
		ctx = context.TODO()
	}
	// 默认超时时间为30秒
	options := &Options{
		timeout: 30,
	}

	for _, opt := range opts {
		opt(options)
	}

	return options, nil
}
