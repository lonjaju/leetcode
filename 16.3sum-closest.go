// package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (33.61%)
 * Total Accepted:    236.6K
 * Total Submissions: 693.5K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 *
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 */
func threeSumClosest(nums []int, target int) int {

	/*
		nums = [-1, 2, 1, -4], and target = 1.
		 排序 [-4,-1,1,2]

	*/
	sort.Ints(nums)

	l := len(nums)
	nearest := nums[0] + nums[1] + nums[2]
	diff := abs(nearest - target)

	for i := 0; i <= l-3; i++ {
		j := i + 1
		k := l - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == target {
				return target
			} else if sum > target {
				k--
			} else {
				j++
			}

			tmp := abs(sum - target)
			if tmp < diff {
				nearest = sum
				diff = tmp
			}
		}
	}

	return nearest
}

func abs(num int) int {
	if num < 0 {
		num = 0 - num
	}
	return num
}

// func main() {
// 	nums := []int{6, -18, -20, -7, -15, 9, 18, 10, 1, -20, -17, -19, -3, -5, -19, 10, 6, -11, 1, -17, -15, 6, 17, -18, -3, 16, 19, -20, -3, -17, -15, -3, 12, 1, -9, 4, 1, 12, -2, 14, 4, -4, 19, -20, 6, 0, -19, 18, 14, 1, -15, -5, 14, 12, -4, 0, -10, 6, 6, -6, 20, -8, -6, 5, 0, 3, 10, 7, -2, 17, 20, 12, 19, -13, -1, 10, -1, 14, 0, 7, -3, 10, 14, 14, 11, 0, -4, -15, -8, 3, 2, -5, 9, 10, 16, -4, -3, -9, -8, -14, 10, 6, 2, -12, -7, -16, -6, 10}
// 	fmt.Println(threeSumClosest(nums, -52))
// }
