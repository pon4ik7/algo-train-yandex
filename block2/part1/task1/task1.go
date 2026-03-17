package main

import (
	"bufio"
	"fmt"
	"os"
)

func countUnique(arr []int) int {

	set := make(map[int]bool)

	for _, v := range arr {
		set[v] = true
	}

	return len(set)

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

	fmt.Println(countUnique(arr))

}
