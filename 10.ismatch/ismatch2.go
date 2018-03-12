package main

import "fmt"

type key struct {
	i int
	j int
}

var match map[key]bool

func isMatch(s string, p string) bool {
	match = make(map[key]bool, 0)
	fmt.Printf("%s  %s\n", s, p)
	src := []byte(s)
	reg := []byte(p)
	l2 := len(p)
	l1 := len(s)

	k := key{
		i: l1,
		j: l2,
	}

	match[k] = true

	for m := l1; m >= 0; m-- {
		for n := l2 - 1; n >= 0; n-- {
			k := key{
				i: m,
				j: n,
			}

			isMatch := m < l1 && (reg[n] == byte('.') || src[m] == reg[n])
			if n+2 <= l2 && reg[n+1] == byte('*') {
				k2 := key{
					i: m,
					j: n + 2,
				}
				k1 := key{
					i: m + 1,
					j: n,
				}
				r2, _ := match[k2]
				r1, _ := match[k1]
				match[k] = (r2 || isMatch && r1)
			} else {
				k1 := key{
					i: m + 1,
					j: n + 1,
				}

				r1, _ := match[k1]
				match[k] = r1 && isMatch
			}
		}
	}
	/*
		for m := 0; m <= l1; m++ {
			for n := 0; n <= l2; n++ {
				k := key{
					i: m,
					j: n,
				}
				if match[k] {
					fmt.Printf("%d %d true\n", m, n)
				} else {
					fmt.Printf("%d %d false\n", m, n)
				}
			}
		}*/
	k = key{
		i: 0,
		j: 0,
	}
	result, _ := match[k]
	return result
}

func main() {
	if isMatch("aacdddc", "f*g*a*c*d*.") {
		fmt.Println("vim-go")
	}
}
