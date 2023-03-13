package main

import (
	"fmt"
	_ "golangDemo/05init1/sub" //只想使用这个包里面的init函数，不使用其他函数，加下划线
)

func main() {
	fmt.Println("这是main函数")
	//r:=sub.Sub(10,6)
	fmt.Println("这是main 函数，没有调用sub文件中之后")
}
