package main

import (
	"fmt"
	"time"
)

func main() {
	//当涉及到多线程时，C语言使用互斥量，上锁来保持资源，避免资源竞争问题
	//go语言也支持这种方式，但是GO语言更好的解决方案，是使用通道
	// 使用通道不需要我们去进行加解锁
	//A 通道里面写数据 B从通道里读取数据  go 自动帮我们做好了数据同步

	//创建管道 ： 创建一个装数字的管道
	//make没有创建空间大小，此时是没有缓存的
	numChan := make(chan int) //使用管道一定要用make，同map一样
	//numChan1:=make(chan string)  //创建一个装字符串的管道
	go func() {
		for i := 0; i < 50; i++ {
			data := <-numChan
			fmt.Println("这是子go程1，读取数据data:", data)
		}
	}()
	go func() {
		for data2 := 20; data2 < 50; data2++ {
			numChan <- data2
			fmt.Println("《《《《《这是子go程2，写入数据:", data2)
		}
	}()
	for i := 0; i < 20; i++ {
		//向管道写数据
		numChan <- i
		fmt.Println("《《《《《《这是主go程,写入数据:", i)

	}
	//注意 写入和读取并不是一一对应的
	time.Sleep(5 * time.Second)
}
