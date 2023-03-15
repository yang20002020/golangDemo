package main

import "fmt"

func main() {
	numChan := make(chan int, 10)

	//写
	go func() {
		for i := 0; i < 10; i++ {
			numChan <- i
			fmt.Println("写入数据：", i)
		}
		close(numChan)
	}()
	//读数据
	//for value := range numChan{
	//  fmt.Println("v:",value)
	//}

	for {
		v, ok := <-numChan
		if !ok {
			fmt.Println("管道已经关闭了，准备退出")
			break
		}
		fmt.Println("v:", v)
	}
	fmt.Println("over!")
}
