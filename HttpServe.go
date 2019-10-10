package main

import (
	"fmt"
	_ "fmt"
	"net/http"
)
import _ "net/http"

func main()  {
	//处理路由为 / 的方法
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintln(w,"Hello Go",r,url.Path)
	})
	//监听3000端口
	http.ListenAndServe(":3000",nil)
}

