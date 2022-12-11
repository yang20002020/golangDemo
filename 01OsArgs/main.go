package main

import (
	"fmt"
	"os"
)

func main() {

	//用go build  执行  例如 go build main.go
	//终端输入 ./main.exe  hello  world
	cmds := os.Args
	for key, value := range cmds {
		fmt.Println("key=", key, "value=", value, "len=", len(cmds))
	}
	fmt.Println()

	if len(cmds) < 2 {
		fmt.Println("请输入足够参数！！")
	}
	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default")
	}

}
