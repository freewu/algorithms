package main

// 3072. Distribute Elements Into Two Arrays II
// You are given a 1-indexed array of integers nums of length n.
// We define a function greaterCount such that greaterCount(arr, val) returns the number of elements in arr that are strictly greater than val.

// You need to distribute all the elements of nums between two arrays arr1 and arr2 using n operations. 
// In the first operation, append nums[1] to arr1. In the second operation, append nums[2] to arr2. 
// Afterwards, in the ith operation:
//     If greaterCount(arr1, nums[i]) > greaterCount(arr2, nums[i]), append nums[i] to arr1.
//     If greaterCount(arr1, nums[i]) < greaterCount(arr2, nums[i]), append nums[i] to arr2.
//     If greaterCount(arr1, nums[i]) == greaterCount(arr2, nums[i]), append nums[i] to the array with a lesser number of elements.
//     If there is still a tie, append nums[i] to arr1.

// The array result is formed by concatenating the arrays arr1 and arr2. 
// For example, if arr1 == [1,2,3] and arr2 == [4,5,6], then result = [1,2,3,4,5,6].

// Return the integer array result.

// Example 1:
// Input: nums = [2,1,3,3]
// Output: [2,3,1,3]
// Explanation: After the first 2 operations, arr1 = [2] and arr2 = [1].
// In the 3rd operation, the number of elements greater than 3 is zero in both arrays. Also, the lengths are equal, hence, append nums[3] to arr1.
// In the 4th operation, the number of elements greater than 3 is zero in both arrays. As the length of arr2 is lesser, hence, append nums[4] to arr2.
// After 4 operations, arr1 = [2,3] and arr2 = [1,3].
// Hence, the array result formed by concatenation is [2,3,1,3].

// Example 2:
// Input: nums = [5,14,3,1,2]
// Output: [5,3,1,2,14]
// Explanation: After the first 2 operations, arr1 = [5] and arr2 = [14].
// In the 3rd operation, the number of elements greater than 3 is one in both arrays. Also, the lengths are equal, hence, append nums[3] to arr1.
// In the 4th operation, the number of elements greater than 1 is greater in arr1 than arr2 (2 > 1). Hence, append nums[4] to arr1.
// In the 5th operation, the number of elements greater than 2 is greater in arr1 than arr2 (2 > 1). Hence, append nums[5] to arr1.
// After 5 operations, arr1 = [5,3,1,2] and arr2 = [14].
// Hence, the array result formed by concatenation is [5,3,1,2,14].

// Example 3:
// Input: nums = [3,3,3,3]
// Output: [3,3,3,3]
// Explanation: At the end of 4 operations, arr1 = [3,3] and arr2 = [3,3].
// Hence, the array result formed by concatenation is [3,3,3,3].

// Constraints:
//     3 <= n <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"
import "sort"
import "slices"

type Bit struct {
    tree []int
    n    int
}

func NewBit(n int) Bit {
    return Bit {
        tree: make([]int, n + 1),
        n:    n,
    }
}

func (this *Bit) update(idx, val int) {
    idx++
    for idx <= this.n {
        this.tree[idx] += val
        idx += idx & -idx
    }
}

func (this *Bit) query(idx int) int {
    idx++
    result := 0
    for idx > 0 {
        result += this.tree[idx]
        idx -= idx & -idx
    }
    return result
}

func compress(nums []int) []int {
    d := make([]int, len(nums))
    copy(d, nums)
    sort.Ints(d)
    cur := 0
    for i := 1; i < len(d); i++ {
        if d[i] != d[cur] {
            cur++
            d[cur] = d[i]
        }
    }
    return d[:cur + 1]
}

func resultArray(nums []int) []int {
    n, d := len(nums), compress(nums)
    m := len(d)
    trans := make(map[int]int)
    for i, val := range d {
        trans[val] = i
    }
    arr1, arr2 := []int{}, []int{}
    bit1, bit2 := NewBit(m), NewBit(m)

    arr1 = append(arr1, nums[0])
    bit1.update(trans[nums[0]], 1)
    arr2 = append(arr2, nums[1])
    bit2.update(trans[nums[1]], 1)

    for i := 2; i < n; i++ {
        count1 := len(arr1) - bit1.query(trans[nums[i]])
        count2 := len(arr2) - bit2.query(trans[nums[i]])
        if count1 > count2 || (count1 == count2 && len(arr1) <= len(arr2)) {
            arr1 = append(arr1, nums[i])
            bit1.update(trans[nums[i]], 1)
        } else {
            arr2 = append(arr2, nums[i])
            bit2.update(trans[nums[i]], 1)
        }
    }
    arr1 = append(arr1, arr2...)
    return arr1
}

