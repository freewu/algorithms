package main 

// 3224. Minimum Array Changes to Make Differences Equal
// You are given an integer array nums of size n where n is even, and an integer k.

// You can perform some changes on the array, where in one change you can replace any element in the array with any integer in the range from 0 to k.

// You need to perform some changes (possibly none) such that the final array satisfies the following condition:
//     There exists an integer X such that abs(a[i] - a[n - i - 1]) = X for all (0 <= i < n).

// Return the minimum number of changes required to satisfy the above condition.

// Example 1:
// Input: nums = [1,0,1,2,4,3], k = 4
// Output: 2
// Explanation:
// We can perform the following changes:
// Replace nums[1] by 2. The resulting array is nums = [1,2,1,2,4,3].
// Replace nums[3] by 3. The resulting array is nums = [1,2,1,3,4,3].
// The integer X will be 2.

// Example 2:
// Input: nums = [0,1,2,3,3,6,5,4], k = 6
// Output: 2
// Explanation:
// We can perform the following operations:
// Replace nums[3] by 0. The resulting array is nums = [0,1,2,0,3,6,5,4].
// Replace nums[4] by 4. The resulting array is nums = [0,1,2,0,4,6,5,4].
// The integer X will be 4.

// Constraints:
//     2 <= n == nums.length <= 10^5
//     n is even.
//     0 <= nums[i] <= k <= 10^5

import "fmt"

func minChanges(nums []int, k int) int {
    res, n := 1 << 31, len(nums)
    diff, mx := make([]int, k + 1),  make([]int, k+1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n / 2; i++ {
        diff[abs(nums[i] - nums[n-i-1])]++
        mx[max(max(nums[i], nums[n-i-1]), k - min(nums[i], nums[n-i-1]))]++
    }
    two := 0
    for i := 0; i <= k; i++ {
        one := n / 2 - two - diff[i]
        res = min(res, 2 * two + one)
        two += mx[i]
    }
    return res
}

func minChanges1(nums []int, k int) int {
    res, n, count := 1 << 31, len(nums), 0
    diff := make([]int, k + 2) // diff[i]表示差为i时, 改动的次数
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n / 2; i++ {
        cur, mx := abs(nums[i] - nums[n - i - 1]), max(max(nums[i], nums[n - i - 1]), max(k - nums[i], k - nums[n - i - 1]))
        // x 代表符合题目条件的绝对差值
        // 1.如果x在[0, diff)范围内，只需要改一次数字
        diff[0]++
        diff[cur]--
        // 2.如果 x == diff，不用改变任何数字
        // 3.如果 x 在 (diff, max] 范围内，只需要改一次数字
        diff[cur + 1]++
        diff[mx + 1]--
        // 4.如果 x 在(max, +infinity) 范围内，需要两个数字都改
        diff[mx + 1] += 2
    }
    for i := 0; i <= k; i++ {
        count += diff[i]
        res = min(res, count)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,0,1,2,4,3], k = 4
    // Output: 2
    // Explanation:
    // We can perform the following changes:
    // Replace nums[1] by 2. The resulting array is nums = [1,2,1,2,4,3].
    // Replace nums[3] by 3. The resulting array is nums = [1,2,1,3,4,3].
    // The integer X will be 2.
    fmt.Println(minChanges([]int{1,0,1,2,4,3}, 4)) // 2
    // Example 2:
    // Input: nums = [0,1,2,3,3,6,5,4], k = 6
    // Output: 2
    // Explanation:
    // We can perform the following operations:
    // Replace nums[3] by 0. The resulting array is nums = [0,1,2,0,3,6,5,4].
    // Replace nums[4] by 4. The resulting array is nums = [0,1,2,0,4,6,5,4].
    // The integer X will be 4.
    fmt.Println(minChanges([]int{0,1,2,3,3,6,5,4}, 6)) // 2

    fmt.Println(minChanges([]int{0,1,2,3,4,5,6,7,8,9}, 9)) // 4
    fmt.Println(minChanges([]int{9,8,7,6,5,4,3,2,1,0}, 9)) // 4

    fmt.Println(minChanges1([]int{1,0,1,2,4,3}, 4)) // 2
    fmt.Println(minChanges1([]int{0,1,2,3,3,6,5,4}, 6)) // 2
    fmt.Println(minChanges1([]int{0,1,2,3,4,5,6,7,8,9}, 9)) // 4
    fmt.Println(minChanges1([]int{9,8,7,6,5,4,3,2,1,0}, 9)) // 4
}