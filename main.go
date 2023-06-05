package main

import (
	"fmt"
	"github.com/walkjohnp/go-demo/global"
	"github.com/walkjohnp/go-demo/router"
	"net/http"
)

func main() {
	// 初始化
	global.Init()

	// 初始化路由
	rr := router.Init()

	server := &http.Server{
		Addr:    ":8080",
		Handler: rr.Handler(),
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("error" + err.Error())
	}
}
