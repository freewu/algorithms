package main

// 2736. Maximum Sum Queries
// You are given two 0-indexed integer arrays nums1 and nums2, each of length n, and a 1-indexed 2D array queries where queries[i] = [xi, yi].

// For the ith query, find the maximum value of nums1[j] + nums2[j] among all indices j (0 <= j < n), where nums1[j] >= xi and nums2[j] >= yi, or -1 if there is no j satisfying the constraints.

// Return an array answer where answer[i] is the answer to the ith query.

// Example 1:
// Input: nums1 = [4,3,1,2], nums2 = [2,4,9,5], queries = [[4,1],[1,3],[2,5]]
// Output: [6,10,7]
// Explanation: 
// For the 1st query xi = 4 and yi = 1, we can select index j = 0 since nums1[j] >= 4 and nums2[j] >= 1. The sum nums1[j] + nums2[j] is 6, and we can show that 6 is the maximum we can obtain.
// For the 2nd query xi = 1 and yi = 3, we can select index j = 2 since nums1[j] >= 1 and nums2[j] >= 3. The sum nums1[j] + nums2[j] is 10, and we can show that 10 is the maximum we can obtain. 
// For the 3rd query xi = 2 and yi = 5, we can select index j = 3 since nums1[j] >= 2 and nums2[j] >= 5. The sum nums1[j] + nums2[j] is 7, and we can show that 7 is the maximum we can obtain.
// Therefore, we return [6,10,7].

// Example 2:
// Input: nums1 = [3,2,5], nums2 = [2,3,4], queries = [[4,4],[3,2],[1,1]]
// Output: [9,9,9]
// Explanation: For this example, we can use index j = 2 for all the queries since it satisfies the constraints for each query.

// Example 3:
// Input: nums1 = [2,1], nums2 = [2,3], queries = [[3,3]]
// Output: [-1]
// Explanation: There is one query in this example with xi = 3 and yi = 3. For every index, j, either nums1[j] < xi or nums2[j] < yi. Hence, there is no solution. 

// Constraints:
//     nums1.length == nums2.length 
//     n == nums1.length 
//     1 <= n <= 10^5
//     1 <= nums1[i], nums2[i] <= 10^9 
//     1 <= queries.length <= 10^5
//     queries[i].length == 2
//     xi == queries[i][1]
//     yi == queries[i][2]
//     1 <= xi, yi <= 10^9

import "fmt"
import "sort"

type BinaryIndexedTree struct {
    n int
    c []int
}

func NewBinaryIndexedTree(n int) BinaryIndexedTree {
    c := make([]int, n+1)
    for i := range c {
        c[i] = -1
    }
    return BinaryIndexedTree{ n: n, c: c }
}

func (bit *BinaryIndexedTree) Update(x, v int) {
    for x <= bit.n {
        bit.c[x] = max(bit.c[x], v)
        x += x & -x
    }
}

func (bit *BinaryIndexedTree) Query(x int) int {
    mx := -1
    for x > 0 {
        mx = max(mx, bit.c[x])
        x -= x & -x
    }
    return mx
}

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
    n, m := len(nums1), len(queries)
    nums := make([][2]int, n)
    for i := range nums {
        nums[i] = [2]int{nums1[i], nums2[i]}
    }
    sort.Slice(nums, func(i, j int) bool { 
        return nums[j][0] < nums[i][0] 
    })
    sort.Ints(nums2)
    index := make([]int, m)
    for i := range index {
        index[i] = i
    }
    sort.Slice(index, func(i, j int) bool { 
        return queries[index[j]][0] < queries[index[i]][0] 
    })
    tree := NewBinaryIndexedTree(n)
    res, j := make([]int, m), 0
    for _, i := range index {
        x, y := queries[i][0], queries[i][1]
        for ; j < n && nums[j][0] >= x; j++ {
            k := n - sort.SearchInts(nums2, nums[j][1])
            tree.Update(k, nums[j][0]+nums[j][1])
        }
        k := n - sort.SearchInts(nums2, y)
        res[i] = tree.Query(k)
    }
    return res
}

