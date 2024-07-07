package main

// LCR 061. 查找和最小的 K 对数字
// 给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k 。
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
// 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。

// 示例 1:
// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [1,2],[1,4],[1,6]
// 解释: 返回序列中的前 3 对数：
//     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

// 示例 2:
// 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
// 输出: [1,1],[1,1]
// 解释: 返回序列中的前 2 对数：
//      [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]

// 示例 3:
// 输入: nums1 = [1,2], nums2 = [3], k = 3 
// 输出: [1,3],[2,3]
// 解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]

// 提示:
//     1 <= nums1.length, nums2.length <= 10^4
//     -10^9 <= nums1[i], nums2[i] <= 10^9
//     nums1, nums2 均为升序排列
//     1 <= k <= 1000

import "fmt"
import "container/heap"
import "sort"

// min heap
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    res := make([][]int, 0)
    minHeap := &SetHeap{}
    heap.Init(minHeap)
    for i, v := range nums1 {
        heap.Push(minHeap, Set{sum: v + nums2[0], i: i, j: 0})
    }
    for !minHeap.Empty() && k > 0 {
        cur := heap.Pop(minHeap).(Set)
        i, j := cur.i, cur.j
        res = append(res, []int{ nums1[i], nums2[j] })
        next := j + 1
        if next < len(nums2) {
            heap.Push(minHeap, Set{sum: nums1[i] + nums2[next], i: i, j: next})
        }
        k--
    } 
    return res
}

type Set struct {
    sum int
    i int
    j int
}
type SetHeap []Set
func (h SetHeap)  Len() int           { return len(h) }
func (h SetHeap)  Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h SetHeap)  Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h SetHeap)  Empty() bool        { return len(h) == 0 }
func (h *SetHeap) Push(x interface{}) { *h = append(*h, x.(Set)) }
func (h *SetHeap) Pop() interface{}   { old := *h; n := len(old); x := old[n-1]; *h = old[0 : n-1]; return x; }

// 二分
func kSmallestPairs1(nums1 []int, nums2 []int, k int) [][]int {
    m, n := len(nums1), len(nums2)
    l, r := nums1[0] + nums2[0], nums1[m-1] + nums2[n-1]
    check := func (t int, nums1, nums2 []int) int {
        m, n := len(nums1), len(nums2)
        i, j := 0, n - 1
        cnt := 1
        for i < m && j >= 0 {
            if nums1[i] + nums2[j] >= t {
                j--
            } else {
                i++
                cnt += j + 1
            }
        }
        return cnt
    }
    for l <= r {
        mid := (r + l) / 2
        if check(mid, nums1, nums2) <= k {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    i, j := 0, n - 1
    res := make([][]int, 0)
    for i < m && j >= 0 {
        if nums1[i] + nums2[j] >= r {
            j--
        } else {
            for c := 0; c <= j; c++ {
                cur := []int{nums1[i], nums2[c]}
                res = append(res, cur)
            }
            i++
        }
    }
    if len(res) == k {
        return res
    }
    i,j = 0, n - 1
    for i < m && j >= 0 {
        if nums1[i]+nums2[j] > r {
            j--
        } else {
            for c := 0; c <= j; c++ {
                if nums1[i] + nums2[c] == r {
                    cur := []int{nums1[i], nums2[c]}
                    res = append(res, cur)
                }
                if len(res) == k {
                    return res
                }
            }
            i++
        }
    }
    sort.Slice(res, func(i, j int) bool {
        if res[i][0] + res[i][1] == res[j][0] + res[j][1] {
            return res[i][0] < res[j][0]
        } else {
            return res[i][0] + res[i][1] < res[j][0] + res[j][1]
        }
    })
    return res
}

func main() {
    // Explanation: The first 3 pairs are returned from the sequence: [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
    fmt.Println(kSmallestPairs([]int{1,7,11},[]int{2,4,6}, 3)) // [[1,2],[1,4],[1,6]]
    // Explanation: The first 2 pairs are returned from the sequence: [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
    fmt.Println(kSmallestPairs([]int{1,1,2},[]int{1,2,3}, 2)) // [[1,1],[1,1]]

    fmt.Println(kSmallestPairs1([]int{1,7,11},[]int{2,4,6}, 3)) // [[1,2],[1,4],[1,6]]
    fmt.Println(kSmallestPairs1([]int{1,1,2},[]int{1,2,3}, 2)) // [[1,1],[1,1]]
}