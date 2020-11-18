package main

import "fmt"

func RotateArray(data []int, k int) {
	n := len(data)
	ReverseArray(data, 0, k-1)
	ReverseArray(data, k, n-1)
	ReverseArray(data, 0, n-1)
}
func ReverseArray(data []int, start int, end int) {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Input array :", arr)
	RotateArray(arr, 2)
	fmt.Println("Rotated array :", arr)
}