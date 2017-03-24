func checkSubarraySum(nums []int, k int) bool {
	// if a[i] % k == t AND (a[i]+a[i+1]+...+a[j]) % k == t
	//    => sum a(i,j] = n*k.
	bkt := make(map[int64]int)
	bkt[0] = -1 // 首元素为n*k的情况
	sum := int64(0)
	for i, v := range nums {
		sum += int64(v)
		cur := sum
		if k != 0 {
			cur = cur % int64(k)
		}
		if pre, ok := bkt[cur]; ok {
			if i-pre > 1 { // subarray size >= 2
				return true
			}
		} else {
			bkt[cur] = i
		}
	}
	return false
} 
