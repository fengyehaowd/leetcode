package main

import "fmt"

func quickSort(nums []int) {
	length := len(nums)
	if len(nums) <= 1 {
		return
	}

	guard := nums[length-1]
	index := length - 1
	left := 0
	right := length - 2
	for left <= right {

		for nums[left] <= guard && left <= right {
			left++
		}
		if nums[left] > guard {
			nums[left], nums[index] = nums[index], nums[left]
			index = left
			left++
		} else {
			break
		}

		//	fmt.Println(nums)
		for nums[right] >= guard && left <= right {
			right--
		}
		if nums[right] < guard {
			nums[right], nums[index] = nums[index], nums[right]
			index = right
			right--

			//		fmt.Println(nums)
		} else {
			break
		}
	}
	//	fmt.Println(nums)
	quickSort(nums[:index])
	quickSort(nums[index+1:])
}

func isSame(n1 []int, n2 []int) bool {
	if len(n1) != len(n2) {
		return false
	}
	for i := 0; i < len(n1); i++ {
		if n1[i] != n2[i] {
			return false
		}
	}
	return true
}

func getStr(n []int) string {
	return fmt.Sprintf("%d+%d+%d", n[0], n[1], n[2])
}

func threeSum(nums []int) [][]int {
	quickSort(nums)

	//fmt.Println(nums)
	if len(nums) < 3 {
		return nil
	}

	myset := make(map[int]int)
	var nums2 []int
	for _, i := range nums {
		t, ok := myset[i]
		if !ok {
			nums2 = append(nums2, i)
			myset[i] = 1
		} else if t == 1 {
			nums2 = append(nums2, i)
			myset[i] = 2
		} else if t == 2 {
			nums2 = append(nums2, i)
			myset[i] = 3
		}

	}
	nums = nums2
	length := len(nums)
	//fmt.Println(nums)
	var result [][]int
	for i := 0; i < length; i++ {
		sum := 0 - nums[i]
		m := i + 1
		n := length - 1
		for m < n {
			for nums[m]+nums[n] < sum && m < n {
				m++
			}
			if m >= n {
				break
			}
			if nums[m]+nums[n] == sum {
				answer := []int{nums[i], nums[m], nums[n]}
				//				fmt.Println(answer)
				result = append(result, answer)
				m++
			}

			for nums[m]+nums[n] > sum && m < n {
				n--
			}

			if m >= n {
				break
			}

			if nums[m]+nums[n] == sum {
				answer := []int{nums[i], nums[m], nums[n]}

				//				fmt.Println(answer)
				result = append(result, answer)
				n--
			}
		}
	}
	var result2 [][]int
	//fmt.Println(result)
	myset2 := make(map[string]int)
	for _, i := range result {
		str := getStr(i)
		_, ok := myset2[str]
		if !ok {
			result2 = append(result2, i)
			myset2[str] = 1
		}
	}
	return result2

}

func main() {

	var n []int
	n = []int{-1, 0, 0, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(n))
}
