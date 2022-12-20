package  main
 import "fmt"
 func main (){

 	label:
 		for i:=0;i<5;i++{
 			for j:=0;j<5;j++{
 				if j==3{
 					break label
				}
				fmt.Println("i=",i,"j=",j)
			}

		}
	 fmt.Println("over")

 }

