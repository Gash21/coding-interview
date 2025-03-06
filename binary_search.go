package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// TreeNode represents a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST converts a sorted array into a balanced BST
func sortedArrayToBST(nums []int) *TreeNode {
	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left >= right {
			return nil
		}
		mid := (left + right) / 2
		return &TreeNode{
			Val:   nums[mid],
			Left:  build(left, mid),
			Right: build(mid+1, right),
		}
	}
	return build(0, len(nums))
}

// printTree recursively prints the tree structure
func printTree(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}
	if isLeft {
		fmt.Printf("%s├── %d\n", prefix, root.Val)
	} else {
		fmt.Printf("%s└── %d\n", prefix, root.Val)
	}
	newPrefix := prefix
	if isLeft {
		newPrefix += "│   "
	} else {
		newPrefix += "    "
	}
	printTree(root.Left, newPrefix, true)
	printTree(root.Right, newPrefix, false)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	numbers, _, _ := reader.ReadLine()

	stringLists := strings.Split(string(numbers), " ")

	var listOfNumbers []int
	for i := 0; i < len(stringLists); i++ {
		numb, err := strconv.Atoi(stringLists[i])
		if err != nil {
			fmt.Println("input is not a number")
			return
		}
		listOfNumbers = append(listOfNumbers, numb)
	}

	nums := listOfNumbers
	sort.Ints(nums) // Ensure the input is sorted
	root := sortedArrayToBST(nums)
	printTree(root, "", false)
}
