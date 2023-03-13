package main

import "fmt"

func main() {

	array := make([]interface{}, 3)
	array[0] = 1
	array[1] = "hello world"
	array[2] = true

	for _, value := range array {
		//可以获取当前接口的真正数据类型
		switch v := value.(type) {
		case int:
			fmt.Printf("这个类型是int,内容为：%d\n", value)
		case bool:
			fmt.Printf("这个类型是bool类型,内容为：%v\n", v)
		case string:
			fmt.Printf("这个类型是string类型,内容为：%s\n", v)
		}

	}

	fmt.Println()
}
