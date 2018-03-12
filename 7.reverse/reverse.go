package main

import "fmt"

func reverse(x int) int {
	if x < 0 {
		return -1 * reverse(-1*x)
	} else if x == 0 {
		return 0
	}

	result := 0
	remain := x
	res := 0
	for remain != 0 {
		res = remain % 10
		remain = (remain - res) / 10
		result = result*10 + res
	}
	if result > 0x80000000 {
		return 0
	}
	return result
}

func main() {
	fmt.Println(reverse(0x5B729735))
}
