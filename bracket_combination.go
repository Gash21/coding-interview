package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
* Have the function BracketCombinations(num) read num which will be an integer greater than or equal to zero,
* and return the number of valid combinations that can be formed with num pairs of parentheses. For example,
* if the input is 3, then the possible combinations of 3 pairs of parenthesis,
* namely: ()()(), are ()()(), ()(()), (())(), ((())), and (()()).
* There are 5 total combinations when the input is 3, so your program should return 5.
 */

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a number of brackets:")
	brackets, _, _ := reader.ReadLine()

	bracketNumber, _ := strconv.ParseInt(string(brackets), 10, 64)
	fmt.Println(solution(bracketNumber))
}

func solution(brackets int64) int64 {
	count := int64(0)

	dfs(0, 0, brackets, &count)
	return count
}

func dfs(left, right, brackets int64, count *int64) {
	if left < right {
		return
	}

	if left == brackets && right == brackets {
		*count++
		return
	}

	if left < brackets {
		dfs(left+1, right, brackets, count)
	}
	if right < brackets {
		dfs(left, right+1, brackets, count)
	}
}
