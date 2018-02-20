package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	for indexF, itemF := range nums[:len(nums)-1] {
		for indexS, itemS := range nums[indexF+1:] {
			if itemF+itemS == target {
				result = append(result, indexF)
				result = append(result, indexS+indexF+1)
				return result
			}
		}
	}
	return result
}
