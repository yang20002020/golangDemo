package main

import (
	"fmt"
	"time"
)

// 这个将用于子go程使用
func display() {
	count := 1
	for {
		fmt.Println("==============》这是子go程：", count)
		count++
		time.Sleep(1 * time.Second)
	}
}
func main() {
	//子go程
	go display()
	//主go程
	count := 1
	for {
		fmt.Println("这是主go程：", count)
		count++
		time.Sleep(1 * time.Second)
	}

}
