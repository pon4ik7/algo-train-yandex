package main

import (
	"bufio"
	"fmt"
	"os"
)

func Task2Solution(arr []int) (num1, num2 int) {
	maxPositive := 0
	maxNegative := 0

	maxValue := 0

	if len(arr) == 2 {
		num1 = arr[0]
		num2 = arr[1]

		if num1 <= num2 {
			return num1, num2
		}
		return num2, num1

	}

	for i := 0; i < len(arr); i++ {

		if arr[i]*maxNegative >= maxValue {
			maxValue = arr[i] * maxNegative
			num1 = arr[i]
			num2 = maxNegative
		}

		if arr[i]*maxPositive >= maxValue {
			maxValue = arr[i] * maxPositive
			num1 = arr[i]
			num2 = maxPositive
		}

		if arr[i] > maxPositive {
			maxPositive = arr[i]
		}
		if arr[i] < maxNegative {
			maxNegative = arr[i]
		}
	}
	if num1 <= num2 {
		return num1, num2
	}
	return num2, num1
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

	a, b := Task2Solution(arr)
	fmt.Println(a, b)

}
