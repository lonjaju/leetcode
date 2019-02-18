// package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (29.22%)
 * Total Accepted:    208.2K
 * Total Submissions: 705.4K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers and an integer target, are there elements
 * a, b, c, and d in nums such that a + b + c + d = target? Find all unique
 * quadruplets in the array which gives the sum of target.
 *
 * Note:
 *
 * The solution set must not contain duplicate quadruplets.
 *
 * Example:
 *
 *
 * Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
 *
 * A solution set is:
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	solutions := [][]int{}
	l := len(nums)

	for i := 0; i <= l-4; i++ {
		if nums[i] > 0 && nums[i] > target {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j <= l-3; j++ {
			if nums[i]+nums[j] > 0 && nums[i]+nums[j] > target {
				break
			}

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k := j + 1
			m := l - 1

			for k < m {
				sum := nums[i] + nums[j] + nums[k] + nums[m]

				if k > j+1 && nums[k] == nums[k-1] {
					k++
					continue
				}

				if m < l-1 && nums[m+1] == nums[m] {
					m--
					continue
				}

				if sum > target {
					m--
				} else if sum < target {
					k++
				} else {
					solutions = append(solutions, []int{nums[i], nums[j], nums[k], nums[m]})
					k++
					m--
				}
			}
		}
	}
	return solutions
}

// func main() {
// 	nums := []int{1, -2, -5, -4, -3, 3, 3, 5}
// 	//[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// 	sort.Ints(nums)

// 	fmt.Println(nums)
// 	fmt.Println(len(nums))

// 	result := fourSum(nums, -11)
// 	for _, v := range result {
// 		fmt.Println(v)
// 	}
// }
