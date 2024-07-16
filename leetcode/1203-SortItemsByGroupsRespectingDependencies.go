package main

// 1203. Sort Items by Groups Respecting Dependencies
// There are n items each belonging to zero or one of m groups 
// where group[i] is the group that the i-th item belongs to and it's equal to -1 if the i-th item belongs to no group. 
// The items and the groups are zero indexed. A group can have no item belonging to it.

// Return a sorted list of the items such that:
//     The items that belong to the same group are next to each other in the sorted list.
//     There are some relations between these items where beforeItems[i] is a list containing all the items that should come before the i-th item in the sorted array (to the left of the i-th item).
//     Return any solution if there is more than one solution and return an empty list if there is no solution.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/09/11/1359_ex1.png" />
// Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3,6],[],[],[]]
// Output: [6,3,4,1,5,2,0,7]

// Example 2:
// Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3],[],[4],[]]
// Output: []
// Explanation: This is the same as example 1 except that 4 needs to be before 6 in the sorted list.

// Constraints:
//     1 <= m <= n <= 3 * 10^4
//     group.length == beforeItems.length == n
//     -1 <= group[i] <= m - 1
//     0 <= beforeItems[i].length <= n - 1
//     0 <= beforeItems[i][j] <= n - 1
//     i != beforeItems[i][j]
//     beforeItems[i] does not contain duplicates elements.

import "fmt"

// topo sort
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
    for i := 0; i < len(group); i++ {
        if group[i] == -1 {
            group[i] = m
            m++
        }
    }      
    gc := make([]int, m)
    for i := 0; i < len(group); i++ {
        gc[group[i]]++
    }
    d1, d2 := make([]int, m), make([]int, n) // group indegree, item indegree
    g1, g2 := make([]map[int]struct{}, m), make([]map[int]struct{}, n)
    for i := range g1 { g1[i] = map[int]struct{}{} }
    for i := range g2 { g2[i] = map[int]struct{}{} }
    for i := 0; i < len(beforeItems); i++ {
        for j := 0; j < len(beforeItems[i]); j++ {
            d2[i]++
            g2[beforeItems[i][j]][i] = struct{}{}
            gi := group[i]
            gj := group[beforeItems[i][j]]
            if  _, ok := g1[gj][gi]; !ok && gi != gj {
                d1[gi]++
                g1[gj][gi] = struct{}{}
            }
        }
    }
    queue := []int{}
    for i, count := range d1 {
        if count == 0 {
            queue = append(queue, i)
        }
    }
    res, count, pos := make([]int, n), 0, make([]int, m)
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        pos[node] = count
        count += gc[node]
        for j, _ := range g1[node] {
            d1[j]--
            if d1[j] == 0 {
                queue = append(queue, j)
            }
        }
    }
    if count != n {
        return nil
    }
    queue = []int{}
    for i, count := range d2 {
        if count == 0 {
            queue = append(queue, i)
        }
    }
    count = 0 
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        g := group[node]
        res[pos[g]] = node
        pos[g]++
        count++
        for j, _ := range g2[node] {
            d2[j]--
            if d2[j] == 0 {
                queue = append(queue, j)
            }
        }
    }
    if count != n {
        return nil
    }
    return res
}


