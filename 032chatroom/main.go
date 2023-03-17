package main

import (
	"fmt"
	"net"
)

type User struct {
	//名字
	name string
	//唯一的id
	id string
	//管道
	msg chan string
}

// 创建一个全局的map结构，用于保存所有的用户
var allUsers = make(map[string]User)

// 定义一个message全局通道，用于接收任何人发送过来消息
var message = make(chan string, 10)

func main() {
	//1. 创建服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//启动全局唯一的go程，负责监听message通道，给所有的用户
	go broadcast()
	fmt.Println("服务器启动成功!")
	for {

		fmt.Println("主go程监听中.....")
		//2. 监听
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		fmt.Println("建立链接成功！")
		//建立链接
		go handler(conn)
	}

	// q启动处理业务的go程
}

// 处理具体业务
func handler(conn net.Conn) {

	fmt.Println("启动业务...")
	//客户端与服务器建立链接的时候，会有IP和port==>当成user的id
	clientAddr := conn.RemoteAddr().String()
	fmt.Println("clientAdrr:", clientAddr)
	//创建user
	newUser := User{
		name: clientAddr,              //id,我们不会修改，这个作为在map中的key
		id:   clientAddr,              //可以修改，会提供rename命令修改，建立链接时，初试值与id相同
		msg:  make(chan string, 1024), //注意需要make空间，否则无法写入数据
	}
	//添加user到map结构
	allUsers[newUser.id] = newUser

	//启动go程，负责将msg信息返回给客户端
	go writeBackToClient(&newUser, conn)

	//向message写入数据，当前用户上线的消息，用于通知所有人（广播）
	loginInfo := fmt.Sprintf("[%s]:[%s]===>上线了login!!", newUser.id, newUser.name)
	message <- loginInfo
	//TODO 代表这里以后再具体实现，当前保留

	for {
		//具体业务逻辑
		buf := make([]byte, 1024)
		//读取客户端发送过来的请求数据)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("服务器接收客户端发送过来的数据为：", string(buf[:cnt-1]), ",cnt:", cnt)
		//~~~~~~~~~~~~业务逻辑处理 开始~~~~~~~~~~~~~~~~~~~~~~
		//1. 查询当前所有的用户，who
		//a.先判断 接收的数据是不是who==》长度，字符串
		userInput := buf[:cnt-1] //这是用户输入的数据，最后一个是回车，需要去掉
		if len(userInput) == 3 && string(userInput) == "who" {
			//b. 遍历allUsers这个map：key：userid value ：user本身 将id和name拼接成一个字符串，返回给客户端
			fmt.Println("用户即将查询所有用户信息!")
		} else {
			//若果用户输入的不是命令，只是普通的聊天信息，那么只需要写到广播通道中即可，由其他的go程进行常规转发
			message <- string(userInput)
		}

		//~~~~~~~~~~~~业务逻辑处理 结束~~~~~~~~~~~~~~~~~~~~~~

	}
}

// 向所有的用户广播消息，启动一个全局唯一go程
func broadcast() {
	fmt.Println("广播go程启动成功...")
	defer fmt.Println("broadcast程序退出！")
	for {
		//1.从message中读取数据
		fmt.Println("broadcast监听message中...")
		info := <-message
		fmt.Println("message接收到消息：", info)
		//2.将数据写入到每一个用户的msg管道中
		for _, user := range allUsers {
			//如果msg是非缓冲的，那么会在这里阻塞
			user.msg <- info
		}

	}
}

// 每一个用户应该还有一个用来监听自己msg管道的go程，负责将数据返回给客户端
func writeBackToClient(user *User, conn net.Conn) {
	//TODO
	fmt.Printf("user:%s 的go程正在监听自己的msg管道:", user.name)
	for data := range user.msg {
		fmt.Printf("user:%s写回给客户端数据为:%s\n", user.name, data)

		//write(b []byte)(n int,err error)
		_, _ = conn.Write([]byte(data))
	}

}
