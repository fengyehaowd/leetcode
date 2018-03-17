package main

import "fmt"

var t []int

func preProcess(str string) {
	t = append(t, 0)
	for i := 1; i < len(str); i++ {
		for j := 1; j < i+1; j++ {
			str2 := str[j:i]
			flag := false
			for index := 0; index < i-j; index++ {
				if str[index] != str2[index] {
					flag = true
					break
				}
			}
			if str[i-j] != str[i] && !flag {
				t = append(t, j-1)
				fmt.Println(i, j)
				break
			}
		}
		if len(t) < i+1 {
			t = append(t, i)
		}
	}
	fmt.Println(t)
	return
}

func getNums(str string) int {
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] != str[i-1] {
			count++
		}
	}
	return count
}
func strStr(haystack string, needle string) int {

	fmt.Printf("%s %s\n", haystack, needle)
	preProcess(needle)
	l1 := len(haystack)
	l2 := len(needle)
	if l2 == 0 {
		return 0
	}
	if l2 > l1 {
		return -1
	}
	for i := 0; i <= l1-l2; i++ {
		find := true
		for j := 0; j < l2; j++ {
			if needle[j] != haystack[i+j] {
				i = i + t[j]
				fmt.Printf("index:%d skip:%d\n", i, t[j])
				find = false
				break
			}
		}
		if find {
			return i
		}
	}
	return -1

}

func main() {
	fmt.Println("vim-go")
	fmt.Println(strStr("mississippi", "pi"))
}
