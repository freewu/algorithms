package main

// 4076. Longest Resilient Subarray II
// You are given an integer array nums and an integer k.

// A subarray is resilient if, for every position in it, deleting the element at that position leaves the remaining elements with a sum divisible by k.

// For a subarray of length 1, the remaining sum is 0, which is divisible by k.

// Return the length of the longest resilient subarray.

// Example 1:
// Input: nums = [2,4,6,3], k = 2
// Output: 3
// Explanation:
// The subarray [2, 4, 6] sums to 12. 
// Deleting 2, 4, or 6 leaves a sum of 10, 8, or 6, respectively, each of which is divisible by 2. 
// Therefore, the answer is 3.

// Example 2:
// Input: nums = [1,4,7,1], k = 3
// Output: 4
// Explanation:
// The subarray [1, 4, 7, 1] sums to 13. 
// Deleting the first 1, the 4, the 7, or the last 1 leaves a sum of 12, 9, 6, or 12, respectively, each of which is divisible by 3. 
// Therefore, the answer is 4.

// Example 3:
// Input: nums = [5,5,4,8], k = 4
// Output: 2
// Explanation:
// The subarray [5, 5] is not resilient, since deleting either 5 leaves a sum of 5, which is not divisible by 4.
// The subarray [4, 8] sums to 12. Deleting 4 or 8 leaves a sum of 8 or 4, respectively, each of which is divisible by 4. 
// Therefore, the answer is 2.

// Constraints:
//     - 1 <= nums.length <= 10^{5}
//     - 1 <= nums[i] <= 10^{5}
//     - 1 <= k <= 10^{5}

import "fmt"

func resilientSubarray(nums []int, k int) int {
    res, count := 0, 0
    for i, v := range nums {
        r := v % k
        if i > 0 && r == nums[i - 1] % k {
            count++
        } else { // 新的模 k 同余段
            count = 1
        }
        // 去掉一个数后，剩余元素之和是 k 的倍数
        if r * (count - 1) % k == 0 {
            res = max(res, count)
        }
    }
    return res
}

func resilientSubarray1(nums []int, k int) int {
    res, n := 0, len(nums)
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i := 0; i < n; {
        r, j := nums[i] % k, i
        for j < n && nums[j]%k == r {
            j++
        }
        length := j - i
        t := k / gcd(r, k)
        best := 1 + ((length - 1) / t) * t
        if best > res {
            res = best
        }
        i = j
    }
    return res
}

func resilientSubarray2(nums []int, k int) int {
    for i := range nums {
        nums[i] = nums[i] % k
    }
    calc := func(arr []int, k int) int {
        res, sum := 1, 0
        for i, v := range arr {
            sum += v
            if sum%k == 0 {
                if v == 0 {
                    res = max(res, i+1)
                } else {
                    if i < len(arr)-1 {
                        res = max(res, i+2)
                    }
                }
            }
        }
        return res
    }
    res, start, curr := 0, 0, nums[0]
    for i := 0; i < len(nums); i++ {
        if curr != nums[i] {
            res = max(res, calc(nums[start:i], k))
            curr = nums[i]
            start = i
        }
    }
    return max(res, calc(nums[start:len(nums)], k))
}

func main() {
    // Example 1:
    // Input: nums = [2,4,6,3], k = 2
    // Output: 3
    // Explanation:
    // The subarray [2, 4, 6] sums to 12. 
    // Deleting 2, 4, or 6 leaves a sum of 10, 8, or 6, respectively, each of which is divisible by 2. 
    // Therefore, the answer is 3.
    fmt.Println(resilientSubarray([]int{2,4,6,3}, 2)) // 3
    // Example 2:
    // Input: nums = [1,4,7,1], k = 3
    // Output: 4
    // Explanation:
    // The subarray [1, 4, 7, 1] sums to 13. 
    // Deleting the first 1, the 4, the 7, or the last 1 leaves a sum of 12, 9, 6, or 12, respectively, each of which is divisible by 3. 
    // Therefore, the answer is 4.
    fmt.Println(resilientSubarray([]int{1,4,7,1}, 3)) // 4
    // Example 3:
    // Input: nums = [5,5,4,8], k = 4
    // Output: 2
    // Explanation:
    // The subarray [5, 5] is not resilient, since deleting either 5 leaves a sum of 5, which is not divisible by 4.
    // The subarray [4, 8] sums to 12. Deleting 4 or 8 leaves a sum of 8 or 4, respectively, each of which is divisible by 4. 
    // Therefore, the answer is 2.
    fmt.Println(resilientSubarray([]int{5,5,4,8}, 4)) // 2

    fmt.Println(resilientSubarray([]int{1,2,3,4,5,6,7,8,9}, 4)) // 1
    fmt.Println(resilientSubarray([]int{9,8,7,6,5,4,3,2,1}, 4)) // 1

    fmt.Println(resilientSubarray1([]int{2,4,6,3}, 2)) // 3
    fmt.Println(resilientSubarray1([]int{1,4,7,1}, 3)) // 4
    fmt.Println(resilientSubarray1([]int{5,5,4,8}, 4)) // 2
    fmt.Println(resilientSubarray1([]int{1,2,3,4,5,6,7,8,9}, 4)) // 1
    fmt.Println(resilientSubarray1([]int{9,8,7,6,5,4,3,2,1}, 4)) // 1

    fmt.Println(resilientSubarray2([]int{2,4,6,3}, 2)) // 3
    fmt.Println(resilientSubarray2([]int{1,4,7,1}, 3)) // 4
    fmt.Println(resilientSubarray2([]int{5,5,4,8}, 4)) // 2
    fmt.Println(resilientSubarray2([]int{1,2,3,4,5,6,7,8,9}, 4)) // 1
    fmt.Println(resilientSubarray2([]int{9,8,7,6,5,4,3,2,1}, 4)) // 1
}