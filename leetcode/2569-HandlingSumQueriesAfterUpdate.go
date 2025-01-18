package main

// 2569. Handling Sum Queries After Update
// You are given two 0-indexed arrays nums1 and nums2 and a 2D array queries of queries. 
// There are three types of queries:
//     1. For a query of type 1, queries[i] = [1, l, r]. 
//        Flip the values from 0 to 1 and from 1 to 0 in nums1 from index l to index r. 
//        Both l and r are 0-indexed.
//     2. For a query of type 2, queries[i] = [2, p, 0]. 
//        For every index 0 <= i < n, set nums2[i] = nums2[i] + nums1[i] * p.
//     3. For a query of type 3, queries[i] = [3, 0, 0]. 
//        Find the sum of the elements in nums2.

// Return an array containing all the answers to the third type queries.

// Example 1:
// Input: nums1 = [1,0,1], nums2 = [0,0,0], queries = [[1,1,1],[2,1,0],[3,0,0]]
// Output: [3]
// Explanation: After the first query nums1 becomes [1,1,1]. After the second query, nums2 becomes [1,1,1], so the answer to the third query is 3. Thus, [3] is returned.

// Example 2:
// Input: nums1 = [1], nums2 = [5], queries = [[2,0,0],[3,0,0]]
// Output: [5]
// Explanation: After the first query, nums2 remains [5], so the answer to the second query is 5. Thus, [5] is returned.

// Constraints:
//     1 <= nums1.length,nums2.length <= 10^5
//     nums1.length = nums2.length
//     1 <= queries.length <= 10^5
//     queries[i].length = 3
//     0 <= l <= r <= nums1.length - 1
//     0 <= p <= 10^6
//     0 <= nums1[i] <= 1
//     0 <= nums2[i] <= 10^9

import "fmt"

type SegmentTree struct {
    tree []int
    lazy []int
}

func NewSegmentTree(n int) *SegmentTree {
    return &SegmentTree{ tree: make([]int, 4*n),  lazy: make([]int, 4*n), }
}

func (s *SegmentTree) Build(treeIndex int, treeLow int, treeHigh int, arr []int) {
    if treeLow == treeHigh {
        s.tree[treeIndex] = arr[treeLow]
        return
    }
    mid := (treeHigh - treeLow) / 2 + treeLow
    s.Build(treeIndex*2, treeLow, mid, arr)
    s.Build(treeIndex*2+1, mid+1, treeHigh, arr)
    s.tree[treeIndex] = s.tree[treeIndex*2] + s.tree[treeIndex*2+1]
}

func (s *SegmentTree) Update(treeIndex int, treeLow int, treeHigh int, updateLow int, updateHigh int) {
    if s.lazy[treeIndex] > 0 {
        s.tree[treeIndex] = treeHigh - treeLow + 1-s.tree[treeIndex]
        if treeLow != treeHigh {
            s.lazy[treeIndex*2] ^= 1
            s.lazy[treeIndex*2+1] ^= 1
        }
        s.lazy[treeIndex] = 0
    }
    if treeLow > updateHigh || treeHigh < updateLow { return }
    if updateLow <= treeLow && treeHigh <= updateHigh {
        s.tree[treeIndex] = treeHigh - treeLow + 1 - s.tree[treeIndex]
        if treeLow != treeHigh {
            s.lazy[treeIndex*2] ^= 1
            s.lazy[treeIndex*2+1] ^= 1
        }
        return
    }
    mid := (treeHigh - treeLow) / 2 + treeLow
    s.Update(treeIndex * 2, treeLow, mid, updateLow, updateHigh)
    s.Update(treeIndex * 2 + 1, mid + 1, treeHigh, updateLow, updateHigh)
    s.tree[treeIndex] = s.tree[treeIndex*2] + s.tree[treeIndex*2+1]
}

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
    res, n, sum := []int64{}, len(nums1), int64(0)
    for _, v := range nums2 {
        sum += int64(v)
    }
    segmentTree := NewSegmentTree(n)
    segmentTree.Build(1, 0, n - 1, nums1)
    for _, q := range queries {
        switch q[0] {
            case 1:
                segmentTree.Update(1, 0, n - 1, q[1], q[2])
            case 2:
                sum += int64(segmentTree.tree[1] * q[1])
            case 3:
                res = append(res, sum)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [1,0,1], nums2 = [0,0,0], queries = [[1,1,1],[2,1,0],[3,0,0]]
    // Output: [3]
    // Explanation: After the first query nums1 becomes [1,1,1]. After the second query, nums2 becomes [1,1,1], so the answer to the third query is 3. Thus, [3] is returned.
    fmt.Println(handleQuery([]int{1,0,1}, []int{0,0,0}, [][]int{{1,1,1},{2,1,0},{3,0,0}})) // [3]
    // Example 2:
    // Input: nums1 = [1], nums2 = [5], queries = [[2,0,0],[3,0,0]]
    // Output: [5]
    // Explanation: After the first query, nums2 remains [5], so the answer to the second query is 5. Thus, [5] is returned.
    fmt.Println(handleQuery([]int{1}, []int{5}, [][]int{{2,0,0},{3,0,0}})) // [5]
}