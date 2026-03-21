package main

import "fmt"

func solve(a, b int64) int64 {
	var ans int64 = 0
	for a > b {
		a = (a + 1) / 2
		ans++
	}
	return ans
}

func main() {
	var hight, lenght, hPoc, wPoc int64
	fmt.Scan(&hight, &lenght, &hPoc, &wPoc)

	ans1 := solve(hight, hPoc) + solve(lenght, wPoc)
	ans2 := solve(hight, wPoc) + solve(lenght, hPoc)

	if ans1 < ans2 {
		fmt.Println(ans1)
	} else {
		fmt.Println(ans2)
	}
}
