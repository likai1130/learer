package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 自定义处理静态资源的路由处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取静态资源路径

		path := "./web" + r.URL.Path

		// 通过 http.ServeFile 函数加载静态资源文件
		http.ServeFile(w, r, path)
	})

	fmt.Println("访问 http://localhost:8080 来查看地图")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
