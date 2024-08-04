package main

// 1508. Range Sum of Sorted Subarray Sums
// You are given the array nums consisting of n positive integers. 
// You computed the sum of all non-empty continuous subarrays from the array and then sorted them in non-decreasing order, creating a new array of n * (n + 1) / 2 numbers.

// Return the sum of the numbers from index left to index right (indexed from 1), inclusive, in the new array. 
// Since the answer can be a huge number return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [1,2,3,4], n = 4, left = 1, right = 5
// Output: 13 
// Explanation: 
// All subarray sums are 1, 3, 6, 10, 2, 5, 9, 3, 7, 4. 
// After sorting them in non-decreasing order we have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. 
// The sum of the numbers from index le = 1 to ri = 5 is 1 + 2 + 3 + 3 + 4 = 13. 

// Example 2:
// Input: nums = [1,2,3,4], n = 4, left = 3, right = 4
// Output: 6
// Explanation: 
// The given array is the same as example 1.
// We have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. 
// The sum of the numbers from index le = 3 to ri = 4 is 3 + 3 = 6.

// Example 3:
// Input: nums = [1,2,3,4], n = 4, left = 1, right = 10
// Output: 50

// Constraints:
//     n == nums.length
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 100
//     1 <= left <= right <= n * (n + 1) / 2

import "fmt"
import "sort"

// 模拟
func rangeSum(nums []int, n int, left int, right int) int {
    sumsLength, mod := n * (n + 1) / 2, 1_000_000_007
    sums := make([]int, sumsLength)
    res, index := 0, 0
    for i := 0; i < n; i++ {
        sum := 0
        for j := i; j < n; j++ {
            sum += nums[j]
            sums[index] = sum
            index++
        }
    }
    sort.Ints(sums)
    for i := left - 1; i < right; i++ {
        res = (res + sums[i]) % mod
    }
    return res
}

// 前缀和 + 二分法
func rangeSum1(nums []int, n int, left int, right int) int {
    prefixSum := make([]int, n+1)
    prefixSum[0] = 0
    for i := 0; i < n; i++ {
        prefixSum[i+1] = prefixSum[i] + nums[i]
    }
    prefixPrefixSum := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefixPrefixSum[i+1] = prefixPrefixSum[i] + prefixSum[i+1]
    }
    getCount := func (prefixSum []int, n int, x int) int { // 统计前缀和中值小于等于x的数目
        count, j := 0, 1
        for i := 0; i < n; i++ {
            for j <= n && prefixSum[j] - prefixSum[i] <= x {
                j++
            }
            j -= 1
            count += j - i
        }
        return count
    }
    // 二分法求有序矩阵的第K小的值
    // 求前缀和列表表示中的第K小的值，prefixSum是首行的前缀和列表
    getKth := func (prefixSum []int, n int, k int) int {
        left, right := 0, prefixSum[n]
        for left < right {
            mid := left + (right - left) / 2
            count := getCount(prefixSum, n, mid)
            if count < k {
                left = mid + 1
            } else {
                right = mid
            }
        }
        return left // 最后得到的left也一定会在前缀和矩阵中
    }
    getSum := func (prefixSum, prefixPrefixSum []int, n int, k int) int { // 求有序前缀和中前K个元素的和
        num := getKth(prefixSum, n, k)
        // 求严格小于num的数
        sum, cnt, j := 0, 0, 1
        for i := 0; i < n; i++ {
            for j <= n && prefixSum[j] - prefixSum[i] < num {
                j++
            }
            j -= 1
            sum += (prefixPrefixSum[j] - prefixPrefixSum[i] - (j - i) * prefixSum[i])
            cnt += j - i
        }
        sum = (sum + (k - cnt) * num) % 1_000_000_007
        return sum
    }
    return getSum(prefixSum, prefixPrefixSum, n, right) - getSum(prefixSum, prefixPrefixSum, n, left - 1)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4], n = 4, left = 1, right = 5
    // Output: 13 
    // Explanation: 
    // All subarray sums are 1, 3, 6, 10, 2, 5, 9, 3, 7, 4. 
    // After sorting them in non-decreasing order we have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. 
    // The sum of the numbers from index le = 1 to ri = 5 is 1 + 2 + 3 + 3 + 4 = 13. 
    fmt.Println(rangeSum([]int{1,2,3,4},4,1,5)) // 13
    // Example 2:
    // Input: nums = [1,2,3,4], n = 4, left = 3, right = 4
    // Output: 6
    // Explanation: 
    // The given array is the same as example 1.
    // We have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. 
    // The sum of the numbers from index le = 3 to ri = 4 is 3 + 3 = 6.
    fmt.Println(rangeSum([]int{1,2,3,4},4,3,4)) // 6
    // Example 3:
    // Input: nums = [1,2,3,4], n = 4, left = 1, right = 10
    // Output: 50
    fmt.Println(rangeSum([]int{1,2,3,4},4,1,10)) // 50

    fmt.Println(rangeSum1([]int{1,2,3,4},4,1,5)) // 13
    fmt.Println(rangeSum1([]int{1,2,3,4},4,3,4)) // 6
    fmt.Println(rangeSum1([]int{1,2,3,4},4,1,10)) // 50
}