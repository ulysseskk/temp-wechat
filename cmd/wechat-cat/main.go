package main

import "gitlab.ulyssesk.top/other/wechat-cat/internal/router"

func main() {
	r := router.InitRouter()
	err := r.Run(":19919")
	if err != nil {
		panic(err)
	}

}
