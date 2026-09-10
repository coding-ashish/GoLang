package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("The current time is : ")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01 - 02 - 2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2028, time.April, 23, 23, 23, 0, 0, time.UTC)
	// createdDate.Format("01-02-2006 Monday")

	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))
}
