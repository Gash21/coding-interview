package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a string:")
	str, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	fmt.Println(IsPalindrome(string(str)))
}

func IsPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
