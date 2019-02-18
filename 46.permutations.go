// package main

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (52.42%)
 * Total Accepted:    326.2K
 * Total Submissions: 619.3K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a collection of distinct integers, return all possible permutations.
 *
 * Example:
 *
 *
 * Input: [1,2,3]
 * Output:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 *
 */
func permute(nums []int) [][]int {
	var results [][]int
	dfs(nums, 0, &results)
	return results
}

func dfs(nums []int, start int, res *[][]int) {
	if start >= len(nums) {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		*res = append(*res, numsCopy)
		return
	}

	for i := start; i < len(nums); i++ {
		swap(&nums, i, start)
		dfs(nums, start+1, res)
		swap(&nums, i, start)
	}
}

func swap(nums *[]int, i int, j int) {
	(*nums)[j], (*nums)[i] = (*nums)[i], (*nums)[j]
}

// func main() {
// 	nums := []int{1, 2}
// 	ret := permute(nums)
// 	fmt.Println(ret)
// }
