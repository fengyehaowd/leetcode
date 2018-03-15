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
		for nums[right] >= guard && left <= right {
			right--
		}
		if nums[right] < guard {
			nums[right], nums[index] = nums[index], nums[right]
			index = right
			right--
		} else {
			break
		}
	}
	quickSort(nums[:index])
	quickSort(nums[index+1:])
}

func abs(n int) int {
	if n < 0 {
		return n * (-1)
	} else {
		return n
	}
}

func threeSumClosest(nums []int, target int) int {
	quickSort(nums)
	if len(nums) < 3 {
		return 0
	}

	length := len(nums)
	min := abs(nums[0] + nums[1] + nums[length-1] - target)
	result := nums[0] + nums[1] + nums[length-1]
	for i := 0; i < length-2; i++ {
		m := i + 1
		n := length - 1
		for m < n {
			flag := true
			for nums[m]+nums[n]+nums[i] < target && m < n {
				m++
				flag = false
			}
			if m == n {
				if abs(nums[m-1]+nums[n]+nums[i]-target) < min {
					result = nums[m-1] + nums[n] + nums[i]
					min = abs(nums[m-1] + nums[n] + nums[i] - target)
				}
			}

			if m >= n {
				continue
			}

			if abs(nums[m]+nums[n]+nums[i]-target) < min {
				min = abs(nums[m] + nums[n] + nums[i] - target)
				result = nums[m] + nums[n] + nums[i]
			}
			if !flag {

				if abs(nums[m-1]+nums[n]+nums[i]-target) < min {
					min = abs(nums[m-1] + nums[n] + nums[i] - target)
					result = nums[m-1] + nums[n] + nums[i]
				}
			}

			flag = true
			for nums[m]+nums[n]+nums[i] >= target && m < n {
				n--
				flag = false
			}
			if m == n {
				if abs(nums[m]+nums[n+1]+nums[i]-target) < min {
					result = nums[m] + nums[n+1] + nums[i]
					min = abs(nums[m] + nums[n+1] + nums[i] - target)
				}
			}
			if m >= n {
				continue
			}

			if abs(nums[m]+nums[n]+nums[i]-target) < min {
				min = abs(nums[m] + nums[n] + nums[i] - target)
				result = nums[m] + nums[n] + nums[i]
			}
			if !flag {
				if abs(nums[m]+nums[n+1]+nums[i]-target) < min {
					min = abs(nums[m] + nums[n+1] + nums[i] - target)
					result = nums[m] + nums[n+1] + nums[i]
				}
			}
		}
	}
	return result

}

func main() {

	var n []int
	n = []int{0, 1, 2}
	fmt.Println(threeSumClosest(n, 1))
}