func resultArray1(nums []int) []int {
    sorted := slices.Clone(nums)
    slices.Sort(sorted)
    sorted = slices.Compact(sorted)

    m, a, b := len(sorted), nums[:1],[]int{nums[1]}
    t1, t2 := make(fenwick, m+1), make(fenwick, m+1)
    t1.add(sort.SearchInts(sorted, nums[0]) + 1)
    t2.add(sort.SearchInts(sorted, nums[1]) + 1)
    for _, x := range nums[2:] {
        v := sort.SearchInts(sorted, x) + 1
        gc1 := len(a) - t1.pre(v) // greaterCount(a, v)
        gc2 := len(b) - t2.pre(v) // greaterCount(b, v)
        if gc1 > gc2 || gc1 == gc2 && len(a) <= len(b) {
            a = append(a, x)
            t1.add(v)
        } else {
            b = append(b, x)
            t2.add(v)
        }
    }
    return append(a, b...)
}

// 先定义树状数组
type fenwick []int

func lowBit(x int) int{
    return x&(-x)
}

func(t fenwick) add(p int){
    for p < len(t) {
        t[p] += 1
        p += lowBit(p)
    }
}

func(t fenwick) pre(p int) int{
    if p == 0{
        return 0
    }
    var res int
    for p > 0{
        res += t[p]
        p -= lowBit(p)
    }
    return res
}


type BinaryIndexedTree struct {
    n int
    c []int
}

func NewBinaryIndexedTree(n int) *BinaryIndexedTree {
    return &BinaryIndexedTree{n: n, c: make([]int, n+1)}
}

func (this *BinaryIndexedTree) update(x, delta int) {
    for ; x <= this.n; x += x & -x {
        this.c[x] += delta
    }
}

func (this *BinaryIndexedTree) query(x int) int {
    s := 0
    for ; x > 0; x -= x & -x {
        s += this.c[x]
    }
    return s
}

func resultArray2(nums []int) []int {
    st := make([]int, len(nums))
    copy(st, nums)
    sort.Ints(st)
    n := len(st)
    tree1 := NewBinaryIndexedTree(n + 1)
    tree2 := NewBinaryIndexedTree(n + 1)
    tree1.update(sort.SearchInts(st, nums[0])+1, 1)
    tree2.update(sort.SearchInts(st, nums[1])+1, 1)
    arr1 := []int{nums[0]}
    arr2 := []int{nums[1]}
    for _, x := range nums[2:] {
        i := sort.SearchInts(st, x) + 1
        a := len(arr1) - tree1.query(i)
        b := len(arr2) - tree2.query(i)
        if a > b {
            arr1 = append(arr1, x)
            tree1.update(i, 1)
        } else if a < b {
            arr2 = append(arr2, x)
            tree2.update(i, 1)
        } else if len(arr1) <= len(arr2) {
            arr1 = append(arr1, x)
            tree1.update(i, 1)
        } else {
            arr2 = append(arr2, x)
            tree2.update(i, 1)
        }
    }
    arr1 = append(arr1, arr2...)
    return arr1
}

func main() {
    // Example 1:
    // Input: nums = [2,1,3,3]
    // Output: [2,3,1,3]
    // Explanation: After the first 2 operations, arr1 = [2] and arr2 = [1].
    // In the 3rd operation, the number of elements greater than 3 is zero in both arrays. Also, the lengths are equal, hence, append nums[3] to arr1.
    // In the 4th operation, the number of elements greater than 3 is zero in both arrays. As the length of arr2 is lesser, hence, append nums[4] to arr2.
    // After 4 operations, arr1 = [2,3] and arr2 = [1,3].
    // Hence, the array result formed by concatenation is [2,3,1,3].
    fmt.Println(resultArray([]int{2,1,3,3})) // [2,3,1,3]
    // Example 2:
    // Input: nums = [5,14,3,1,2]
    // Output: [5,3,1,2,14]
    // Explanation: After the first 2 operations, arr1 = [5] and arr2 = [14].
    // In the 3rd operation, the number of elements greater than 3 is one in both arrays. Also, the lengths are equal, hence, append nums[3] to arr1.
    // In the 4th operation, the number of elements greater than 1 is greater in arr1 than arr2 (2 > 1). Hence, append nums[4] to arr1.
    // In the 5th operation, the number of elements greater than 2 is greater in arr1 than arr2 (2 > 1). Hence, append nums[5] to arr1.
    // After 5 operations, arr1 = [5,3,1,2] and arr2 = [14].
    // Hence, the array result formed by concatenation is [5,3,1,2,14].
    fmt.Println(resultArray([]int{5,14,3,1,2})) // [5,3,1,2,14]
    // Example 3:
    // Input: nums = [3,3,3,3]
    // Output: [3,3,3,3]
    // Explanation: At the end of 4 operations, arr1 = [3,3] and arr2 = [3,3].
    // Hence, the array result formed by concatenation is [3,3,3,3].
    fmt.Println(resultArray([]int{3,3,3,3})) // [3,3,3,3]

    fmt.Println(resultArray1([]int{2,1,3,3})) // [2,3,1,3]
    fmt.Println(resultArray1([]int{5,14,3,1,2})) // [5,3,1,2,14]
    fmt.Println(resultArray1([]int{3,3,3,3})) // [3,3,3,3]

    fmt.Println(resultArray2([]int{2,1,3,3})) // [2,3,1,3]
    fmt.Println(resultArray2([]int{5,14,3,1,2})) // [5,3,1,2,14]
    fmt.Println(resultArray2([]int{3,3,3,3})) // [3,3,3,3]
}