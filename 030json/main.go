package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Name   string
	Age    int
	gender string //小写字符开始的字段，在json编码的时候会忽略掉数值
}

func main() {
	Lily := Student{
		Id:     1,
		Name:   "Lily",
		Age:    10,
		gender: "女士",
	}

	//编码
	//func Marshal(v interface{}) ([]byte, error)
	encodeInfo, err := json.Marshal(&Lily)
	if err != nil {
		fmt.Println("json.marshal err", err)
		return
	}
	fmt.Println("encodeInfor:", string(encodeInfo))
	////////////////////////////////////
	//解码  字符串转换成结构体
	var stu Student
	//func Unmarshal(data []byte, v interface{}) error
	if json.Unmarshal([]byte(encodeInfo), &stu); err != nil {
		fmt.Println("Json.Unmarshal err:", err)
		return
	}
	fmt.Println("Id:", stu.Id)
	fmt.Println("Name:", stu.Name)
	fmt.Println("Age:", stu.Age)
	fmt.Println("gender:", stu.gender)
}
