package main

// 2092. Find All People With Secret
// You are given an integer n indicating there are n people numbered from 0 to n - 1. 
// You are also given a 0-indexed 2D integer array meetings where meetings[i] = [xi, yi, timei] indicates that person xi and person yi have a meeting at timei. 
// A person may attend multiple meetings at the same time. Finally, you are given an integer firstPerson.

// Person 0 has a secret and initially shares the secret with a person firstPerson at time 0. 
// This secret is then shared every time a meeting takes place with a person that has the secret. 
// More formally, for every meeting, if a person xi has the secret at timei, then they will share the secret with person yi, and vice versa.

// The secrets are shared instantaneously. That is, a person may receive the secret and share it with people in other meetings within the same time frame.
// Return a list of all the people that have the secret after all the meetings have taken place. You may return the answer in any order.

// Example 1:
// Input: n = 6, meetings = [[1,2,5],[2,3,8],[1,5,10]], firstPerson = 1
// Output: [0,1,2,3,5]
// Explanation:
// At time 0, person 0 shares the secret with person 1.
// At time 5, person 1 shares the secret with person 2.
// At time 8, person 2 shares the secret with person 3.
// At time 10, person 1 shares the secret with person 5.​​​​
// Thus, people 0, 1, 2, 3, and 5 know the secret after all the meetings.

// Example 2:
// Input: n = 4, meetings = [[3,1,3],[1,2,2],[0,3,3]], firstPerson = 3
// Output: [0,1,3]
// Explanation:
// At time 0, person 0 shares the secret with person 3.
// At time 2, neither person 1 nor person 2 know the secret.
// At time 3, person 3 shares the secret with person 0 and person 1.
// Thus, people 0, 1, and 3 know the secret after all the meetings.

// Example 3:
// Input: n = 5, meetings = [[3,4,2],[1,2,1],[2,3,1]], firstPerson = 1
// Output: [0,1,2,3,4]
// Explanation:
// At time 0, person 0 shares the secret with person 1.
// At time 1, person 1 shares the secret with person 2, and person 2 shares the secret with person 3.
// Note that person 2 can share the secret at the same time as receiving it.
// At time 2, person 3 shares the secret with person 4.
// Thus, people 0, 1, 2, 3, and 4 know the secret after all the meetings.

// Constraints:
//         2 <= n <= 10^5
//         1 <= meetings.length <= 10^5
//         meetings[i].length == 3
//         0 <= xi, yi <= n - 1
//         xi != yi
//         1 <= timei <= 10^5
//         1 <= firstPerson <= n - 1

// meetings[i] = [xi, yi, timei] 表示专家 xi 和专家 yi 在时间 timei

import "fmt"
import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
    // 按照时间排序
    sort.Slice(meetings, func(i, j int) bool { 
        return meetings[i][2] < meetings[j][2] 
    }) 
    res := []int{}
    haveSecret := map[int]bool{ 0: true, firstPerson: true} // 一开始 0 和 firstPerson 都知道秘密
    for i, m := 0, len(meetings); i < m; {
        g := map[int][]int{}
        time := meetings[i][2]
        // 遍历时间相同的会议。注意这里的 i 和外层循环的 i 是同一个变量，所以整个循环部分的时间复杂度是线性的
        for ; i < m && meetings[i][2] == time; i++ {
            v, w := meetings[i][0], meetings[i][1]
            g[v] = append(g[v], w) // 建图
            g[w] = append(g[w], v)
        }

        vis := map[int]bool{} // 避免重复访问专家
        var dfs func(int)
        dfs = func(v int) {
            vis[v] = true
            haveSecret[v] = true
            for _, w := range g[v] {
                if !vis[w] {
                    dfs(w)
                }
            }
        }
        for v := range g {
            if haveSecret[v] && !vis[v] { // 从在图上且知道秘密的专家出发，DFS 标记所有能到达的专家
                dfs(v)
            }
        }
    }
    for i := range haveSecret {
        res = append(res, i) // 注意可以按任何顺序返回答案
    }
    return res
}

// 并查集
func findAllPeople1(n int, meetings [][]int, firstPerson int) []int {
    const  MX = 10001
    // 集合的标签信息: 设置集合的一些属性
    // 属性在哪？knew[代表元素] 代表集合的属性
    father, knew := make([]int,MX), make([]bool,MX) 
    sort.Slice(meetings, func(a, b int) bool {
        return meetings[a][2] < meetings[b][2]
    })
    build := func(n, firstPerson int) {
        for i := 0; i < n; i++ {
            father[i] = i
            knew[i] = false
        }
        // 0 知道秘密，并且告诉了 firstPerson
        knew[0] = true
        father[firstPerson] = 0
    }
    var find func(i int) int
    find = func(i int) int {
        if i != father[i] {
            father[i] = find(father[i])
        }
        return father[i]
    }
    union := func(a, b int) {
        af, bf := find(a), find(b)
        if af == bf { return }
        father[af] = bf
        if knew[af] {
            knew[bf] = true
        }
    }
    m := len(meetings)
    build(n, firstPerson)
    for l, r := 0, 0; l < m; {
        r = l
        for r+1 < m && meetings[l][2] == meetings[r+1][2] {
            r++
        }
        // l...r 这些会议，是一个时刻
        for i := l; i <= r; i++ {
            union(meetings[i][0], meetings[i][1])
        }
        // 有小的撤销行为，但这不是可撤销并查集
        // 只是每一批没有知道秘密的专家重新建立集合而已
        for i := l; i <= r; i++ {
            // 不知道密码就解散，知道密码也可以解散
            // 但是知道密码解不解散不影响答案
            // 为了不运行无用逻辑，就不解散了
            a := meetings[i][0]
            b := meetings[i][1]
            if !knew[find(a)] {
                father[a] = a
            }
            if !knew[find(b)] {
                father[b] = b
            }
        }
        l = r + 1
    }
    res := []int{}
    for i := 0; i < n; i++ {
        if knew[find(i)] {
            res = append(res, i)
        }
    }
    return res
}

