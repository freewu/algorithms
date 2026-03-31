package main

// 3886. Sum of Sortable Integers
// You are given an integer array nums of length n.

// An integer k is called sortable if k divides n and you can sort nums in non-decreasing order by sequentially performing the following operations:
//     1. Partition nums into consecutive subarrays of length k.
//     2. Cyclically rotate each subarray independently any number of times to the left or to the right.

// Return an integer denoting the sum of all possible sortable integers k.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [3,1,2]
// Output: 3
// Explanation:​​​​​​​
// For n = 3, possible divisors are 1 and 3.
// For k = 1: each subarray has one element. No rotation can sort the array.
// For k = 3: the single subarray [3, 1, 2] can be rotated once to produce [1, 2, 3], which is sorted.
// Only k = 3 is sortable. Hence, the answer is 3.

// Example 2:
// Input: nums = [7,6,5]
// Output: 0
// Explanation:
// For n = 3, possible divisors are 1 and 3.
// For k = 1: each subarray has one element. No rotation can sort the array.
// For k = 3: the single subarray [7, 6, 5] cannot be rotated into non-decreasing order.
// No k is sortable. Hence, the answer is 0.

// Example 3:
// Input: nums = [5,8]
// Output: 3
// Explanation:​​​​​​​
// For n = 2, possible divisors are 1 and 2.
// Since [5, 8] is already sorted, every divisor is sortable. Hence, the answer is 1 + 2 = 3.
 
// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"


func sortableIntegers(nums []int) int {
    res := 0
    getDivisors := func(n int) []int {
        res := []int{}
        for i := 1; i * i <= n; i++ {
            if n % i == 0 {
                res = append(res, i)
                if i != n/i {
                    res = append(res, n/i)
                }
            }
        }
        return res
    }
    check := func(nums []int, k int) bool {
        n, last:= len(nums), -1000000000
        for i := 0; i < n; i += k {
            count, mn, mx := 0, 1000000000, -1000000000
            for j := i; j < i+k; j++ {
                if nums[j] < mn { mn = nums[j] }
                if nums[j] > mx { mx = nums[j] }
                next := nums[i]
                if j != i+k-1 {
                    next = nums[j+1]
                }
                if nums[j] > next { count++ }
            }
            if count > 1 || last > mn {
                return false
            }
            last = mx
        }
        return true
    }
    for _, k := range getDivisors(len(nums)) {
        if check(nums, k) {
            res += k
        }
    }
    return res
}

func sortableIntegers1(nums []int) int {
    res, n := 0, len(nums)
    divisors := []int{}
    for i := 1; i * i <= n; i++ {
        if n % i == 0 {
            divisors = append(divisors, i)
            if i * i != n {
                divisors = append(divisors, n / i)
            }
        }
    }
    checkIncr := func(i, j int) (bool, int, int) {
        breakPoint := -1
        for p := i + 1; p <= j; p++ {
            if nums[p] < nums[p-1] {
                if nums[p] > nums[i] {
                    return false, 0, 0
                }
                breakPoint = p
                break
            }
        }
        if breakPoint == -1 {
            return true, nums[i], nums[j]
        }
        for p := breakPoint + 1; p <= j; p++ {
            if nums[p] < nums[p-1] {
                return false, 0, 0
            }
            if nums[p] > nums[i] {
                return false, 0, 0
            }
        }
        return true, nums[breakPoint], nums[breakPoint-1]
    }
    check := func(x int) bool {
        last := -1 << 61
        for i := 0; i < n; i += x {
            ok, mn, mx := checkIncr(i, i + x - 1)
            if !ok {
                return false
            }
            if mn < last {
                return false
            }
            last = mx
        }
        return true
    }
    for _, div := range divisors {
        if check(div) {
            res += div
        }
    }
    return res
}

func sortableIntegers2(nums []int) int {
    res, n, p := 0, len(nums), len(nums)
    next := make([]int, n)
    next[n - 1] = n
    for i := n - 2; i >= 0; i -- {
        if nums[i] > nums[i+1] {
            p = i
        }
        next[i] = p
    }
    calc := func(v int) int {
        last := 0
        for r := v - 1; r < n; r += v {
            l := r - v + 1  
            m := next[l]
            if m >= r {
                if nums[l] < last {
                    return 0
                }
                last = nums[r]
            } else {
                if next[m + 1] < r || nums[m + 1] < last || nums[l] < nums[r] {
                    return 0
                }
                last = nums[m]
            }
        }
        return v
    }
    for i := 1; i * i <= n; i ++ {
        if n % i == 0 {
            res += calc(i)
            if i != n / i {
                res += calc(n / i)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,1,2]
    // Output: 3
    // Explanation:​​​​​​​
    // For n = 3, possible divisors are 1 and 3.
    // For k = 1: each subarray has one element. No rotation can sort the array.
    // For k = 3: the single subarray [3, 1, 2] can be rotated once to produce [1, 2, 3], which is sorted.
    // Only k = 3 is sortable. Hence, the answer is 3.
    fmt.Println(sortableIntegers([]int{3,1,2})) // 3
    // Example 2:
    // Input: nums = [7,6,5]
    // Output: 0
    // Explanation:
    // For n = 3, possible divisors are 1 and 3.
    // For k = 1: each subarray has one element. No rotation can sort the array.
    // For k = 3: the single subarray [7, 6, 5] cannot be rotated into non-decreasing order.
    // No k is sortable. Hence, the answer is 0.
    fmt.Println(sortableIntegers([]int{7,6,5})) // 0
    // Example 3:
    // Input: nums = [5,8]
    // Output: 3
    // Explanation:​​​​​​​
    // For n = 2, possible divisors are 1 and 2.
    // Since [5, 8] is already sorted, every divisor is sortable. Hence, the answer is 1 + 2 = 3.
    fmt.Println(sortableIntegers([]int{5,8})) // 3

    fmt.Println(sortableIntegers([]int{1,2,3,4,5,6,7,8,9})) // 13
    fmt.Println(sortableIntegers([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(sortableIntegers1([]int{3,1,2})) // 3
    fmt.Println(sortableIntegers1([]int{7,6,5})) // 0
    fmt.Println(sortableIntegers1([]int{5,8})) // 3
    fmt.Println(sortableIntegers1([]int{1,2,3,4,5,6,7,8,9})) // 13
    fmt.Println(sortableIntegers1([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(sortableIntegers2([]int{3,1,2})) // 3
    fmt.Println(sortableIntegers2([]int{7,6,5})) // 0
    fmt.Println(sortableIntegers2([]int{5,8})) // 3
    fmt.Println(sortableIntegers2([]int{1,2,3,4,5,6,7,8,9})) // 13
    fmt.Println(sortableIntegers2([]int{9,8,7,6,5,4,3,2,1})) // 0
}