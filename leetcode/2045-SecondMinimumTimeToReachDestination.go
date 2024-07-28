package main

// 2045. Second Minimum Time to Reach Destination
// A city is represented as a bi-directional connected graph with n vertices where each vertex is labeled from 1 to n (inclusive). 
// The edges in the graph are represented as a 2D integer array edges, where each edges[i] = [ui, vi] denotes a bi-directional edge between vertex ui and vertex vi. 
// Every vertex pair is connected by at most one edge, and no vertex has an edge to itself. The time taken to traverse any edge is time minutes.

// Each vertex has a traffic signal which changes its color from green to red and vice versa every change minutes. 
// All signals change at the same time. You can enter a vertex at any time, but can leave a vertex only when the signal is green. 
// You cannot wait at a vertex if the signal is green.

// The second minimum value is defined as the smallest value strictly larger than the minimum value.
//     For example the second minimum value of [2, 3, 4] is 3, and the second minimum value of [2, 2, 4] is 4.

// Given n, edges, time, and change, return the second minimum time it will take to go from vertex 1 to vertex n.

// Notes:
//     You can go through any vertex any number of times, including 1 and n.
//     You can assume that when the journey starts, all signals have just turned green.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/09/29/e1.png" />
// <img src="https://assets.leetcode.com/uploads/2021/09/29/e2.png" />
// Input: n = 5, edges = [[1,2],[1,3],[1,4],[3,4],[4,5]], time = 3, change = 5
// Output: 13
// Explanation:
// The figure on the left shows the given graph.
// The blue path in the figure on the right is the minimum time path.
// The time taken is:
// - Start at 1, time elapsed=0
// - 1 -> 4: 3 minutes, time elapsed=3
// - 4 -> 5: 3 minutes, time elapsed=6
// Hence the minimum time needed is 6 minutes.
// The red path shows the path to get the second minimum time.
// - Start at 1, time elapsed=0
// - 1 -> 3: 3 minutes, time elapsed=3
// - 3 -> 4: 3 minutes, time elapsed=6
// - Wait at 4 for 4 minutes, time elapsed=10
// - 4 -> 5: 3 minutes, time elapsed=13
// Hence the second minimum time is 13 minutes.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/09/29/eg2.png" />
// Input: n = 2, edges = [[1,2]], time = 3, change = 2
// Output: 11
// Explanation:
// The minimum time path is 1 -> 2 with time = 3 minutes.
// The second minimum time path is 1 -> 2 -> 1 -> 2 with time = 11 minutes.
 
// Constraints:
//     2 <= n <= 10^4
//     n - 1 <= edges.length <= min(2 * 10^4, n * (n - 1) / 2)
//     edges[i].length == 2
//     1 <= ui, vi <= n
//     ui != vi
//     There are no duplicate edges.
//     Each vertex can be reached directly or indirectly from every other vertex.
//     1 <= time, change <= 10^3

import "fmt"

func secondMinimum(n int, edges [][]int, time int, change int) int {
    // 首先新建一个graph二维切片，每一行记录了第i个节点所连接的节点序号，以便后续bfs
    graph, inf := make([][]int, n+1), 1 << 32 - 1 // 考虑到原题节点下标从1开始，方便起见，分配多一行，第0行不作使用
    for _, e := range edges { // 遍历每一个边
        x, y := e[0], e[1] // 边的两个端点
        graph[x] = append(graph[x], y)
        graph[y] = append(graph[y], x)
    }
    // 求所需时间，实际上是求所需步数
    // dist为n+1行2列切片，第i行第0列表示从1到i的最少步数，第1列表示次少步数
    // 而dist[n][1]就是题目需要的从1到n的次少步数
    dist := make([][2]int, n+1)
    // 除了dist[1][0]为0，其他值默认为正无穷，表示从1到i的所需步数未知
    dist[1][1] = inf
    for i := 2; i <= n; i++ {
        dist[i] = [2]int{inf, inf}
    }
    // 定义一个bfs遍历的节点结构体
    type pair struct {
        x int // 当前遍历所处节点
        d int // 从1步行至当前节点的步数
    }
    queue := make([]pair, 0)          // 定义一个队列，用于bfs
    queue = append(queue, pair{1, 0}) // 初始节点为1，节点1途径自己后抵达自己所需步数为0
    // 开始bfs。在dist[n][1]，也就是1到n的次短步数还没求出来前，一直遍历
    for dist[n][1] == inf {
        node := queue[0]
        queue = queue[1:] // 队首出列
        // 考察当前节点的邻接节点，更新节点1至各邻接节点的所需步数
        for _, y := range graph[node.x] { // y是node.x的其中一个邻接节点
            d := node.d + 1     // d是节点1途径节点x抵达y的步数
            if d < dist[y][0] { // 如果d比此前已记录的1到y的最少步数还要低，则更新，并将y加入队列
                dist[y][0] = d
                queue = append(queue, pair{y, d})
            } else if d > dist[y][0] && d < dist[y][1] {
                // 如果d并非1到y的最少步数，但是次少，同样更新
                // 这里要留意，必须有d>dist[y][0]的判断，防止d==dist[y][0]的情况却依旧去覆盖次少值
                dist[y][1] = d
                queue = append(queue, pair{y, d})
            }
        }
        // 在这一层循环中，如果有y没有入队，说明节点1途径x到达y的这条路比之前记录都要远
    } // 循环退出，说明节点1到n的次少步数已经获知
    // 已知从节点1到n的次少步数，下面模拟实际的用时，把等红灯的时间考虑进去
    t := 0 // 总时长
    for i := 0; i < dist[n][1]; i++ {
        if (t / change) % 2 == 1 {
            // t表示已经走了的时长
            // 如果t除以红灯时长change的商是奇数
            // 说明需要等红灯
            // 等灯的时长为change-(t%change)
            t += change - (t % change)
        }
        t += time // 等完红灯或者处在绿灯的情况下，走一步，花费time时间
    }
    return t
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/09/29/e1.png" />
    // <img src="https://assets.leetcode.com/uploads/2021/09/29/e2.png" />
    // Input: n = 5, edges = [[1,2],[1,3],[1,4],[3,4],[4,5]], time = 3, change = 5
    // Output: 13
    // Explanation:
    // The figure on the left shows the given graph.
    // The blue path in the figure on the right is the minimum time path.
    // The time taken is:
    // - Start at 1, time elapsed=0
    // - 1 -> 4: 3 minutes, time elapsed=3
    // - 4 -> 5: 3 minutes, time elapsed=6
    // Hence the minimum time needed is 6 minutes.
    // The red path shows the path to get the second minimum time.
    // - Start at 1, time elapsed=0
    // - 1 -> 3: 3 minutes, time elapsed=3
    // - 3 -> 4: 3 minutes, time elapsed=6
    // - Wait at 4 for 4 minutes, time elapsed=10
    // - 4 -> 5: 3 minutes, time elapsed=13
    // Hence the second minimum time is 13 minutes.
    fmt.Println(secondMinimum(5,[][]int{{1,2},{1,3},{1,4},{3,4},{4,5}},3,5)) // 13
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/09/29/eg2.png" />
    // Input: n = 2, edges = [[1,2]], time = 3, change = 2
    // Output: 11
    // Explanation:
    // The minimum time path is 1 -> 2 with time = 3 minutes.
    // The second minimum time path is 1 -> 2 -> 1 -> 2 with time = 11 minutes.
    fmt.Println(secondMinimum(2,[][]int{{1,2}},3,2)) // 11
}