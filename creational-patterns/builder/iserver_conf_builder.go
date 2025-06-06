package main

type IServerConfBuilder interface {
	SetHost(host string) *ServerConfBuilder
	SetPort(port int) *ServerConfBuilder
	SetProtocol(protocol string) *ServerConfBuilder
	SetMaxConn(maxConn int) *ServerConfBuilder
	SetTimeout(timeout int) *ServerConfBuilder

	GetServConf() *ServerConf
}

func NewServerConfBuilder() *ServerConfBuilder {
	return &ServerConfBuilder{
		config: &ServerConf{
			// 默认值
			host:     "0.0.0.0",
			protocol: "http",
			port:     8080,
			maxConn:  100,
			timeout:  30,
		},
	}
}

type ServerConfBuilder struct {
	config *ServerConf
}

func (s *ServerConfBuilder) SetHost(host string) *ServerConfBuilder {
	s.config.host = host
	return s
}

func (s *ServerConfBuilder) SetPort(port int) *ServerConfBuilder {
	s.config.port = port
	return s
}

func (s *ServerConfBuilder) SetProtocol(protocol string) *ServerConfBuilder {
	s.config.protocol = protocol
	return s
}

func (s *ServerConfBuilder) SetMaxConn(maxConn int) *ServerConfBuilder {
	s.config.maxConn = maxConn
	return s
}

func (s *ServerConfBuilder) SetTimeout(timeout int) *ServerConfBuilder {
	s.config.timeout = timeout
	return s
}

func (s *ServerConfBuilder) GetServConf() *ServerConf {
	return s.config
}
