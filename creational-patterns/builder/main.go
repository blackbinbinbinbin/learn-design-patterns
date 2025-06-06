package main

import (
	"fmt"
)

func main() {
	confBuilder := NewServerConfBuilder()
	config := confBuilder.SetHost("127.0.0.1").
		SetPort(8080).
		SetProtocol("TCP").
		SetMaxConn(1000).
		SetTimeout(500).
		GetServConf()

	fmt.Printf("服务器配置：%v\n", config)
}
