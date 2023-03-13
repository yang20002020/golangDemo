package main

import "fmt"

type Human struct {
	name   string
	age    int
	gender string
}

func (this *Human) Eating() {
	fmt.Println("this is :" + this.name)
}

type Student struct {
	hum    Human // 在结构体中嵌套结构体
	school string
}
type Teacher struct {
	Human   // 结构体继承
	subject string
}

func main() {
	s1 := Student{
		hum: Human{
			name:   "yang",
			age:    17,
			gender: "男生",
		},
		school: "藁城一中",
	}
	fmt.Println("s1=", s1)
	fmt.Println("s1.hum.name:", s1.hum.name)
	fmt.Println("s1.School:", s1.school)
	/***************************************/
	s2 := Teacher{
		Human: Human{
			name:   "陈",
			age:    32,
			gender: "男",
		},
		subject: "英语",
	}

	fmt.Println("s2=", s2)
	fmt.Println("s2.name:", s2.name)
	fmt.Println("s2.School:", s2.subject)
	/***************************************/
	s3 := Teacher{}
	s3.name = "su"
	s3.subject = "语文"
	s3.gender = "女"
	fmt.Println("s3=", s3)
	fmt.Println("s3.name:", s3.name)
	fmt.Println("s3.human.name:", s3.Human.name) //为了防止父类和子类出现同名的字段
	fmt.Println("s3.School:", s3.subject)
	s3.Eating()
}
