package main

import (
	"fmt"
	"golangDemo/04init/sub"
)

func main() {
	fmt.Println("这是main 函数")
	r := sub.Sub(10, 6)
	fmt.Println("这是main 函数，调用的其他文件中函数sub：", r)
}
