package main

// LCP 54. 夺回据点
// 欢迎各位勇者来到力扣城，本次试炼主题为「夺回据点」。
// 魔物了占领若干据点，这些据点被若干条道路相连接，roads[i] = [x, y] 表示编号 x、y 的两个据点通过一条道路连接。
// 现在勇者要将按照以下原则将这些据点逐一夺回：
//     在开始的时候，勇者可以花费资源先夺回一些据点，初始夺回第 j 个据点所需消耗的资源数量为 cost[j]
//     接下来，勇者在不消耗资源情况下，每次可以夺回一个和「已夺回据点」相连接的魔物据点，并对其进行夺回

// 注：为了防止魔物暴动，勇者在每一次夺回据点后（包括花费资源夺回据点后），需要保证剩余的所有魔物据点之间是相连通的（不经过「已夺回据点」）。
// 请返回勇者夺回所有据点需要消耗的最少资源数量。

// 注意：
//     输入保证初始所有据点都是连通的，且不存在重边和自环
    
// 示例 1：
// <img src="https://pic.leetcode-cn.com/1648706944-KJstUN-image.png" />
// 输入： cost = [1,2,3,4,5,6] roads = [[0,1],[0,2],[1,3],[2,3],[1,2],[2,4],[2,5]]
// 输出：6
// 解释： 勇者消耗资源 6 夺回据点 0 和 4，魔物据点 1、2、3、5 相连通； 第一次夺回据点 1，魔物据点 2、3、5 相连通； 第二次夺回据点 3，魔物据点 2、5 相连通； 第三次夺回据点 2，剩余魔物据点 5； 第四次夺回据点 5，无剩余魔物据点； 因此最少需要消耗资源为 6，可占领所有据点。image.png

// 示例 2：
// <img src="https://pic.leetcode-cn.com/1648707186-LJRwzU-image.png" />
// 输入： cost = [3,2,1,4] roads = [[0,2],[2,3],[3,1]]
// 输出：2
// 解释： 勇者消耗资源 2 夺回据点 1，魔物据点 0、2、3 相连通； 第一次夺回据点 3，魔物据点 2、0 相连通； 第二次夺回据点 2，剩余魔物据点 0； 第三次夺回据点 0，无剩余魔物据点； 因此最少需要消耗资源为 2，可占领所有据点。image.png

// 提示：
//     1 <= roads.length, cost.length <= 10^5
//     0 <= roads[i][0], roads[i][1] < cost.length
//     1 <= cost[i] <= 10^9

import "fmt"
import "math"
// // https://leetcode.cn/problems/s5kipK/solutions/2354044/lcp-54-duo-hui-ju-dian-jie-ti-si-lu-lao-rpxjy/
// from typing import List
// from collections import defaultdict

// class Solution:
//     def minimumCost(self, cost: List[int], roads: List[List[int]]) -> int:
//         ans = 0
//         n = len(cost)
        
//         adj = defaultdict(list)
//         for u, v in roads:
//             adj[u].append(v)
//             adj[v].append(u)
        
//         def tarjan(u, dfn, low, tim, root, cut, vdcc, st):
//             tim[0] += 1
//             dfn[u] = low[u] = tim[0]
//             st.append(u)
//             child = 0
            
//             if not adj[u]:
//                 vdcc.append([u])
//                 return
            
//             for v in adj[u]:
//                 if dfn[v]:
//                     low[u] = min(low[u], dfn[v])
//                     continue
                
//                 tarjan(v, dfn, low, tim, root, cut, vdcc, st)
//                 low[u] = min(low[u], low[v])
                
//                 if low[v] >= dfn[u]:
//                     child += 1
//                     if u != root or child > 1:
//                         cut[u] = 1
//                     vdcc.append([])
//                     x = None
//                     while x != v:
//                         x = st.pop()
//                         vdcc[-1].append(x)
//                     vdcc[-1].append(u)
        
//         dfn = [0] * n
//         low = [0] * n
//         tim = [0]
//         root = 0
//         cut = {}
//         vdcc = []
//         st = []
        
//         tarjan(root, dfn, low, tim, root, cut, vdcc, st)
        
//         cnt = [0] * len(vdcc)
//         minCost = [float('inf')] * len(vdcc)
//         for i in range(len(vdcc)):
//             for x in vdcc[i]:
//                 if x in cut:
//                     cnt[i] += 1
//                 else:
//                     minCost[i] = min(minCost[i], cost[x])
        
