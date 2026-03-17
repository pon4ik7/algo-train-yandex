package main

import (
	"bufio"
	"fmt"
	"os"
)

func solveEquation(a, b, c int) (int, error) {

	if c < 0 {
		return -1, fmt.Errorf("NO SOLUTION")
	}

	partAns := c*c - b

	if a == 0 {
		if partAns != 0 {
			return -1, fmt.Errorf("NO SOLUTION")
		}
		return -1, fmt.Errorf("MANY SOLUTIONS")
	}

	if partAns%a != 0 {
		return -1, fmt.Errorf("NO SOLUTION")
	}

	return partAns / a, nil

}

func main() {

	in := bufio.NewReader(os.Stdin)

	var a, b, c int

	fmt.Fscan(in, &a, &b, &c)

	ans, err := solveEquation(a, b, c)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ans)

}
