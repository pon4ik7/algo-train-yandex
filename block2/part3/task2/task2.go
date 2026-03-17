package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pair struct {
	item   string
	number int
}

func NewPair(item string, number int) Pair {
	return Pair{item: item, number: number}
}

func main() {

	in := bufio.NewReader(os.Stdin)

	set := make(map[string]map[string]int)

	var name, item string
	var number int

	for {
		_, err := fmt.Fscan(in, &name, &item, &number)
		if err != nil {
			break
		}
		if _, ok := set[name]; !ok {
			set[name] = make(map[string]int)
		}
		if _, ok := set[name][item]; !ok {
			set[name][item] = number
		} else {
			set[name][item] += number
		}

	}

	names := make([]string, 0)

	for k, _ := range set {
		names = append(names, k)
	}

	sort.Strings(names)

	for _, v := range names {
		fmt.Println(v + ":")

		items := make([]Pair, 0)

		for k1, v1 := range set[v] {
			items = append(items, NewPair(k1, v1))
		}

		sort.Slice(items, func(i, j int) bool {
			return items[i].item < items[j].item
		})

		for _, v1 := range items {
			fmt.Println(v1.item, v1.number)
		}

	}

}
