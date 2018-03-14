package main

import "fmt"

var series []int

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}

func maxArea1(height []int) int {
	length := len(height)
	if len(height) == 2 {
		if height[0] >= height[1] {
			series = append(series, 0)
			return height[1]
		} else {
			series = append(series, 0)
			series = append(series, 1)
			return height[0]
		}
	}

	area := 0
	flag := true
	old := maxArea(height[:length-1])

	for _, i := range series {
		if height[i] >= height[length-1] {
			new := height[length-1] * (length - 1 - i)

			flag = false
			if new > area {
				area = new
				break
			}
		} else {
			new := height[i] * (length - 1 - i)
			if new > area {
				area = new
			}
		}

	}

	if flag {
		series = append(series, length-1)
	}

	fmt.Println(series)

	return max(old, area)
}

func maxArea(height []int) int {
	series = make([]int, 0)
	return maxArea1(height)
}

func main() {
	fmt.Println(maxArea([]int{1, 2}))
}
