package main

// 1719. Number Of Ways To Reconstruct A Tree
// You are given an array pairs, where pairs[i] = [xi, yi], and:
//     There are no duplicates.
//     xi < yi

// Let ways be the number of rooted trees that satisfy the following conditions:
//     The tree consists of nodes whose values appeared in pairs.
//     A pair [xi, yi] exists in pairs if and only if xi is an ancestor of yi or yi is an ancestor of xi.
//     Note: the tree does not have to be a binary tree.
    
// Two ways are considered to be different if there is at least one node that has different parents in both ways.

// Return:
//     0 if ways == 0
//     1 if ways == 1
//     2 if ways > 1

// A rooted tree is a tree that has a single root node, and all edges are oriented to be outgoing from the root.

// An ancestor of a node is any node on the path from the root to that node (excluding the node itself).
// The root has no ancestors.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/03/trees2.png" />
// Input: pairs = [[1,2],[2,3]]
// Output: 1
// Explanation: There is exactly one valid rooted tree, which is shown in the above figure.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/03/tree.png" />
// Input: pairs = [[1,2],[2,3],[1,3]]
// Output: 2
// Explanation: There are multiple valid rooted trees. Three of them are shown in the above figures.

// Example 3:
// Input: pairs = [[1,2],[2,3],[2,4],[1,5]]
// Output: 0
// Explanation: There are no valid rooted trees.

// Constraints:
//     1 <= pairs.length <= 10^5
//     1 <= xi < yi <= 500
//     The elements in pairs are unique.

import "fmt"

func checkWays(pairs [][]int) int {
    // 递归求图的分支,并检查与其余所有点邻接的结点
    type branch struct {
        vals             []int
        neighbors        map[int][]int
        neighborsCounter map[int]int
    } // 邻接表形式储存分支
    valid := true    //存在树
    hasMult := false //存在多种树

    root := new(branch)
    root.neighbors = map[int][]int{}
    root.neighborsCounter = map[int]int{}
    for _, v := range pairs {
        if root.neighbors[v[0]] == nil {
            root.vals = append(root.vals, v[0])
        }
        root.neighbors[v[0]] = append(root.neighbors[v[0]], v[1])
        root.neighborsCounter[v[0]]++

        if root.neighbors[v[1]] == nil {
            root.vals = append(root.vals, v[1])
        }
        root.neighbors[v[1]] = append(root.neighbors[v[1]], v[0])
        root.neighborsCounter[v[1]]++
    } // 构造基本分支
    var currentQueue []*branch
    currentQueue = append(currentQueue, root)
    for currentQueue != nil { //当前分支队列不为空
        var nextQueue []*branch
        for _, v := range currentQueue { //对于每个分支
            vlen := len(v.vals) //确定分支元素数目，判定是否存在与所有结点均邻接的结点
            if vlen > 1 {       //分支仅有一个结点时为边界条件
                multCounter := 0              //确定与所有结点均邻接的结点有几个，不存在即不存在树结构，存在多个则存在多种组织方式
                for _, node := range v.vals { //遍历分支结点
                    if v.neighborsCounter[node] >= vlen-1 { //如果是全邻接结点则抛弃，并计入全邻接结点计数
                        multCounter++
                        v.neighborsCounter[node] = 0
                    } else if v.neighborsCounter[node] > 0 { //抛弃已经在之前dfs中检查过的结点(从属于之前已建立的分支)
                        var dfsArray []int
                        newBranch := new(branch)
                        newBranch.neighbors = map[int][]int{}
                        newBranch.neighborsCounter = map[int]int{}
                        newBranch.vals = append(newBranch.vals, node) //由于每次dfs开始时，总能保证是一个新分支，所以新建分支
                        nextQueue = append(nextQueue, newBranch)
                        for _, neighEle := range v.neighbors[node] { //遍历结点的邻接结点
                            if v.neighborsCounter[neighEle] > 0 && v.neighborsCounter[neighEle] < vlen-1 { //未被遍历过且不为全邻接结点
                                dfsArray = append(dfsArray, neighEle)          //加入下一级dfs(其实应该是bfs)，脑抽了
                                if newBranch.neighborsCounter[neighEle] <= 0 { //如果结点没有被加入过该分支的结点队列
                                    newBranch.vals = append(newBranch.vals, neighEle)
                                }
                                newBranch.neighbors[neighEle] = append(newBranch.neighbors[neighEle], node) //加入边
                                newBranch.neighborsCounter[neighEle]++
                                newBranch.neighbors[node] = append(newBranch.neighbors[node], neighEle) //加入边
                                newBranch.neighborsCounter[node]++
                            }
                        }
                        v.neighborsCounter[node] = 0
                        for dfsArray != nil { //bfs队列不为空，即该分支还有点未遍历过
                            var nextDfsArray []int
                            for _, nv := range dfsArray { //检查bfs队列中的结点
                                if v.neighborsCounter[nv] > 0 { //避免出现先出队的结点与同一队列中的后出队的结点
                                    for _, dfsEle := range v.neighbors[nv] { //检查bfs队列中的结点的邻接结点
                                        if v.neighborsCounter[dfsEle] > 0 && v.neighborsCounter[dfsEle] < vlen-1 { //不为全邻接结点且未被检查过
                                            nextDfsArray = append(nextDfsArray, dfsEle)  //加入下一级bfs队列
                                            if newBranch.neighborsCounter[dfsEle] <= 0 { //如果未被加入分支结点队列
                                                newBranch.vals = append(newBranch.vals, dfsEle)
                                            }
                                            newBranch.neighbors[dfsEle] = append(newBranch.neighbors[dfsEle], nv) //加入边
                                            newBranch.neighborsCounter[dfsEle]++
                                            newBranch.neighbors[nv] = append(newBranch.neighbors[nv], dfsEle) //加入边
                                            newBranch.neighborsCounter[nv]++

                                        }
                                    }
                                    v.neighborsCounter[nv] = 0 //检查过的结点抛弃
                                }
                            }
                            dfsArray = nextDfsArray //更新bfs队列
                        }
                    }
                }
                if multCounter > 1 { //若某一分支存在多个全邻接结点，则有多构
                    hasMult = true
                } else if multCounter < 1 { //若某一分支不存在全邻接结点，则无法构成树
                    valid = false
                    return 0
                }

            }
        }
        currentQueue = nextQueue //更新分支队列
    }
    if valid {
        if hasMult {
            return 2
        }
        return 1
    }
    return 0
}

