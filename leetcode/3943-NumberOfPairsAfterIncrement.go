package main

// 3943. Number of Pairs After Increment
// You are given two integer arrays nums1 and nums2, and a 2D integer array queries.

// Each queries[i] is one of the following types:
//     1. [1, x, y, val] – Add val to every element in nums2[x..y].
//     2. [2, tot] – Compute the number of pairs (j, k) such that nums1[j] + nums2[k] == tot.

// Return an integer array answer, where answer[j] is the number of pairs for the jth query of type 2.

// Example 1:
// Input: nums1 = [1,2], nums2 = [3,4], queries = [[2,5],[1,0,0,2],[2,5]]
// Output: [2,1]
// Explanation:
// queries[0] = [2, 5]: Valid pairs are nums1[0] + nums2[1] = 1 + 4 = 5 and nums1[1] + nums2[0] = 2 + 3 = 5.
// queries[1] = [1, 0, 0, 2]: Add 2 to nums2[0], resulting in nums2 = [5, 4].
// queries[2] = [2, 5]: Valid pair is nums1[0] + nums2[1] = 1 + 4 = 5.
// Thus, the answer = [2, 1].

// Example 2:
// Input: nums1 = [1,1], nums2 = [2,2,3], queries = [[2,4],[1,0,1,1],[2,4]]
// Output: [2,6]
// Explanation:
// queries[0] = [2, 4]: Valid pairs are nums1[0] + nums2[2] = 1 + 3 and nums1[1] + nums2[2] = 1 + 3.
// queries[1] = [1, 0, 1, 1]: Add 1 to nums2[0] and nums2[1], resulting in nums2 = [3, 3, 3].
// queries[2] = [2, 4]: Every element of nums1 = [1, 1] pairs with every element of nums2 = [3, 3, 3] as 1 + 3 = 4. That gives 2 × 3 = 6 pairs in total.
// Thus, the answer = [2, 6].

// Example 3:
// Input: nums1 = [2,5,8,4], nums2 = [1,3,8], queries = [[2,9],[1,1,2,1],[2,10]]
// Output: [1,0]
// Explanation:
// queries[0] = [2, 9]: Only valid pair is nums1[2] + nums2[0] = 8 + 1 = 9.
// queries[1] = [1, 1, 2, 1]: Add 1 to nums2[1] and nums2[2], resulting in​​​​​​​ nums2 = [1, 4, 9].
// queries[2] = [2, 10]: No pair sums to 10.
// Thus, the answer = [1, 0].

// Constraints:
//     1 <= nums1.length <= 5
//     1 <= nums2.length <= 5 * 10^4
//     1 <= nums1[i], nums2[i] <= 10^5
//     1 <= queries.length <= 5 * 10^4
//     queries[i].length == 2 or 4
//     queries[i] == [1, x, y, val], or
//     queries[i] == [2, tot]
//     0 <= x <= y < nums2.length
//     1 <= val <= 10^5
//     1 <= tot <= 10^9

import "fmt"
import "math"

func numberOfPairs(nums1, nums2 []int, queries [][]int) []int {
    res, m, n := make([]int, 0), len(nums1), len(nums2)
    b := int(math.Sqrt(float64(m * n)))
    type Block struct {
        l, r int         // 这一段对应 nums2 的子数组 [l, r)，注意是左闭右开区间
        count  map[int]int // 这一段每个元素的出现次数
        add  int         // 这一段整体要增加 add
    }
    blocks := make([]Block, (n - 1) / b + 1)
    for i := 0; i < n; i += b {
        r := min(i + b, n)
        count := map[int]int{}
        for _, x := range nums2[i:r] {
            count[x]++
        }
        blocks[i/b] = Block{i, r, count, 0}
    }
    for _, q := range queries {
        if q[0] == 1 {
            l, r, val := q[1], q[2]+1, q[3]
            for i := range blocks {
                b := &blocks[i]
                if b.r <= l {
                    continue
                }
                if b.l >= r {
                    break
                }
                // b 在 [l, r) 中
                if l <= b.l && b.r <= r {
                    b.add += val
                    continue
                }
                // b 的一部分在 [l, r) 中
                bl := max(b.l, l)
                br := min(b.r, r)
                // 暴力更新 nums2 的子数组 [bl, br) 的元素值及其出现次数
                for j := bl; j < br; j++ {
                    b.count[nums2[j]]-- // 撤销旧的
                    nums2[j] += val
                    b.count[nums2[j]]++ // 添加新的
                }
            }
        } else {
            val := 0
            for _, b := range blocks {
                target := q[1] - b.add
                for _, x := range nums1 {
                    val += b.count[target-x]
                }
            }
            res = append(res, val)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [1,2], nums2 = [3,4], queries = [[2,5],[1,0,0,2],[2,5]]
    // Output: [2,1]
    // Explanation:
    // queries[0] = [2, 5]: Valid pairs are nums1[0] + nums2[1] = 1 + 4 = 5 and nums1[1] + nums2[0] = 2 + 3 = 5.
    // queries[1] = [1, 0, 0, 2]: Add 2 to nums2[0], resulting in nums2 = [5, 4].
    // queries[2] = [2, 5]: Valid pair is nums1[0] + nums2[1] = 1 + 4 = 5.
    // Thus, the answer = [2, 1].
    fmt.Println(numberOfPairs([]int{1,2}, []int{3,4}, [][]int{{2,5},{1,0,0,2},{2,5}})) // [2,1]
    // Example 2:
    // Input: nums1 = [1,1], nums2 = [2,2,3], queries = [[2,4],[1,0,1,1],[2,4]]
    // Output: [2,6]
    // Explanation:
    // queries[0] = [2, 4]: Valid pairs are nums1[0] + nums2[2] = 1 + 3 and nums1[1] + nums2[2] = 1 + 3.
    // queries[1] = [1, 0, 1, 1]: Add 1 to nums2[0] and nums2[1], resulting in nums2 = [3, 3, 3].
    // queries[2] = [2, 4]: Every element of nums1 = [1, 1] pairs with every element of nums2 = [3, 3, 3] as 1 + 3 = 4. That gives 2 × 3 = 6 pairs in total.
    // Thus, the answer = [2, 6].
    fmt.Println(numberOfPairs([]int{1,1}, []int{2,2,3}, [][]int{{2,4},{1,0,1,1},{2,4}})) // [2,6]
    // Example 3:
    // Input: nums1 = [2,5,8,4], nums2 = [1,3,8], queries = [[2,9],[1,1,2,1],[2,10]]
    // Output: [1,0]
    // Explanation:
    // queries[0] = [2, 9]: Only valid pair is nums1[2] + nums2[0] = 8 + 1 = 9.
    // queries[1] = [1, 1, 2, 1]: Add 1 to nums2[1] and nums2[2], resulting in​​​​​​​ nums2 = [1, 4, 9].
    // queries[2] = [2, 10]: No pair sums to 10.
    // Thus, the answer = [1, 0].
    fmt.Println(numberOfPairs([]int{2,5,8,4}, []int{1,3,8}, [][]int{{2,9},{1,1,2,1},{2,10}})) // [1,0]

    fmt.Println(numberOfPairs([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, [][]int{{2,9},{1,1,2,1},{2,10}})) // [8 9]
    fmt.Println(numberOfPairs([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, [][]int{{2,9},{1,1,2,1},{2,10}})) // [8 9]
    fmt.Println(numberOfPairs([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, [][]int{{2,9},{1,1,2,1},{2,10}})) // [8 9]
    fmt.Println(numberOfPairs([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, [][]int{{2,9},{1,1,2,1},{2,10}})) // [8 9]
}