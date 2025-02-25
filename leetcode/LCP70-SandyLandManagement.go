package main

// LCP 70. 沙地治理
// 在力扣城的沙漠分会场展示了一种沙柳树，这种沙柳树能够将沙地转化为坚实的绿地。 
// 展示的区域为正三角形，这片区域可以拆分为若干个子区域，每个子区域都是边长为 1 的小三角形，其中第 i 行有 2i - 1 个小三角形。

// 初始情况下，区域中的所有位置都为沙地，你需要指定一些子区域种植沙柳树成为绿地，以达到转化整片区域为绿地的最终目的，规则如下：
//     1. 若两个子区域共用一条边，则视为相邻；

// 如下图所示，(1,1)和(2,2)相邻，(3,2)和(3,3)相邻；(2,2)和(3,3)不相邻，因为它们没有共用边。
//     1. 若至少有两片绿地与同一片沙地相邻，则这片沙地也会转化为绿地
//     2. 转化为绿地的区域会影响其相邻的沙地
//        <img src="https://pic.leetcode-cn.com/1662692397-VlvErS-image.png" />

// 现要将一片边长为 size 的沙地全部转化为绿地，请找到任意一种初始指定 最少 数量子区域种植沙柳的方案，并返回所有初始种植沙柳树的绿地坐标。

// 示例 1：
// 输入：size = 3 
// 输出：[[1,1],[2,1],[2,3],[3,1],[3,5]] 
// 解释：如下图所示，一种方案为： 指定所示的 5 个子区域为绿地。 
//       相邻至少两片绿地的 (2,2)，(3,2) 和 (3,4) 演变为绿地。 
//       相邻两片绿地的 (3,3) 演变为绿地。
//       <img src="https://pic.leetcode-cn.com/1662692503-ncjywh-image.png" />

// 示例 2：
// 输入：size = 2 
// 输出：[[1,1],[2,1],[2,3]] 
// 解释：如下图所示： 指定所示的 3 个子区域为绿地。 
//      相邻三片绿地的 (2,2) 演变为绿地。
//      <img src="https://pic.leetcode-cn.com/1662692507-mgFXRj-image.png" />

// 提示：
//     1 <= size <= 1000

import "fmt"

func sandyLandManagement(size int) [][]int {
    res := [][]int{ []int{1, 1} } // 添加第一个点 [1, 1]
    if size == 1 { return res }
    for j := 1; j <= 2 * size - 1; j += 2 { // 添加最后一行的所有奇数位置的点
        res = append(res, []int{ size, j })
    }
    flag := 2
    for i := size - 1; i > 1; { // 从倒数第二行开始，交替添加点
        res = append(res, []int{ i, flag })
        i--
        if i == 1 { break }
        start := 3
        if flag == 1 {
            start = 1
        }
        for j := start; j <= 2 * i  -1; j += 2 {
            res = append(res, []int{ i, j })
        }
        if flag == 1 {
            flag = 2
        } else {
            flag = 1
        }
        i--
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：size = 3 
    // 输出：[[1,1],[2,1],[2,3],[3,1],[3,5]] 
    // 解释：如下图所示，一种方案为： 指定所示的 5 个子区域为绿地。 
    //       相邻至少两片绿地的 (2,2)，(3,2) 和 (3,4) 演变为绿地。 
    //       相邻两片绿地的 (3,3) 演变为绿地。
    //       <img src="https://pic.leetcode-cn.com/1662692503-ncjywh-image.png" />
    fmt.Println(sandyLandManagement(3)) // [[1,1],[2,1],[2,3],[3,1],[3,5]]
    // 示例 2：
    // 输入：size = 2 
    // 输出：[[1,1],[2,1],[2,3]] 
    // 解释：如下图所示： 指定所示的 3 个子区域为绿地。 
    //      相邻三片绿地的 (2,2) 演变为绿地。
    //      <img src="https://pic.leetcode-cn.com/1662692507-mgFXRj-image.png" />
    fmt.Println(sandyLandManagement(2)) // [[1,1],[2,1],[2,3]]

    fmt.Println(sandyLandManagement(1)) // [[1 1]]
    fmt.Println(sandyLandManagement(5)) // [[1 1] [5 1] [5 3] [5 5] [5 7] [5 9] [4 2] [3 3] [3 5] [2 1]]
    fmt.Println(sandyLandManagement(8)) // [[1 1] [8 1] [8 3] [8 5] [8 7] [8 9] [8 11] [8 13] [8 15] [7 2] [6 3] [6 5] [6 7] [6 9] [6 11] [5 1] [4 1] [4 3] [4 5] [4 7] [3 2] [2 3]]
    // fmt.Println(sandyLandManagement(100)) //
    // fmt.Println(sandyLandManagement(999)) //
    //fmt.Println(sandyLandManagement(1000)) //
}