package algorithm

//最长上升子序列
func MaxAscSubSeq(seq []int) int {
	seqLen := len(seq)

	//已i结尾的最长子序列
	dp := make([]int, seqLen)
	dp[0] = 1
	max := dp[0]

	for i := 1; i < seqLen; i++ {
		dp[i] = 1 //如果之前的元素都比i节点元素大
		for j := 0; j < i; j++ {
			if seq[i] > seq[j] && dp[j] >= dp[i] {
				dp[i] = dp[j] + 1
			}
		}

		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
