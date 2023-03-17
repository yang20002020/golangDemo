package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name    string `json:"-"`            //在使用json编码的时候，这个编码不参与
	Subject string `json:"Subject_name"` //json编码时，这个字段会编码成Subject_name
	Age     int    `json:"age,string"`   //json 编码的时候，将age转成string类型；两个字段之间必须用逗号隔开，不能有空格
	gender  string
	Adress  string `json:"adress,omitempty"` //在json编码的时候，如果字段是空的，就不参与编码
}

func main() {
	t1 := Teacher{
		Name:    "Duke",
		Subject: "Golang",
		Age:     0,
		gender:  "Man",
	}
	//func Marshal(v interface{}) ([]byte, error)
	information, err := json.Marshal(&t1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("information:", string(information))
}