// 有 n 个项目，每个项目或者不属于任何小组，或者属于 m 个小组之一。group[i] 表示第 i 个项目所属的小组，如果第 i 个项目不属于任何小组，则 group[i] 等于 -1。
// 项目和小组都是从零开始编号的。可能存在小组不负责任何项目，即没有任何项目属于这个小组。
// 请你帮忙按要求安排这些项目的进度，并返回排序后的项目列表：
// 同一小组的项目，排序后在列表中彼此相邻。
// 项目之间存在一定的依赖关系，我们用一个列表 beforeItems 来表示，其中 beforeItems[i] 表示在进行第 i 个项目前（位于第 i 个项目左侧）应该完成的所有项目。
// 如果存在多个解决方案，只需要返回其中任意一个即可。如果没有合适的解决方案，就请返回一个 空列表
func sortItems1(n int, m int, group []int, beforeItems [][]int) []int {
    // 拓扑排序(双重)
    // 因为同一组的项目要呆在一起,造成组与组之间也有依赖
    // trick: 给没有组的项目设置一个组编号,使得其可以参与组的拓扑排序
    topoSort := func(objs []int, degree []int, g [][]int) []int { // k个items进行排序,返回排序后的items
        orders, k := []int{}, len(objs)
        queue := make([]int, 0, k)
        for _, id := range objs {
            if degree[id] == 0 {
                queue = append(queue, id)
            }
        }
        cur := 0
        for len(queue) > 0 {
            cur, queue = queue[0], queue[1:]
            orders = append(orders, cur) // 出的时机收集(如果是入的时机收集,有两处)
            for _, nx := range g[cur] {
                degree[nx]--
                if degree[nx] == 0 {
                    queue = append(queue, nx)
                }
            }
        }
        return orders
    }
    // 预处理,让没有组的项目设置一个组
    groupItems := make([][]int, m+n) // 一个group有哪些item
    for id, gID := range group {
        if gID == -1 {
            gID = m
            group[id] = m
            m++
        }
        groupItems[gID] = append(groupItems[gID], id) // 组的范围[0,m)
    }
    groupItems = groupItems[:m]
    // 建立组和项目的依赖
    // trick!!!!, 如果是跨组的项目依赖,只建立组之间的依赖(如果前置组的依赖全部满足,那么就是前置组项目的依赖全部满足),
    // 这样的好处是项目只需考虑组内的依赖,因为最终需要一组组的分批处理, 当某个项目前置满足时,并不是处理的时机,而是要等到处理此组数据时才处理),还有一个好处就是可以共用一份topoSort代码
    groupGraph, groupDegree := make([][]int, m), make([]int, m)
    itemGraph, itemDegree := make([][]int, n), make([]int, n)
    for id, items := range beforeItems {
        gID := group[id]
        for _, preID := range items {
            gPreID := group[preID]
            if gID != gPreID { // 跨组依赖,只需建立组之间的依赖
                groupGraph[gPreID] = append(groupGraph[gPreID], gID)
                groupDegree[gID]++
            } else { // 同组依赖
                itemGraph[preID] = append(itemGraph[preID], id)
                itemDegree[id]++
            }
        }
    }
    // 先处理组中依赖,如果可以全部满足,那跨组的项目依赖也全部满足
    groupIDs := make([]int, m)
    for i := 0; i < m; i++ {
        groupIDs[i] = i
    }
    groupOrders := topoSort(groupIDs, groupDegree, groupGraph)
    if len(groupOrders) != m { // 组中有依赖循环
        return []int{}
    }
    res := make([]int, 0, n)
    for _, gID := range groupOrders {
        items := groupItems[gID]
        // 一组组分批处理,好让同组项目放在一起
        itemOrdersInOneGroup := topoSort(items, itemDegree, itemGraph)
        if len(itemOrdersInOneGroup) != len(items) { // 注意!! 同组依赖也可能有循环
            return []int{}
        }
        res = append(res, itemOrdersInOneGroup...)
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/09/11/1359_ex1.png" />
    // Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3,6],[],[],[]]
    // Output: [6,3,4,1,5,2,0,7]
    fmt.Println(sortItems(8,2,[]int{-1,-1,1,0,0,1,0,-1},[][]int{{},{6},{5},{6},{3,6},{},{},{}})) // [6,3,4,1,5,2,0,7]
    // Example 2:
    // Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3],[],[4],[]]
    // Output: []
    // Explanation: This is the same as example 1 except that 4 needs to be before 6 in the sorted list.
    fmt.Println(sortItems(8,2,[]int{-1,-1,1,0,0,1,0,-1},[][]int{{},{6},{5},{6},{3},{},{4},{}})) // []

    fmt.Println(sortItems1(8,2,[]int{-1,-1,1,0,0,1,0,-1},[][]int{{},{6},{5},{6},{3,6},{},{},{}})) // [6,3,4,1,5,2,0,7]
    fmt.Println(sortItems1(8,2,[]int{-1,-1,1,0,0,1,0,-1},[][]int{{},{6},{5},{6},{3},{},{4},{}})) // []
}