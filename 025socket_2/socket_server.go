package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//创建监听
	ip := "127.0.0.1"
	port := 8848

	address := fmt.Sprintf("%s:%d", ip, port)
	fmt.Println("address:", address)

	listener, err := net.Listen("tcp", address) //11111111111111111111111111
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	//需求：server可以接收多个链接==》主go程负责监听，子go程负责数据处理
	// 每个链接可以接收多次数据请求

	for {
		fmt.Println("监听中*****")
		Conn, err := listener.Accept() //222222222222222222222222
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		fmt.Println("链接建立成功!")
		go handleFunc(Conn)
	}

}

// 处理具体业务的逻辑，需要conn传递进来，每一新链接，conn是不同的
func handleFunc(Conn net.Conn) {
	//创建一个容器，用于接收读取到的数据
	buf := make([]byte, 1024) //使用make 创建切片 byte==》unit8
	cnt, err := Conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("Client ===》server，长度：", cnt, ",数据：", string(buf[0:cnt]))
	//服务器对客户端请求进行响应，将数据转成大写
	upperData := strings.ToUpper(string(buf[0:cnt]))
	//Write(b []byte) (n int, err error)
	cnt, err = Conn.Write([]byte(upperData))
	fmt.Println("Client <======Server,长度：", cnt, ",数据：", upperData)
	//关闭链接
	_ = Conn.Close()
}
