package main

import "fmt"

// 定义一个接口，注意类型是interface
type IAttack interface {
	//接口的函数有多个，但是只能有函数声明，不能有函数实现
	Attack()
}

// 结构体
type HumanLevel struct {
	name  string
	level int
}

func (this *HumanLevel) Attack() {
	fmt.Println("我是：", this.name, "等级为：", this.level)
}
func main() {
	Lowlevel := HumanLevel{
		name:  "David",
		level: 1,
	}
	Lowlevel.Attack()
	Highlevel := HumanLevel{
		name:  "jack",
		level: 10,
	}
	Highlevel.Attack()
	/************************/
	var player IAttack //定义一个含有Attack函数的接口变量
	player = &Lowlevel //接口需要地址指针来赋值
	player.Attack()
	/***********************/

	player = &Highlevel //接口需要地址指针来赋值
	player.Attack()
}
