package aign

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	arr := [][]int{}
	for i := 0; i < len(nums); i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			n2, n3 := nums[left], nums[right]
			sum := n1 + n2 + n3
			if sum == 0 {
				arr = append(arr, []int{n1, n2, n3})
				for left < right && n2 == nums[left] {
					left++
				}
				for left < right && n3 == nums[right] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return arr
}
