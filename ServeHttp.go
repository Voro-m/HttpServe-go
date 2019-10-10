package main

import (
	"fmt"
	_ "fmt"
	"go/types"
	"net/http"
)
import _ "net/http"

type MyHandler struct{

}

//实现接口的serverHttp方法
func (this *MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Handler`s ServeHTTP create server")
}

func main(){
	MyHandler := &MyHandler{}
	http.Handle("/",MyHandler)
	http.ListenAndServe(":8888",nil)
}
