import "sort"

// 典型的DP问题
func lengthOfLIS(nums []int) int {
	var dp []int
	for _, v := range nums {
		i := sort.Search(len(dp), func(i int) bool { return dp[i] >= v })
		if i == len(dp) {
			dp = append(dp, v)
		} else {
			dp[i] = v
		}
	}
	return len(dp)
}
