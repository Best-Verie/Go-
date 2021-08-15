//Go has only one looping construct, the for loop.

package main
import "fmt"
func main(){
	sum:= 0 
	for i:=0; i< 10; i++ {
		sum += i
	}
	fmt.Println("The sum is " , sum)

	//infinite

	// for{}

	//

	// for as go's while
	total := 1
	for total < 1000 {
		total += total
	}
	fmt.Println(total)

}