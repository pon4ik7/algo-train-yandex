package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Prefix struct {
	prefSum   []int64
	maxSubSum int64
}

func NewPrefix(arr []int64) *Prefix {
	prefSum := make([]int64, len(arr)+1)
	prefSum[0] = 0
	for i := 0; i < len(arr); i++ {
		prefSum[i+1] = prefSum[i] + arr[i]
	}
	return &Prefix{
		prefSum:   prefSum,
		maxSubSum: math.MaxInt64,
	}
}

func (p *Prefix) GetMaxSubSum() int64 {
	if len(p.prefSum) <= 1 {
		p.maxSubSum = 0
		return 0
	}

	minPref := p.prefSum[0]
	best := p.prefSum[1] - p.prefSum[0]

	for i := 1; i < len(p.prefSum); i++ {
		cur := p.prefSum[i] - minPref
		if cur > best {
			best = cur
		}
		if p.prefSum[i] < minPref {
			minPref = p.prefSum[i]
		}
	}

	p.maxSubSum = best
	return best
}

func main() {

	in := bufio.NewReader(os.Stdin)

	var n int

	fmt.Fscan(in, &n)

	arr := make([]int64, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	prefix := NewPrefix(arr)

	fmt.Println(prefix.GetMaxSubSum())

}
