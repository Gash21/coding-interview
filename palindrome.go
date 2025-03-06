package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Enter a string:")
	str, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	alphanumeric := regexp.MustCompile("[^a-zA-Z0-9]")
	strgs := alphanumeric.ReplaceAllString(string(str), "")
	fmt.Println(IsPalindrome(strings.ToLower(string(strgs))))
}

func IsPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
