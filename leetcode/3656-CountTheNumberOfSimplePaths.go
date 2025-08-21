package main

// 3656. Determine if a Simple Graph Exists
// You are given an integer array degrees, where degrees[i] represents the desired degree of the ith vertex.

// Your task is to determine if there exists an undirected simple graph with exactly these vertex degrees.

// A simple graph has no self-loops or parallel edges between the same pair of vertices.

// Return true if such a graph exists, otherwise return false.

// Example 1:
// Input: degrees = [3,1,2,2]
// Output: true
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/08/12/screenshot-2025-08-13-at-24347-am.png" />
// One possible undirected simple graph is:
// Edges: (0, 1), (0, 2), (0, 3), (2, 3)
// Degrees: deg(0) = 3, deg(1) = 1, deg(2) = 2, deg(3) = 2.

// Example 2:
// Input: degrees = [1,3,3,1]
// Output: false
// Explanation:​​​​​​​
// degrees[1] = 3 and degrees[2] = 3 means they must be connected to all other vertices.
// This requires degrees[0] and degrees[3] to be at least 2, but both are equal to 1, which contradicts the requirement.
// Thus, the answer is false.

// Constraints:
//     1 <= n == degrees.length <= 10^​​​​​​​5
//     0 <= degrees[i] <= n - 1

import "fmt"
import "sort"

func simpleGraphExists(degrees []int) bool {
    sum, n := 0, len(degrees)
    for _, v := range degrees { // 检查每个度数是否在合理范围内（0到n-1之间）
        if v < 0 || v >= n {
            return false
        }
        sum += v
    }
    if sum % 2 != 0 { // 检查度数之和是否为偶数（无向图中所有顶点度数和必为偶数）
        return false
    }
    // 对度数进行降序排序
    sort.Sort(sort.Reverse(sort.IntSlice(degrees)))
    // 计算前缀和，方便快速计算任意区间的和
    prefix := make([]int, n + 1)
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + degrees[i]
    }
    // 应用Erdős–Gallai定理检查每个k的条件
    for k := 1; k <= n; k++ {
        // 计算前k个度数之和
        sumK := prefix[k]
        // 计算右边第一项k(k-1)
        term1 := k * (k - 1)
        // 找到第一个小于k的位置（使用二分查找优化）
        left, right := k, n
        for left < right {
            mid := (left + right) / 2
            if degrees[mid] < k {
                right = mid
            } else {
                left = mid + 1
            }
        }
        // 计算右边第二项sum(min(di, k)) for i > k
        // 对于i从k到left-1，min(di, k) = k
        // 对于i从left到n-1，min(di, k) = di
        term2 := k*(left-k) + (prefix[n] - prefix[left])
        // 检查条件是否满足
        if sumK > term1 + term2 {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: degrees = [3,1,2,2]
    // Output: true
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/08/12/screenshot-2025-08-13-at-24347-am.png" />
    // One possible undirected simple graph is:
    // Edges: (0, 1), (0, 2), (0, 3), (2, 3)
    // Degrees: deg(0) = 3, deg(1) = 1, deg(2) = 2, deg(3) = 2.
    fmt.Println(simpleGraphExists([]int{3,1,2,2})) // true
    // Example 2:
    // Input: degrees = [1,3,3,1]
    // Output: false
    // Explanation:​​​​​​​
    // degrees[1] = 3 and degrees[2] = 3 means they must be connected to all other vertices.
    // This requires degrees[0] and degrees[3] to be at least 2, but both are equal to 1, which contradicts the requirement.
    // Thus, the answer is false.
    fmt.Println(simpleGraphExists([]int{1,3,3,1})) // false

    fmt.Println(simpleGraphExists([]int{0,1,2,3,4,5,6,7,8,9})) // false
    fmt.Println(simpleGraphExists([]int{9,8,7,6,5,4,3,2,1,0})) // false
}