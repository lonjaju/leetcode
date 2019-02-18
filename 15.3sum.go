package main

// package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (22.84%)
 * Total Accepted:    470.9K
 * Total Submissions: 2M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 *
 * Note:
 *
 * The solution set must not contain duplicate triplets.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 *
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */
func threeSum(nums []int) [][]int {
	targets := [][]int{}

	l := len(nums)

	sort.Ints(nums)

	for i := 0; i <= l-3; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := l - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if nums[i]+nums[j] > 0 {
				break
			}

			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			if k < l-1 && nums[k] == nums[k+1] {
				k--
				continue
			}

			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				targets = append(targets, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
		}
	}
	return targets
}

// func main() {
// 	nums := []int{-1, 0, 1, 2, -1, -4}
// 	//重复问题
// 	// nums
// 	res := threeSum(nums)
// 	fmt.Println(res)
// }
