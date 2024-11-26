package main

// 2040. Kth Smallest Product of Two Sorted Arrays
// Given two sorted 0-indexed integer arrays nums1 and nums2 as well as an integer k, 
// return the kth (1-based) smallest product of nums1[i] * nums2[j] where 0 <= i < nums1.length and 0 <= j < nums2.length.

// Example 1:
// Input: nums1 = [2,5], nums2 = [3,4], k = 2
// Output: 8
// Explanation: The 2 smallest products are:
// - nums1[0] * nums2[0] = 2 * 3 = 6
// - nums1[0] * nums2[1] = 2 * 4 = 8
// The 2nd smallest product is 8.

// Example 2:
// Input: nums1 = [-4,-2,0,3], nums2 = [2,4], k = 6
// Output: 0
// Explanation: The 6 smallest products are:
// - nums1[0] * nums2[1] = (-4) * 4 = -16
// - nums1[0] * nums2[0] = (-4) * 2 = -8
// - nums1[1] * nums2[1] = (-2) * 4 = -8
// - nums1[1] * nums2[0] = (-2) * 2 = -4
// - nums1[2] * nums2[0] = 0 * 2 = 0
// - nums1[2] * nums2[1] = 0 * 4 = 0
// The 6th smallest product is 0.

// Example 3:
// Input: nums1 = [-2,-1,0,1,2], nums2 = [-3,-1,2,4,5], k = 3
// Output: -6
// Explanation: The 3 smallest products are:
// - nums1[0] * nums2[4] = (-2) * 5 = -10
// - nums1[0] * nums2[3] = (-2) * 4 = -8
// - nums1[4] * nums2[0] = 2 * (-3) = -6
// The 3rd smallest product is -6.

// Constraints:
//     1 <= nums1.length, nums2.length <= 5 * 10^4
//     -10^5 <= nums1[i], nums2[j] <= 10^5
//     1 <= k <= nums1.length * nums2.length
//     nums1 and nums2 are sorted.

import "fmt"
import "sort"

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    m, n := len(nums1), len(nums2)
    a, b := max(abs(nums1[0]), abs(nums1[m - 1])), max(abs(nums2[0]), abs(nums2[n - 1]))
    r := int64(a) * int64(b)
    l := -r
    count := func(p int64) int64 {
        res := int64(0)
        for _, x := range nums1 {
            if x > 0 {
                l, r := 0, n
                for l < r {
                    mid := (l + r) >> 1
                    if int64(x) * int64(nums2[mid]) > p {
                        r = mid
                    } else {
                        l = mid + 1
                    }
                }
                res += int64(l)
            } else if x < 0 {
                l, r := 0, n
                for l < r {
                    mid := (l + r) >> 1
                    if int64(x) * int64(nums2[mid]) <= p {
                        r = mid
                    } else {
                        l = mid + 1
                    }
                }
                res += int64(n - l)
            } else if p >= 0 {
                res += int64(n)
            }
        }
        return res
    }
    for l < r {
        mid := (l + r) >> 1
        if count(mid) >= k {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}

func kthSmallestProduct1(nums1 []int, nums2 []int, k int64) int64 {
    n1, n2 := len(nums1), len(nums2)
    p1 := sort.Search(n1, func(v int) bool { return nums1[v] >= 0 })
    p2 := sort.Search(n2, func(v int) bool { return nums2[v] >= 0 })
    res := sort.Search(1e11, func(v int) bool {
        v -= 1e10
        c := 0
        if v < 0 {
            j := n2 - 1
            for i := p1-1; i >=0; i-- {
                for j >= 0 && nums2[j] * nums1[i] <= v {
                    j--
                }
                c += n2 - 1 - j
            }
            i := n1 - 1
            for j := p2 - 1; j >=0; j-- {
                for i >= 0 && nums2[j] * nums1[i] <= v {
                    i--
                }
                c += n1 - 1 - i
            }
        } else {
            c += p1 * (n2 - p2) + p2 * (n1 - p1)
            j := p2
            for i := n1 - 1; i >= p1; i-- {
                for j < n2 && nums1[i] * nums2[j] <= v {
                    j++
                }
                c += j - p2
            }
            j = p2 - 1
            for i := 0; i < p1; i++ {
                for j >= 0 && nums1[i] * nums2[j] <= v {
                    j--
                }
                c += p2 - j - 1
            }
        }
        return int64(c) >= k
    }) - 1e10
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums1 = [2,5], nums2 = [3,4], k = 2
    // Output: 8
    // Explanation: The 2 smallest products are:
    // - nums1[0] * nums2[0] = 2 * 3 = 6
    // - nums1[0] * nums2[1] = 2 * 4 = 8
    // The 2nd smallest product is 8.
    fmt.Println(kthSmallestProduct([]int{2,5}, []int{3,4}, 2)) // 8
    // Example 2:
    // Input: nums1 = [-4,-2,0,3], nums2 = [2,4], k = 6
    // Output: 0
    // Explanation: The 6 smallest products are:
    // - nums1[0] * nums2[1] = (-4) * 4 = -16
    // - nums1[0] * nums2[0] = (-4) * 2 = -8
    // - nums1[1] * nums2[1] = (-2) * 4 = -8
    // - nums1[1] * nums2[0] = (-2) * 2 = -4
    // - nums1[2] * nums2[0] = 0 * 2 = 0
    // - nums1[2] * nums2[1] = 0 * 4 = 0
    // The 6th smallest product is 0.
    fmt.Println(kthSmallestProduct([]int{-4,-2,0,3}, []int{2,4}, 6)) // 0
    // Example 3:
    // Input: nums1 = [-2,-1,0,1,2], nums2 = [-3,-1,2,4,5], k = 3
    // Output: -6
    // Explanation: The 3 smallest products are:
    // - nums1[0] * nums2[4] = (-2) * 5 = -10
    // - nums1[0] * nums2[3] = (-2) * 4 = -8
    // - nums1[4] * nums2[0] = 2 * (-3) = -6
    // The 3rd smallest product is -6.
    fmt.Println(kthSmallestProduct([]int{-2,-1,0,1,2}, []int{-3,-1,2,4,5}, 3)) // -6

    fmt.Println(kthSmallestProduct1([]int{2,5}, []int{3,4}, 2)) // 8
    fmt.Println(kthSmallestProduct1([]int{-4,-2,0,3}, []int{2,4}, 6)) // 0
    fmt.Println(kthSmallestProduct1([]int{-2,-1,0,1,2}, []int{-3,-1,2,4,5}, 3)) // -6
}