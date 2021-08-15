//An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

//The type []T is a slice with elements of type T

package main

import "fmt"

func main(){
	s:= make([]string, 3)
	fmt.Println("empty: ", s)

	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
	fmt.Println("new values: " , s)
	fmt.Println("slice at index 2 : ", s[2] )

	// yet they are dynamic we can append

	s = append(s, "d")
	fmt.Println("appended: ", s)

	

}