package main

import (
	"fmt"
	"time"
)

/*
主动通知停止 go程
使用select +chan
*/

func main() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")
	stop <- true
	time.Sleep(5 * time.Second)
}
