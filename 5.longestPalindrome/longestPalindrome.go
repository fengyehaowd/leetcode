package main

import "fmt"

type keys struct {
	v1 int
	v2 int
}

var palindromeMap map[keys]bool

func initPalindromeMap(s string, i int, j int) bool {
	key := keys{
		v1: i,
		v2: j,
	}
	_, ok := palindromeMap[key]
	if !ok {
		if i == j {
			palindromeMap[key] = true
		} else if j == i+1 {
			if s[i] == s[j] {
				palindromeMap[key] = true
			} else {
				palindromeMap[key] = false
			}
		} else {
			palindromeMap[key] = (initPalindromeMap(s, i+1, j-1) && (s[i] == s[j]))
		}
	}
	result, _ := palindromeMap[key]
	return result
}
func longestPalindrome(s string) string {
	length := len(s)
	maxLen := 0
	left := 0
	right := 0
	palindromeMap = make(map[keys]bool)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if initPalindromeMap(s, i, j) && ((j - i + 1) > maxLen) {
				left = i
				right = j + 1
				maxLen = j - i + 1
			}
		}
	}
	return s[left:right]
}

func main() {

	fmt.Println(longestPalindrome("cyyoacmjwjubfkzrrbvquqkwhsxvmytmjvbborrtoiyotobzjmohpadfrvmxuagbdczsjuekjrmcwyaovpiogspbslcppxojgbfxhtsxmecgqjfuvahzpgprscjwwutwoiksegfreortttdotgxbfkisyakejihfjnrdngkwjxeituomuhmeiesctywhryqtjimwjadhhymydlsmcpycfdzrjhstxddvoqprrjufvihjcsoseltpyuaywgiocfodtylluuikkqkbrdxgjhrqiselmwnpdzdmpsvbfimnoulayqgdiavdgeiilayrafxlgxxtoqskmtixhbyjikfmsmxwribfzeffccczwdwukubopsoxliagenzwkbiveiajfirzvngverrbcwqmryvckvhpiioccmaqoxgmbwenyeyhzhliusupmrgmrcvwmdnniipvztmtklihobbekkgeopgwipihadswbqhzyxqsdgekazdtnamwzbitwfwezhhqznipalmomanbyezapgpxtjhudlcsfqondoiojkqadacnhcgwkhaxmttfebqelkjfigglxjfqegxpcawhpihrxydprdgavxjygfhgpcylpvsfcizkfbqzdnmxdgsjcekvrhesykldgptbeasktkasyuevtxrcrxmiylrlclocldmiwhuizhuaiophykxskufgjbmcmzpogpmyerzovzhqusxzrjcwgsdpcienkizutedcwrmowwolekockvyukyvmeidhjvbkoortjbemevrsquwnjoaikhbkycvvcscyamffbjyvkqkyeavtlkxyrrnsmqohyyqxzgtjdavgwpsgpjhqzttukynonbnnkuqfxgaatpilrrxhcqhfyyextrvqzktcrtrsbimuokxqtsbfkrgoiznhiysfhzspkpvrhtewthpbafmzgchqpgfsuiddjkhnwchpleibavgmuivfiorpteflholmnxdwewj"))
}
