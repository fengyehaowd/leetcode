package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int][]int)
	for index, item := range nums {
		info, ok := numMap[item]
		if !ok {
			var info []int
			info = append(info, index)
			numMap[item] = info
		} else {
			info = append(info, index)
		}
	}
	var result []int
	for indexF, itemF := range nums {
		dst := target - itemF
		info, ok := numMap[dst]
		if ok {
			if indexF != info[0] {
				result = append(result, indexF)
				result = append(result, info[0])

				return result
			}
		}
	}
	return result
}
