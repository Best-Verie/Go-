package main
import "fmt"
var a, isBoy, isGirl bool = true, false, true;

func main(){
	var i int = 10;
	fmt.Println(i,a,isBoy,isGirl);

	// short variable declarations

	k := 100;
	c,python,java := true,false,"stronger!";

	fmt.Println(k,c,python,java);

	// constants 

	// Constants cannot be declared using the := syntax.


	const Pi = 3.14

	const country = "RWANDA"

	fmt.Println("Hello",country);
	fmt.Println("Happy", Pi, "Day")


}