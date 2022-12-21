package  main

import (
	"fmt"
	"os"
)


func readfile(fileName string){
	//func Open(name string) (*File, error) {
	//	return OpenFile(name, O_RDONLY, 0)
	//}
	//go语言一般会将错误信息作为最后一个参数返回

	//err==nil 表示执行成功；err！=nil 表示有执行失败
	f1,err:=os.Open(fileName)
	//匿名函数,大括号后面的()是用来调用匿名函数
	defer func (){
		fmt.Println("准备关闭文件！！！")

		 f1.Close()
	}()


	if err!=nil{
		fmt.Println("打开文件失败！err:",err)

	}
	buf:=make([]byte,1024)
	//func (f *File) Read(b []byte) (n int, err error)
	n,_:=f1.Read(buf)
	fmt.Println("读取文件的实际长度：",n)
	fmt.Println("读取的文件内容：",string(buf))

}

func main(){
	fileName:="D:\\golandProject\\src\\00import\\main.go"
	readfile(fileName)
}


