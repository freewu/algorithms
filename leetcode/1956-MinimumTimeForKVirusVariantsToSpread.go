package main

// 1956. Minimum Time For K Virus Variants to Spread
// There are n unique virus variants in an infinite 2D grid. 
// You are given a 2D array points, where points[i] = [xi, yi] represents a virus originating at (xi, yi) on day 0. 
// Note that it is possible for multiple virus variants to originate at the same point.

// Every day, each cell infected with a virus variant will spread the virus to all neighboring points in the four cardinal directions (i.e. up, down, left, and right). 
// If a cell has multiple variants, all the variants will spread without interfering with each other.

// Given an integer k, return the minimum integer number of days for any point to contain at least k of the unique virus variants.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/30/case-1.png" />
// Input: points = [[1,1],[6,1]], k = 2
// Output: 3
// Explanation: On day 3, points (3,1) and (4,1) will contain both virus variants. Note that these are not the only points that will contain both virus variants.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/30/case-2.png" />
// Input: points = [[3,3],[1,2],[9,2]], k = 2
// Output: 2
// Explanation: On day 2, points (1,3), (2,3), (2,2), and (3,2) will contain the first two viruses. Note that these are not the only points that will contain both virus variants.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/06/30/case-2.png" />
// Input: points = [[3,3],[1,2],[9,2]], k = 3
// Output: 4
// Explanation: On day 4, the point (5,2) will contain all 3 viruses. Note that this is not the only point that will contain all 3 virus variants.

// Constraints:
//     n == points.length
//     2 <= n <= 50
//     points[i].length == 2
//     1 <= xi, yi <= 100
//     2 <= k <= n

import "fmt"
import "sort"

// 超出时间限制 47 / 70 
func minDayskVariants(points [][]int, k int) int {
    inf := 1 << 31
    res := inf
    popcount := func (n uint) int { // 计算无符号整数中1的个数
        res := 0
        for n != 0 {
            res += int(n & 1)
            n >>= 1
        }
        return res
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < (1 << len(points)); i++ {
        if popcount(uint(i)) != k { continue }
        lefttop, leftdown, righttop, rightdown := inf, inf, -inf, -inf
        for j := 0; j < len(points); j++ {
            if ((1 << j) & i) == 0 { continue }
            lefttop = min(lefttop, points[j][0] - points[j][1])
            leftdown = min(leftdown, points[j][0] + points[j][1])
            righttop = max(righttop, points[j][0] + points[j][1])
            rightdown = max(rightdown, points[j][0] - points[j][1])
        }
        res = min(res, (max(rightdown - lefttop, righttop - leftdown) + 1) / 2)
    }
    return res
}

// 超出时间限制 47 / 70 
func minDayskVariants1(points [][]int, k int) int {
    res, n := 1 << 31, len(points)
    npoints, distx := make([][]int, k), make([]int, k)
    start, end := (1 << k) - 1, (1 << n) - (1 << (n - k))
    x1, y1, x2, y2 := 1_000_000_000, 1_000_000_000, 1, 1
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        x1, y1 = min(x1, points[i][0]), min(y1, points[i][1])
        x2, y2 = max(x2, points[i][0]), max(y2, points[i][1])
    }
    count := func(digit, n, k int) int {
        res := 0
        for i := 0; i < n; i++ {
            if ((1 << i) & digit) != 0 {
                res++
                if res > k { return res }
            }
        }
        return res
    }
    prob := func(points [][]int) int {
        res, n := 0, len(points)
        for i := 0; i < n; i++ {
            for j := i + 1; j < n; j++ {
                tmp := abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1])
                if res < tmp {
                    res = tmp
                }
            }
        }
        return res
    }
    dist := func(points [][]int, distx []int, y int) int {
        res = 0
        for i := 0; i < len(points); i++ {
            tmp := distx[i] + abs(points[i][1] - y)
            if res < tmp {
                res = tmp
            }
        }
        return res
    }
    dist1 := func(points [][]int, x int, y1 int, y2 int, distx []int) int {
        res, n, left, right := 1 << 31, len(points), y1, y2;
        for i := 0; i < n; i++{
            distx[i] = abs(x - points[i][0])
        }
        for left <= right {
            mid := (left + right) / 2
            mmid := (mid + right) / 2
            md, mmd := dist(points, distx, mid),dist(points, distx, mmid);
            if md >= mmd {
                left = mid + 1
                if res > mmd {
                    res = mmd
                }
            } else {
                right = mmid - 1
                if res > md {
                    res = md
                }
            }
        }
        return res
    }
    dist2 := func(points [][]int, distx []int, x1 int, x2 int, y1 int, y2 int) int {
        res, left, right := 1 << 31, x1, x2
        for left <= right {
            mid := (left + right) >> 1
            mmid := (mid + right) >> 1
            md, mmd := dist1(points, mid, y1, y2, distx), dist1(points, mmid, y1, y2, distx)
            if md >= mmd {
                if res > mmd {
                    res = mmd
                }
                left = mid + 1
            } else {
                if res > md {
                    res = md
                }
                right = mmid - 1
            }
        }
        return res
    }
    for i := start; i <= end; i++ {
        if count(i, n, k) != k { continue }
        digit, index := i, 0
        for j := 0; j < n; j++ {
            if ((1 << j) & digit) != 0 {
                npoints[index] = points[j]
                index++
            }
        }
        a := prob(npoints)
        tmp1 := a / 2
        if a % 2 == 1 { tmp1++ }
        if tmp1 > res { continue }
        tmp := dist2(npoints, distx, x1, x2, y1, y2)
        if res > tmp {
            res = tmp
        }
    }
    return res
}

