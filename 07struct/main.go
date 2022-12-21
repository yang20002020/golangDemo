package  main

import "fmt"

type Person struct {
	name string
	age  int
	gender string
}

func (this Person) Eat1(){
	fmt.Println("person is eating")
}
func (this *Person) Eat2(){
fmt.Println(this.name+" is eating")
}
func (this Person) Eat3(){
	fmt.Println("person is eating")
	this.name="Duke"
}
func (this *Person) Eat4(){
	this.name="Duke"
	fmt.Println(this.name+" is eating")
}
func main(){
	lily:=Person{
		name:   "lily",
		age:    20,
		gender: "女生",
	}
	fmt.Println("lily:",lily)
	lily.Eat1()  //person is eating
	lily.Eat2()  //lily is eating
	lily.Eat3()  //person is eating
	lily.Eat4()  //Duke is eating
	
}

