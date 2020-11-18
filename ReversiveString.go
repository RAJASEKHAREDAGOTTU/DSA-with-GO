package main

import (
	"fmt"
	"strings"
)

//Approach: 1
func reverseWord(word string) string {
	var wtr string
	for _, c := range word {
		wtr = string(c) + wtr
	}

	return wtr
}

//Approach: 2
func Reverse(word string) string {
	var n string
	for i := len(word) - 1; i >= 0; i-- {
		n = n + string(word[i])
	}
	return n
}

//Approach: 3
func ReverseW(word string) string {
	var sb strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}
	return sb.String()
}
func main() {
	fmt.Println(reverseWord("manmohan"))

	fmt.Println(Reverse("abdfgh"))

	fmt.Println(ReverseW("Yeah!"))
}