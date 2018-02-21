package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {

	var buf bytes.Buffer
	length := len(s)
	if numRows == 1 {
		return s
	}
	k := 0
	for ; 2*k*(numRows-1) < length; k++ {
		index := 2 * k * (numRows - 1)
		buf.WriteString(string(s[index]))
	}
	if numRows > 2 {
		for n := 1; n < numRows-1; n++ {
			k = 1
			var index1, index2 int
			for ; ; k = k + 2 {
				index1 = (k-1)*(numRows-1) + n
				index2 = (k+1)*(numRows-1) - n
				if index1 >= length {
					break
				}

				buf.WriteString(string(s[index1]))
				if index2 >= length {
					break
				}
				buf.WriteString(string(s[index2]))
			}
		}
	}
	k = 1
	for ; k*(numRows-1) < length; k = k + 2 {
		index := k * (numRows - 1)
		buf.WriteString(string(s[index]))
	}
	return buf.String()

}

func main() {
	fmt.Println(convert("P", 3))
}