func minDayskVariants2(points [][]int, k int) int {
    n, multer := len(points), 2000000001
    cp, mp := make([][2]int, n), make(map[int]bool)
    for i := 0; i < n; i++ {
        cp[i][0], cp[i][1] = (points[i][0] + points[i][1]), (points[i][0] - points[i][1])
        mp[cp[i][0] * multer + cp[i][1]] = true
    }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if cp[i][0] > cp[j][0] {
                cp[i][0], cp[j][0] = cp[j][0], cp[i][0]
                cp[i][1], cp[j][1] = cp[j][1], cp[i][1]
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := 1000000000
    for i := 0; i <= n - k; i++ {
        for j := k; j <= n - i; j++ {
            xDis := cp[i+j-1][0] - cp[i][0]
            if xDis < res {
                yAxis := []int{}
                for t := 0; t < j; t++ {
                    yAxis = append(yAxis, cp[i + t][1])
                }
                sort.Ints(yAxis)
                for t := 0; t <= j - k; t++ {
                    yDis, yStart, xStart := (yAxis[t+k-1] - yAxis[t]), yAxis[t], cp[i][0]
                    dis := max(xDis, yDis)
                    if dis % 2 == 0 && xDis == yDis && (xStart + yStart) % 2 != 0 {
                        dis++
                    }
                    res = min(res, dis)
                }
            }
        }
    }
    if res == 0 { return 0 }
    return ((res - 1) / 2) + 1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/30/case-1.png" />
    // Input: points = [[1,1],[6,1]], k = 2
    // Output: 3
    // Explanation: On day 3, points (3,1) and (4,1) will contain both virus variants. Note that these are not the only points that will contain both virus variants.
    fmt.Println(minDayskVariants([][]int{{1,1},{6,1}}, 2)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/30/case-2.png" />
    // Input: points = [[3,3],[1,2],[9,2]], k = 2
    // Output: 2
    // Explanation: On day 2, points (1,3), (2,3), (2,2), and (3,2) will contain the first two viruses. Note that these are not the only points that will contain both virus variants.
    fmt.Println(minDayskVariants([][]int{{3,3},{1,2},{9,2}}, 2)) // 2
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/06/30/case-2.png" />
    // Input: points = [[3,3],[1,2],[9,2]], k = 3
    // Output: 4
    // Explanation: On day 4, the point (5,2) will contain all 3 viruses. Note that this is not the only point that will contain all 3 virus variants.
    fmt.Println(minDayskVariants([][]int{{3,3},{1,2},{9,2}}, 3)) // 4
    
    //fmt.Println(minDayskVariants([][]int{{35,43},{41,70},{11,18},{20,30},{50,89},{20,91},{28,9},{54,53},{43,70},{60,54},{8,27},{54,50},{99,75},{90,3},{98,74},{49,62},{1,46},{39,97},{50,54},{69,96},{95,70},{78,29},{63,29},{35,56},{63,4},{50,44},{86,87},{52,93},{22,60},{17,80},{69,4},{51,76}}, 28)) // 4

    fmt.Println(minDayskVariants1([][]int{{1,1},{6,1}}, 2)) // 3
    fmt.Println(minDayskVariants1([][]int{{3,3},{1,2},{9,2}}, 2)) // 2
    fmt.Println(minDayskVariants1([][]int{{3,3},{1,2},{9,2}}, 3)) // 4
    //fmt.Println(minDayskVariants1([][]int{{35,43},{41,70},{11,18},{20,30},{50,89},{20,91},{28,9},{54,53},{43,70},{60,54},{8,27},{54,50},{99,75},{90,3},{98,74},{49,62},{1,46},{39,97},{50,54},{69,96},{95,70},{78,29},{63,29},{35,56},{63,4},{50,44},{86,87},{52,93},{22,60},{17,80},{69,4},{51,76}}, 28)) // 4

    fmt.Println(minDayskVariants2([][]int{{1,1},{6,1}}, 2)) // 3
    fmt.Println(minDayskVariants2([][]int{{3,3},{1,2},{9,2}}, 2)) // 2
    fmt.Println(minDayskVariants2([][]int{{3,3},{1,2},{9,2}}, 3)) // 4
    fmt.Println(minDayskVariants2([][]int{{35,43},{41,70},{11,18},{20,30},{50,89},{20,91},{28,9},{54,53},{43,70},{60,54},{8,27},{54,50},{99,75},{90,3},{98,74},{49,62},{1,46},{39,97},{50,54},{69,96},{95,70},{78,29},{63,29},{35,56},{63,4},{50,44},{86,87},{52,93},{22,60},{17,80},{69,4},{51,76}}, 28)) // 68
}