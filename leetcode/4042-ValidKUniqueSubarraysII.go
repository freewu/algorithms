package main

// 4042. Valid K-Unique Subarrays II
// You are given an integer array nums of length n and an integer k.

// You are also given integers l0 and r0, which define the first query, and an integer q, representing the total number of queries to process.

// A subarray nums[li..ri] is considered valid if:
//     1. It contains exactly k distinct numbers, and
//     2. Every distinct number in it occurs an even number of times.

// For query 0, set l0 = l0 and r0 = r0.

// Let ansi denote the result of the ith query, where ansi = 1 if nums[li..ri] is valid, and ansi = 0 otherwise.

// For each i > 0, generate the next query as follows:
//     1. If ansi-1 = 1, set gi-1 = li-1 + ri-1. Otherwise, set gi-1 = ri-1 - li-1.
//     2. Compute li = (li-1 XOR gi-1) % n and ri = (ri-1 XOR gi-1) % n.
//     3. If li > ri, swap them.
    
// Return a boolean array ans, where ans[i] is true if ansi = 1, and false otherwise.

// Example 1:
// Input: nums = [1,2,2,1], k = 2, l0 = 1, r0 = 2, q = 2
// Output: [false,true]
// Explanation:
// i	[li, ri]	Subarray	Distinct numbers	Counts	Validity check	ans[i]	[li+1, ri+1]
// 0	[1, 2]	[2, 2]	{2} → 1	{2:2}	false: The subarray contains fewer than k distinct numbers.	ans0 = 0	g0 = 2 - 1 = 1
// 			l1 = (1 XOR 1) % 4 = 0
// 			r1 = (2 XOR 1) % 4 = 3
// 1	[0, 3]	[1, 2, 2, 1]	{1,2} → 2	{1:2,2:2}	true: The subarray contains exactly k distinct numbers, each occurring an even number of times.	ans1 = 1	-
// Thus, ans = [false, true].

// Example 2:
// Input: nums = [1,2,3,3,4], k = 1, l0 = 2, r0 = 3, q = 2
// Output: [true,false]
// Explanation:
// i	[li, ri]	Subarray	Distinct numbers	Counts	Validity check	ans[i]	[li+1, ri+1]
// 0	[2, 3]	[3, 3]	{3} → 1	{3:2}	true: The subarray contains exactly k distinct numbers, each occurring an even number of times.	ans0 = 1	g0 = 2 + 3 = 5
// 			l1 = (2 XOR 5) % 5 = 7 % 5 = 2
// 			r1 = (3 XOR 5) % 5 = 6 % 5 = 1
// Since l1 > r1, swap them to obtain [l1, r1] = [1, 2].
// 1	[1, 2]	[2, 3]	{2,3} → 2	{2:1,3:1}	false: The subarray contains 2 distinct numbers instead of exactly k = 1.	ans1 = 0	-
// Thus, ans = [true, false].

// Constraints:
//     2 <= n == nums.length <= 5 × 10^5
//     1 <= nums[i] <= 5 × 10^5
//     1 <= k <= n
//     0 <= l0 < r0 <= n - 1
//     1 <= q <= 5 × 10^5

import "fmt"
import "math/rand"

func validSubarrays(nums []int, k int, l0 int, r0 int, q int) []bool {
    res, n := make([]bool, q), len(nums)
    sum, mp := make([]uint64, n + 1), make(map[int]uint64)
    for i, v := range nums {
        // 把 nums[i] 映射成一个随机的 uint64
        if _, ok := mp[v]; !ok {
            mp[v] = rand.Uint64()
        }
        sum[i+1] = sum[i] ^ mp[v]
    }
    calcLeft := func(k int) []int {
        left, lefts, count := 0, make([]int, n), make(map[int]int)
        for i, v := range nums {
            count[v]++
            for len(count) >= k {
                val := nums[left]
                if count[val] > 1 {
                    count[val]--
                } else {
                    delete(count, val) // 保证 len(count) 是窗口内的不同元素个数
                }
                left++
            }
            lefts[i] = left
        }
        return lefts
    }
    l1, l2 := calcLeft(k + 1), calcLeft(k)
    l, r := l0, r0
    for i := range res {
        if i > 0 {
            g := r - l
            if res[i-1] {
                g = l + r
            }
            l = (l ^ g) % n
            r = (r ^ g) % n
            if l > r {
                l, r = r, l
            }
        }
        res[i] = sum[r+1] == sum[l] && l1[r] <= l && l < l2[r]
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,1], k = 2, l0 = 1, r0 = 2, q = 2
    // Output: [false,true]
    // Explanation:
    // i	[li, ri]	Subarray	Distinct numbers	Counts	Validity check	ans[i]	[li+1, ri+1]
    // 0	[1, 2]	[2, 2]	{2} → 1	{2:2}	false: The subarray contains fewer than k distinct numbers.	ans0 = 0	g0 = 2 - 1 = 1
    // 			l1 = (1 XOR 1) % 4 = 0
    // 			r1 = (2 XOR 1) % 4 = 3
    // 1	[0, 3]	[1, 2, 2, 1]	{1,2} → 2	{1:2,2:2}	true: The subarray contains exactly k distinct numbers, each occurring an even number of times.	ans1 = 1	-
    // Thus, ans = [false, true].
    fmt.Println(validSubarrays([]int{1,2,2,1}, 2, 1, 2, 2)) // [false,true]
    // Example 2:
    // Input: nums = [1,2,3,3,4], k = 1, l0 = 2, r0 = 3, q = 2
    // Output: [true,false]
    // Explanation:
    // i	[li, ri]	Subarray	Distinct numbers	Counts	Validity check	ans[i]	[li+1, ri+1]
    // 0	[2, 3]	[3, 3]	{3} → 1	{3:2}	true: The subarray contains exactly k distinct numbers, each occurring an even number of times.	ans0 = 1	g0 = 2 + 3 = 5
    // 			l1 = (2 XOR 5) % 5 = 7 % 5 = 2
    // 			r1 = (3 XOR 5) % 5 = 6 % 5 = 1
    // Since l1 > r1, swap them to obtain [l1, r1] = [1, 2].
    // 1	[1, 2]	[2, 3]	{2,3} → 2	{2:1,3:1}	false: The subarray contains 2 distinct numbers instead of exactly k = 1.	ans1 = 0	-
    // Thus, ans = [true, false].
    fmt.Println(validSubarrays([]int{1,2,3,3,4}, 1, 2, 3, 2)) // [true,false]

    fmt.Println(validSubarrays([]int{1,2,3,4,5,6,7,8,9}, 2, 1, 2, 2)) // [false false]
    fmt.Println(validSubarrays([]int{9,8,7,6,5,4,3,2,1}, 2, 1, 2, 2)) // [false false]
}