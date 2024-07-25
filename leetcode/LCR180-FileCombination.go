package main

// LCR 180. 文件组合
// 待传输文件被切分成多个部分，按照原排列顺序，每部分文件编号均为一个 正整数（至少含有两个文件）。
// 传输要求为：连续文件编号总和为接收方指定数字 target 的所有文件。
// 请返回所有符合该要求的文件传输组合列表。

// 注意，返回时需遵循以下规则：
//     每种组合按照文件编号 升序 排列；
//     不同组合按照第一个文件编号 升序 排列。

// 示例 1：
// 输入：target = 12
// 输出：[[3, 4, 5]]
// 解释：在上述示例中，存在一个连续正整数序列的和为 12，为 [3, 4, 5]。

// 示例 2：
// 输入：target = 18
// 输出：[[3,4,5,6],[5,6,7]]
// 解释：在上述示例中，存在两个连续正整数序列的和分别为 18，分别为 [3, 4, 5, 6] 和 [5, 6, 7]。

// 提示：
//     1 <= target <= 10^5

import "fmt"

// 滑动窗口
func fileCombination(target int) [][]int {
    res, sum := [][]int{}, 0
    for left, right := 1, 1; right <= target; {
        for ;right <= target && sum < target; right++ {
            sum += right
        }
        if sum == target {
            temp := []int{}
            for i:=left; i<right;i++{
                temp = append(temp, i)
            }
            res = append(res, append([]int{}, temp...))
            left++
            right=left
            sum = 0
            continue
        }
        for right <=target && sum > target {
            sum -= left
            left++
        }
    }
    return res
}


// 回溯
func fileCombination1(target int) [][]int {
    res, path := [][]int{}, []int{}
    mx := target / 2
    var backtrack func(int, int)
    backtrack = func(startIdx int, target int) {
        if len(path) > 1 && path[len(path)-1] != path[len(path)-2] + 1 {
            return 
        }
        if target == 0 {
            res = append(res, append([]int{}, path...))
            return
        }
        if target < 0 {
            return 
        }
        for i := startIdx; i <= mx; i++ {
            path = append(path, i+1)
            backtrack(i+1, target-i-1)
            path = path[:len(path)-1]
        }
    }
    backtrack(0, target)
    return res
}

func main() {
    // 示例 1：
    // 输入：target = 12
    // 输出：[[3, 4, 5]]
    // 解释：在上述示例中，存在一个连续正整数序列的和为 12，为 [3, 4, 5]。
    fmt.Println(fileCombination(12)) // [[3, 4, 5]]
    // 示例 2：
    // 输入：target = 18
    // 输出：[[3,4,5,6],[5,6,7]]
    // 解释：在上述示例中，存在两个连续正整数序列的和分别为 18，分别为 [3, 4, 5, 6] 和 [5, 6, 7]。
    fmt.Println(fileCombination(18)) // [[3,4,5,6],[5,6,7]]

    fmt.Println(fileCombination1(12)) // [[3, 4, 5]]
    fmt.Println(fileCombination1(18)) // [[3,4,5,6],[5,6,7]]
}