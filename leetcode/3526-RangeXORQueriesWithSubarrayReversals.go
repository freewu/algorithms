package main

// 3526. Range XOR Queries with Subarray Reversals
// You are given an integer array nums of length n and a 2D integer array queries of length q, 
// where each query is one of the following three types:

//     1 Update: queries[i] = [1, index, value]
//         Set nums[index] = value.

//     2. Range XOR Query: queries[i] = [2, left, right]
//         Compute the bitwise XOR of all elements in the subarray nums[left...right], and record this result.

//     3. Reverse Subarray: queries[i] = [3, left, right]
//         Reverse the subarray nums[left...right] in place.

// Return an array of the results of all range XOR queries in the order they were encountered.

// Example 1:
// Input: nums = [1,2,3,4,5], queries = [[2,1,3],[1,2,10],[3,0,4],[2,0,4]]
// Output: [5,8]
// Explanation:
// Query 1: [2, 1, 3] – Compute XOR of subarray [2, 3, 4] resulting in 5.
// Query 2: [1, 2, 10] – Update nums[2] to 10, updating the array to [1, 2, 10, 4, 5].
// Query 3: [3, 0, 4] – Reverse the entire array to get [5, 4, 10, 2, 1].
// Query 4: [2, 0, 4] – Compute XOR of subarray [5, 4, 10, 2, 1] resulting in 8.

// Example 2:
// Input: nums = [7,8,9], queries = [[1,0,3],[2,0,2],[3,1,2]]
// Output: [2]
// Explanation:
// Query 1: [1, 0, 3] – Update nums[0] to 3, updating the array to [3, 8, 9].
// Query 2: [2, 0, 2] – Compute XOR of subarray [3, 8, 9] resulting in 2.
// Query 3: [3, 1, 2] – Reverse the subarray [8, 9] to get [9, 8].

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9
//     1 <= queries.length <= 10^5
//     queries[i].length == 3​
//     queries[i][0] ∈ {1, 2, 3}​
//     If queries[i][0] == 1:​
//         0 <= index < nums.length​
//         0 <= value <= 10^9
//     If queries[i][0] == 2 or queries[i][0] == 3:​
//         0 <= left <= right < nums.length​

import "fmt"

type Node struct {
    ch  [2]*Node
    sz  int
    key int
    sum int
    flipTodo bool
    }

// 设置如下返回值是为了方便使用 node 中的 ch 数组
func (o *Node) cmpKth(k int) int {
    d := k - o.ch[0].size() - 1
    switch {
    case d < 0:
        return 0 // 左儿子
    case d > 0:
        return 1 // 右儿子
    default:
        return -1
    }
}

func (o *Node) size() int {
    if o != nil { return o.sz }
    return 0
}

func (o *Node) xorSum() int {
    if o != nil { return o.sum }
    return 0
}

func (o *Node) maintain() {
    o.sz = 1 + o.ch[0].size() + o.ch[1].size()
    o.sum = o.key ^ o.ch[0].xorSum() ^ o.ch[1].xorSum()
}

func (o *Node) apply() {
    if o != nil {
        o.flipTodo = !o.flipTodo
    }
}

func (o *Node) pushDown() {
    if !o.flipTodo { return }
    o.ch[0].apply()
    o.ch[1].apply()
    o.ch[0], o.ch[1] = o.ch[1], o.ch[0]
    o.flipTodo = false
}

// 旋转，并维护子树大小
// d=0：左旋，返回 o 的右儿子
// d=1：右旋，返回 o 的左儿子
func (o *Node) rotate(d int) *Node {
    x := o.ch[d^1]
    o.ch[d^1] = x.ch[d]
    x.ch[d] = o
    o.maintain()
    x.maintain()
    return x
}

// 将子树 o（中序遍历）的第 k 个节点伸展到 o，并返回该节点
// 1 <= k <= o.size()
func (o *Node) splay(k int) (kth *Node) {
    o.pushDown()
    d := o.cmpKth(k)
    if d < 0 {
        return o
    }
    k -= d * (o.ch[0].size() + 1)
    c := o.ch[d]
    c.pushDown()
    if d2 := c.cmpKth(k); d2 >= 0 {
        c.ch[d2] = c.ch[d2].splay(k - d2*(c.ch[0].size()+1))
        if d2 == d {
            o = o.rotate(d ^ 1)
        } else {
            o.ch[d] = c.rotate(d)
        }
    }
    return o.rotate(d ^ 1)
}

