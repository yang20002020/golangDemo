package main

import (
	"fmt"
	"net"
	"strings"
)

// 只接收一个链接，只能发送一次数据
func main() {
	//创建监听
	ip := "127.0.0.1"
	port := 8848
	//func Sprintf(format string, a ...interface{}) string
	address := fmt.Sprintf("%s:%d", ip, port)
	fmt.Println("address:", address)
	//func Listen(network, address string) (Listener, error)
	listener, err := net.Listen("tcp", address) //11111111111111111111111111 设置服务器
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("监听中*****")
	//Accept() (Conn, error)
	//err 是第一个err,并不是新的变量
	Conn, err := listener.Accept() //222222222222222222222222 链接
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	fmt.Println("链接建立成功!")
	/************************************************************************/
	//创建一个容器，用于接收读取到的数据
	buf := make([]byte, 1024) //使用make 创建切片 byte==》unit8

	//Read(b []byte) (n int, err error)
	//cnt:真正读取client发来的数据的长度
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
	Conn.Close()

}
