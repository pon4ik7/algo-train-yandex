package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)

	var n int

	fmt.Fscan(in, &n)

	set := make(map[string]string)

	var first, second string

	for i := 0; i < n; i++ {

		fmt.Fscan(in, &first, &second)

		set[first] = second
		set[second] = first

	}

	fmt.Fscan(in, &first)

	fmt.Println(set[first])

}