// 分裂子树 o，把 o（中序遍历）的前 k 个节点放在 lo 子树，其余放在 ro 子树
// 返回的 lo 节点为 o（中序遍历）的第 k 个节点
// 1 <= k <= o.size()
// 特别地，k = o.size() 时 ro 为 nil
func (o *Node) split(k int) (lo, ro *Node) {
    lo = o.splay(k)
    ro = lo.ch[1]
    lo.ch[1] = nil
    lo.maintain()
    return
}

// 把子树 ro 合并进子树 o，返回合并前 o（中序遍历）的最后一个节点
// 相当于把 ro 的中序遍历 append 到 o 的中序遍历之后
// ro 可以为 nil，但 o 不能为 nil
func (o *Node) merge(ro *Node) *Node {
    // 把最大节点伸展上来，这样会空出一个右儿子用来合并 ro
    o = o.splay(o.size())
    o.ch[1] = ro
    o.maintain()
    return o
}

// 构建一棵中序遍历为 a 的 splay 树
// 比如，给你一个序列和一些修改操作，每次取出一段子区间，cut 掉然后 append 到末尾，输出完成所有操作后的最终序列：
//     我们可以 buildSplay(1,n)，每次操作调用两次 split 来 cut 区间，得到三棵子树 a b c
//     append 之后应该是 a c b，那么我们可以 a.merge(c.merge(b)) 来完成这一操作
//     注意 merge 后可能就不满足搜索树的性质了，但是没有关系，中序遍历的结果仍然是正确的，我们只要保证这一点成立，就能正确得到完成所有操作后的最终序列
func buildSplay(a []int) *Node {
    if len(a) == 0 { return nil }
    m := len(a) / 2
    o := &Node{key: a[m]}
    o.ch[0] = buildSplay(a[:m])
    o.ch[1] = buildSplay(a[m+1:])
    o.maintain()
    return o
}

func getResults(nums []int, queries [][]int) []int {
    nums = append([]int{0}, nums...) // 前面插个哨兵，保证 t.split(l) 的返回值 lo 非空
    res, t := []int{}, buildSplay(nums)
    for _, q := range queries {
        if q[0] == 1 {
            t = t.splay(q[1] + 2)
            t.sum ^= t.key ^ q[2]
            t.key = q[2]
        } else {
            l, r := q[1]+1, q[2]+1
            // 先拆开
            lo, o := t.split(l)
            mo, ro := o.split(r - l + 1)
            if q[0] == 2 {
                res = append(res, mo.sum)
            } else {
                mo.apply() // 把中间这段反转
            }
            t = lo.merge(mo).merge(ro) // 再拼回去
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5], queries = [[2,1,3],[1,2,10],[3,0,4],[2,0,4]]
    // Output: [5,8]
    // Explanation:
    // Query 1: [2, 1, 3] – Compute XOR of subarray [2, 3, 4] resulting in 5.
    // Query 2: [1, 2, 10] – Update nums[2] to 10, updating the array to [1, 2, 10, 4, 5].
    // Query 3: [3, 0, 4] – Reverse the entire array to get [5, 4, 10, 2, 1].
    // Query 4: [2, 0, 4] – Compute XOR of subarray [5, 4, 10, 2, 1] resulting in 8.
    fmt.Println(getResults([]int{1,2,3,4,5}, [][]int{{2,1,3},{1,2,10},{3,0,4},{2,0,4}})) // [5,8]
    // Example 2:
    // Input: nums = [7,8,9], queries = [[1,0,3],[2,0,2],[3,1,2]]
    // Output: [2]
    // Explanation:
    // Query 1: [1, 0, 3] – Update nums[0] to 3, updating the array to [3, 8, 9].
    // Query 2: [2, 0, 2] – Compute XOR of subarray [3, 8, 9] resulting in 2.
    // Query 3: [3, 1, 2] – Reverse the subarray [8, 9] to get [9, 8].
    fmt.Println(getResults([]int{7,8,9}, [][]int{{1,0,3},{2,0,2},{3,1,2}})) // [2]

    fmt.Println(getResults([]int{1,2,3,4,5,6,7,8,9}, [][]int{{1,0,3},{2,0,2},{3,1,2}})) // [2]
    fmt.Println(getResults([]int{9,8,7,6,5,4,3,2,1}, [][]int{{1,0,3},{2,0,2},{3,1,2}})) // [12]
}