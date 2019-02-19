//package main

import "fmt"

/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 *
 * https://leetcode.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (37.39%)
 * Total Accepted:    127.3K
 * Total Submissions: 340.4K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 * such that no two queens attack each other.
 *
 *
 *
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 *
 * Each solution contains a distinct board configuration of the n-queens'
 * placement, where 'Q' and '.' both indicate a queen and an empty space
 * respectively.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: [
 * ⁠[".Q..",  // Solution 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 *
 * ⁠["..Q.",  // Solution 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as
 * shown above.
 *
 *
 */
func solveNQueens(n int) [][]string {
	var cols []int
	var res [][]int
	search(n, cols, &res)
	var ret [][]string
	ret = trans2Str(res)
	return ret
}

func trans2Str(xs [][]int) [][]string {
	var ret [][]string
	for i := 0; i < len(xs); i++ {
		var b []string
		for _, j := range xs[i] {
			line := make([]byte, len(xs[i]))
			for k := range line {
				if j != k {
					line[k] = '.'
				} else {
					line[k] = 'Q'
				}
			}
			b = append(b, string(line))
		}
		ret = append(ret, b)
	}
	return ret
}

func search(n int, cols []int, res *[][]int) {
	if len(cols) == n {
		colsCopy := make([]int, len(cols))
		copy(colsCopy, cols)
		*res = append(*res, colsCopy)
		return
	}

	for col := 0; col < n; col++ {
		if !isValid(cols, col) {
			continue
		}
		cols = append(cols, col)
		search(n, cols, res)
		cols = cols[0 : len(cols)-1]
	}
}

func isValid(cols []int, col int) bool {
	for k, v := range cols {
		if v == col {
			return false
		}
		colK := len(cols)
		if colK-k == col-v {
			return false
		}
		if v-col == colK-k {
			return false
		}
	}
	return true
}