package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan int)
	chan2 := make(chan int)
	//启动一个go程，负责监听两个channel
	go func() {
		for {
			fmt.Println("监听中**************")
			select {
			case data1 := <-chan1:
				fmt.Println("从chan1读取数据成功，data1：", data1)
			case data2 := <-chan2:
				fmt.Println("*************从chan2读取数据成功，data2：", data2)
			default:
				fmt.Println("select default 分支 called！")
				time.Sleep(time.Second)
			}

		}
	}()

	//启动go1，写chan1
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("chan1写入：", i)
			chan1 <- i
		}
		time.Sleep(1 * time.Second / 2)
	}()
	//启动 go2 写chan2
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("chan2写入：", i)
			chan2 <- i
		}
		time.Sleep(1 * time.Second)
	}()

	for {
		fmt.Println("OVER")
		time.Sleep(5 * time.Second)
	}

}
