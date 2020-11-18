package main

import "fmt"

func SmallestPositiveMissingNumber(arr []int, size int) int {
	found := 0
	for i := 1; i < size+1; i++ {
		found = 0
		for j := 0; j < size; j++ {
			if arr[j] == i {
				found = 1
				break
			}
		}
		if found == 0 {
			return i
		}
	}
	return -1
}

func SmallestPositiveMissingNumber2(arr []int, size int) int {
	hs := make(map[int]int)
	for i := 0; i < size; i++ {
		hs[arr[i]] = 1
	}
	for i := 1; i < size+1; i++ {
		_, ok := hs[i]
		if ok == false {
			return i
		}
	}
	return -1
}

func SmallestPositiveMissingNumber3(arr []int, size int) int {
	aux := make([]int, size)

	for index, _ := range aux {
		aux[index] = -1
	}

	for i := 0; i < size; i++ {
		if arr[i] > 0 && arr[i] <= size {
			aux[arr[i]-1] = arr[i]
		}
	}
	for i := 0; i < size; i++ {
		if aux[i] != i+1 {
			return i + 1
		}
	}
	return -1
}

func SmallestPositiveMissingNumber4(arr []int, size int) int {
	temp := 0
	for i := 0; i < size; i++ {
		for arr[i] != i+1 && arr[i] > 0 && arr[i] <= size {
			temp = arr[i]
			arr[i] = arr[temp-1]
			arr[temp-1] = temp
		}
	}
	for i := 0; i < size; i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}
	return -1
}

func main() {
	arr := []int{8, 5, 6, 1, 9, 11, 2, 7, 3, 10}
	size := len(arr)

	fmt.Println("Missing Number :", SmallestPositiveMissingNumber(arr, size))
	fmt.Println("Missing Number :", SmallestPositiveMissingNumber2(arr, size))
	fmt.Println("Missing Number :", SmallestPositiveMissingNumber3(arr, size))
	fmt.Println("Missing Number :", SmallestPositiveMissingNumber4(arr, size))
}
