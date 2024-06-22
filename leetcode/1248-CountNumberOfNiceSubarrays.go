package main

// 1248. Count Number of Nice Subarrays
// Given an array of integers nums and an integer k. 
// A continuous subarray is called nice if there are k odd numbers on it.

// Return the number of nice sub-arrays.

// Example 1:
// Input: nums = [1,1,2,1,1], k = 3
// Output: 2
// Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].

// Example 2:
// Input: nums = [2,4,6], k = 1
// Output: 0
// Explanation: There are no odd numbers in the array.

// Example 3:
// Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
// Output: 16
 
// Constraints:
//     1 <= nums.length <= 50000
//     1 <= nums[i] <= 10^5
//     1 <= k <= nums.length

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
    prefixSums := map[int]int{0: 1}
    sum, res := 0, 0
    for _, v := range nums {
        sum += v % 2 // 统计奇数
        res += prefixSums[sum - k]
        prefixSums[sum]++
    }
    // fmt.Println(prefixSums)
    return res
}

func numberOfSubarrays1(nums []int, k int) int {
    for i, x := range nums {
        nums[i] = x&1
    }
    res, left1, left2, sum1, sum2 := 0, 0, 0, 0, 0
    for right, x := range nums {
        sum1 += x
        sum2 += x
        for left1 <= right && sum1 > k {
            sum1 -= nums[left1]
            left1++
        }
        for left2 <= right && sum2 >= k {
            sum2 -= nums[left2]
            left2++
        }
        res += left2-left1
    }
    return res
}

func numberOfSubarrays2(nums []int, k int) int {
    res, q, k_in_q, lr := 0, []int{}, 0, []int{0}
    for _, v := range nums {
        if v%2 != 0 && k_in_q == k {
            res += (lr[0]+1) * (lr[len(lr)-1]+1)
            q = append(q[lr[0]+1:], v)
            lr = lr[1:]
            lr = append(lr, 0)
        } else if v%2 != 0 {
            k_in_q++
            q = append(q, v)
            lr = append(lr, 0)
        } else {
            q = append(q, v)
            lr[len(lr)-1]++
        }
    }
    if len(lr) > k {
        res += (lr[0]+1) * (lr[len(lr)-1]+1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,2,1,1], k = 3
    // Output: 2
    // Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].
    fmt.Println(numberOfSubarrays([]int{1,1,2,1,1}, 3)) // 2
    // Example 2: 
    // Input: nums = [2,4,6], k = 1
    // Output: 0
    // Explanation: There are no odd numbers in the array.
    fmt.Println(numberOfSubarrays([]int{2,4,6}, 1)) // 0
    // Example 3:
    // Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
    // Output: 16
    fmt.Println(numberOfSubarrays([]int{2,2,2,1,2,2,1,2,2,2}, 2)) // 16

    fmt.Println(numberOfSubarrays1([]int{1,1,2,1,1}, 3)) // 2
    fmt.Println(numberOfSubarrays1([]int{2,4,6}, 1)) // 0
    fmt.Println(numberOfSubarrays1([]int{2,2,2,1,2,2,1,2,2,2}, 2)) // 16

    fmt.Println(numberOfSubarrays2([]int{1,1,2,1,1}, 3)) // 2
    fmt.Println(numberOfSubarrays2([]int{2,4,6}, 1)) // 0
    fmt.Println(numberOfSubarrays2([]int{2,2,2,1,2,2,1,2,2,2}, 2)) // 16
}