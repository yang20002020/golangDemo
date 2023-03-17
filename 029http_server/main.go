package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	/*******************************************************************/
	//注册路由 router
	//xxxx/user ===>func1
	//XXXX/name===>func2
	//xxxx/id====>func3

	//http://127.0.0.1:8080/user,func是回调函数，用于路由的相应；这个回调函数是固定的
	//第一个参数对应路由 ，第二个参数对应回调函数
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {

		//request:==>请求 客户端发来的数据
		fmt.Println("用户请求详细:")
		fmt.Println("客户端发过来的数据 request:", request)

		//write:==>通过write 将数据写给客户端  服务端到客户端
		_, _ = io.WriteString(writer, "这是/user 请求返回的数据！")
	})

	//http://127.0.0.1:8080/name
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/name 请求返回的数据！")
	})
	//http://127.0.0.1:8080/id
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/id 请求返回的数据！")
	})

	/***************************************************************/
	fmt.Println("http server start..........")
	//func ListenAndServe(addr string, handler Handler) error
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("http start failed,err", err)
		return
	}

}
