package main

import "fmt"

type stack struct {
	top   int
	array []byte
}

func (s *stack) push(b byte) byte {
	s.array = append(s.array[:s.top], b)
	s.top = s.top + 1
	return b
}

func (s *stack) pop() byte {
	if s.top <= 0 {
		return byte(' ')
	}
	b := s.array[s.top-1]
	s.top = s.top - 1
	return b
}

func isValid(s string) bool {
	cm := make(map[byte]byte)
	cm[byte(')')] = byte('(')
	cm[byte('}')] = byte('{')
	cm[byte(']')] = byte('[')

	f := make(map[byte]byte)
	f[byte('(')] = byte(')')
	f[byte('{')] = byte('}')
	f[byte('[')] = byte(']')
	sk := stack{
		top:   0,
		array: []byte{},
	}

	for _, c := range []byte(s) {
		b, ok := f[c]
		if ok {
			sk.push(c)
			continue
		}
		b, ok = cm[c]
		if ok {
			k := sk.pop()
			if k != b || k == byte(' ') {
				return false
			}
		}
	}
	if sk.top == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	if isValid("()") {
		fmt.Println("vim-go")
	}
}
