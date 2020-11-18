package main

import "fmt"

func Sum_list(lst []int) int {
	var sum int
	sum = 0
	for _, val := range lst {
		sum = sum + val
	}
	return sum
}

func Sum_num(lst []int) int {
	if len(lst) == 0 {
		return 0
	}
	return lst[0] + Sum_num(lst[1:])
}
func main() {
	lst := []int{67, 45, 3, 2, 1, 89, 6, 7, 65}
	fmt.Println("Sum using iteration:", Sum_list(lst))
	fmt.Println("Sum of num of list using recursion:", Sum_num(lst))
}