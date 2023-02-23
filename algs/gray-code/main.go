package main

import "fmt"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	prevSeq := grayCode(n - 1)
	mask := 1 << (n - 1)
	seqLen := len(prevSeq)
	res := make([]int, seqLen*2)
	for i := 0; i < seqLen; i++ {
		res[i] = prevSeq[i]
		res[seqLen+i] = prevSeq[seqLen-1-i] | mask
	}
	return res
}

func main() {
	fmt.Println(grayCode(1))
	fmt.Println(grayCode(2))
}
