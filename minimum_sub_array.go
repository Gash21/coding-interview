package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	target, _, _ := reader.ReadLine()
	targetNumber, _ := strconv.Atoi(string(target))

	listNumber, _, _ := reader.ReadLine()
	stringLists := strings.Split(string(listNumber), " ")

	var listOfNumbers []int
	for i := 0; i < len(stringLists); i++ {
		numb, _ := strconv.Atoi(stringLists[i])
		listOfNumbers = append(listOfNumbers, numb)
	}

	fmt.Println(minSubArrayLen(int(targetNumber), listOfNumbers))
}

func minSubArrayLen(target int, nums []int) int {
	minLength := math.MaxInt32
	sum := 0
	j := 0
	fmt.Println(target, nums)

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			fmt.Println(minLength, i, j)
			minLength = Minimal(minLength, i-j+1)
			sum -= nums[j]
			j++
		}

	}

	if minLength == math.MaxInt32 {
		return 0
	}

	return minLength
}

func Minimal(a, b int) int {
	if a > b {
		return b
	}
	return a
}
