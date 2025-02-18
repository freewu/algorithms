package main

// LCP 51. 烹饪料理
// 欢迎各位勇者来到力扣城，城内设有烹饪锅供勇者制作料理，为自己恢复状态。

// 勇者背包内共有编号为 0 ~ 4 的五种食材，其中 materials[j] 表示第 j 种食材的数量。
// 通过这些食材可以制作若干料理，cookbooks[i][j] 表示制作第 i 种料理需要第 j 种食材的数量，而 attribute[i] = [x,y] 表示第 i 道料理的美味度 x 和饱腹感 y。

// 在饱腹感不小于 limit 的情况下，请返回勇者可获得的最大美味度。
// 如果无法满足饱腹感要求，则返回 -1。

// 注意：
//     每种料理只能制作一次。

// 示例 1：
// 输入：materials = [3,2,4,1,2] cookbooks = [[1,1,0,1,2],[2,1,4,0,0],[3,2,4,1,0]] attribute = [[3,2],[2,4],[7,6]] limit = 5
// 输出：7
// 解释： 食材数量可以满足以下两种方案： 
// 方案一：制作料理 0 和料理 1，可获得饱腹感 2+4、美味度 3+2 
// 方案二：仅制作料理 2， 可饱腹感为 6、美味度为 7 因此在满足饱腹感的要求下，可获得最高美味度 7

// 示例 2：
// 输入：materials = [10,10,10,10,10] cookbooks = [[1,1,1,1,1],[3,3,3,3,3],[10,10,10,10,10]] attribute = [[5,5],[6,6],[10,10]] limit = 1
// 输出：11
// 解释：通过制作料理 0 和 1，可满足饱腹感，并获得最高美味度 11

// 提示：
//     materials.length == 5
//     1 <= cookbooks.length == attribute.length <= 8
//     cookbooks[i].length == 5
//     attribute[i].length == 2
//     0 <= materials[i], cookbooks[i][j], attribute[i][j] <= 20
//     1 <= limit <= 100

import "fmt"

func perfectMenu(materials []int, cookbooks [][]int, attribute [][]int, limit int) int {
    res := -1
    check := func(arr1, arr2 []int)bool{
        for i := range arr1 {
            if arr1[i] < arr2[i] { return false }
        }
        return true
    }
    var dfs func(mat, attr []int,index int)
    dfs = func(mat, attr []int,index int) {
        if !check(materials, mat) { return }
        if attr[1] >= limit && attr[0] > res {
            res = attr[0]
        }
        for i := index; i <len(cookbooks); i++ {
            tmp1, tmp2 := make([]int,len(mat)), make([]int, len(attr))
            copy(tmp1, mat)
            copy(tmp2,attr)
            for j := range cookbooks[i] {
                mat[j] += cookbooks[i][j]
            }
            for k := range attribute[i] {
                attr[k] += attribute[i][k]
            }
            dfs(mat,attr,i + 1)
            mat, attr = tmp1, tmp2
        }
    }
    dfs(make([]int, len(materials)), make([]int, len(attribute[0])), 0)
    return res
}

func perfectMenu1(materials []int, cookbooks [][]int, attribute [][]int, limit int) int {
    res, n := -1, len(cookbooks)
    path := []int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(i int)
    dfs = func(i int) {
        if i == n {
            t := make([]int, len(materials))
            yum, lim := 0, 0
            for _, v := range path {
                for j, x := range cookbooks[v] {
                    t[j] += x
                }
                yum += attribute[v][0]
                lim += attribute[v][1]
            }
            // fmt.Println(lim, yum, ans)
            if lim < limit {
                res = max(res, -1)
                return 
            }
            for j, v := range t {
                if materials[j] < v {
                    res = max(res, -1)
                    return
                } 
            }
            res = max(res, yum)
            return
        }
        path = append(path, i)
        dfs(i + 1)
        path = path[:len(path) - 1]
        dfs(i + 1)
    }
    dfs(0)
    return res
}

func main() {
    // 示例 1：
    // 输入：materials = [3,2,4,1,2] cookbooks = [[1,1,0,1,2],[2,1,4,0,0],[3,2,4,1,0]] attribute = [[3,2],[2,4],[7,6]] limit = 5
    // 输出：7
    // 解释： 食材数量可以满足以下两种方案： 
    // 方案一：制作料理 0 和料理 1，可获得饱腹感 2+4、美味度 3+2 
    // 方案二：仅制作料理 2， 可饱腹感为 6、美味度为 7 因此在满足饱腹感的要求下，可获得最高美味度 7
    fmt.Println(perfectMenu([]int{3,2,4,1,2}, [][]int{{1,1,0,1,2},{2,1,4,0,0},{3,2,4,1,0}}, [][]int{{3,2},{2,4},{7,6}}, 5)) // 7
    // 示例 2：
    // 输入：materials = [10,10,10,10,10] cookbooks = [[1,1,1,1,1],[3,3,3,3,3],[10,10,10,10,10]] attribute = [[5,5],[6,6],[10,10]] limit = 1
    // 输出：11
    // 解释：通过制作料理 0 和 1，可满足饱腹感，并获得最高美味度 11
    fmt.Println(perfectMenu([]int{10,10,10,10,10}, [][]int{{1,1,1,1,1},{3,3,3,3,3},{10,10,10,10,10}}, [][]int{{5,5},{6,6},{10,10}}, 1)) // 11

    fmt.Println(perfectMenu1([]int{3,2,4,1,2}, [][]int{{1,1,0,1,2},{2,1,4,0,0},{3,2,4,1,0}}, [][]int{{3,2},{2,4},{7,6}}, 5)) // 7
    fmt.Println(perfectMenu1([]int{10,10,10,10,10}, [][]int{{1,1,1,1,1},{3,3,3,3,3},{10,10,10,10,10}}, [][]int{{5,5},{6,6},{10,10}}, 1)) // 11
}