package  main
import "fmt"
type Student struct{
	name string
	age int
	gender string
	score float64
}
func main(){

	lily:=Student{
		name:   "lily",
		age:    20,
		gender: "女生",
		score:  80,
	}
	fmt.Println("lily:",lily.name,lily.age,lily.gender,lily.score)
   ptr:=&lily
	fmt.Println("lily:",(*ptr).name,(*ptr).age,(*ptr).gender,(*ptr).score)

   hanmeimei:=Student{
	   name:   "hanmeimei",
	   age:    20,
   }
   fmt.Println("hanmeimei：",hanmeimei)
}

