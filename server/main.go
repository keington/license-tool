package main

import (
	"net/http"
	"server/config"
	"server/router"
)

func main() {

	port := config.GetConfig("server", "port")

	r := router.SetupRouter()
	if r == nil {
		// 路由器配置失败，无法启动服务器
		return
	}

	// 启动服务器
	if err := http.ListenAndServe(":"+port, r); err != nil {
		panic(err)
	}
}
