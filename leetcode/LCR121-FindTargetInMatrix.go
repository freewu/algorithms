package main

// LCR 121. 寻找目标值 - 二维数组
// m*n 的二维数组 plants 记录了园林景观的植物排布情况，具有以下特性：
//     每行中，每棵植物的右侧相邻植物不矮于该植物；
//     每列中，每棵植物的下侧相邻植物不矮于该植物。

// 请判断 plants 中是否存在目标高度值 target。

// 示例 1：
// 输入：plants = [[2,3,6,8],[4,5,8,9],[5,9,10,12]], target = 8
// 输出：true

// 示例 2：
// 输入：plants = [[1,3,5],[2,5,7]], target = 4
// 输出：false

// 提示：
//     0 <= n <= 1000
//     0 <= m <= 1000

import "fmt"

// 模拟，时间复杂度 O(m+n)
func findTargetIn2DPlants(plants [][]int, target int) bool {
    if len(plants) == 0 {
        return false
    }
    row, col := 0, len(plants[0])-1
    for col >= 0 && row <= len(plants)-1 {
        if target == plants[row][col] {
            return true
        } else if target > plants[row][col] { // 如果当前值小于目标值向下一层走，否则向里走一步
            row++
        } else {
            col--
        }
    }
    return false
}

// 二分搜索，时间复杂度 O(n log n)
func findTargetIn2DPlants1(plants [][]int, target int) bool {
    if len(plants) == 0 {
        return false
    }
    for _, row := range plants { // 循环每行
        // 每行做二分法查找
        low, high := 0, len(plants[0])-1
        for low <= high {
            mid := low + (high - low) >> 1
            if row[mid] > target {
                high = mid - 1
            } else if row[mid] < target {
                low = mid + 1
            } else { // row[mid] == target 找到了
                return true
            }
        }
    }
    return false
}

func main() {
    matrix := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
    fmt.Printf("findTargetIn2DPlants(matrix,5) = %v\n",findTargetIn2DPlants(matrix,5)) // true
    fmt.Printf("findTargetIn2DPlants(matrix,20) = %v\n",findTargetIn2DPlants(matrix,20)) // false

    fmt.Printf("findTargetIn2DPlants1(matrix,5) = %v\n",findTargetIn2DPlants1(matrix,5)) // true
    fmt.Printf("findTargetIn2DPlants1(matrix,20) = %v\n",findTargetIn2DPlants1(matrix,20)) // false
}