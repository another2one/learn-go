//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回
// -1。
//
// 你可以认为每种硬币的数量是无限的。
//
//
//
// 示例 1：
//
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
// 示例 2：
//
//
//输入：coins = [2], amount = 3
//输出：-1
//
// 示例 3：
//
//
//输入：coins = [1], amount = 0
//输出：0
//
//
// 示例 4：
//
//
//输入：coins = [1], amount = 1
//输出：1
//
//
// 示例 5：
//
//
//输入：coins = [1], amount = 2
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 231 - 1
// 0 <= amount <= 104
//
// Related Topics 动态规划
// 👍 1203 👎 0

package q322

import "math"

// c = [2, 3, 5]
// 假设f(n) 为n的最小银币数， 第n等于n的上一个最小银币值加1
// f(n) = min(f(n-2), f(n-3), f(n-5)) + 1
func coinChange(coins []int, amount int) int {

	if len(coins) < 1 && amount < 1 {
		return -1
	}

	// 数组存储达到每个总数的最小硬币数
	coinMinArray := make([]int, amount+1)

	for now := 1; now <= amount; now++ {
		coinMinArray[now] = math.MaxInt32
		for _, coin := range coins {
			if now >= coin && coinMinArray[now] > coinMinArray[now-coin]+1 {
				// 当前步数不是最小值
				coinMinArray[now] = coinMinArray[now-coin] + 1
			}
		}
	}

	if coinMinArray[amount] == math.MaxInt32 {
		return -1
	}

	return coinMinArray[amount]
}
