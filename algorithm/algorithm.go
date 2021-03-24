package algorithm

//两人从上下两端走楼梯
var steps = []int{1, 2, 3, 4}

func WaysMoving(ap int, bp int) int {
	if ap > bp {
		return 0
	} else if ap == bp {
		return 1
	} else {
		count := 0
		for _, stepA := range steps {
			for _, stepB := range steps {
				count += WaysMoving(ap+stepA, bp-stepB)
			}
		}
		return count
	}
}

//只展开到能走三级阶梯
func WaysMovingRC(n int) int {
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	} else if n == 4 {
		return 4
	} else if n == 5 {
		return 6
	} else if n == 6 {
		return 12
	} else {
		return WaysMovingRC(n-2) +
			2*WaysMovingRC(n-3) +
			3*WaysMovingRC(n-4) +
			2*WaysMovingRC(n-5) +
			1*WaysMovingRC(n-6)
	}
}

//动态规划，滚动数组算法
func WaysMovingByEvenCount(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	count := 0
	for i := 0; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			if dp[j] > 0 {
				for _, step := range steps {
					if j+step <= n {
						dp[j+step] += dp[j]
						if i%2 == 1 && j+step == n {
							count += dp[j]
						}
					}
				}
				dp[j] -= dp[j]
			}
		}
	}
	return count
}

/*
0 0
1 0
2 1
3 2
4 3
5 2
6 1
*/
