package main

import (
	"bytes"
	"fmt"
)

var trans map[byte]byte

func tenTimes(num string) string {
	var buf bytes.Buffer
	for _, c := range []byte(num) {
		r, _ := trans[c]
		buf.Write([]byte{r})
	}
	return buf.String()
}

func getRoman(num int) string {

	var buf bytes.Buffer
	if num <= 3 {
		for i := 1; i <= num; i++ {
			buf.Write([]byte{'I'})
		}
	} else if num > 3 && num <= 5 {
		for i := num; i < 5; i++ {
			buf.Write([]byte{'I'})
		}
		buf.Write([]byte{'V'})
	} else if num > 5 && num <= 8 {
		buf.Write([]byte{'V'})
		for i := 6; i <= num; i++ {
			buf.Write([]byte{'I'})
		}
	} else if num > 8 && num < 10 {
		buf.Write([]byte{'I', 'X'})
	}
	return buf.String()
}

func intToRoman(num int) string {
	trans = make(map[byte]byte)
	trans[byte('I')] = byte('X')
	trans[byte('X')] = byte('C')
	trans[byte('C')] = byte('M')
	trans[byte('V')] = byte('L')
	trans[byte('L')] = byte('D')

	n := num
	k := 1000
	result := ""
	for {
		if n < k {
			k = k / 10
		} else {
			break
		}
	}

	for {
		if k == 0 {
			return result
		}

		remain := n % k
		c := (n - remain) / k
		str1 := tenTimes(result)
		str2 := getRoman(c)
		result = str1 + str2
		n = remain
		k = k / 10
	}
	return result
}

func main() {
	fmt.Println(intToRoman(11))
}
