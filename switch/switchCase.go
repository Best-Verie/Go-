package main
import (
	"fmt"
	"time"
)

func main(){
	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Today is Weekday")
	default:
		fmt.Println("Today is weekend")
	}
}