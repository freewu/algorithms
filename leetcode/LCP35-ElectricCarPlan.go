package main

// LCP 35. 电动车游城市
// 小明的电动车电量充满时可行驶距离为 cnt，每行驶 1 单位距离消耗 1 单位电量，且花费 1 单位时间。
// 小明想选择电动车作为代步工具。地图上共有 N 个景点，景点编号为 0 ~ N-1。
// 他将地图信息以 [城市 A 编号,城市 B 编号,两城市间距离] 格式整理在在二维数组 paths，表示城市 A、B 间存在双向通路。
// 初始状态，电动车电量为 0。每个城市都设有充电桩，charge[i] 表示第 i 个城市每充 1 单位电量需要花费的单位时间。
// 请返回小明最少需要花费多少单位时间从起点城市 start 抵达终点城市 end。

// 示例 1：
// 输入：paths = [[1,3,3],[3,2,1],[2,1,3],[0,1,4],[3,0,5]], cnt = 6, start = 1, end = 0, charge = [2,10,4,1]
// 输出：43
// 解释：最佳路线为：1->3->0。 在城市 1 仅充 3 单位电至城市 3，然后在城市 3 充 5 单位电，行驶至城市 0。 
//       充电用时共 3*10 + 5*1= 35 行驶用时 3 + 5 = 8，此时总用时最短 43。
//       <img src="https://pic.leetcode-cn.com/1616125304-mzVxIV-image.png" />

// 示例 2：
// 输入：paths = [[0,4,2],[4,3,5],[3,0,5],[0,1,5],[3,2,4],[1,2,8]], cnt = 8, start = 0, end = 2, charge = [4,1,1,3,2]
// 输出：38
// 解释：最佳路线为：0->4->3->2。 城市 0 充电 2 单位，行驶至城市 4 充电 8 单位，行驶至城市 3 充电 1 单位，最终行驶至城市 2。 
//       充电用时 4*2+2*8+3*1 = 27 行驶用时 2+5+4 = 11，总用时最短 38。

// 提示：
//     1 <= paths.length <= 200
//     paths[i].length == 3
//     2 <= charge.length == n <= 100
//     0 <= path[i][0],path[i][1],start,end < n
//     1 <= cnt <= 100
//     1 <= path[i][2] <= cnt
//     1 <= charge[i] <= 100
//     题目保证所有城市相互可以到达

import "fmt"
import "container/heap"

// type MinHeap [][3]int
// func (h *MinHeap) Len() int           { return len(*h) }
// func (h *MinHeap) Less(i, j int) bool { return (*h)[i][1] < (*h)[j][1] }
// func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
// func (h *MinHeap) Push(x any)         { *h = append(*h, x.([3]int)) }
// func (h *MinHeap) Pop() any {
//     x := (*h)[h.Len()-1]
//     *h = (*h)[:h.Len()-1]
//     return x
// }

// func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
//     //cnt 电动车最大电量 初始电量为0
//     //start end 起点和终点
//     //charge 单位电量充电时间 长度为城市数量
//     //paths[i][j]:城市i到城市j的距离也即行驶用时
//     //堆里放的东西：点 起点到该点的距离/行驶用时
//     type Tuple struct {
//         location int
//         time     int
//         curCnt   int
//     }
//     n := len(charge)
//     // 建图 无向图
//     graph := make([][][2]int, n)
//     for _, path := range paths {
//         u, v, w := path[0], path[1], path[2]
//         graph[u] = append(graph[u], [2]int{v, w})
//         graph[v] = append(graph[v], [2]int{u, w})
//     }
//     //从起点到终点的最短用时初始化
//     cost := make([][]int, n)
//     for i := 0; i < n; i++ {
//         cost[i] = make([]int, cnt + 1)
//     }
//     for i := 0; i < n; i++ {
//         for j := 0; j <= cnt; j++ {
//             cost[i][j] = 1 << 31
//         }
//     }
//     cost[start][0] = 0
//     // 所在位置+用时+目前电量表示一个状态 地图上的一个广义点
//     visited := make(map[Tuple]bool)
//     hp := &MinHeap{}
//     heap.Push(hp, [3]int{start, 0, 0})
//     for hp.Len() > 0 {
//         x := heap.Pop(hp).([3]int)
//         curLocation, curTime, curCnt := x[0], x[1], x[2]
//         if curLocation == end { // 到达终点 返回结果
//             return curTime
//         }
//         if !visited[Tuple{curLocation, curTime, curCnt}] {
//             visited[Tuple{curLocation, curTime, curCnt}] = true
//             //可以充电
//             if curCnt < cnt {
//                 chargedTime := curTime + charge[curLocation]
//                 chargedCnt := curCnt + 1
//                 //如果在这个城市、当前时间、当前电量的状态之前没有访问过
//                 if !visited[Tuple{curLocation, chargedTime, chargedCnt}] {
//                     if cost[curLocation][chargedCnt] > chargedTime {
//                         cost[curLocation][chargedCnt] = chargedTime
//                         heap.Push(hp, [3]int{curLocation, chargedTime, chargedCnt})
//                     }
//                 }
//             }
//             // 不充电
//             for _, to := range graph[curLocation] {
//                 nextLocation := to[0]
//                 timeCost := to[1]
//                 // 电量足够的情况下才能去下一个城市
//                 if curCnt >= timeCost {
//                     arriveTime := curTime + timeCost
//                     arriveCnt := curCnt - timeCost
//                     if !visited[Tuple{nextLocation, arriveTime, arriveCnt}] {
//                         if cost[nextLocation][arriveCnt] > arriveTime {
//                             cost[nextLocation][arriveCnt] = arriveTime
//                             heap.Push(hp, [3]int{nextLocation, arriveTime, arriveCnt})
//                         }
//                     }
//                 }
//             }
//         }
//     }
//     return -1
// }