func checkWays1(pairs [][]int) int {
    adj := map[int]map[int]bool{}
    for _, p := range pairs {
        x, y := p[0], p[1]
        if adj[x] == nil { adj[x] = map[int]bool{} }
        adj[x][y] = true
        if adj[y] == nil { adj[y] = map[int]bool{} }
        adj[y][x] = true
    }
    // 检测是否存在根节点
    root := -1
    for node, neighbours := range adj {
        if len(neighbours) == len(adj)-1 {
            root = node
            break
        }
    }
    if root == -1 { return 0 }
    res := 1
    for node, neighbours := range adj {
        if node == root { continue }
        currDegree, parent, parentDegree := len(neighbours), -1, 1 << 31
        // 根据 degree 的大小找到 node 的父节点 parent
        for neighbour := range neighbours {
            if len(adj[neighbour]) < parentDegree && len(adj[neighbour]) >= currDegree {
                parent = neighbour
                parentDegree = len(adj[neighbour])
            }
        }
        if parent == -1 { return 0 }
        // 检测 neighbours 是否为 adj[parent] 的子集
        for neighbour := range neighbours {
            if neighbour != parent && !adj[parent][neighbour] { return 0 }
        }
        if parentDegree == currDegree { res = 2 }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/12/03/trees2.png" />
    // Input: pairs = [[1,2],[2,3]]
    // Output: 1
    // Explanation: There is exactly one valid rooted tree, which is shown in the above figure.
    fmt.Println(checkWays([][]int{{1,2},{2,3}})) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/12/03/tree.png" />
    // Input: pairs = [[1,2],[2,3],[1,3]]
    // Output: 2
    // Explanation: There are multiple valid rooted trees. Three of them are shown in the above figures.
    fmt.Println(checkWays([][]int{{1,2},{2,3},{1,3}})) // 2
    // Example 3:
    // Input: pairs = [[1,2],[2,3],[2,4],[1,5]]
    // Output: 0
    // Explanation: There are no valid rooted trees.
    fmt.Println(checkWays([][]int{{1,2},{2,3},{2,4},{1,5}})) // 0

    fmt.Println(checkWays1([][]int{{1,2},{2,3}})) // 1
    fmt.Println(checkWays1([][]int{{1,2},{2,3},{1,3}})) // 2
    fmt.Println(checkWays1([][]int{{1,2},{2,3},{2,4},{1,5}})) // 0
}