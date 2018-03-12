package main

import "fmt"

func isMatch(s string, p string) bool {
	fmt.Printf("%s  %s\n", s, p)
	src := []byte(s)
	reg := []byte(p)
	l2 := len(p)
	l1 := len(s)

	//边界条件1
	if s == p {
		return true
	}

	// l2 == 0
	if s != "" && p == "" {
		return false
	}
	// l2 == 1
	if l2 == 1 {
		if l1 == 1 && (reg[0] == byte('.') || reg[0] == src[0]) {
			return true
		} else {
			return false
		}
	}

	// 边界条件2
	if s == "" {
		if reg[1] == byte('*') {
			return isMatch(s, p[2:])
		} else {
			return false
		}
	}

	if reg[1] == byte('*') {
		if reg[0] == byte('.') {
			return isMatch(s[1:], p) || isMatch(s, p[2:])
		} else if reg[0] == src[0] {
			return isMatch(s[1:], p) || isMatch(s, p[2:])
		} else {
			return isMatch(s, p[2:])
		}
	} else {
		if reg[0] == src[0] || reg[0] == byte('.') {
			return isMatch(s[1:], p[1:])
		} else {
			return false
		}
	}
}

func main() {
	if isMatch("aacdddc", "f*g*a*c*d*.") {
		fmt.Println("vim-go")
	}
}
