package main

// 1710. Maximum Units on a Truck
// You are assigned to put some amount of boxes onto one truck. 
// You are given a 2D array boxTypes, where boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi]:
//     numberOfBoxesi is the number of boxes of type i.
//     numberOfUnitsPerBoxi is the number of units in each box of the type i.

// You are also given an integer truckSize, which is the maximum number of boxes that can be put on the truck. 
// You can choose any boxes to put on the truck as long as the number of boxes does not exceed truckSize.

// Return the maximum total number of units that can be put on the truck.

// Example 1:
// Input: boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
// Output: 8
// Explanation: There are:
// - 1 box of the first type that contains 3 units.
// - 2 boxes of the second type that contain 2 units each.
// - 3 boxes of the third type that contain 1 unit each.
// You can take all the boxes of the first and second types, and one box of the third type.
// The total number of units will be = (1 * 3) + (2 * 2) + (1 * 1) = 8.

// Example 2:
// Input: boxTypes = [[5,10],[2,5],[4,7],[3,9]], truckSize = 10
// Output: 91

// Constraints:
//     1 <= boxTypes.length <= 1000
//     1 <= numberOfBoxesi, numberOfUnitsPerBoxi <= 1000
//     1 <= truckSize <= 10^6

import "fmt"
import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
    sort.Slice(boxTypes, func(i, j int) bool { // 按每个箱子可以装载的单元数量从大到小排序
        return boxTypes[i][1] > boxTypes[j][1]
    })
    res, trucks := 0, 0
    for _, b := range boxTypes {
        if trucks < truckSize {
            if trucks + b[0] > truckSize {
                res += (truckSize - trucks) * b[1]
                break
            } else {
                res += b[0] * b[1]
                trucks += b[0]
            }
        }
    }
    return res
}

func maximumUnits1(boxTypes [][]int, truckSize int) int {
    sort.Slice(boxTypes , func(i,j int) bool {
        return boxTypes[i][1] > boxTypes[j][1]
    })
    res, i := 0, 0
    for truckSize > 0 && i < len(boxTypes) {
        if truckSize >= boxTypes[i][0] {
            res += boxTypes[i][0] * boxTypes[i][1]
            truckSize -= boxTypes[i][0]
        } else {
            res += truckSize * boxTypes[i][1]
            truckSize = 0
        }
        i++
    }
    return res
}

func main() {
    // Example 1:
    // Input: boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
    // Output: 8
    // Explanation: There are:
    // - 1 box of the first type that contains 3 units.
    // - 2 boxes of the second type that contain 2 units each.
    // - 3 boxes of the third type that contain 1 unit each.
    // You can take all the boxes of the first and second types, and one box of the third type.
    // The total number of units will be = (1 * 3) + (2 * 2) + (1 * 1) = 8.
    fmt.Println(maximumUnits([][]int{{1,3},{2,2},{3,1}}, 4)) // 8  (1 * 3) + (2 * 2) + (1 * 1) = 8.
    // Example 2:
    // Input: boxTypes = [[5,10],[2,5],[4,7],[3,9]], truckSize = 10
    // Output: 91
    fmt.Println(maximumUnits([][]int{{5,10},{2,5},{4,7},{3,9}}, 10)) // 91 (5 * 10) + (3 * 9) + ((10-8) * 7) = 91

    fmt.Println(maximumUnits1([][]int{{1,3},{2,2},{3,1}}, 4)) // 8  (1 * 3) + (2 * 2) + (1 * 1) = 8.
    fmt.Println(maximumUnits1([][]int{{5,10},{2,5},{4,7},{3,9}}, 10)) // 91 (5 * 10) + (3 * 9) + ((10-8) * 7) = 91
}