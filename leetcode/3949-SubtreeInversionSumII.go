package main

// 3949. Subtree Inversion Sum II
// You are given an undirected tree rooted at node 0, with n nodes numbered from 0 to n - 1. 
// The tree is represented by a 2D integer array edges of length n - 1, where edges[i] = [ui, vi] indicates an edge between nodes ui and vi.

// You are also given an integer array nums of length n, where nums[i] represents the value at node i, and an integer k.

// You may perform inversion operations on a subset of nodes subject to the following rules:
//     1. Subtree Inversion Operation:
//         1.1 When you invert a node, every value in the subtree rooted at that node is multiplied by -1.
//     2. Distance Constraint on Inversions:
//         2.1 You may only invert a node if it is “sufficiently far” from any other inverted node.
//         2.2 If you invert two nodes a and b, the distance (the number of edges on the unique path between them) must be at least k.

// Return the maximum possible sum of the tree’s node values after applying inversion operations.

// Example 1:
// Input: edges = [[0,1],[0,2],[0,3],[1,4],[1,5]], nums = [1,0,-10,3,4,5], k = 2
// Output: 23
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/21/4183example1drawio.png" />
// After inverting the subtree rooted at node 2, the maximum sum becomes 1 + 0 + 10 + 3 + 4 + 5 = 23.

// Example 2:
// Input: edges = [[0,1],[1,2]], nums = [5,-10,-10], k = 1
// Output: 25
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/21/4183example2drawio.png" />
// After inverting the subtree rooted at node 1, the maximum sum becomes 5 + 10 + 10 = 25.

// Example 3:
// Input: edges = [[0,1],[0,2]], nums = [1,-5,-6], k = 2
// Output: 12
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/21/4183example3drawio.png" />
// After inverting the subtrees rooted at nodes 1 and 2, nums = [1, 5, 6].
// This is valid because nodes 1 and 2 are two edges apart (1 → 0 and 0 → 2), which is at least k.
// The maximum sum is 1 + 5 + 6 = 12.

// Example 4:
// Input: edges = [[0,1],[0,2]], nums = [1,-5,-6], k = 3
// Output: 10
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/21/4183example4drawio.png" />
// After inverting the subtree rooted at nodes 0, nums = [-1, 5, 6].
// The maximum sum is (-1) + 5 + 6 = 10.
// Note that we cannot invert nodes 1 and 2 because their distance is 2 < k = 3.
 
// Constraints:
//     nums.length == n
//     edges.length == n - 1
//     2 <= n <= 5 * 10^4
//     edges[i].length == 2
//     0 <= edges[i][0], edges[i][1] < n
//     -5 * 10^4 <= nums[i] <= 5 * 10^4
//     1 <= k <= 50
//     It is guaranteed that edges forms a tree.

import "fmt"
import "math"

func subtreeInversionSum(branches [][]int, values []int, limit int) int {
    const INF = int64(1e18)
    n := len(values)
    forest := make([][]int, n)
    for i := range forest {
        forest[i] = make([]int, 0)
    }
    // 构建无向树
    for _, twig := range branches {
        u, v := twig[0], twig[1]
        forest[u] = append(forest[u], v)
        forest[v] = append(forest[v], u)
    }
    capacity := limit + 1
    harmony := make([]int64, n)
    for i := range harmony {
        harmony[i] = int64(values[i])
    }
    var wander func(leaf, root int, wood [][]int, harmony []int64, cap, bound int) ([]int64, []int64) 
    wander = func(leaf, root int, wood [][]int, harmony []int64, cap, bound int) ([]int64, []int64) {
        var saplings [][2][]int64
        // 遍历子节点
        for _, bud := range wood[leaf] {
            if bud == root {
                continue
            }
            sp, ab := wander(bud, leaf, wood, harmony, cap, bound)
            saplings = append(saplings, [2][]int64{sp, ab})
        }
        // 初始化 ascent/descent
        ascent := make([]int64, cap+1)
        descent := make([]int64, cap+1)
        for i := range ascent {
            ascent[i] = int64(math.MinInt64)
            descent[i] = INF
        }
        ascent[cap] = harmony[leaf]
        descent[cap] = harmony[leaf]
        // 合并子树 DP
        for _, sprout := range saplings {
            childPeak, childValley := sprout[0], sprout[1]
            freshRise, freshFall := make([]int64, cap + 1), make([]int64, cap + 1)
            for i := range freshRise {
                freshRise[i] = int64(math.MinInt64)
                freshFall[i] = INF
            }
            for parentTurn := 1; parentTurn <= cap; parentTurn++ {
                parentHeight := ascent[parentTurn]
                if parentHeight == int64(math.MinInt64) {
                    continue
                }
                for childTurn := 0; childTurn <= cap; childTurn++ {
                    childHeight := childPeak[childTurn]
                    if childHeight == int64(math.MinInt64) {
                        continue
                    }
                    childStep := min(childTurn+1, cap)
                    unison := min(parentTurn, childStep)
                    if parentTurn+childStep < bound {
                        continue
                    }
                    // 更新峰值
                    peak := parentHeight + childHeight
                    if peak > freshRise[unison] {
                        freshRise[unison] = peak
                    }
                    // 更新谷值
                    trough := descent[parentTurn] + childValley[childTurn]
                    if trough < freshFall[unison] {
                        freshFall[unison] = trough
                    }
                }
            }
            ascent, descent = freshRise, freshFall
        }
        // 计算 valley 和 crest
        valley := INF
        crest := int64(math.MinInt64)
        graceful := true
        var gatherLow, gatherHigh int64 = 0, 0
        for _, sprout := range saplings {
            childPeak := sprout[0]
            childValley := sprout[1]
            bestLow := INF
            bestHigh := int64(math.MinInt64)
            for turn := 0; turn <= cap; turn++ {
                if turn <= bound && turn < bound-1 {
                    continue
                }
                if childValley[turn] < bestLow {
                    bestLow = childValley[turn]
                }
                if childPeak[turn] > bestHigh {
                    bestHigh = childPeak[turn]
                }
            }
            if bestLow == INF {
                graceful = false
                break
            }
            gatherLow += bestLow
            gatherHigh += bestHigh
        }
        if graceful {
            valley = -harmony[leaf] - gatherLow
            crest = -harmony[leaf] - gatherHigh
        }
        // 复制并更新结果
        summit := make([]int64, cap+1)
        abyss := make([]int64, cap+1)
        copy(summit, ascent)
        copy(abyss, descent)
        if valley > summit[0] {
            summit[0] = valley
        }
        if crest < abyss[0] {
            abyss[0] = crest
        }
        return summit, abyss
    }
    // 深度优先遍历
    glory, _ := wander(0, -1, forest, harmony, capacity, limit)
    // 找最大值
    res := int64(math.MinInt64)
    for _, radiance := range glory {
        if radiance > res {
            res = radiance
        }
    }
    return int(res)
}

