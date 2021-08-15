package main

import ("fmt")
func main(){
	var fnumber int = 5;
	var snumber int = 10;

	if snumber%fnumber == 0{
		fmt.Println(snumber, "is divisible by", fnumber);
	}else{
		fmt.Println(snumber, "is not divisible by", fnumber);
	}
	
	if snumber%2 == 0 {
		fmt.Println(snumber, "is even");
	}else{
		fmt.Println(snumber, "is odd");
	}
	if fnumber:=5; fnumber < 0{
		fmt.Println(fnumber, "is negative");
	} else if fnumber > 0 {
		fmt.Println(fnumber, "is positive");
	}else{
		fmt.Println(fnumber, "is zero");
	}
}


// theres no ternary with Go