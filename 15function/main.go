package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(adder(a, b))

	proRes, message := solve(4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is :", proRes)
	fmt.Println("The sum is :", message)

}

// Writing a function to print the values from 1 to n
func adder(a int, b int) int {
	return a + b
}

//The function definition when we dont know how many arguements are there to come to function (called a Pro-Function)
func solve(values ...int) (int, string) {
	total := 0
	for i := range values {
		total += values[i]
	}
	return total, "Hello Everyone, My Name is Garvit Sharma"
}
