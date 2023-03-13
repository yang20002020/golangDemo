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

// 定义一个多态的通用接口，传入不同的对象，调用同样的函数，实现不同效果》》》多态
func DoAttack(a IAttack) {
	a.Attack()
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

	fmt.Println(">>>>>>>多态>>>>>>>>")
	DoAttack(&Lowlevel)
	DoAttack(&Highlevel)
}
