package main

// 3801. Minimum Cost to Merge Sorted Lists
// You are given a 2D integer array lists, where each lists[i] is a non-empty array of integers sorted in non-decreasing order.

// You may repeatedly choose two lists a = lists[i] and b = lists[j], where i != j, and merge them. 
// The cost to merge a and b is:

// len(a) + len(b) + abs(median(a) - median(b)), where len and median denote the list length and median, respectively.

// After merging a and b, remove both a and b from lists and insert the new merged sorted list in any position. 
// Repeat merges until only one list remains.

// Return an integer denoting the minimum total cost required to merge all lists into one single sorted list.

// The median of an array is the middle element after sorting it in non-decreasing order. 
// If the array has an even number of elements, the median is the left middle element.
 
// Example 1:
// Input: lists = [[1,3,5],[2,4],[6,7,8]]
// Output: 18
// Explanation:
// Merge a = [1, 3, 5] and b = [2, 4]:
// len(a) = 3 and len(b) = 2
// median(a) = 3 and median(b) = 2
// cost = len(a) + len(b) + abs(median(a) - median(b)) = 3 + 2 + abs(3 - 2) = 6
// So lists becomes [[1, 2, 3, 4, 5], [6, 7, 8]].
// Merge a = [1, 2, 3, 4, 5] and b = [6, 7, 8]:
// len(a) = 5 and len(b) = 3
// median(a) = 3 and median(b) = 7
// cost = len(a) + len(b) + abs(median(a) - median(b)) = 5 + 3 + abs(3 - 7) = 12
// So lists becomes [[1, 2, 3, 4, 5, 6, 7, 8]], and total cost is 6 + 12 = 18.

// Example 2:
// Input: lists = [[1,1,5],[1,4,7,8]]
// Output: 10
// Explanation:
// Merge a = [1, 1, 5] and b = [1, 4, 7, 8]:
// len(a) = 3 and len(b) = 4
// median(a) = 1 and median(b) = 4
// cost = len(a) + len(b) + abs(median(a) - median(b)) = 3 + 4 + abs(1 - 4) = 10
// So lists becomes [[1, 1, 1, 4, 5, 7, 8]], and total cost is 10.

// Example 3:
// Input: lists = [[1],[3]]
// Output: 4
// Explanation:
// Merge a = [1] and b = [3]:
// len(a) = 1 and len(b) = 1
// median(a) = 1 and median(b) = 3
// cost = len(a) + len(b) + abs(median(a) - median(b)) = 1 + 1 + abs(1 - 3) = 4
// So lists becomes [[1, 3]], and total cost is 4.

// Example 4:
// Input: lists = [[1],[1]]
// Output: 2
// Explanation:
// The total cost is len(a) + len(b) + abs(median(a) - median(b)) = 1 + 1 + abs(1 - 1) = 2.

// Constraints:
//     2 <= lists.length <= 12
//     1 <= lists[i].length <= 500
//     -10^9 <= lists[i][j] <= 10^9
//     lists[i] is sorted in non-decreasing order.
//     The sum of lists[i].length will not exceed 2000.

import "fmt"
import "sort"

func minMergeCost(lists [][]int) int64 {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    merge := func (a, b []int) []int {
        i, j, n, m := 0, 0, len(a), len(b)
        res := make([]int, 0, n + m)
        for {
            if i == n {
                return append(res, b[j:]...)
            }
            if j == m {
                return append(res, a[i:]...)
            }
            if a[i] < b[j] {
                res = append(res, a[i])
                i++
            } else {
                res = append(res, b[j])
                j++
            }
        }
    }
    u := 1 << len(lists)
    sorted := make([][]int, u)
    for i, a := range lists { // 枚举不在 s 中的下标 i
        highBit := 1 << i
        for s, b := range sorted[:highBit] {
            sorted[highBit|s] = merge(a, b)
        }
    }
    f := make([]int, u)
    for i := range f {
        if i&(i-1) == 0 { // i 只包含一个元素，无法分解成两个非空子集
            continue // f[i] = 0
        }
        f[i] = 1 << 61
        // 枚举 i 的非空真子集 j
        for j := i & (i - 1); j > i^j; j = (j - 1) & i {
            k := i ^ j // j 关于 i 的补集是 k
            lenJ, lenK := len(sorted[j]), len(sorted[k])
            medJ, medK := sorted[j][(lenJ - 1) / 2], sorted[k][(lenK - 1) / 2]
            f[i] = min(f[i], f[j] + f[k] + lenJ + lenK+ abs(medJ - medK))
        }
    }
    return int64(f[u - 1])
}

