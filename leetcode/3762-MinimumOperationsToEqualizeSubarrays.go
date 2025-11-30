package main

// 3762. Minimum Operations to Equalize Subarrays
// You are given an integer array nums and an integer k.

// In one operation, you can increase or decrease any element of nums by exactly k.

// You are also given a 2D integer array queries, where each queries[i] = [li, ri].

// For each query, find the minimum number of operations required to make all elements in the subarray nums[li..ri] equal. 
// If it is impossible, the answer for that query is -1.

// Return an array ans, where ans[i] is the answer for the ith query.

// Example 1:
// Input: nums = [1,4,7], k = 3, queries = [[0,1],[0,2]]
// Output: [1,2]
// Explanation:
// One optimal set of operations:
// i	| [li, ri] | nums[li..ri] | Possibility | Operations                         | Final nums[li..ri] | ans[i]
// 0   | [0, 1]   | [1, 4]	      | Yes	        | nums[0] + k = 1 + 3 = 4 = nums[1]	 |  [4, 4]            | 1
// 1   | [0, 2]   | [1, 4, 7]    | Yes	        | nums[0] + k = 1 + 3 = 4 = nums[1]  |  [4, 4, 4]         | 2   
//                                              | nums[2] - k = 7 - 3 = 4 = nums[1]	 |
// Thus, ans = [1, 2].

// Example 2:
// Input: nums = [1,2,4], k = 2, queries = [[0,2],[0,0],[1,2]]
// Output: [-1,0,1]
// Explanation:
// One optimal set of operations:
// i   | [li, ri] | nums[li..ri]  | Possibility | Operations                         | Final nums[li..ri] | ans[i]  
// 0   | [0, 2]   | [1, 2, 4]	  | No	        | -	                                 | [1, 2, 4]	      |  -1
// 1   | [0, 0]   | [1]	          | Yes	        | Already equal	                     | [1]	              |  0
// 2   | [1, 2]   | [2, 4]	      | Yes	        | nums[1] + k = 2 + 2 = 4 = nums[2]  | [4, 4]	          |  1
// Thus, ans = [-1, 0, 1].

// Constraints:
//     1 <= n == nums.length <= 4 × 10^4
//     1 <= nums[i] <= 10^9​​​​​​​
//     1 <= k <= 10^9
//     1 <= queries.length <= 4 × 10^4
//     ​​​​​​​queries[i] = [li, ri]
//     0 <= li <= ri <= n - 1

import "fmt"
import "runtime/debug"
import "slices"
import "sort"

// 有大量指针的题目，关闭 GC 更快
func init() { debug.SetGCPercent(-1) } 

type Node struct {
    lo, ro   *Node
    l, r     int
    count, sum int
}

func (o *Node) maintain() {
    o.count = o.lo.count + o.ro.count
    o.sum = o.lo.sum + o.ro.sum
}

func build(l, r int) *Node {
    o := &Node{l: l, r: r}
    if l == r {
        return o
    }
    mid := (l + r) / 2
    o.lo = build(l, mid)
    o.ro = build(mid + 1, r)
    return o
}

// 在线段树的位置 i 添加 val
// 注意这里传的不是指针，会把 node 复制一份，而这正好是我们需要的
func (o Node) add(i, val int) *Node {
    if o.l == o.r {
        o.count++
        o.sum += val
        return &o
    }
    mid := (o.l + o.r) / 2
    if i <= mid {
        o.lo = o.lo.add(i, val)
    } else {
        o.ro = o.ro.add(i, val)
    }
    o.maintain()
    return &o
}

// 查询 old 和 o 对应区间的第 k 小，k 从 1 开始
func (o *Node) kth(old *Node, k int) int {
    if o.l == o.r {
        return o.l
    }
    diff := o.lo.count - old.lo.count
    if k <= diff { // 答案在左子树中
        return o.lo.kth(old.lo, k)
    }
    return o.ro.kth(old.ro, k - diff) // 答案在右子树中
}

