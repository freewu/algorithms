package main

// 3901. Good Subsequence Queries
// You are given an integer array nums of length n and an integer p.

// A non-empty subsequence of nums is called good if:
//     1. Its length is strictly less than n.
//     2. The greatest common divisor (GCD) of its elements is exactly p.

// You are also given a 2D integer array queries of length q, where each queries[i] = [indi, vali] indicates that you should update nums[indi] to vali.

// After each query, determine whether there exists any good subsequence in the current array.

// Return the number of queries for which a good subsequence exists.

// A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

// The term gcd(a, b) denotes the greatest common divisor of a and b.

// Example 1:
// Input: nums = [4,8,12,16], p = 2, queries = [[0,3],[2,6]]
// Output: 1
// Explanation:
// i	| [indi, vali]  | Operation	Updated nums	            | Any good Subsequence
// 0	| [0, 3]	    | Update nums[0] to 3	[3, 8, 12, 16]	| No, as no subsequence has GCD exactly p = 2
// 1	| [2, 6]	    | Update nums[2] to 6	[3, 8, 6, 16]	| Yes, subsequence [8, 6] has GCD exactly p = 2
// Thus, the answer is 1.

// Example 2:
// Input: nums = [4,5,7,8], p = 3, queries = [[0,6],[1,9],[2,3]]
// Output: 2
// Explanation:
// i	| [indi, vali]  | Operation	Updated nums	            | Any good Subsequence    
// 0	| [0, 6]	    | Update nums[0] to 6	[6, 5, 7, 8]	| No, as no subsequence has GCD exactly p = 3
// 1	| [1, 9]	    | Update nums[1] to 9	[6, 9, 7, 8]	| Yes, subsequence [6, 9] has GCD exactly p = 3
// 2	| [2, 3]	    | Update nums[2] to 3	[6, 9, 3, 8]	| Yes, subsequence [6, 9, 3] has GCD exactly p = 3
// Thus, the answer is 2.

// Example 3:
// Input: nums = [5,7,9], p = 2, queries = [[1,4],[2,8]]
// Output: 0
// Explanation:
// i   | [indi, vali]  | Operation	Updated nums	        | Any good Subsequence
// 0	| [1, 4]	    | Update nums[1] to 4	[5, 4, 9]	| No, as no subsequence has GCD exactly p = 2
// 1	| [2, 8]	    | Update nums[2] to 8	[5, 4, 8]	| No, as no subsequence has GCD exactly p = 2
// Thus, the answer is 0.

// Constraints:
//     2 <= n == nums.length <= 5 * 10^4
//     1 <= nums[i] <= 5 * 10^4
//     1 <= queries.length <= 5 * 10^4
//     queries[i] = [indi, vali]
//     1 <= vali, p <= 5 * 10^4
//     0 <= indi <= n - 1

import "fmt"
import "math/bits"

var TargetGcd int

func gcd(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }

type SegmentTree []struct{ l, r, gcd int }

func (t SegmentTree) maintain(o int) {
    t[o].gcd = gcd(t[o << 1].gcd, t[o << 1 | 1].gcd)
}

func (t SegmentTree) build(a []int, o, l, r int) {
    t[o].l, t[o].r = l, r
    if l == r {
        if a[l] % TargetGcd == 0 {
            t[o].gcd = a[l]
        }
        return
    }
    m := (l + r) >> 1
    t.build(a, o << 1, l, m)
    t.build(a, o << 1 | 1, m + 1, r)
    t.maintain(o)
}

func (t SegmentTree) update(o, i, val int) {
    cur := &t[o]
    if cur.l == cur.r {
        if val%TargetGcd == 0 { 
            cur.gcd = val
        } else {
            cur.gcd = 0
        }
        return
    }
    m := (cur.l + cur.r) >> 1
    if i <= m {
        t.update(o<<1, i, val)
    } else {
        t.update(o<<1|1, i, val)
    }
    t.maintain(o)
}

