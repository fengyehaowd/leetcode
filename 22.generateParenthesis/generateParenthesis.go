package main

import "fmt"

func generateParenthesis(n int) []string {
	if n < 1 {
		return []string{}
	}

	if n == 1 {
		return []string{"()"}
	}

	pre := generateParenthesis(n - 1)
	var answer []string
	set := make(map[string]int)

	for _, str := range pre {
		for i := 0; i <= (2 * (n - 1)); i++ {
			s := str[0:i] + "(" + str[i:2*(n-1)] + ")"
			_, ok := set[s]
			if !ok {
				set[s] = 1
				answer = append(answer, s)
			}
		}
	}
	return answer
}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(2))
}