type Pair struct {
    to, weight, cnt int
}
type MinHeap []Pair
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Pair)) }
func (h *MinHeap) Pop() any          {
    old := *h 
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
    n := len(charge)
    g := make([][][2]int, n)
    for _, v := range paths {
        a, b, w := v[0], v[1], v[2]
        g[a] = append(g[a], [2]int{b, w})
        g[b] = append(g[b], [2]int{a, w})
    }
    dis := make([][]int, n)
    for i := range dis {
        dis[i] = make([]int, cnt + 1)
        for j := range dis[i] {
            dis[i][j] = 1 << 31
        }
    }
    h := &MinHeap{{start, 0, 0}}
    heap.Init(h)
    for h.Len() > 0 {
        cur := heap.Pop(h).(Pair)
        u, w, t := cur.to, cur.weight, cur.cnt
        if w > dis[u][t] {
            continue
        } else if u == end {
            return w
        }
        if t < cnt {
            if w + charge[u] < dis[u][t+1] {
                dis[u][t+1] = w + charge[u]
                heap.Push(h, Pair{u, dis[u][t+1], t + 1})
            }
        }
        for _, v := range g[u] {
            if t < v[1] {
                continue
            }
            cur_w := w + v[1]
            if cur_w < dis[v[0]][t-v[1]] {
                dis[v[0]][t-v[1]] = cur_w
                heap.Push(h, Pair{v[0], cur_w, t - v[1]})
            }
        }
    }
    return -1
}

func main() {
    // 示例 1：
    // 输入：paths = [[1,3,3],[3,2,1],[2,1,3],[0,1,4],[3,0,5]], cnt = 6, start = 1, end = 0, charge = [2,10,4,1]
    // 输出：43
    // 解释：最佳路线为：1->3->0。 在城市 1 仅充 3 单位电至城市 3，然后在城市 3 充 5 单位电，行驶至城市 0。 
    //       充电用时共 3*10 + 5*1= 35 行驶用时 3 + 5 = 8，此时总用时最短 43。
    //       <img src="https://pic.leetcode-cn.com/1616125304-mzVxIV-image.png" />
    fmt.Println(electricCarPlan([][]int{{1,3,3},{3,2,1},{2,1,3},{0,1,4},{3,0,5}}, 6, 1, 0, []int{2,10,4,1})) // 43
    // 示例 2：
    // 输入：paths = [[0,4,2],[4,3,5],[3,0,5],[0,1,5],[3,2,4],[1,2,8]], cnt = 8, start = 0, end = 2, charge = [4,1,1,3,2]
    // 输出：38
    // 解释：最佳路线为：0->4->3->2。 城市 0 充电 2 单位，行驶至城市 4 充电 8 单位，行驶至城市 3 充电 1 单位，最终行驶至城市 2。 
    //       充电用时 4*2+2*8+3*1 = 27 行驶用时 2+5+4 = 11，总用时最短 38。
    fmt.Println(electricCarPlan([][]int{{0,4,2},{4,3,5},{3,0,5},{0,1,5},{3,2,4},{1,2,8}}, 8, 0, 2, []int{4,1,1,3,2})) // 38
}