func (t SegmentTree) query(o, l, r int) int {
    if l > r {
        return 0
    }
    if l <= t[o].l && t[o].r <= r {
        return t[o].gcd
    }
    m := (t[o].l + t[o].r) >> 1
    if r <= m {
        return t.query(o<<1, l, r)
    }
    if m < l {
        return t.query(o<<1|1, l, r)
    }
    return gcd(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func (t SegmentTree) check(n int) bool {
    for i := range n {
        if gcd(t.query(1, 0, i-1), t.query(1, i+1, n-1)) == TargetGcd {
            return true
        }
    }
    return false
}

func countGoodSubseq(nums []int, p int, queries [][]int) int {
    TargetGcd = p
    res, count := 0, 0   
    for _, v := range nums {
        if v % p == 0 {
            count++
        }
    }
    n := len(nums)
    t := make(SegmentTree, 2 <<bits.Len(uint(n-1)))
    t.build(nums, 1, 0, n-1)
    for _, q := range queries {
        i, v := q[0], q[1]
        if nums[i] % p == 0 {
            count--
        }
        if v % p == 0 {
            count++
        }
        nums[i] = v
        t.update(1, i, v)
        if t[1].gcd == p && (count < n || n > 7 || t.check(n)) {
            res++
        }
    }
    return res
}

func countGoodSubseq1(nums []int,p int,queries [][]int) int {  
    res, n := 0, len(nums)
    tree := make([]int,2 * n)
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    build := func() {
        for i := 0; i < n; i++ {
            if nums[i] % p == 0 {
                tree[n + i] = nums[i]
            } else {
                tree[n + i] = 0
            }
        }
        for i := n - 1; i > 0; i-- {
            tree[i] = gcd(tree[i << 1], tree[i << 1 | 1])
        }
    }
    update := func(index, val int) {
        i := n + index
        if val % p == 0 {
            tree[i] = val
        } else {
            tree[i] = 0
        }
        for i := i >> 1; i > 0; i = i >> 1 {
            tree[i] = gcd(tree[i << 1], tree[i << 1 | 1])   
        }
    }
    build()
    nonp := 0
    for _, v := range nums {
        if v % p != 0 {
            nonp++
        }
    }
    for _, q := range queries {
        idx,val:=q[0],q[1]
        if nums[idx]%p!=0{
            nonp--
        }
        nums[idx]=val
        if nums[idx] % p != 0 {
            nonp++
        }
        update(idx,val)
        total := tree[1]
        if total == p {
            if nonp > 0 {
                res++
            } else {
                if n >= 8 {
                    res++
                } else {
                    flag := false
                    for i := 0; i < n; i++ {
                        g := 0
                        for j := 0; j < n; j++ {
                            if i != j {
                                g = gcd(g, nums[j])
                            }
                        }
                        if g == p {
                            flag = true
                            break
                        }
                    }
                    if flag {
                        res++
                    }
                }
            }
        }
    }
    return res
}

func countGoodSubseq2(nums []int, p int, queries [][]int) int {
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    res, countDivisible, n := 0, 0, len(nums)
    transformed := make([]int, n)
    // Initialize transformed array and count of divisible elements
    for i, x := range nums {
        if x % p == 0 {
            transformed[i] = x / p
            countDivisible++
        } else {
            transformed[i] = 0
        }
    }
    // Build Segment Tree for Range GCD
    // Size of segment tree array
    size := 1
    for size < n {
        size <<= 1
    }
    segment := make([]int, 2 * size)
    // Initialize leaves
    for i := 0; i < n; i++ {
        segment[size + i] = transformed[i]
    }
    // Initialize unused leaves to 0 (GCD(x, 0) = x)
    for i := size + n; i < 2 * size; i++ {
        segment[i] = 0
    }
    // Build internal nodes
    for i := size - 1; i > 0; i-- {
        segment[i] = gcd(segment[2*i], segment[2*i+1])
    }
    for _, q := range queries {
        i, val := q[0], q[1]
        oldVal := transformed[i]
        // Calculate new transformed value
        var newVal int
        if val % p == 0 {
            newVal = val / p
        } else {
            newVal = 0
        }
        // Update count of divisible elements
        if oldVal == 0 && newVal > 0 {
            countDivisible++
        } else if oldVal > 0 && newVal == 0 {
            countDivisible--
        }
        // Update transformed array and segment tree
        if oldVal != newVal {
            transformed[i] = newVal
            pos := size + i
            segment[pos] = newVal
            for pos > 1 {
                pos >>= 1
                segment[pos] = gcd(segment[2*pos], segment[2 * pos + 1])
            }
        }
        // Check for good subsequence
        // 1. Must have at least one multiple of p
        // 2. GCD of all multiples of p (normalized by p) must be 1
        //    Because any subsequence of multiples of p has GCD = p * (sub-GCD).
        //    We need sub-GCD = 1.
        //    If GCD of all normalized multiples is g > 1, any sub-GCD is multiple of g, so != 1.
        //    If GCD of all normalized multiples is 1, then there exists a subset with GCD 1.
        totalGCD := segment[1] // Root of segment tree
        if totalGCD == 0 { continue } // No multiples of p
        if totalGCD > 1  { continue } // GCD of all multiples is > 1, so any subsequence GCD will be multiple of p * totalGCD > p
        // totalGCD == 1
        // There exists a subsequence with GCD exactly p.
        // We need to ensure its length is strictly less than n.
        if countDivisible < n {
            // The set of all multiples of p has size < n.
            // This set has GCD p. Its size is countDivisible < n.
            // So this set is a good subsequence.
            res++
        } else {
            // countDivisible == n
            // All elements are multiples of p. The set of all elements has size n.
            // We need a proper subset with GCD 1.
            // For values up to 50000, the minimum subset size to achieve GCD 1 is small (<= 6 or 7).
            // If n > 6, a proper subset definitely exists.
            if n > 6 {
                res++
            } else {
                // n is small (<= 6), check all subsets.
                // There are 2^n - 1 non-empty subsets.
                // We need a subset of size < n with GCD 1.
                // Since totalGCD is 1, the full set has GCD 1. We check proper subsets.
                found := false
                limit := 1 << n
                // Iterate all masks except 0 and the full set
                for mask := 1; mask < limit-1; mask++ {
                    currentGCD := 0
                    first := true
                    tempMask := mask
                    bitPos := 0
                    // Iterate bits to compute GCD
                    // Optimization: if currentGCD becomes 1, we can stop for this mask
                    for tempMask > 0 {
                        if tempMask&1 != 0 {
                            if first {
                                currentGCD = transformed[bitPos]
                                first = false
                            } else {
                                currentGCD = gcd(currentGCD, transformed[bitPos])
                            }
                            
                            if currentGCD == 1 {
                                found = true
                                break
                            }
                        }
                        tempMask >>= 1
                        bitPos++
                    }
                    if found {
                        break
                    }
                }
                if found {
                    res++
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,8,12,16], p = 2, queries = [[0,3],[2,6]]
    // Output: 1
    // Explanation:
    // i	| [indi, vali]  | Operation	Updated nums	            | Any good Subsequence
    // 0	| [0, 3]	    | Update nums[0] to 3	[3, 8, 12, 16]	| No, as no subsequence has GCD exactly p = 2
    // 1	| [2, 6]	    | Update nums[2] to 6	[3, 8, 6, 16]	| Yes, subsequence [8, 6] has GCD exactly p = 2
    // Thus, the answer is 1.
    fmt.Println(countGoodSubseq([]int{4,8,12,16}, 2, [][]int{{0,3},{2,6}})) // 1
    // Example 2:
    // Input: nums = [4,5,7,8], p = 3, queries = [[0,6],[1,9],[2,3]]
    // Output: 2
    // Explanation:
    // i	| [indi, vali]  | Operation	Updated nums	            | Any good Subsequence    
    // 0	| [0, 6]	    | Update nums[0] to 6	[6, 5, 7, 8]	| No, as no subsequence has GCD exactly p = 3
    // 1	| [1, 9]	    | Update nums[1] to 9	[6, 9, 7, 8]	| Yes, subsequence [6, 9] has GCD exactly p = 3
    // 2	| [2, 3]	    | Update nums[2] to 3	[6, 9, 3, 8]	| Yes, subsequence [6, 9, 3] has GCD exactly p = 3
    // Thus, the answer is 2.
    fmt.Println(countGoodSubseq([]int{4,5,7,8}, 3, [][]int{{0,6},{1,9},{2,3}})) // 2
    // Example 3:
    // Input: nums = [5,7,9], p = 2, queries = [[1,4],[2,8]]
    // Output: 0
    // Explanation:
    // i   | [indi, vali]  | Operation	Updated nums	        | Any good Subsequence
    // 0	| [1, 4]	    | Update nums[1] to 4	[5, 4, 9]	| No, as no subsequence has GCD exactly p = 2
    // 1	| [2, 8]	    | Update nums[2] to 8	[5, 4, 8]	| No, as no subsequence has GCD exactly p = 2
    // Thus, the answer is 0.
    fmt.Println(countGoodSubseq([]int{5,7,9}, 2, [][]int{{1,4},{2,8}})) // 0

    fmt.Println(countGoodSubseq([]int{1,2,3,4,5,6,7,8,9}, 2, [][]int{{1,4},{2,8}})) // 2
    fmt.Println(countGoodSubseq([]int{9,8,7,6,5,4,3,2,1}, 2, [][]int{{1,4},{2,8}})) // 2

    fmt.Println(countGoodSubseq1([]int{4,8,12,16}, 2, [][]int{{0,3},{2,6}})) // 1
    fmt.Println(countGoodSubseq1([]int{4,5,7,8}, 3, [][]int{{0,6},{1,9},{2,3}})) // 2
    fmt.Println(countGoodSubseq1([]int{5,7,9}, 2, [][]int{{1,4},{2,8}})) // 0
    fmt.Println(countGoodSubseq1([]int{1,2,3,4,5,6,7,8,9}, 2, [][]int{{1,4},{2,8}})) // 2
    fmt.Println(countGoodSubseq1([]int{9,8,7,6,5,4,3,2,1}, 2, [][]int{{1,4},{2,8}})) // 2

    fmt.Println(countGoodSubseq2([]int{4,8,12,16}, 2, [][]int{{0,3},{2,6}})) // 1
    fmt.Println(countGoodSubseq2([]int{4,5,7,8}, 3, [][]int{{0,6},{1,9},{2,3}})) // 2
    fmt.Println(countGoodSubseq2([]int{5,7,9}, 2, [][]int{{1,4},{2,8}})) // 0
    fmt.Println(countGoodSubseq2([]int{1,2,3,4,5,6,7,8,9}, 2, [][]int{{1,4},{2,8}})) // 2
    fmt.Println(countGoodSubseq2([]int{9,8,7,6,5,4,3,2,1}, 2, [][]int{{1,4},{2,8}})) // 2
}