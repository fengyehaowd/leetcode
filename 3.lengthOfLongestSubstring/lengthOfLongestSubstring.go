package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var result int
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	chIndex := make(map[byte]int)
	for index, c := range []byte(s) {

		i, ok := chIndex[c]
		if !ok {
			result = result + 1
			chIndex[c] = index
		} else {
			length := lengthOfLongestSubstring(s[i+1:])
			if length > result {
				return length
			} else {
				return result
			}
		}
	}
	return result
}
func main() {
	fmt.Println("vim-go")
}
