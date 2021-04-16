//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
// 问总共有多少条不同的路径？
//
//
//
// 示例 1：
//
//
//输入：m = 3, n = 7
//输出：28
//
// 示例 2：
//
//
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
//
//
// 示例 3：
//
//
//输入：m = 7, n = 3
//输出：28
//
//
// 示例 4：
//
//
//输入：m = 3, n = 3
//输出：6
//
//
//
// 提示：
//
//
// 1 <= m, n <= 100
// 题目数据保证答案小于等于 2 * 109
//
// Related Topics 数组 动态规划
// 👍 957 👎 0

package q62

var Count1 = 1
var Count2 = 1

// m,n格只能从左上和右下过i来
// f(m, n) = f(m, n-1) + f(m-1, n)
func uniquePaths(m int, n int) int {
	Count1++
	if m <= 1 || n <= 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

// 上下为0,1 则走到m,n其实就是m-1+n-1个数，n-1个0的组合:C(m-1+n-1, n-1) = (m+n-2)!/ ((m+n-2-(n-1))! * (n-1)!)
func uniquePaths1(m int, n int) int {
	if m <= 1 || n <= 1 {
		return 1
	}
	var factorial = func(s int, n int) int {
		if s <= 0 {
			return 0
		}
		if n < 1 {
			n = 1
		}
		sum := 1
		for s >= n {
			sum *= s
			s--
		}
		return sum
	}
	min := m
	if m > n {
		min = n
	}
	return factorial(m+n-2, m+n-min) / factorial(min-1, 1)
}

func uniquePaths2(m int, n int) int {
	if m <= 1 || n <= 1 {
		return 1
	}
	memo := make([][]int, m+1)

	for ms := 1; ms <= m; ms++ {
		memo[ms] = make([]int, n+1)
		for ns := 1; ns <= n; ns++ {
			Count2++
			if ms == 1 && ns == 1 {
				// 起始位置
				memo[ms][ns] = 0
			} else {
				memo[ms][ns] = 1
			}

			if ns > 1 && ms > 1 {
				memo[ms][ns] = memo[ms][ns-1] + memo[ms-1][ns]
			}
		}
	}
	return memo[m][n]
}
