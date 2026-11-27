package main

// 4124. Minimum Operations to Make Mountain Array II
// You are given an integer array nums of length n.

// In one operation, you may increment or decrement any element of nums by 1. 
// Elements are not required to stay positive.

// An array is called a mountain array if there exists an index peak such that 0 < peak < n - 1 and:
//     1. nums[0] < nums[1] < ... < nums[peak]
//     2. nums[peak] > nums[peak + 1] > ... > nums[n - 1]

// Return the minimum number of operations required to transform nums into a mountain array. If it is impossible, return -1.

// Example 1:
// Input: nums = [2,2,3]
// Output: 2
// Explanation:
// One optimal way is to:
// Increment nums[1] by 1 and decrement nums[2] by 1.
// The resulting array is [2, 3, 2], which is a mountain array with peak = 1.
// Therefore, the total number of operations is 1 + 1 = 2.

// Example 2:
// Input: nums = [3,1,2,5,4]
// Output: 3
// Explanation:
// One optimal way is to:
// Decrement nums[0] by 1 three times, making it 0.
// The resulting array is [0, 1, 2, 5, 4], which is a mountain array with peak = 3.
// Therefore, the total number of operations is 3.

// Example 3:
// Input: nums = [5,7,7,4]
// Output: 1
// Explanation:
// One optimal way is to:
// Decrement nums[1] by 1.
// The resulting array is [5, 6, 7, 4], which is a mountain array with peak = 2.
// Therefore, the total number of operations is 1.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"
import "math/bits"

type Shoal struct {
    reed   []uint32
    ripple []uint64
    swell  []uint64
    sky    uint64
}

func NewShoal(span int) *Shoal {
    w0 := (span + 63) >> 6
    w1 := (w0 + 63) >> 6
    return &Shoal{
        reed:   make([]uint32, span),
        ripple: make([]uint64, w0),
        swell:  make([]uint64, w1),
        sky:    0,
    }
}

func (s *Shoal) Sow(p int) {
    s.reed[p]++
    if s.reed[p] == 1 {
        q := p >> 6
        bit := p & 63
        s.ripple[q] |= 1 << bit
        q2 := q >> 6
        s.swell[q2] |= 1 << (q & 63)
        s.sky |= 1 << q2
    }
}

func (s *Shoal) Cull(p int) {
    s.reed[p]--
    if s.reed[p] == 0 {
        q := p >> 6
        bit := p & 63
        s.ripple[q] &^= 1 << bit
        if s.ripple[q] == 0 {
            q2 := q >> 6
            s.swell[q2] &^= 1 << (q & 63)
            if s.swell[q2] == 0 {
                s.sky &^= 1 << q2
            }
        }
    }
}

func (s *Shoal) Crest() int {
    t := 63 - bits.LeadingZeros64(s.sky)
    q := (t << 6) + (63 - bits.LeadingZeros64(s.swell[t]))
    return (q << 6) + (63 - bits.LeadingZeros64(s.ripple[q]))
}

func (s *Shoal) Reset() {
    for q := range s.ripple {
        w := s.ripple[q]
        for w != 0 {
            b := bits.TrailingZeros64(w)
            w &^= 1 << b
            pos := (q << 6) + b
            s.reed[pos] = 0
        }
        s.ripple[q] = 0
    }
    for i := range s.swell {
        s.swell[i] = 0
    }
    s.sky = 0
}

func minOperations(nums []int) int64 {
    n := len(nums)
    if n < 3 {
        return -1
    }
    bed, top := int64(1 << 61), int64(-1 << 61)
    for i := 0; i < n; i++ {
        a := int64(nums[i]) - int64(i)
        b := int64(nums[i]) - int64(n-1-i)
        if a < bed {
            bed = a
        }
        if b < bed {
            bed = b
        }
        if a > top {
            top = a
        }
        if b > top {
            top = b
        }
    }
    span := int(top - bed + 1)
    shoal := NewShoal(span)
    descent, shelf := make([]int64, n), make([]int32, n)
    tide := int64(0)
    for j := 0; j < n; j++ {
        i := n - 1 - j
        p := int(int64(nums[i]) - int64(j) - bed)
        shoal.Sow(p)
        m := shoal.Crest()
        if m > p {
            tide += int64(m - p)
            shoal.Cull(m)
            shoal.Sow(p)
        }
        descent[i] = tide
        shelf[i] = int32(bed + int64(shoal.Crest()) + int64(j))
    }
    shoal.Reset()
    tide = 0
    res := int64(1 << 61)
    for i := 0; i < n-1; i++ {
        p := int(int64(nums[i]) - int64(i) - bed)
        shoal.Sow(p)
        m := shoal.Crest()
        if m > p {
            tide += int64(m - p)
            shoal.Cull(m)
            shoal.Sow(p)
        }
        pk := i + 1
        if pk <= n - 2 {
            ridge := bed + int64(shoal.Crest()) + int64(i)
            opt1 := ridge + 1
            opt2 := int64(shelf[pk+1]) + 1
            opt3 := int64(nums[pk])
            crestVal := opt1
            if opt2 > crestVal {
                crestVal = opt2
            }
            if opt3 > crestVal {
                crestVal = opt3
            }
            cost := tide + descent[pk+1] + (crestVal - int64(nums[pk]))
            if cost < res {
                res = cost
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,2,3]
    // Output: 2
    // Explanation:
    // One optimal way is to:
    // Increment nums[1] by 1 and decrement nums[2] by 1.
    // The resulting array is [2, 3, 2], which is a mountain array with peak = 1.
    // Therefore, the total number of operations is 1 + 1 = 2.
    fmt.Println(minOperations([]int{2,2,3})) // 2 
    // Example 2:
    // Input: nums = [3,1,2,5,4]
    // Output: 3
    // Explanation:
    // One optimal way is to:
    // Decrement nums[0] by 1 three times, making it 0.
    // The resulting array is [0, 1, 2, 5, 4], which is a mountain array with peak = 3.
    // Therefore, the total number of operations is 3.
    fmt.Println(minOperations([]int{3,1,2,5,4})) // 3
    // Example 3:
    // Input: nums = [5,7,7,4]
    // Output: 1
    // Explanation:
    // One optimal way is to:
    // Decrement nums[1] by 1.
    // The resulting array is [5, 6, 7, 4], which is a mountain array with peak = 2.
    // Therefore, the total number of operations is 1.
    fmt.Println(minOperations([]int{5,7,7,4})) // 1

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 2
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 2
}