package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a number of brackets:")
	brackets, _, _ := reader.ReadLine()

	bracketNumber, _ := strconv.ParseInt(string(brackets), 10, 64)

	result := generateBracket("", bracketNumber, 0, 0)
	fmt.Println(result)
}

func generateBracket(brackets string, bracketNumber int64, left, right int64) string {
	if right == bracketNumber {
		return brackets
	}

	if left < bracketNumber {
		return generateBracket(brackets+"(", bracketNumber, left+1, right)
	}

	if right < bracketNumber {
		return generateBracket(brackets+")", bracketNumber, left, right+1)
	}
	return brackets
}
