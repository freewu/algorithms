package main

// 2280. Minimum Lines to Represent a Line Chart
// You are given a 2D integer array stockPrices where stockPrices[i] = [dayi, pricei] indicates the price of the stock on day dayi is pricei. 
// A line chart is created from the array by plotting the points on an XY plane 
// with the X-axis representing the day and the Y-axis representing the price and connecting adjacent points. 
// One such example is shown below:
// <img src="https://assets.leetcode.com/uploads/2022/03/30/1920px-pushkin_population_historysvg.png" />

// Return the minimum number of lines needed to represent the line chart.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/30/ex0.png" />
// Input: stockPrices = [[1,7],[2,6],[3,5],[4,4],[5,4],[6,3],[7,2],[8,1]]
// Output: 3
// Explanation:
// The diagram above represents the input, with the X-axis representing the day and Y-axis representing the price.
// The following 3 lines can be drawn to represent the line chart:
// - Line 1 (in red) from (1,7) to (4,4) passing through (1,7), (2,6), (3,5), and (4,4).
// - Line 2 (in blue) from (4,4) to (5,4).
// - Line 3 (in green) from (5,4) to (8,1) passing through (5,4), (6,3), (7,2), and (8,1).
// It can be shown that it is not possible to represent the line chart using less than 3 lines.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/30/ex1.png" />
// Input: stockPrices = [[3,4],[1,2],[7,8],[2,3]]
// Output: 1
// Explanation:
// As shown in the diagram above, the line chart can be represented with a single line.

// Constraints:
//     1 <= stockPrices.length <= 10^5
//     stockPrices[i].length == 2
//     1 <= dayi, pricei <= 10^9
//     All dayi are distinct.

import "fmt"
import "slices"

func minimumLines(stockPrices [][]int) int {
    slices.SortFunc(stockPrices, func(a1, a2 []int) int {
        return a2[0] - a1[0]
    })
    x1, y1, res, i, n := stockPrices[0][0], stockPrices[0][1], 0, 1, len(stockPrices)
    for i < n {
        x2, y2 := stockPrices[i][0], stockPrices[i][1]
        for i + 1 < n && (x2 - stockPrices[i+1][0]) * (y2 - y1) == (x2 - x1) * (y2 - stockPrices[i+1][1]) {
            i++
        }
        x1, y1 = stockPrices[i][0], stockPrices[i][1]
        i++
        res++
    }
    return res
}

func minimumLines1(stockPrices [][]int) int {
    if len(stockPrices) == 1 { return 0 }
    slices.SortFunc(stockPrices, func(a, b []int) int { 
        return a[0] - b[0]
    })
    res, x, y, pdx, pdy := 1, stockPrices[0][0], stockPrices[0][1], 0, 0
    for _, v := range stockPrices[1:] {
        dx, dy := v[0] - x, v[1] - y
        if dx * pdy != dy * pdx {
            res++
        }
        x, y = v[0], v[1]
        pdx, pdy = dx, dy
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/30/ex0.png" />
    // Input: stockPrices = [[1,7],[2,6],[3,5],[4,4],[5,4],[6,3],[7,2],[8,1]]
    // Output: 3
    // Explanation:
    // The diagram above represents the input, with the X-axis representing the day and Y-axis representing the price.
    // The following 3 lines can be drawn to represent the line chart:
    // - Line 1 (in red) from (1,7) to (4,4) passing through (1,7), (2,6), (3,5), and (4,4).
    // - Line 2 (in blue) from (4,4) to (5,4).
    // - Line 3 (in green) from (5,4) to (8,1) passing through (5,4), (6,3), (7,2), and (8,1).
    // It can be shown that it is not possible to represent the line chart using less than 3 lines.
    fmt.Println(minimumLines([][]int{{1,7},{2,6},{3,5},{4,4},{5,4},{6,3},{7,2},{8,1}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/30/ex1.png" />
    // Input: stockPrices = [[3,4],[1,2],[7,8],[2,3]]
    // Output: 1
    // Explanation:
    // As shown in the diagram above, the line chart can be represented with a single line.
    fmt.Println(minimumLines([][]int{{3,4},{1,2},{7,8},{2,3}})) // 1

    fmt.Println(minimumLines1([][]int{{1,7},{2,6},{3,5},{4,4},{5,4},{6,3},{7,2},{8,1}})) // 3
    fmt.Println(minimumLines1([][]int{{3,4},{1,2},{7,8},{2,3}})) // 1
}