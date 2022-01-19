package main

import (
	"go-restful/routers"
	"go-restful/routers/requestmethod"
)


func main() {
	// 设置需要加载的路由配置
	r := routers.Init(
		requestmethod.Router,
		//hello.Routers, // 后面还可以有多个
	)
	_ = r.Run(":8080")
}
