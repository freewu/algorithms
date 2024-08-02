package main

// 1548. The Most Similar Path in a Graph
// We have n cities and m bi-directional roads where roads[i] = [ai, bi] connects city ai with city bi. 
// Each city has a name consisting of exactly three upper-case English letters given in the string array names. 
// Starting at any city x, you can reach any city y where y != x (i.e., the cities and the roads are forming an undirected connected graph).

// You will be given a string array targetPath. 
// You should find a path in the graph of the same length and with the minimum edit distance to targetPath.

// You need to return the order of the nodes in the path with the minimum edit distance. 
// The path should be of the same length of targetPath and should be valid (i.e., there should be a direct road between ans[i] and ans[i + 1]). 
// If there are multiple answers return any one of them.

// The edit distance is defined as follows:
// <img src="https://assets.leetcode.com/uploads/2020/08/08/edit.jpg" />

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/08/08/e1.jpg" />
// Input: n = 5, roads = [[0,2],[0,3],[1,2],[1,3],[1,4],[2,4]], names = ["ATL","PEK","LAX","DXB","HND"], targetPath = ["ATL","DXB","HND","LAX"]
// Output: [0,2,4,2]
// Explanation: [0,2,4,2], [0,3,0,2] and [0,3,1,2] are accepted answers.
// [0,2,4,2] is equivalent to ["ATL","LAX","HND","LAX"] which has edit distance = 1 with targetPath.
// [0,3,0,2] is equivalent to ["ATL","DXB","ATL","LAX"] which has edit distance = 1 with targetPath.
// [0,3,1,2] is equivalent to ["ATL","DXB","PEK","LAX"] which has edit distance = 1 with targetPath.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/08/08/e2.jpg" />
// Input: n = 4, roads = [[1,0],[2,0],[3,0],[2,1],[3,1],[3,2]], names = ["ATL","PEK","LAX","DXB"], targetPath = ["ABC","DEF","GHI","JKL","MNO","PQR","STU","VWX"]
// Output: [0,1,0,1,0,1,0,1]
// Explanation: Any path in this graph has edit distance = 8 with targetPath.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/08/09/e3.jpg" />
// Input: n = 6, roads = [[0,1],[1,2],[2,3],[3,4],[4,5]], names = ["ATL","PEK","LAX","ATL","DXB","HND"], targetPath = ["ATL","DXB","HND","DXB","ATL","LAX","PEK"]
// Output: [3,4,5,4,3,2,1]
// Explanation: [3,4,5,4,3,2,1] is the only path with edit distance = 0 with targetPath.
// It's equivalent to ["ATL","DXB","HND","DXB","ATL","LAX","PEK"]

// Constraints:
//     2 <= n <= 100
//     m == roads.length
//     n - 1 <= m <= (n * (n - 1) / 2)
//     0 <= ai, bi <= n - 1
//     ai != bi
//     The graph is guaranteed to be connected and each pair of nodes may have at most one direct road.
//     names.length == n
//     names[i].length == 3
//     names[i] consists of upper-case English letters.
//     There can be two cities with the same name.
//     1 <= targetPath.length <= 100
//     targetPath[i].length == 3
//     targetPath[i] consists of upper-case English letters.
    
// Follow up: If each node can be visited only once in the path, What should you change in your solution?

import "fmt"

