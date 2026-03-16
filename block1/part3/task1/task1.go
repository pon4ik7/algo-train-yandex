package main

import (
	"bufio"
	"fmt"
	"os"
)

func task1Solution(arr []int) int {

	helpArr := make([]int, len(arr))

	maxLen := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			if maxLen == -1 {
				helpArr[i] = 1e6
			} else {
				helpArr[i] = i - maxLen
			}
		}
		if arr[i] == 2 {
			maxLen = i
		}
	}

	maxLen = -1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 1 {
			if maxLen != -1 {
				helpArr[i] = min(maxLen-i, helpArr[i])
			}
		}
		if arr[i] == 2 {
			maxLen = i
		}
	}

	ans := 0

	for i := 0; i < len(helpArr); i++ {
		if helpArr[i] != 1e6 && helpArr[i] != 0 {
			ans = max(ans, helpArr[i])
		}
	}
	return ans
}

func main() {

	in := bufio.NewReader(os.Stdin)

	var arr []int
	for {
		var x int
		_, err := fmt.Fscan(in, &x)
		if err != nil {
			break
		}
		arr = append(arr, x)
	}

	fmt.Println(task1Solution(arr))

}
