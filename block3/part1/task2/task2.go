package main

import (
	"bufio"
	"fmt"
	"os"
)

type Randomizer struct {
	curr int64
}

func NewRandomizer(x int64) *Randomizer {
	return &Randomizer{x}
}

func (r *Randomizer) Next() int64 {
	prev := r.curr
	r.curr = (11173*r.curr + 1) % 1000000007
	return prev
}

type Ans struct {
	result int64
}

func NewAns() *Ans {
	return &Ans{0}
}

func (a *Ans) Add(x int64) {
	a.result = (a.result + x) % 1000000007
}

func (a *Ans) Get() int64 {
	return a.result
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int

	fmt.Fscan(in, &n)

	arr := make([]int64, n)

	prefix := make([]int64, n+1)
	prefix[0] = 0

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		prefix[i+1] = prefix[i] + arr[i]
	}

	var m int
	fmt.Fscan(in, &m)

	var x0 int64
	fmt.Fscan(in, &x0)
	rand := NewRandomizer(x0)
	ans := NewAns()

	for i := 0; i < m; i++ {
		x1 := rand.Next()
		x2 := rand.Next()
		left := min(x1%int64(n), x2%int64(n))
		right := max(x1%int64(n), x2%int64(n))
		curr := prefix[right+1] - prefix[left]
		ans.Add(curr)
	}
	fmt.Println(ans.Get())
}
