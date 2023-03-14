package main

import (
	"fmt"
	"time"
)

func main() {
	//子go程
	//采用匿名函数
	go func() {
		count := 1
		for {
			fmt.Println("==============》这是子go程：", count)
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	//主go程
	count := 1
	for {
		fmt.Println("这是主go程：", count)
		count++
		time.Sleep(1 * time.Second)
	}

}
