package main

import "fmt"

var cNum map[rune]int
var cIndex map[rune]int

func romanToInt1(s []rune, plus bool) int {
	if plus {

		result := 0
		for _, c := range s {
			result = result + cNum[c]
		}

		return result
	} else {
		result := 0
		length := len(s)
		for _, c := range s[:length-1] {
			result = result + cNum[c]
		}
		result = cNum[s[length-1]] - result

		return result
	}
}

func romanToInt(s string) int {
	cNum = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	cIndex = map[rune]int{
		'I': 1,
		'V': 2,
		'X': 3,
		'L': 4,
		'C': 5,
		'D': 6,
		'M': 7,
	}
	start := 0
	result := 0
	state := -1
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if state == -1 {
			state = cIndex[c]
			continue
		}
		if cIndex[c] == state {
			continue
		}

		if cIndex[c] < state {
			result = result + romanToInt1([]rune(s[start:i]), true)
			start = i
			state = cIndex[c]
		}

		if cIndex[c] > state {
			i++
			result = result + romanToInt1([]rune(s[start:i]), false)
			start = i
			if i >= len(s) {
				break
			}

			state = cIndex[rune(s[i])]
		}
	}
	result = result + romanToInt1([]rune(s[start:]), true)
	return result
}

func main() {
	fmt.Println(romanToInt("CCXLIV"))
}
