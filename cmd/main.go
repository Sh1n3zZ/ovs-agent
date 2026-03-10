package main

import (
	"os"
)

func main() {
	// 如果没有参数，运行 HTTP 服务器
	// If no arguments, run HTTP server
	if len(os.Args) == 1 {
		runServer(false)
		return
	}

	// 否则执行 cobra 命令
	// Otherwise execute cobra commands
	Execute()
}