func maximumSumQueries1(nums1 []int, nums2 []int, queries [][]int) []int {
    n := len(nums1)
    //	nums1与nums2合并，并按第一维度降序
    nums := make([][]int, n)
    for i := 0; i < n; i++ {
        nums[i] = []int{nums1[i], nums2[i]}
    }
    sort.Slice(nums, func(i, j int) bool {
        return nums[i][0] > nums[j][0]
    })
    //	queries下标按第一维度降序
    index := make([]int, len(queries))
    for i := 0; i < len(index); i++ {
        index[i] = i
    }
    sort.Slice(index, func(i, j int) bool {
        return queries[index[i]][0] > queries[index[j]][0]
    })
    //	从左往右遍历queries，遍历出的nums入单减栈
    stack, pos, i := make([][]int, n), -1, 0
    res := make([]int, len(queries))
    binarySearch := func(stack [][]int, R, t int) int {
        l, r := 0, R
        if l > r { return -1 }
        if t <= stack[l][1] { return stack[l][0] + stack[l][1] }
        if t > stack[r][1]  { return -1 }
        for l < r {
            mid := l + (r - l) >> 1
            flag := t <= stack[mid][1]
            if l == mid {
                if !flag {
                    return stack[r][0] + stack[r][1]
                } else {
                    return stack[l][0] + stack[l][1]
                }
            }
            if flag {
                r = mid
            } else {
                l = mid
            }
        }
        return -1
    }
    for _, v := range index {
        x, y := queries[v][0], queries[v][1]
        for i < n && nums[i][0] >= x {
            n := nums[i]
            sum := n[0] + n[1]
            i++
            for pos >= 0 && sum >= stack[pos][0] + stack[pos][1] {
                pos--
            }
            if pos == -1 || stack[pos][1] < n[1] {
                stack[pos+1], pos = n, pos+1
            }
        }
        //	二分搜（栈内是按第二维度升序的）
        res[v] = binarySearch(stack, pos, y)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [4,3,1,2], nums2 = [2,4,9,5], queries = [[4,1],[1,3],[2,5]]
    // Output: [6,10,7]
    // Explanation: 
    // For the 1st query xi = 4 and yi = 1, we can select index j = 0 since nums1[j] >= 4 and nums2[j] >= 1. The sum nums1[j] + nums2[j] is 6, and we can show that 6 is the maximum we can obtain.
    // For the 2nd query xi = 1 and yi = 3, we can select index j = 2 since nums1[j] >= 1 and nums2[j] >= 3. The sum nums1[j] + nums2[j] is 10, and we can show that 10 is the maximum we can obtain. 
    // For the 3rd query xi = 2 and yi = 5, we can select index j = 3 since nums1[j] >= 2 and nums2[j] >= 5. The sum nums1[j] + nums2[j] is 7, and we can show that 7 is the maximum we can obtain.
    // Therefore, we return [6,10,7].
    fmt.Println(maximumSumQueries([]int{4,3,1,2}, []int{2,4,9,5}, [][]int{{4,1},{1,3},{2,5}})) // [6,10,7]
    // Example 2:
    // Input: nums1 = [3,2,5], nums2 = [2,3,4], queries = [[4,4],[3,2],[1,1]]
    // Output: [9,9,9]
    // Explanation: For this example, we can use index j = 2 for all the queries since it satisfies the constraints for each query.
    fmt.Println(maximumSumQueries([]int{3,2,5}, []int{2,3,4}, [][]int{{4,4},{3,2},{1,1}})) // [9,9,9]
    // Example 3:
    // Input: nums1 = [2,1], nums2 = [2,3], queries = [[3,3]]
    // Output: [-1]
    // Explanation: There is one query in this example with xi = 3 and yi = 3. For every index, j, either nums1[j] < xi or nums2[j] < yi. Hence, there is no solution. 
    fmt.Println(maximumSumQueries([]int{2,1}, []int{2,3}, [][]int{{3,3}})) // [-1]

    fmt.Println(maximumSumQueries1([]int{4,3,1,2}, []int{2,4,9,5}, [][]int{{4,1},{1,3},{2,5}})) // [6,10,7]
    fmt.Println(maximumSumQueries1([]int{3,2,5}, []int{2,3,4}, [][]int{{4,4},{3,2},{1,1}})) // [9,9,9]
    fmt.Println(maximumSumQueries1([]int{2,1}, []int{2,3}, [][]int{{3,3}})) // [-1]
}