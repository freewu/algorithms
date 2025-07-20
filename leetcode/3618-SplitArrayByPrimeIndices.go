package main

// 3618. Split Array by Prime Indices
// You are given an integer array nums.

// Split nums into two arrays A and B using the following rule:
//     1. Elements at prime indices in nums must go into array A.
//     2. All other elements must go into array B.

// Return the absolute difference between the sums of the two arrays: |sum(A) - sum(B)|.

// A prime number is a natural number greater than 1 with only two factors, 1 and itself.

// Note: An empty array has a sum of 0.

// Example 1:
// Input: nums = [2,3,4]
// Output: 1
// Explanation:
// The only prime index in the array is 2, so nums[2] = 4 is placed in array A.
// The remaining elements, nums[0] = 2 and nums[1] = 3 are placed in array B.
// sum(A) = 4, sum(B) = 2 + 3 = 5.
// The absolute difference is |4 - 5| = 1.

// Example 2:
// Input: nums = [-1,5,7,0]
// Output: 3
// Explanation:
// The prime indices in the array are 2 and 3, so nums[2] = 7 and nums[3] = 0 are placed in array A.
// The remaining elements, nums[0] = -1 and nums[1] = 5 are placed in array B.
// sum(A) = 7 + 0 = 7, sum(B) = -1 + 5 = 4.
// The absolute difference is |7 - 4| = 3.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"
import "math"

func splitArray(nums []int) int64 {
    n := len(nums)
    if n == 1 { return int64(nums[0]) }
    prime := make([]bool,n)
    prime[0],  prime[1] = true, true
    for i := 2; i <= int(math.Sqrt(float64(n + 1))); i++ { // generate prime number
        for j := i + i; j < n; j = j + i{
            prime[j]=true
        }
    }
    sum1, sum2 := 0, 0
    for i := 0; i < n; i++ {
        if prime[i] {
            sum1 += nums[i]
        } else {
            sum2 += nums[i]
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    return int64(abs(sum2 - sum1))
}

func splitArray1(nums []int) int64 {
    res, n := 0, len(nums)
    prime := make([]bool, n + 1)
    for i := range prime {
        prime[i] = true
    }
    prime[0], prime[1] = false, false
    for i := 2; i <= n; i++ {
        if !prime[i] { continue }
        for j := 2 * i; j < n; j += i {
            prime[j] = false
        }
    }
    for i, v := range nums {
        if prime[i] {
            res += v
        } else {
            res -= v
        }
    }
    if res < 0 { return int64(-res) }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [2,3,4]
    // Output: 1
    // Explanation:
    // The only prime index in the array is 2, so nums[2] = 4 is placed in array A.
    // The remaining elements, nums[0] = 2 and nums[1] = 3 are placed in array B.
    // sum(A) = 4, sum(B) = 2 + 3 = 5.
    // The absolute difference is |4 - 5| = 1.
    fmt.Println(splitArray([]int{2,3,4})) // 1
    // Example 2:
    // Input: nums = [-1,5,7,0]
    // Output: 3
    // Explanation:
    // The prime indices in the array are 2 and 3, so nums[2] = 7 and nums[3] = 0 are placed in array A.
    // The remaining elements, nums[0] = -1 and nums[1] = 5 are placed in array B.
    // sum(A) = 7 + 0 = 7, sum(B) = -1 + 5 = 4.
    // The absolute difference is |7 - 4| = 3.
    fmt.Println(splitArray([]int{-1,5,7,0})) // 3

    fmt.Println(splitArray([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(splitArray([]int{9,8,7,6,5,4,3,2,1})) // 7

    fmt.Println(splitArray1([]int{2,3,4})) // 1
    fmt.Println(splitArray1([]int{-1,5,7,0})) // 3
    fmt.Println(splitArray1([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(splitArray1([]int{9,8,7,6,5,4,3,2,1})) // 7
}