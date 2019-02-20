// package main

/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 *
 * https://leetcode.com/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (50.45%)
 * Total Accepted:    92K
 * Total Submissions: 182.1K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 * such that no two queens attack each other.
 *
 *
 *
 * Given an integer n, return the number of distinct solutions to the n-queens
 * puzzle.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: 2
 * Explanation: There are two distinct solutions to the 4-queens puzzle as
 * shown below.
 * [
 * [".Q..",  // Solution 1
 * "...Q",
 * "Q...",
 * "..Q."],
 *
 * ["..Q.",  // Solution 2
 * "Q...",
 * "...Q",
 * ".Q.."]
 * ]
 *
 *
 */
func totalNQueens(n int) int {
	var cols []int
	var ret int
	search(cols, n, &ret)

	return ret
}

func search(x []int, n int, count *int) {
	if len(x) == n {
		(*count)++
		return
	}
	for i := 0; i < n; i++ {
		if !isValid(x, i) {
			continue
		}
		x = append(x, i)
		search(x, n, count)
		x = x[0 : len(x)-1]
	}
}

func isValid(x []int, y int) bool {
	l := len(x)
	for k := range x {
		if x[k] == y {
			return false
		}

		if l-k == x[k]-y {
			return false
		}

		if l-k == y-x[k] {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(totalNQueens(4))
// }
