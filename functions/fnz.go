package main
import "fmt"

func addMe(a int, b int){
	return a+b
}

func main(){
	res:= addMe(4,5)
	fmt.Println(res)

}