func mostSimilar(n int, roads [][]int, names []string, targetPath []string) []int {
    // targetPath最长是100，每个位置的可能性是100个城市
    // dp[i][j] 表示当targetPath的第i个位置是城市j的时候，的最短编辑距离
    // 最后求的是min(dp[len(targetPath)][j]), 其中j >= 0 && j < len(names)
    // 加上一个pre[i][j]表示前一个城市，记录路径
    getAdjMap := func (roads [][]int) map[int]map[int]bool { // 注意，是双向可达
        am := make(map[int]map[int]bool, 0)
        for _, r := range roads {
            _, e := am[r[0]]
            if !e {
                am[r[0]] = make(map[int]bool)
            }
            am[r[0]][r[1]] = true
    
            _, e1 := am[r[1]]
            if !e1 {
                am[r[1]] = make(map[int]bool)
            }
            am[r[1]][r[0]] = true
        }
        return am
    }

    am := getAdjMap(roads)
    dp, pre, maxNum := make([][]int, 0), make([][]int, 0),101
    for i := 0; i < len(targetPath); i ++ {
        tmp := make([]int, 0)
        tmpPre := make([]int, 0)
        for j := 0; j < len(names); j ++ {
            tmp = append(tmp, maxNum)
            tmpPre = append(tmpPre, -1)
        }
        dp = append(dp, tmp)
        pre = append(pre, tmpPre)
    }
    for j := 0; j < len(names); j ++ {
        if targetPath[0] == names[j] {
            dp[0][j] = 0
        } else {
            dp[0][j] = 1 // 1个编辑距离
        }
    }
    for i := 1; i < len(targetPath); i ++ {
        for j := 0; j < len(names); j ++ {
            tmpCost := maxNum
            for cand, _ := range am[j] {
                curCost := dp[i - 1][cand]
                if targetPath[i] != names[j] { // 如果j不相等，那就加上一个编辑距离
                    curCost += 1
                }
                if curCost < tmpCost {
                    tmpCost = curCost
                    pre[i][j] = cand // 往前追路径
                }
            }
            dp[i][j] = tmpCost
        }
    }
    res, minIndex, minCost := make([]int, 0), -1, maxNum
    for j := 0; j < len(names); j ++ { // 先找到min
        if dp[len(targetPath) - 1][j] < minCost {
            minIndex = j
            minCost = dp[len(targetPath) - 1][j]
        }
    }
    for i := len(targetPath) - 1; i >= 0; i -- { // 这里是逆序的，从头部插入
        res = append([]int{minIndex}, res...)
        minIndex = pre[i][minIndex]
    }
    return res
}


func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/08/08/e1.jpg" />
    // Input: n = 5, roads = [[0,2],[0,3],[1,2],[1,3],[1,4],[2,4]], names = ["ATL","PEK","LAX","DXB","HND"], targetPath = ["ATL","DXB","HND","LAX"]
    // Output: [0,2,4,2]
    // Explanation: [0,2,4,2], [0,3,0,2] and [0,3,1,2] are accepted answers.
    // [0,2,4,2] is equivalent to ["ATL","LAX","HND","LAX"] which has edit distance = 1 with targetPath.
    // [0,3,0,2] is equivalent to ["ATL","DXB","ATL","LAX"] which has edit distance = 1 with targetPath.
    // [0,3,1,2] is equivalent to ["ATL","DXB","PEK","LAX"] which has edit distance = 1 with targetPath.
    fmt.Println(mostSimilar(5, [][]int{{0,2},{0,3},{1,2},{1,3},{1,4},{2,4}},[]string{"ATL","PEK","LAX","DXB","HND"},[]string{"ATL","DXB","HND","LAX"})) // [0,2,4,2]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/08/08/e2.jpg" />
    // Input: n = 4, roads = [[1,0],[2,0],[3,0],[2,1],[3,1],[3,2]], names = ["ATL","PEK","LAX","DXB"], targetPath = ["ABC","DEF","GHI","JKL","MNO","PQR","STU","VWX"]
    // Output: [0,1,0,1,0,1,0,1]
    // Explanation: Any path in this graph has edit distance = 8 with targetPath.
    fmt.Println(mostSimilar(4, [][]int{{1,0},{2,0},{3,0},{2,1},{3,1},{3,2}},[]string{"ATL","PEK","LAX","DXB"},[]string{"ABC","DEF","GHI","JKL","MNO","PQR","STU","VWX"})) // [0,1,0,1,0,1,0,1]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/08/09/e3.jpg" />
    // Input: n = 6, roads = [[0,1],[1,2],[2,3],[3,4],[4,5]], names = ["ATL","PEK","LAX","ATL","DXB","HND"], targetPath = ["ATL","DXB","HND","DXB","ATL","LAX","PEK"]
    // Output: [3,4,5,4,3,2,1]
    // Explanation: [3,4,5,4,3,2,1] is the only path with edit distance = 0 with targetPath.
    // It's equivalent to ["ATL","DXB","HND","DXB","ATL","LAX","PEK"]
    fmt.Println(mostSimilar(6, [][]int{{0,1},{1,2},{2,3},{3,4},{4,5}},[]string{"ATL","PEK","LAX","ATL","DXB","HND"},[]string{"ATL","DXB","HND","DXB","ATL","LAX","PEK"})) // [3,4,5,4,3,2,1]
}