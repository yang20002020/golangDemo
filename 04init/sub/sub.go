package sub
import"fmt"
func init(){
	fmt.Println("这是sub.go文件中的init函数1")
}
func init(){
	fmt.Println("这是sub.go文件中的init函数2")
}

func Sub(a ,b int) int{
	fmt.Println("这是sub.go文件中的sub函数")
	k:=test()
	fmt.Println("这是sub.go文件中的sub函数调用test",k)
	return a+b
}