//         Max = float('-inf')
//         for i in range(len(cnt)):
//             if cnt[i] <= 1:
//                 ans += minCost[i]
//                 Max = max(Max, minCost[i])
        
//         return ans if len(cnt) == 1 else ans - Max

// func minimumCost(cost []int, roads [][]int) int64 {
//     res, n := 0, len(cost)
//     g := [][]int{}
//     for _,v := range roads {
//         g[v[0]] = append(g[v[0]],v[1])
//         g[v[1]] = append(g[v[1]],v[0])
//     }
//     fmt.Println(n)
//     return int64(res)
// }

// def minimumCost(self, cost: List[int], roads: List[List[int]]) -> int:
//     n = len(cost)
//     G = [[] for i in range(n)]
//     for i,j in roads:
//         G[i].append(j)
//         G[j].append(i)
//     low = [n] * n
//     seen = {-1: -1}
//     cut = [0] * n
//     res = []
//     inf = float('inf')

//     def tarjan(i, pre):
//         seen[i] = len(seen) + 1
//         children = 0
//         min_cost = inf
//         count_cut = 0
//         for j in G[i]:
//             if j in seen:
//                 low[i] = min(low[i], seen[j])
//                 continue
//             children += 1
//             cur_cost, cur_cut = tarjan(j, i)
//             low[i] = min(low[i], low[j])
//             if seen[i] <= low[j]:
//                 if i != root or children > 1:
//                     cut[i] = 1
//                     cost[i] = inf
//             if i == root or seen[i] > low[j]:
//                 min_cost = min(min_cost, cur_cost)
//                 count_cut += cur_cut

//         min_cost = min(min_cost, cost[i])
//         count_cut += cut[i] > 0

//         if count_cut + (i != root) < 2 and seen[pre] <= low[i] and cut[i] == 0:
//             res.append(min_cost)
//         return [min_cost, count_cut]

//     tarjan(root:= 0, -1)
//     return sum(res) - max(res)

// func minimumCost(cost []int, roads [][]int) int64 {
//     n := len(cost)
//     graph := make([][]int,len(roads) + 1)
//     for _,v := range roads {
//         graph[v[0]] = append(graph[v[0]],v[1])
//         graph[v[1]] = append(graph[v[1]],v[0])
//     }
//     low, cut, seen, res := make([]int,n), make([]int,n), make(map[int]int), []int{}
//     seen[-1] = -1
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     var tarjan func(i, j int) (int, int) 
//     tarjan = func(i, j int) (int, int) {
//         seen[i] = len(seen) + 1
//         children, count_cut, min_cost := 0, 0, math.MaxInt32
//         for _, v := range graph[i] {
//             if seen[v] != 0 {
//                 low[i] = min(low[i], seen[j])
//                 continue
//             }
//             children += 1
//             cur_cost, cur_cut = tarjan(j, i)
//             low[i] = min(low[i], low[j])
//             if seen[i] <= low[j] {
//                 if i != root || children > 1 {
//                     cut[i] = 1
//                     cost[i] = math.MaxInt32
//                 }
//             }
//             if i == root || seen[i] > low[j] {
//                 min_cost = min(min_cost, cur_cost)
//                 count_cut += cur_cut
//             }
//         }
//         min_cost = min(min_cost, cost[i])
//         count_cut += cut[i] > 0
//         if count_cut + (i != root) < 2 && seen[pre] <= low[i] && cut[i] == 0 {
//             res = append(res, min_cost)
//         }
//         return min_cost, count_cut
//     }
//     tarjan(0, -1)
//     sum := func(nums []int) {
//         res := 0
//         for _,v := range nums {
//             res += v
//         }
//         return res
//     }
//     max := func(nums []int) {
//         res := math.MinInt32
//         for _,v := range nums {
//             if v > res {
//                 res = v
//             }
//         }
//         return res
//     }
//     return int64(sum(res) - max(res))
// }

// def minimumCost(self, cost: List[int], roads: List[List[int]]) -> int:
//     n = len(cost)
//     G = [[] for i in range(n)]
//     for i,j in roads:
//         G[i].append(j)
//         G[j].append(i)
//     low = [n] * n
//     seen = {-1: -1}
//     cut = [0] * n
//     res = []
//     inf = float('inf')

