package main

import (
	"fmt"
	"strings"
)

func isMinus(b byte) bool {
	if b == byte('-') {
		return true
	}
	return false
}

func isPlus(b byte) bool {
	if b == byte('+') {
		return true
	}
	return false
}

func isNumber(b byte) bool {
	if b >= byte('0') && b <= byte('9') {
		return true
	}
	return false
}

func getNum(b byte) int64 {
	return int64(b) - int64(byte('0'))
}

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	var max, min, minus int64
	minus = 1
	bytes := []byte(str)
	var result int64
	result = 0
	max = 2147483647
	min = 2147483648
	if isMinus(bytes[0]) {
		minus = -1
		str = str[1:]
		if len(str) == 0 {
			return 0
		}
	}

	if isPlus(bytes[0]) {
		str = str[1:]
		if len(str) == 0 {
			return 0
		}
	}

	bytes = []byte(str)
	if isNumber(bytes[0]) {
		for _, c := range []byte(str) {
			if isNumber(c) {
				result = result*int64(10) + getNum(c)
				if result >= max && minus == 1 {
					return int(max)
				}
				if result >= min && minus == -1 {
					return int(int64(-1) * min)
				}
			} else {
				return int(result * minus)
			}
		}
	} else {
		return 0
	}
	return int(result * minus)
}

func main() {
	fmt.Println(getNum(byte('9')))
	fmt.Println(myAtoi("+3434343dfdf"))
	fmt.Println(myAtoi("-2147483649"))
	fmt.Println(myAtoi("   - 3333333333333333333333333333333333333333ssssssssssssss3ddddddddddddddddddddddddfds"))

}
