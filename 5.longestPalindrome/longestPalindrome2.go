package main

import "fmt"

func newString(s string) string {
	result := "#"
	for _, c := range s {
		result = result + string(c)
		result = result + "#"
	}
	return result
}

func oldString(s string) string {
	result := ""
	for i := 1; i < len(s); i = i + 2 {
		result = result + string(s[i])
	}
	return result
}

func longestPalindrome(s string) string {
	str := newString(s)
	fmt.Println(str)
	p := make([]int, 2000)
	max := 0
	id := 0
	longest := 0
	longestIndex := 0
	for i := 1; i < len(str)-1; i++ {
		if max > i {
			if p[2*id-i] > max-i {
				p[i] = max - i
			} else {
				p[i] = p[2*id-i]
			}
		}
		p[i] = 0
		index := 1
		for ; index <= i && i+index < len(str) && str[i-index] == str[i+index]; index++ {
		}
		p[i] = index - 1
		if p[i] > longest {
			longest = p[i]
			longestIndex = i
		}
		if i+p[i] > max {
			max = i + p[i]
			id = i
		}

		fmt.Printf("i:%d p[i]:%d\n", i, p[i])
	}
	fmt.Printf("index:%d length:%d\n", longestIndex, longest)
	return oldString(str[longestIndex-longest : longestIndex+longest+1])
}

func main() {
	fmt.Println(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
	fmt.Println(longestPalindrome("A"))
}