func findAllPeople2(n int, meetings [][]int, firstPerson int) []int {
    var row []int
    arr := make([]bool, n)
    arr[0] = true
    arr[firstPerson] = true
    sort.Slice(meetings, func(a, b int) bool {
        return meetings[a][2] < meetings[b][2]
    })
    f := make([]int, n)
    var find func(int) int
    find = func(i int) int {
        if f[i] != i {
            f[i] = find(f[i])
        }
        return f[i]
    }
    var temp [][]int
    var times []int = make([]int, n)
    for len(meetings) > 0 {
        time := meetings[0][2]
        x := 0
        for x < len(meetings) && meetings[x][2] == time {
            x++
        }
        temp, meetings = meetings[:x], meetings[x:]
        for _, t := range temp {
            if times[t[0]] < time {
                row = append(row, t[0])
                times[t[0]] = time
            }
            if times[t[1]] < time {
                row = append(row, t[1])
                times[t[1]] = time
            }
        }
        for _, t := range row {
            f[t] = t
        }
        for _, t := range temp {
            f[find(t[0])] = find(t[1])
        }
        for _, t := range row {
            if arr[t] {
                f[find(t)] = t
                f[t] = t
            }
        }
        for _, t := range row {
            arr[t] = arr[find(t)]
        }
        row = row[:0]
    }
    pos := 0
    for i := range n {
        if arr[i] {
            f[pos] = i
            pos++
        }
    }
    return f[:pos]
}

func main() {
    // Example 1:
    // Input: n = 6, meetings = [[1,2,5],[2,3,8],[1,5,10]], firstPerson = 1
    // Output: [0,1,2,3,5]
    // Explanation:
    // At time 0, person 0 shares the secret with person 1.
    // At time 5, person 1 shares the secret with person 2.
    // At time 8, person 2 shares the secret with person 3.
    // At time 10, person 1 shares the secret with person 5.​​​​
    // Thus, people 0, 1, 2, 3, and 5 know the secret after all the meetings.
    fmt.Println(findAllPeople(6, [][]int{[]int{1,2,5},[]int{2,3,8},[]int{1,5,10}}, 1)) // [0,1,2,3,5]

    // Example 2:
    // Input: n = 4, meetings = [[3,1,3],[1,2,2],[0,3,3]], firstPerson = 3
    // Output: [0,1,3]
    // Explanation:
    // At time 0, person 0 shares the secret with person 3.
    // At time 2, neither person 1 nor person 2 know the secret.
    // At time 3, person 3 shares the secret with person 0 and person 1.
    // Thus, people 0, 1, and 3 know the secret after all the meetings.
    fmt.Println(findAllPeople( 4,  [][]int{[]int{3,1,3},[]int{1,2,2},[]int{0,3,3}},  3)) // [0,1,3]

    // Example 3:
    // Input: n = 5, meetings = [[3,4,2],[1,2,1],[2,3,1]], firstPerson = 1
    // Output: [0,1,2,3,4]
    // Explanation:
    // At time 0, person 0 shares the secret with person 1.
    // At time 1, person 1 shares the secret with person 2, and person 2 shares the secret with person 3.
    // Note that person 2 can share the secret at the same time as receiving it.
    // At time 2, person 3 shares the secret with person 4.
    // Thus, people 0, 1, 2, 3, and 4 know the secret after all the meetings.
    fmt.Println(findAllPeople(5,[][]int{[]int{3,4,2},[]int{1,2,1},[]int{2,3,1}},1,)) // [0,1,2,3,4]

    fmt.Println(findAllPeople1(6, [][]int{[]int{1,2,5},[]int{2,3,8},[]int{1,5,10}}, 1)) // [0,1,2,3,5]
    fmt.Println(findAllPeople1(4, [][]int{[]int{3,1,3},[]int{1,2,2},[]int{0,3,3}},  3)) // [0,1,3]
    fmt.Println(findAllPeople1(5, [][]int{[]int{3,4,2},[]int{1,2,1},[]int{2,3,1}},  1)) // [0,1,2,3,4]

    fmt.Println(findAllPeople2(6, [][]int{[]int{1,2,5},[]int{2,3,8},[]int{1,5,10}}, 1)) // [0,1,2,3,5]
    fmt.Println(findAllPeople2(4, [][]int{[]int{3,1,3},[]int{1,2,2},[]int{0,3,3}},  3)) // [0,1,3]
    fmt.Println(findAllPeople2(5, [][]int{[]int{3,4,2},[]int{1,2,1},[]int{2,3,1}},  1)) // [0,1,2,3,4]
}