func minMergeCost1(lists [][]int) int64 {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    merge := func (a, b []int) []int { // 合并两个有序数组（创建一个新数组）
        i, n := 0, len(a)
        j, m := 0, len(b)
        res := make([]int, 0, n+m)
        for {
            if i == n {
                return append(res, b[j:]...)
            }
            if j == m {
                return append(res, a[i:]...)
            }
            if a[i] < b[j] {
                res = append(res, a[i])
                i++
            } else {
                res = append(res, b[j])
                j++
            }
        }
    }
    calcSorted := func(lists [][]int) [][]int {
        u := 1 << len(lists)
        sorted := make([][]int, u)
        for i, a := range lists {
            highBit := 1 << i
            for s, b := range sorted[:highBit] {
                sorted[highBit|s] = merge(a, b)
            }
        }
        return sorted
    }
    findMedianSortedArrays := func(a, b []int) int { // 寻找两个正序数组的中位数
        if len(a) > len(b) {
            a, b = b, a
        }
        m, n := len(a), len(b)
        i := sort.Search(m, func(i int) bool {
            j := (m+n+1)/2 - i - 2
            return a[i] > b[j+1]
        }) - 1
        j := (m+n+1)/2 - i - 2
        if i < 0 { return b[j] }
        if j < 0 { return a[i] }
        return max(a[i], b[j])
    }
    n := len(lists)
    m := n / 2
    sorted1, sorted2 := calcSorted(lists[:m]), calcSorted(lists[m:])
    u := 1 << n
    half := 1<<m - 1
    sumLen, median := make([]int, u), make([]int, u) // 可以省略，但预处理出来，相比直接在后面 DP 中计算更快
    for i := 1; i < u; i++ {
        // 把 i 分成低 m 位和高 n-m 位
        // 低 half 位去 sorted1 中找合并后的数组
        // 高 n-half 位去 sorted2 中找合并后的数组
        sumLen[i] = len(sorted1[i&half]) + len(sorted2[i>>m])
        median[i] = findMedianSortedArrays(sorted1[i&half], sorted2[i>>m])
    }
    f := make([]int, u)
    for i := range f {
        if i&(i-1) == 0 { continue }
        f[i] = 1 << 61
        for j := i & (i - 1); j > i^j; j = (j - 1) & i {
            k := i ^ j
            f[i] = min(f[i], f[j]+f[k] + sumLen[i] + abs(median[j]-median[k]))
        }
    }
    return int64(f[u-1])
}

func main() {
    // Example 1:
    // Input: lists = [[1,3,5],[2,4],[6,7,8]]
    // Output: 18
    // Explanation:
    // Merge a = [1, 3, 5] and b = [2, 4]:
    // len(a) = 3 and len(b) = 2
    // median(a) = 3 and median(b) = 2
    // cost = len(a) + len(b) + abs(median(a) - median(b)) = 3 + 2 + abs(3 - 2) = 6
    // So lists becomes [[1, 2, 3, 4, 5], [6, 7, 8]].
    // Merge a = [1, 2, 3, 4, 5] and b = [6, 7, 8]:
    // len(a) = 5 and len(b) = 3
    // median(a) = 3 and median(b) = 7
    // cost = len(a) + len(b) + abs(median(a) - median(b)) = 5 + 3 + abs(3 - 7) = 12
    // So lists becomes [[1, 2, 3, 4, 5, 6, 7, 8]], and total cost is 6 + 12 = 18.
    fmt.Println(minMergeCost([][]int{{1,3,5},{2,4},{6,7,8}})) // 18
    // Example 2:
    // Input: lists = [[1,1,5],[1,4,7,8]]
    // Output: 10
    // Explanation:
    // Merge a = [1, 1, 5] and b = [1, 4, 7, 8]:
    // len(a) = 3 and len(b) = 4
    // median(a) = 1 and median(b) = 4
    // cost = len(a) + len(b) + abs(median(a) - median(b)) = 3 + 4 + abs(1 - 4) = 10
    // So lists becomes [[1, 1, 1, 4, 5, 7, 8]], and total cost is 10.
    fmt.Println(minMergeCost([][]int{{1,1,5},{1,4,7,8}})) // 10 
    // Example 3:
    // Input: lists = [[1],[3]]
    // Output: 4
    // Explanation:
    // Merge a = [1] and b = [3]:
    // len(a) = 1 and len(b) = 1
    // median(a) = 1 and median(b) = 3
    // cost = len(a) + len(b) + abs(median(a) - median(b)) = 1 + 1 + abs(1 - 3) = 4
    // So lists becomes [[1, 3]], and total cost is 4.
    fmt.Println(minMergeCost([][]int{{1},{3}})) // 4
    // Example 4:
    // Input: lists = [[1],[1]]
    // Output: 2
    // Explanation:
    // The total cost is len(a) + len(b) + abs(median(a) - median(b)) = 1 + 1 + abs(1 - 1) = 2.
    fmt.Println(minMergeCost([][]int{{1},{1}})) // 2

    fmt.Println(minMergeCost([][]int{{1,2,3,4,5,6,7,8,9},{1,2,3,4,5,6,7,8,9}})) // 18
    fmt.Println(minMergeCost([][]int{{1,2,3,4,5,6,7,8,9},{9,8,7,6,5,4,3,2,1}})) // 18
    fmt.Println(minMergeCost([][]int{{9,8,7,6,5,4,3,2,1},{1,2,3,4,5,6,7,8,9}})) // 18
    fmt.Println(minMergeCost([][]int{{9,8,7,6,5,4,3,2,1},{9,8,7,6,5,4,3,2,1}})) // 18

    fmt.Println(minMergeCost1([][]int{{1,3,5},{2,4},{6,7,8}})) // 18
    fmt.Println(minMergeCost1([][]int{{1,1,5},{1,4,7,8}})) // 10 
    fmt.Println(minMergeCost1([][]int{{1},{3}})) // 4
    fmt.Println(minMergeCost1([][]int{{1},{1}})) // 2
    fmt.Println(minMergeCost1([][]int{{1,2,3,4,5,6,7,8,9},{1,2,3,4,5,6,7,8,9}})) // 18
    fmt.Println(minMergeCost1([][]int{{1,2,3,4,5,6,7,8,9},{9,8,7,6,5,4,3,2,1}})) // 18
    fmt.Println(minMergeCost1([][]int{{9,8,7,6,5,4,3,2,1},{1,2,3,4,5,6,7,8,9}})) // 18
    fmt.Println(minMergeCost1([][]int{{9,8,7,6,5,4,3,2,1},{9,8,7,6,5,4,3,2,1}})) // 18
}