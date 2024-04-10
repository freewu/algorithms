package main

// 315. Count of Smaller Numbers After Self
// Given an integer array nums, 
// return an integer array counts where counts[i] is the number of smaller elements to the right of nums[i].

// Example 1:
// Input: nums = [5,2,6,1]
// Output: [2,1,1,0]
// Explanation:
// To the right of 5 there are 2 smaller elements (2 and 1).
// To the right of 2 there is only 1 smaller element (1).
// To the right of 6 there is 1 smaller element (1).
// To the right of 1 there is 0 smaller element.

// Example 2:
// Input: nums = [-1]
// Output: [0]

// Example 3:
// Input: nums = [-1,-1]
// Output: [0,0]
 
// Constraints:
//     1 <= nums.length <= 10^5
//     -10^4 <= nums[i] <= 10^4

import "fmt"
import "math"

// 暴力法 Time Limit Exceeded 62/66 
func countSmaller1(nums []int) []int {
    l := len(nums)
    res := make([]int,l)
    for i := 0; i < l; i++ {
        count := 0
        for j := i + 1; j < l; j++ { // ->
            if nums[j] < nums[i] { // 右边的数小于 nums[i] 则累加
                count++
            }
        }
        res[i] = count
    }
    return res
}

// Segment Tree
const (
    numMin   = -10000
    numMax   = 10000
    rangeMax = numMax - numMin
)

func insert(st []int, n int, rfrom, rto int, pos int) {
    if n < rfrom || n > rto {
        return
    }
    st[pos]++
    if rfrom == rto {
        return
    }
    mid := (rfrom + rto) / 2
    insert(st, n, rfrom, mid, 2*pos+1)
    insert(st, n, mid+1, rto, 2*pos+2)
}

// count numbers lower than n
func merge(st []int, n int, rfrom, rto int, pos int) int {
    if n > rto {
        return st[pos]
    }
    if n < rfrom || rto == rfrom {
        return 0
    }
    mid := (rfrom + rto) / 2
    return merge(st, n, rfrom, mid, 2*pos+1) + merge(st, n, mid+1, rto, 2*pos+2)
}

func countSmaller(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    st := make([]int, 4*rangeMax)
    counts := make([]int, len(nums))
    for i := len(nums) - 1; i >= 0; i-- {
        insert(st, nums[i]-  numMin, 0, rangeMax, 0)
        counts[i] = merge(st, nums[i]-numMin, 0, rangeMax, 0)
    }
    return counts
}

// 
type FenwickTree struct {
    sums []int
}

func NewFenwickTree(n int) *FenwickTree {
    return &FenwickTree{
        sums: make([]int, n+1),
    }
}

func (t *FenwickTree) update(index, delta int) {
    for i := index; i < len(t.sums); i += i & -i {
        t.sums[i] += delta
    }
}

func (t *FenwickTree) query(index int) int {
    sum := 0
    for i := index; i > 0; i -= i & -i {
        sum += t.sums[i]
    }
    return sum
}

func countSmaller2(nums []int) []int {
    minValue, maxValue := math.MaxInt32, math.MinInt32
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, num := range nums {
        minValue = min(minValue, num)
    }
    adjustedNums := make([]int, len(nums))
    for i, num := range nums {
        adjustedNums[i] = num - minValue + 1
        maxValue = max(maxValue, adjustedNums[i])
    }
    tree := NewFenwickTree(maxValue)
    counts := make([]int, len(nums))
    
    for i := len(nums) - 1; i >= 0; i-- {
        counts[i] = tree.query(adjustedNums[i] - 1)
        tree.update(adjustedNums[i], 1)
    }
    return counts
}

func main() {
    // To the right of 5 there are 2 smaller elements (2 and 1).
    // To the right of 2 there is only 1 smaller element (1).
    // To the right of 6 there is 1 smaller element (1).
    // To the right of 1 there is 0 smaller element.
    fmt.Println(countSmaller([]int{5,2,6,1})) // [2,1,1,0]
    fmt.Println(countSmaller([]int{-1})) // [0]
    fmt.Println(countSmaller([]int{-1,-1})) // [0,0]

    fmt.Println(countSmaller1([]int{5,2,6,1})) // [2,1,1,0]
    fmt.Println(countSmaller1([]int{-1})) // [0]
    fmt.Println(countSmaller1([]int{-1,-1})) // [0,0]

    fmt.Println(countSmaller2([]int{5,2,6,1})) // [2,1,1,0]
    fmt.Println(countSmaller2([]int{-1})) // [0]
    fmt.Println(countSmaller2([]int{-1,-1})) // [0,0]
}