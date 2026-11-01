package main

// 4098. Minimize the Maximum Value of GCD-Connected Array
//You are given three integers n, index, and value.

// An integer array nums of length n is called GCD-Connected if:
//     1. nums[i] > 0 for 0 <= i <= n - 1.
//     2. nums[i] < nums[i + 1] for 0 <= i <= n - 2.
//     3. nums[index] = value.
//     4. gcd(nums[i], nums[i + 1]) > 1 for 0 <= i <= n - 2.

// Return the minimum possible maximum value, nums[n - 1], among all GCD-Connected arrays nums.
// If no such array exists, return -1.

// Example 1:
// Input: n = 4, index = 2, value = 8
// Output: 10
// Explanation:
// - The GCD-Connected array will look like: ​​​​​​​nums = [?, ?, 8, ?] since nums[2] = 8.
// - Since nums[3] > 8 and gcd(8, 9) = 1, every valid GCD-Connected array must have nums[3] >= 10.
// - One optimal array satisfying all conditions is [2, 4, 8, 10].
// - Thus, the minimum possible value of nums[3] is 10.

// Example 2:
// Input: n = 2, index = 1, value = 3
// Output: -1
// Explanation:
// - The GCD-Connected array will look like: nums = [?, 3] since nums[1] = 3.
// - Since we need 0 < nums[0] < nums[1] = 3, there are only two possible arrays: [1, 3] and [2, 3]. Neither of them satisfies gcd(nums[0], nums[1]) > 1.
// - Therefore, no valid GCD-Connected array exists, and the answer is -1.

// Example 3:
// Input: n = 2, index = 0, value = 3
// Output: 6
// Explanation:
// - The GCD-Connected array will look like: nums = [3, ?] since nums[0] = 3.
// - Since nums[1] > 3 and gcd(3, 4) = gcd(3, 5) = 1, every valid GCD-Connected array must have nums[1] >= 6.
// - One optimal array is [3, 6].
// - Thus, the minimum possible value of nums[1] is 6.

// Constraints:
//     - 2 <= n <= 10^{9}
//     - 1 <= value <= 10^{9}
//     - 0 <= index <= n - 1

import "fmt"

func minimumMaxValue(n, index, value int) int64 {
	minPrimeDivisor := func(value int) int { // 返回 value 的最小质因子
		for p := 2; p*p <= value; p++ {
			if value%p == 0 {
				return p
			}
		}
		return value // value 是质数或者 1
	}
	mn := minPrimeDivisor(value)
	if value == 1 || value-mn-(index-1)*2 <= 0 { // nums[0] <= 0
		return -1
	}
	if index == n-1 {
		return int64(value)
	}
	return int64(value + mn + (n-2-index)*2)
}

func main() {
	// Example 1:
	// Input: n = 4, index = 2, value = 8
	// Output: 10
	// Explanation:
	// - The GCD-Connected array will look like: ​​​​​​​nums = [?, ?, 8, ?] since nums[2] = 8.
	// - Since nums[3] > 8 and gcd(8, 9) = 1, every valid GCD-Connected array must have nums[3] >= 10.
	// - One optimal array satisfying all conditions is [2, 4, 8, 10].
	// - Thus, the minimum possible value of nums[3] is 10.
	fmt.Println(minimumMaxValue(4, 2, 8)) // 10
	// Example 2:
	// Input: n = 2, index = 1, value = 3
	// Output: -1
	// Explanation:
	// - The GCD-Connected array will look like: nums = [?, 3] since nums[1] = 3.
	// - Since we need 0 < nums[0] < nums[1] = 3, there are only two possible arrays: [1, 3] and [2, 3]. Neither of them satisfies gcd(nums[0], nums[1]) > 1.
	// - Therefore, no valid GCD-Connected array exists, and the answer is -1.
	fmt.Println(minimumMaxValue(2, 1, 3)) // -1
	// Example 3:
	// Input: n = 2, index = 0, value = 3
	// Output: 6
	// Explanation:
	// - The GCD-Connected array will look like: nums = [3, ?] since nums[0] = 3.
	// - Since nums[1] > 3 and gcd(3, 4) = gcd(3, 5) = 1, every valid GCD-Connected array must have nums[1] >= 6.
	// - One optimal array is [3, 6].
	// - Thus, the minimum possible value of nums[1] is 6.
	fmt.Println(minimumMaxValue(2, 0, 3)) // 6

	fmt.Println(minimumMaxValue(2, 0, 1))                                   // -1
	fmt.Println(minimumMaxValue(2, 0, 1_000_000_000))                       // 1000000002
	fmt.Println(minimumMaxValue(2, 1, 1))                                   // -1
	fmt.Println(minimumMaxValue(2, 1, 1_000_000_000))                       // 1000000000
	fmt.Println(minimumMaxValue(1_000_000_000, 0, 1))                       // -1
	fmt.Println(minimumMaxValue(1_000_000_000, 0, 1_000_000_000))           // 2999999998
	fmt.Println(minimumMaxValue(1_000_000_000, 999_999_999, 1))             // -1
	fmt.Println(minimumMaxValue(1_000_000_000, 999_999_999, 1_000_000_000)) // -1
}
