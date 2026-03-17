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

	var m int

	setAny := make(map[string]bool)

	setAll := make(map[string]bool)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &m)
		var name string
		tempSet := make(map[string]bool)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &name)
			setAny[name] = true
			tempSet[name] = true
		}

		if i == 0 {
			setAll = tempSet
		} else {
			for name := range setAll {
				if !tempSet[name] {
					delete(setAll, name)
				}
			}
		}
	}

	fmt.Println(len(setAll))
	for name := range setAll {
		fmt.Println(name)
	}

	fmt.Println(len(setAny))
	for name := range setAny {
		fmt.Println(name)
	}

}
