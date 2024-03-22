package main

// 321. Create Maximum Number
// You are given two integer arrays nums1 and nums2 of lengths m and n respectively. 
// nums1 and nums2 represent the digits of two numbers. You are also given an integer k.
// Create the maximum number of length k <= m + n from digits of the two numbers. 
// The relative order of the digits from the same array must be preserved.
// Return an array of the k digits representing the answer.

// Example 1:
// Input: nums1 = [3,4,6,5], nums2 = [9,1,2,5,8,3], k = 5
// Output: [9,8,6,5,3]

// Example 2:
// Input: nums1 = [6,7], nums2 = [6,0,4], k = 5
// Output: [6,7,6,0,4]

// Example 3:
// Input: nums1 = [3,9], nums2 = [8,9], k = 3
// Output: [9,8,9]
 
// Constraints:
//     m == nums1.length
//     n == nums2.length
//     1 <= m, n <= 500
//     0 <= nums1[i], nums2[i] <= 9
//     1 <= k <= m + n

import "fmt"

// stack
func maxNumber(nums1 []int, nums2 []int, k int) []int {
    max := func (a, b int) int { if a > b { return a; }; return b; }
    min := func (a, b int) int { if a < b { return a; }; return b; }

    n,m := len(nums1), len(nums2)
    // Each number is definitely smaller than 256 => uint8
    n1 := make([]uint8, n)
    for i, n := range nums1 {
        n1[i] = uint8(n)
    }
    n2 := make([]uint8, m)
    for i, n := range nums2 {
        n2[i] = uint8(n)
    }

    // The goal is to pick k items from the two arrays so that the combination
    // is the greatest possible sequence. The problem is that it is not clear how
    // many items should be picked from each array, let's call this (k1, k2) where
    // k = k1+k2. Assuming k == 2, then (k1, k2) can take on (0,2), (1,1), (2,0).
    // Then combine the max sequence given k1 items in nums1, and k2 items 
    // in nums2.
    fromFirst := min(n, k)
    fromSecond := max(0, k-fromFirst)

    var bestResult []uint8
    for fromFirst >= 0 && fromSecond <= min(k, m) {
        a := maxNumSingle(n1, fromFirst)
        b := maxNumSingle(n2, fromSecond)
        merged := merge(a, b, fromFirst, fromSecond)
        if len(bestResult) == 0 || greater(merged, bestResult, k, k, 0, 0) {
            bestResult = merged
        }
        fromFirst--
        fromSecond++
    }
    res := make([]int, k)
    for i, num := range bestResult {
        res[i] = int(num)
    }
    return res
}

// merge merges the two arrays a and b optimally.
func merge(a, b []uint8, n, m int) []uint8 {
    res := make([]uint8, m+n)
    var i, j int
    for k := 0; k < m+n; k++ {
        if greater(a, b, n, m, i, j) {
            res[k] = a[i]
            i++
        } else {
            res[k] = b[j]
            j++
        }
    }
    return res
}

// greater compares the provided lists of integers. If the length of one array
// is shorted than another, and they are otherwise equal, the longer array is
// returned. The reasoning behind this is that the partial array is always at
// least as good as the full array. For example, given [1,2], [1,2,3], choosing
// the first shorter array would force the use of 1 in the second array, which
// is sub-optimal.
func greater(a, b []uint8, n, m, i, j int) bool {
    for ; i < n && j < m; i, j = i+1, j+1 {
        if a[i] != b[j] {
            return a[i] > b[j]
        }
    }
    return i != n
}

// maxNumSingle calculates the max sequence in nums of length k.
func maxNumSingle(nums []uint8, k int) []uint8 {
    // Stack contains k elements sorted in descending order until it is absolutely
    // necessary to add more elements. For example, if k == len(nums), then the
    // result is simply the entire nums array. If k == len(nums)-1, it is ok to
    // bubble a number at most once.
    stack := make(uint8Stack, 0, k)
    n := len(nums)
    for i, num := range nums {
        // If num is greater than any elements in the stack, and there is enough
        // elements in nums to fill up the remainder, then clean up the stack to
        // make room for num.
        itemsLeft := n - i
        for len(stack) > 0 && itemsLeft > k-len(stack) && num > stack.peek() {
            stack.pop()
        }
        if len(stack) < k {
            stack.push(num)
        }
    }
    return stack
}

type uint8Stack []uint8

func (s uint8Stack) peek() uint8 {
    return s[len(s)-1]
}

func (s *uint8Stack) pop() uint8 {
    n := len(*s)
    it := (*s)[n-1]
    *s = (*s)[:n-1]
    return it
}

func (s *uint8Stack) push(x uint8) {
    *s = append(*s, x)
}


func maxNumber1(nums1, nums2 []int, k int) []int {
    maxSubsequence := func(a []int, k int) (s []int) {
        for i, v := range a {
            for len(s) > 0 && len(s)+len(a)-1-i >= k && v > s[len(s)-1] {
                s = s[:len(s)-1]
            }
            if len(s) < k {
                s = append(s, v)
            }
        }
        return
    }
    lexicographicalLess := func(a, b []int) bool {
        for i := 0; i < len(a) && i < len(b); i++ {
            if a[i] != b[i] {
                return a[i] < b[i]
            }
        }
        return len(a) < len(b)
    }
    merge := func (a, b []int) []int {
        merged := make([]int, len(a)+len(b))
        for i := range merged {
            if lexicographicalLess(a, b) {
                merged[i], b = b[0], b[1:]
            } else {
                merged[i], a = a[0], a[1:]
            }
        }
        return merged
    }

    var res []int
    start := 0
    if k > len(nums2) {
        start = k - len(nums2)
    }
    for i := start; i <= k && i <= len(nums1); i++ {
        s1 := maxSubsequence(nums1, i)
        s2 := maxSubsequence(nums2, k-i)
        merged := merge(s1, s2)
        if lexicographicalLess(res, merged) {
            res = merged
        }
    }
    return res
}

func main() {
    fmt.Println(maxNumber([]int{3,4,6,5}, []int{9,1,2,5,8,3}, 5)) // [9,8,6,5,3]
    fmt.Println(maxNumber([]int{6,7}, []int{6,0,4}, 5)) // [6,7,6,0,4]
    fmt.Println(maxNumber([]int{3,9}, []int{8,9}, 3)) // [9,8,9]

    fmt.Println(maxNumber1([]int{3,4,6,5}, []int{9,1,2,5,8,3}, 5)) // [9,8,6,5,3]
    fmt.Println(maxNumber1([]int{6,7}, []int{6,0,4}, 5)) // [6,7,6,0,4]
    fmt.Println(maxNumber1([]int{3,9}, []int{8,9}, 3)) // [9,8,9]
}