package main

// ConfOption 定义配置函数类型
type ConfOption func(*ServerConf)

func WithHost(host string) ConfOption {
	return func(conf *ServerConf) {
		conf.host = host
	}
}

func WithPort(port int) ConfOption {
	return func(conf *ServerConf) {
		conf.port = port
	}
}

func WithProtocol(protocol string) ConfOption {
	return func(conf *ServerConf) {
		conf.protocol = protocol
	}
}

func WithMaxConn(maxConn int) ConfOption {
	return func(conf *ServerConf) {
		conf.maxConn = maxConn
	}
}

func WithTimeout(timeout int) ConfOption {
	return func(conf *ServerConf) {
		conf.timeout = timeout
	}
}
