package lists

func Problem_1929(nums []int) []int {
	// Given an integer array nums of length n, you want to create an array ans of length 2n 
	// where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).
	// Specifically, ans is the concatenation of two nums arrays.
	// Return the array ans.
	ans := append(nums, nums...)
	return ans
}

func Problem1480(nums []int) []int {
	// RUNNING SUM IN 1D
	// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
	// Return the running sum of nums.
	sumVal := 0
	runningSum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sumVal = 0
		for j := 0; j <= i; j++ {
			sumVal += nums[j]
		}
		runningSum[i] = sumVal
	}
	return runningSum

}

func Problem2011(operations []string) int {
	// There is a programming language with only four operations and one variable X:
	// ++X and X++ increments the value of the variable X by 1.
	// --X and X-- decrements the value of the variable X by 1.
	// Initially, the value of X is 0.
	// Given an array of strings operations containing a list of operations, 
	// return the final value of X after performing all the operations.
	res := 0
	for i := 0; i < len(operations); i++ {
		if operations[i][1] == '+' {
			res++
		} else {
			res--
		}
	}
	return res
}

func Problem1470(nums []int, n int) []int {
	// Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
	// Return the array in the form [x1,y1,x2,y2,...,xn,yn].
	var res = make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		if i % 2 != 0 {
			res[i] = nums[i]
		} else {
			res[i] = nums[n + i]
		}
	}
	return res
}