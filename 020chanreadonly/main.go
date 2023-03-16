package main

import (
	"fmt"
	"time"
)

func main() {
	//1.在主函数创建一个双向通道 numChan
	numChan := make(chan int, 5)
	//双向管道可以赋值给单向管道，单向通道不能转向双向通道
	//将numChan 传递给producer，负责生产
	go producer(numChan)

	//将numChan 传递给consumer 负责消费
	go consumer(numChan)
	time.Sleep(2 * time.Second)
}

//producer 生产者 ===>提供一个只写的通道
//out 为只写数据 ，不能读
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i //out 只能写
		fmt.Println("向管道中写入数据：", i)

	}
}

//consumer 消费者 ==>提供一个只读的通道
func consumer(in <-chan int) {
	for v := range in {
		fmt.Println("从管道读取数据：", v)
	}

}
