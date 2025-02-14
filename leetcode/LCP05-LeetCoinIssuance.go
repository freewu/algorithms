package main

// LCP 05. 发 LeetCoin
// 力扣决定给一个刷题团队发LeetCoin作为奖励。同时，为了监控给大家发了多少LeetCoin，力扣有时候也会进行查询。

// 该刷题团队的管理模式可以用一棵树表示：
//     1. 团队只有一个负责人，编号为1。除了该负责人外，每个人有且仅有一个领导（负责人没有领导）；
//     2. 不存在循环管理的情况，如A管理B，B管理C，C管理A。
 

// 力扣想进行的操作有以下三种：
//     1. 给团队的一个成员（也可以是负责人）发一定数量的LeetCoin；
//     2. 给团队的一个成员（也可以是负责人），以及他/她管理的所有人（即他/她的下属、他/她下属的下属，……），发一定数量的LeetCoin；
//     3. 查询某一个成员（也可以是负责人），以及他/她管理的所有人被发到的LeetCoin之和。

// 输入：
//     1. N表示团队成员的个数（编号为1～N，负责人为1）；
//     2. leadership是大小为(N - 1) * 2的二维数组，其中每个元素[a, b]代表b是a的下属；
//     3. operations是一个长度为Q的二维数组，代表以时间排序的操作，格式如下：
//         3.1 operations[i][0] = 1: 代表第一种操作，operations[i][1]代表成员的编号，operations[i][2]代表LeetCoin的数量；
//         3.2 operations[i][0] = 2: 代表第二种操作，operations[i][1]代表成员的编号，operations[i][2]代表LeetCoin的数量；
//         3.3 operations[i][0] = 3: 代表第三种操作，operations[i][1]代表成员的编号；

// 输出：
//     1. 返回一个数组，数组里是每次查询的返回值（发LeetCoin的操作不需要任何返回值）。
//        由于发的LeetCoin很多，请把每次查询的结果模1e9+7 (1000000007)。

 

// 示例 1：
// 输入：N = 6, leadership = [[1, 2], [1, 6], [2, 3], [2, 5], [1, 4]], operations = [[1, 1, 500], [2, 2, 50], [3, 1], [2, 6, 15], [3, 1]]
// 输出：[650, 665]
// 解释：团队的管理关系见下图。
// 第一次查询时，每个成员得到的LeetCoin的数量分别为（按编号顺序）：500, 50, 50, 0, 50, 0;
// 第二次查询时，每个成员得到的LeetCoin的数量分别为（按编号顺序）：500, 50, 50, 0, 50, 15.
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/09/09/coin_example_1.jpg" />

// 限制：
//     1 <= N <= 50000
//     1 <= Q <= 50000
//     operations[i][0] != 3 时，1 <= operations[i][2] <= 5000

import "fmt"

const mod int = 1e9 + 7

type SegmentTree struct {
    tree []int // 线段树数据
    lazy []int // lazy数据
}

func NewSegmentTree(n int) SegmentTree {
    return SegmentTree{ tree: make([]int, n * 4), lazy: make([]int, n * 4) }
}

func (t *SegmentTree) pushDown(pos, left, right int) {
    if t.lazy[pos] == 0 { return }
    t.tree[pos] += t.lazy[pos] * (right - left + 1)
    if left < right {
        t.lazy[pos * 2 + 1] += t.lazy[pos]
        t.lazy[pos * 2 + 2] += t.lazy[pos]
    }
    t.lazy[pos] = 0
}

func (t *SegmentTree) update(pos, left, right, curLeft, curRight, val int) {
    t.pushDown(pos, left, right) // 如果当前lazy有值，先转换为树的值，这样上面的聚合才能正确
    if right < curLeft || left > curRight { return }
    if curLeft <= left && right <= curRight {
        t.lazy[pos] = val
        t.pushDown(pos, left, right) // 同理，设置了lazy后要转换为树的值，上层的聚合才能正确
        return
    }
    // 递归更新
    mid := (left + right) >> 1
    t.update(pos * 2 + 1, left, mid, curLeft, curRight, val)
    t.update(pos * 2 + 2, mid + 1, right, curLeft, curRight, val)
    t.tree[pos] = (t.tree[pos * 2 + 1] + t.tree[pos * 2 + 2]) % mod
}

func (t *SegmentTree) query(pos, left, right, curLeft, curRight int) int {
    if right < curLeft || left > curRight { return 0 }
    t.pushDown(pos, left, right)
    if curLeft <= left && right <= curRight { return t.tree[pos] }
    mid := (left + right) >> 1
    return t.query(pos * 2 + 1, left, mid, curLeft, curRight) + t.query(pos * 2 + 2, mid + 1, right, curLeft, curRight)
}

func bonus(n int, leadership [][]int, operations [][]int) []int {
    graph := map[int][]int{} // leader -- followers
    for _, l := range leadership {
        graph[l[0]] = append(graph[l[0]], l[1])
    }
    // 使用dfs将题目中的节点转换为线段树里的，题目里的数字只能算是“Value"
    // 而线段树里的需要以map的key为题目中的数字， value为线段树id
    L, R := map[int]int{}, map[int]int{} // left和right
    id := -1 // 用sequence的形式为节点设置线段树的id
    var dfs func(int)
    dfs = func(leader int) {
        id++
        L[leader] = id
        for _, follower := range graph[leader] {
            dfs(follower)
        }
        R[leader] = id
    }
    dfs(1)
    // 使用线段树执行操作
    t := NewSegmentTree(n)
    res := []int{}
    for _, op := range operations {
        if op[0] == 1 {
            t.update(0, 0, n - 1, L[op[1]], L[op[1]], op[2])
        } else if op[0] == 2 {
            t.update(0, 0, n - 1, L[op[1]], R[op[1]], op[2])
        } else {
            res = append(res, t.query(0, 0, n - 1, L[op[1]], R[op[1]]) % mod)
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：N = 6, leadership = [[1, 2], [1, 6], [2, 3], [2, 5], [1, 4]], operations = [[1, 1, 500], [2, 2, 50], [3, 1], [2, 6, 15], [3, 1]]
    // 输出：[650, 665]
    // 解释：团队的管理关系见下图。
    // 第一次查询时，每个成员得到的LeetCoin的数量分别为（按编号顺序）：500, 50, 50, 0, 50, 0;
    // 第二次查询时，每个成员得到的LeetCoin的数量分别为（按编号顺序）：500, 50, 50, 0, 50, 15.
    // <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/09/09/coin_example_1.jpg" />
    fmt.Println(bonus(6, [][]int{{1, 2}, {1, 6}, {2, 3}, {2, 5}, {1, 4}}, [][]int{{1, 1, 500}, {2, 2, 50}, {3, 1}, {2, 6, 15}, {3, 1}})) // [650, 665]
}