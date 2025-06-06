package main

type ServerConf struct {
	host     string
	port     int
	protocol string
	maxConn  int
	timeout  int
}