// func main() {
//     // Example 1:
//     // Input: edges = [[0,1],[0,2],[0,3],[1,4],[1,5]], nums = [1,0,-10,3,4,5], k = 2
//     // Output: 23
//     // Explanation:
//     // <img src="https://assets.leetcode.com/uploads/2025/04/21/4183example1drawio.png" />
//     // After inverting the subtree rooted at node 2, the maximum sum becomes 1 + 0 + 10 + 3 + 4 + 5 = 23.
//     fmt.Println(subtreeInversionSum([][]int{{0,1},{0,2},{0,3},{1,4},{1,5}}, []int{1,0,-10,3,4,5}, 2)) // 23
//     // Example 2:
//     // Input: edges = [[0,1],[1,2]], nums = [5,-10,-10], k = 1
//     // Output: 25
//     // Explanation:
//     // <img src="https://assets.leetcode.com/uploads/2025/04/21/4183example2drawio.png" />
//     // After inverting the subtree rooted at node 1, the maximum sum becomes 5 + 10 + 10 = 25.
//     fmt.Println(subtreeInversionSum([][]int{{0,1},{1,2}}, []int{5,-10,-10}, 1)) // 25
//     // Example 3:
//     // Input: edges = [[0,1],[0,2]], nums = [1,-5,-6], k = 2
//     // Output: 12
//     // Explanation:
//     // <img src="https://assets.leetcode.com/uploads/2025/04/21/4183example3drawio.png" />
//     // After inverting the subtrees rooted at nodes 1 and 2, nums = [1, 5, 6].
//     // This is valid because nodes 1 and 2 are two edges apart (1 → 0 and 0 → 2), which is at least k.
//     // The maximum sum is 1 + 5 + 6 = 12.
//     fmt.Println(subtreeInversionSum([][]int{{0,1},{0,2}}, []int{1,-5,-6}, 2)) // 12
//     // Example 4:
//     // Input: edges = [[0,1],[0,2]], nums = [1,-5,-6], k = 3
//     // Output: 10
//     // Explanation:
//     // <img src="https://assets.leetcode.com/uploads/2025/04/21/4183example4drawio.png" />
//     // After inverting the subtree rooted at nodes 0, nums = [-1, 5, 6].
//     // The maximum sum is (-1) + 5 + 6 = 10.
//     // Note that we cannot invert nodes 1 and 2 because their distance is 2 < k = 3.
//     fmt.Println(subtreeInversionSum([][]int{{0,1},{0,2}}, []int{1,-5,-6}, 3)) // 10
// }

// 测试用例
func runTest(id int, edges [][]int, nums []int, k int, expect int) {
	actual := subtreeInversionSum(edges, nums, k)
	if actual == expect {
		fmt.Printf("✅ 测试用例 %d 成功 | 输出:%d\n", id, actual)
	} else {
		fmt.Printf("❌ 测试用例 %d 失败 | 期望:%d 实际:%d\n", id, expect, actual)
	}
}

func main() {
	runTest(1, [][]int{{0,1},{0,2},{0,3},{1,4},{1,5}}, []int{1,0,-10,3,4,5}, 2, 23)
	runTest(2, [][]int{{0,1},{1,2}}, []int{5,-10,-10}, 1, 25)
	runTest(3, [][]int{{0,1},{0,2}}, []int{1,-5,-6}, 2, 12)
	runTest(4, [][]int{{0,1},{0,2}}, []int{1,-5,-6}, 3, 10)
    runTest(5, [][]int{{2,1},{0,1}}, []int{-5,-5,2}, 3, 8)
}