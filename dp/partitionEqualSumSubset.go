var dp map[int]bool

func canPartition(nums []int) bool {

	total := 0
	dp = make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	if total%2 != 0 {
		return false
	}

	return partition(0, nums[0], total/2, nums)
}

func partition(start, sum, target int, nums []int) bool {

	if sum == target {
		return true
	}

	if _, ok := dp[sum]; ok {
		return dp[sum]
	}

	if sum > target {
		return false
	}

	for i := start + 1; i < len(nums); i++ {

		if partition(i, sum+nums[i], target, nums) {
			return true
		}
		dp[sum] = false
	}

	return false
}