// 查询 old 和 o 对应区间，有多少个数 <= i，这些数的元素和是多少
func (o *Node) query(old *Node, i int) (int, int) {
    if o.r <= i {
        return o.count - old.count, o.sum - old.sum
    }
    count, sum := o.lo.query(old.lo, i)
    mid := (o.l + o.r) / 2
    if i > mid {
        c, t := o.ro.query(old.ro, i)
        count += c
        sum += t
    }
    return count, sum
}

func minOperations(nums []int, k int, queries [][]int) []int64 {
    n := len(nums)
    left := make([]int, n)
    for i := 1; i < n; i++ {
        if nums[i]%k != nums[i-1]%k {
            left[i] = i
        } else {
            left[i] = left[i-1]
        }
    }
    // 准备离散化
    sorted := slices.Clone(nums)
    slices.Sort(sorted)
    sorted = slices.Compact(sorted)
    t := make([]*Node, n+1)
    t[0] = build(0, len(sorted)-1)
    for i, x := range nums {
        j := sort.SearchInts(sorted, x) // 离散化
        t[i+1] = t[i].add(j, x)         // 构建可持久化线段树
    }
    res := make([]int64, len(queries))
    for qi, q := range queries {
        l, r := q[0], q[1]
        if left[r] > l { // 无解
            res[qi] = -1
            continue
        }
        r++ // 改成左闭右开，方便计算
        // 计算区间中位数
        sz := r - l
        i := t[r].kth(t[l], sz/2+1)
        median := sorted[i] // 离散化后的值 -> 原始值
        // 计算区间所有元素到中位数的距离和
        total := t[r].sum - t[l].sum // 区间元素和
        cntLeft, sumLeft := t[r].query(t[l], i)
        sum := median * cntLeft - sumLeft              // 蓝色面积
        sum += total - sumLeft - median*(sz - cntLeft) // 绿色面积
        res[qi] = int64(sum / k)                     // 操作次数 = 距离和 / k
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,4,7], k = 3, queries = [[0,1],[0,2]]
    // Output: [1,2]
    // Explanation:
    // One optimal set of operations:
    // i	| [li, ri] | nums[li..ri] | Possibility | Operations                         | Final nums[li..ri] | ans[i]
    // 0   | [0, 1]   | [1, 4]	      | Yes	        | nums[0] + k = 1 + 3 = 4 = nums[1]	 |  [4, 4]            | 1
    // 1   | [0, 2]   | [1, 4, 7]    | Yes	        | nums[0] + k = 1 + 3 = 4 = nums[1]  |  [4, 4, 4]         | 2   
    //                                              | nums[2] - k = 7 - 3 = 4 = nums[1]	 |
    // Thus, ans = [1, 2].
    fmt.Println(minOperations([]int{1,4,7}, 3, [][]int{{0,1},{0,2}})) // [1, 2]
    // Example 2:
    // Input: nums = [1,2,4], k = 2, queries = [[0,2],[0,0],[1,2]]
    // Output: [-1,0,1]
    // Explanation:
    // One optimal set of operations:
    // i   | [li, ri] | nums[li..ri]  | Possibility | Operations                         | Final nums[li..ri] | ans[i]  
    // 0   | [0, 2]   | [1, 2, 4]	  | No	        | -	                                 | [1, 2, 4]	      |  -1
    // 1   | [0, 0]   | [1]	          | Yes	        | Already equal	                     | [1]	              |  0
    // 2   | [1, 2]   | [2, 4]	      | Yes	        | nums[1] + k = 2 + 2 = 4 = nums[2]  | [4, 4]	          |  1
    // Thus, ans = [-1, 0, 1].
    fmt.Println(minOperations([]int{1,2,4}, 2, [][]int{{0,2},{0,0},{1,2}})) // [-1, 0, 1]

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, 3, [][]int{{0,1},{0,2}})) // [-1 -1]
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, 3, [][]int{{0,1},{0,2}})) // [-1 -1]
}