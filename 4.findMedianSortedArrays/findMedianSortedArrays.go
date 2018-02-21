package main

import "fmt"

func getMedian(nums []int) float64 {
	length := len(nums)
	if length%2 == 1 {
		return float64(nums[(length-1)/2])
	} else {
		return float64((float64(nums[(length-2)/2]) + float64(nums[length/2])) / 2.0)
	}
}

func getLen(length int) int {
	if length%2 == 1 {
		return (length - 1) / 2
	} else {
		return (length) / 2
	}
}

func findMedianSortedArray(num int, nums []int) float64 {
	length := len(nums)
	if length%2 == 1 {
		med := getMedian(nums)
		l := getLen(length) + 1
		if num > nums[l-2] && num < nums[l] {
			return (float64(num) + float64(med)) / 2
		} else if num <= nums[l-2] {
			return (float64(nums[l-2]) + float64(nums[l-1])) / 2
		} else {
			return (float64(nums[l-1]) + float64(nums[l])) / 2
		}
	} else {
		l := getLen(length)
		if num > nums[l-1] && num < nums[l] {
			return float64(num)
		} else if num <= nums[l-1] {
			return float64(nums[l-1])
		} else {
			return float64(nums[l])
		}
	}
}

func isOverlapped(nums1 []int, nums2 []int) (bool, float64) {
	len1 := len(nums1)
	len2 := len(nums2)
	if (len1%2 == 0) && (len2%2 == 0) {
		l1 := getLen(len1)
		l2 := getLen(len2)
		if nums1[l1-1] > nums2[l2-1] && nums1[l1] < nums2[l2] {
			return true, getMedian(nums1)
		}
		if nums1[l1-1] < nums2[l2-1] && nums1[l1] > nums2[l2] {
			return true, getMedian(nums2)
		}
		return false, float64(0)
	}
	return false, float64(0)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 == 0 {
		return getMedian(nums2)
	}
	if len2 == 0 {
		return getMedian(nums1)
	}
	if len1 == 1 && len2 == 1 {
		return float64((float64(nums1[0]) + float64(nums2[0])) / 2)
	}

	med1 := getMedian(nums1)
	med2 := getMedian(nums2)
	if len1 == 1 {
		return findMedianSortedArray(nums1[0], nums2)
	}

	if len2 == 1 {
		return findMedianSortedArray(nums2[0], nums1)
	}

	ok, result := isOverlapped(nums1, nums2)
	if ok {
		return result
	}
	if med1 == med2 {
		return med1
	} else if med1 < med2 {
		if len1 < len2 {
			l1 := getLen(len1)
			return findMedianSortedArrays(nums1[l1:], nums2[:(len(nums2)-l1)])
		} else {
			l1 := getLen(len2)
			return findMedianSortedArrays(nums1[l1:], nums2[:(len(nums2)-l1)])
		}
	} else {
		if len1 < len2 {
			l1 := getLen(len1)
			return findMedianSortedArrays(nums1[:(len(nums1)-l1)], nums2[l1:])
		} else {
			l1 := getLen(len2)
			return findMedianSortedArrays(nums1[:(len(nums1)-l1)], nums2[l1:])
		}
	}

}

func main() {
	var n1 = []int{1, 2}
	var n2 = []int{3, 4}
	var result float64
	result = findMedianSortedArrays(n1, n2)
	fmt.Printf("%f\n", result)
}
