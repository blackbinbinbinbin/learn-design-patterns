package main

import (
	"fmt"
)

func main() {
	config := NewServerConf(
		WithHost("127.0.0.1"),
		WithPort(8080),
		WithProtocol("TCP"),
		WithMaxConn(1000),
		WithTimeout(600))

	fmt.Printf("服务器配置：%v\n", config)
}
