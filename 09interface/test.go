package main

import "fmt"

func main() {
	fmt.Println()
	var i, j, k interface{}
	name := []string{"duke", "lily"}
	i = name
	fmt.Println("i代表切片：", i)
	age := 20
	j = age
	fmt.Println("j代表整数数字：", j)
	str := "hello"
	k = str
	fmt.Println("k代表字符串：", k)

	_, ok := k.(int)
	if ok {
		fmt.Println("k是int")
	} else {
		fmt.Println("k不是int")
	}

}
