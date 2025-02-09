package main

// LCP 80. 生物进化录
// 在永恒之森中，存在着一本生物进化录，以 一个树形结构 记载了所有生物的演化过程。
// 经过观察并整理了各节点间的关系，parents[i] 表示编号 i 节点的父节点编号(根节点的父节点为 -1)。

// 为了探索和记录其中的演化规律，队伍中的炼金术师提出了一种方法，可以以字符串的形式将其复刻下来，规则如下：
//     1. 初始只有一个根节点，表示演化的起点，依次记录 01 字符串中的字符，
//     2. 如果记录 0，则在当前节点下添加一个子节点，并将指针指向新添加的子节点；
//     3. 如果记录 1，则将指针回退到当前节点的父节点处。

// 现在需要应用上述的记录方法，复刻下它的演化过程。
// 请返回能够复刻演化过程的字符串中， 字典序最小 的 01 字符串。

// 注意：
//     节点指针最终可以停在任何节点上，不一定要回到根节点。

// 示例 1：
// 输入：parents = [-1,0,0,2]
// 输出："00110"
// 解释：树结构如下图所示，共存在 2 种记录方案： 
// 第 1 种方案为：0(记录编号 1 的节点) -> 1(回退至节点 0) -> 0(记录编号 2 的节点) -> 0((记录编号 3 的节点)) 
// 第 2 种方案为：0(记录编号 2 的节点) -> 0(记录编号 3 的节点) -> 1(回退至节点 2) -> 1(回退至节点 0) -> 0(记录编号 1 的节点) 返回字典序更小的 "00110"
// <img src="https://pic.leetcode.cn/1682319485-cRVudI-image.png" />
// <img src="https://pic.leetcode.cn/1682412701-waHdnm-%E8%BF%9B%E5%8C%96%20(3).gif" />

// 示例 2：
// 输入：parents = [-1,0,0,1,2,2]
// 输出："00101100"

// 提示：
//     1 <= parents.length <= 10^4
//     -1 <= parents[i] < i (即父节点编号小于子节点)

import "fmt"
import "strings"
import "sort"

func evolutionaryRecord(parents []int) string {
    n := len(parents)
    graph := make([][]int, n)
    for i := 1; i < n; i++ {
        p := parents[i]
        graph[p] = append(graph[p], i) // 建树
    }
    var dfs func(v int) string
    dfs = func(v int) string {
        arr := make([]string, len(graph[v]))
        for i, w := range graph[v] {
            arr[i] = dfs(w)
        }
        sort.Strings(arr)
        return "0" + strings.Join(arr, "") + "1"
    }
    return strings.TrimRight(dfs(0)[1:], "1") // 去掉根节点以及返回根节点的路径
}

func main() {
    // 示例 1：
    // 输入：parents = [-1,0,0,2]
    // 输出："00110"
    // 解释：树结构如下图所示，共存在 2 种记录方案： 
    // 第 1 种方案为：0(记录编号 1 的节点) -> 1(回退至节点 0) -> 0(记录编号 2 的节点) -> 0((记录编号 3 的节点)) 
    // 第 2 种方案为：0(记录编号 2 的节点) -> 0(记录编号 3 的节点) -> 1(回退至节点 2) -> 1(回退至节点 0) -> 0(记录编号 1 的节点) 返回字典序更小的 "00110"
    // <img src="https://pic.leetcode.cn/1682319485-cRVudI-image.png" />
    // <img src="https://pic.leetcode.cn/1682412701-waHdnm-%E8%BF%9B%E5%8C%96%20(3).gif" />
    fmt.Println(evolutionaryRecord([]int{-1,0,0,2})) // "00110"
    // 示例 2：
    // 输入：parents = [-1,0,0,1,2,2]
    // 输出："00101100"
    fmt.Println(evolutionaryRecord([]int{-1,0,0,1,2,2})) // "00101100"
}