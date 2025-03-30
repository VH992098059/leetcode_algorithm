package aign

import (
	"sort"
)

/*最接近的三数之和*/
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	minDiff := abs(result - target)

	for i := 0; i < len(nums)-2; i++ {
		// 跳过重复的 i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum // 提前返回
			}

			// 更新最小差值
			currentDiff := abs(sum - target)
			if currentDiff < minDiff {
				minDiff = currentDiff
				result = sum
			}

			// 移动指针并跳过重复值
			if sum < target {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}
	return result
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
