package main

type ServerConf struct {
	host     string
	port     int
	protocol string
	maxConn  int
	timeout  int
}

func NewServerConf(options ...ConfOption) *ServerConf {
	s := &ServerConf{}
	for _, option := range options {
		option(s)
	}
	return s
}
