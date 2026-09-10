package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the class of slices.")

	var fruitList = []string{"Mango", "Banana", "Orange", "Guava", "Grapes"}
	fmt.Printf("The type of fruitList is : %T\n", fruitList) // prints []string which is known as slices

	//the interesting part about slices is it is like vector : we can change its size and append element and delete element from it.

	fruitList = append(fruitList, "Aam", "Kela")
	fmt.Println("The updated array is : ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("fruitList is : ", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println("FruitList is = ", fruitList)

	score := make([]int, 4) // another syntax to declare a slice

	score[0] = 10
	score[1] = 6
	score[2] = 5
	score[3] = 8
	// score[4] = 5  // this line of code will throw an error

	score = append(score, 5, 6)
	fmt.Println(score)

	//Now there are sort functions also in the gp

	sort.Ints(score)
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))// gives true in this case
}