//     def tarjan(i, pre):
//         seen[i] = len(seen) + 1
//         children = 0
//         min_cost = inf
//         count_cut = 0
//         for j in G[i]:
//             if j in seen:
//                 low[i] = min(low[i], seen[j])
//                 continue
//             children += 1
//             cur_cost, cur_cut = tarjan(j, i)
//             low[i] = min(low[i], low[j])
//             if seen[i] <= low[j]:
//                 if i != root or children > 1:
//                     cut[i] = 1
//                     cost[i] = inf
//             if i == root or seen[i] > low[j]:
//                 min_cost = min(min_cost, cur_cost)
//                 count_cut += cur_cut

//         min_cost = min(min_cost, cost[i])
//         count_cut += cut[i] > 0

//         if count_cut + (i != root) < 2 and seen[pre] <= low[i] and cut[i] == 0:
//             res.append(min_cost)
//         return [min_cost, count_cut]

//     tarjan(root:= 0, -1)
//     return sum(res) - max(res)

func minimumCost(cost []int, roads [][]int) int64 {
    n := len(cost)
    graph := make([][]int,len(roads) + 1)
    for _,v := range roads {
        graph[v[0]] = append(graph[v[0]],v[1])
        graph[v[1]] = append(graph[v[1]],v[0])
    }
    low, cut, seen, res := make([]int,n), make([]int,n), make(map[int]int), []int{}
    seen[-1] = -1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var tarjan func(i, pre, root int) (int, int) 
    tarjan = func(i, pre, root int) (int, int) {
        seen[i] = len(seen) + 1
        children, count_cut, min_cost := 0, 0, math.MaxInt32
        for _, j := range graph[i] {
            if seen[j] != 0 {
                low[i] = min(low[i], seen[j])
                continue
            }
            children += 1
            cur_cost, cur_cut := tarjan(j, i, root)
            low[i] = min(low[i], low[j])
            if seen[i] <= low[j] {
                if i != root || children > 1 {
                    cut[i] = 1
                    cost[i] = math.MaxInt32
                }
            }
            if i == root || seen[i] > low[j] {
                min_cost = min(min_cost, cur_cost)
                count_cut += cur_cut
            }
        }
        min_cost = min(min_cost, cost[i])
        //count_cut += cut[i] > 0
        if cut[i] > 0 {
            count_cut++
        }
        if i != root {
            if count_cut < 3 && seen[pre] <= low[i] && cut[i] == 0 {
                res = append(res, min_cost)
            }
        }
        return min_cost, count_cut
    }
    tarjan(0, -1, 0)
    sum := func(nums []int) int {
        res := 0
        for _,v := range nums {
            res += v
        }
        return res
    }
    max := func(nums []int) int {
        res := math.MinInt32
        for _,v := range nums {
            if v > res {
                res = v
            }
        }
        return res
    }
    return int64(sum(res) - max(res))
}

// 作者：Han3000
// 链接：https://leetcode.cn/problems/s5kipK/solutions/1468501/by-han3000-yb3c/


func main() {
    //  解释： 勇者消耗资源 6 夺回据点 0 和 4，魔物据点 1、2、3、5 相连通； 
    //      第一次夺回据点 1，魔物据点 2、3、5 相连通； 
    //      第二次夺回据点 3，魔物据点 2、5 相连通； 
    //      第三次夺回据点 2，剩余魔物据点 5； 
    //      第四次夺回据点 5，无剩余魔物据点； 
    //  因此最少需要消耗资源为 6，可占领所有据点。
    fmt.Println(minimumCost([]int{1,2,3,4,5,6},[][]int{{0,1},{0,2},{1,3},{2,3},{1,2},{2,4},{2,5}})) // 6
    //  解释： 勇者消耗资源 2 夺回据点 1，魔物据点 0、2、3 相连通； 
    //      第一次夺回据点 3，魔物据点 2、0 相连通； 
    //      第二次夺回据点 2，剩余魔物据点 0； 
    //      第三次夺回据点 0，无剩余魔物据点； 
    //  因此最少需要消耗资源为 2，可占领所有据点。image.png
    fmt.Println(minimumCost([]int{3,2,1,4},[][]int{{0,2},{2,3},{3,1}})) // 2
}