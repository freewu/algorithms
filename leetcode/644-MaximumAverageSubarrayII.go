package main

// 644. Maximum Average Subarray II
// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is greater than or equal to k that has the maximum average value and return this value. 
// Any answer with a calculation error less than 10^-5 will be accepted.

// Example 1:
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation:
// - When the length is 4, averages are [0.5, 12.75, 10.5] and the maximum average is 12.75
// - When the length is 5, averages are [10.4, 10.8] and the maximum average is 10.8
// - When the length is 6, averages are [9.16667] and the maximum average is 9.16667
// The maximum average is when we choose a subarray of length 4 (i.e., the sub array [12, -5, -6, 50]) which has the max average 12.75, so we return 12.75
// Note that we do not consider the subarrays of length < 4.

// Example 2:
// Input: nums = [5], k = 1
// Output: 5.00000
 
// Constraints:
//     n == nums.length
//     1 <= k <= n <= 10^4
//     -10^4 <= nums[i] <= 10^4

import "fmt"
import "math"

func findMaxAverage(nums []int, k int) float64 {
    mx, mn := nums[0], nums[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        mn = min(mn, v)
        mx = max(mx, v)
    }
    check := func (nums []int, mid float64, k int) bool {
        var sum, pre, minSum float64
        for i := 0; i < k; i++ {
            sum += float64(nums[i]) - mid
        }
        if sum >= 0 {
            return true
        }
        for i := k; i < len(nums); i++ {
            sum += float64(nums[i]) - mid
            pre += float64(nums[i-k]) - mid
            if pre < minSum {
                minSum = pre
            }
            if sum >= minSum {
                return true
            }
        }
        return false
    }
    preMid, err, minVal, maxVal := float64(mx), math.MaxFloat64, float64(mn), float64(mx)
    for err > 0.00001 {
        mid := float64(minVal + maxVal) * 0.5
        if check(nums, mid, k) {
            minVal = mid
        } else {
            maxVal = mid
        }
        err = math.Abs(preMid - mid)
        preMid = mid
    }
    return minVal
}

// stack
func findMaxAverage1(nums []int, k int) float64 {
    type Pair struct {
       sum    int
       length int
    }
    stack, sum, n := []Pair{}, 0, len(nums)
    for i := 0; i < k; i++ {
        sum += nums[i]
    }
    ct, res := k, float64(sum) / float64(k)
    for i := 1; i < n - k + 1; i++ {
        sum += nums[i + k - 1]
        ct++
        s1, c1 := nums[i-1], 1
        for len(stack) > 0 && stack[len(stack)-1].sum * c1 >= stack[len(stack)-1].length * s1 {
            cur := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            c1 += cur.length
            s1 += cur.sum
        }
        stack = append(stack, Pair{s1, c1})
        for len(stack) > 0 && stack[0].sum * ct <= stack[0].length * sum {
            cur := stack[0]
            stack = stack[1:]
            ct -= cur.length
            sum -= cur.sum
        }
        res = math.Max(res, float64(sum) / float64(ct))
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,12,-5,-6,50,3], k = 4
    // Output: 12.75000
    // Explanation:
    // - When the length is 4, averages are [0.5, 12.75, 10.5] and the maximum average is 12.75
    // - When the length is 5, averages are [10.4, 10.8] and the maximum average is 10.8
    // - When the length is 6, averages are [9.16667] and the maximum average is 9.16667
    // The maximum average is when we choose a subarray of length 4 (i.e., the sub array [12, -5, -6, 50]) which has the max average 12.75, so we return 12.75
    // Note that we do not consider the subarrays of length < 4.
    fmt.Println(findMaxAverage([]int{1,12,-5,-6,50,3}, 4)) // 12.75000
    // Example 2:
    // Input: nums = [5], k = 1
    // Output: 5.00000
    fmt.Println(findMaxAverage([]int{5}, 1)) // 5.00000

    fmt.Println(findMaxAverage1([]int{1,12,-5,-6,50,3}, 4)) // 12.75000
    fmt.Println(findMaxAverage1([]int{5}, 1)) // 5.00000
}