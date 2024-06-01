package main

// 1231. Divide Chocolate
// You have one chocolate bar that consists of some chunks. 
// Each chunk has its own sweetness given by the array sweetness.

// You want to share the chocolate with your k friends so you start cutting the chocolate bar into k + 1 pieces using k cuts, each piece consists of some consecutive chunks.
// Being generous, you will eat the piece with the minimum total sweetness and give the other pieces to your friends.
// Find the maximum total sweetness of the piece you can get by cutting the chocolate bar optimally.

// Example 1:
// Input: sweetness = [1,2,3,4,5,6,7,8,9], k = 5
// Output: 6
// Explanation: You can divide the chocolate to [1,2,3], [4,5], [6], [7], [8], [9]

// Example 2:
// Input: sweetness = [5,6,7,8,9,1,2,3,4], k = 8
// Output: 1
// Explanation: There is only one way to cut the bar into 9 pieces.

// Example 3:
// Input: sweetness = [1,2,2,1,2,2,1,2,2], k = 2
// Output: 5
// Explanation: You can divide the chocolate to [1,2,2], [1,2,2], [1,2,2]
 
// Constraints:
//     0 <= k < sweetness.length <= 10^4
//     1 <= sweetness[i] <= 10^5

import "fmt"

func maximizeSweetness(sweetness []int, k int) int {
    // 统计总和，用于确定右界
    sum := 0
    for _, v := range sweetness{
        sum += v
    }
    // 计算指定最小值的甜度，可分成的份数
    var helper = func(m int) int {
        t, count := 0, 0
        for _, v := range sweetness {
            t += v
            if t >= m {
                count++
                t = 0
            }
        }
        return count
    }
    left, right := 1, sum / (k + 1) // 使用平均值作为右界
    // 二分查找
    for left < right {
        mid := (left + right + 1) >> 1
        if helper(mid) > k { // 份数满足要求
            left = mid
        } else {
            right = mid - 1
        }
    }
    return right
}

func maximizeSweetness1(sweetness []int, k int) int {
    low, high := 1, int(1e9)
    // 每一块都由一些 连续 的小块组成
    for low <= high {
        mid := (low + high) / 2
        count, sum := 0, 0 
        for i := 0; i < len(sweetness); i++ {
            sum += sweetness[i]
            if sum >= mid {
                sum = 0
                count++
            }
        }
        if count > k { // 份数满足要求
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return low - 1
}

func maximizeSweetness2(sweetness []int, k int) int {
    sum, least := 0, 1 << 31-1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, s := range sweetness {
        sum += s
        least = min(least, s)
    }
    canMaximizeSweetness := func (sweetness []int, k, least int) bool {
        cnt, sum := 0, 0
        for _, s := range sweetness {
            sum += s
            if sum >= least {
                cnt++
                sum = 0
            }
        }
        return cnt >= k+1
    }
    l,r := least,sum / (k+1) + 1
    for l < r-1 {
        mid := (l+r) >> 1
        if canMaximizeSweetness(sweetness,k,mid){
            l = mid
        }else{
            r = mid
        }
    }
    return l
}

func main() {
    // Explanation: You can divide the chocolate to [1,2,3], [4,5], [6], [7], [8], [9]
    fmt.Println(maximizeSweetness([]int{1,2,3,4,5,6,7,8,9},5)) // 6
    // Explanation: There is only one way to cut the bar into 9 pieces.
    fmt.Println(maximizeSweetness([]int{5,6,7,8,9,1,2,3,4},8)) // 1
    // Explanation: You can divide the chocolate to [1,2,2], [1,2,2], [1,2,2]
    fmt.Println(maximizeSweetness([]int{1,2,2,1,2,2,1,2,2},2)) // 5

    fmt.Println(maximizeSweetness1([]int{1,2,3,4,5,6,7,8,9},5)) // 6
    fmt.Println(maximizeSweetness1([]int{5,6,7,8,9,1,2,3,4},8)) // 1
    fmt.Println(maximizeSweetness1([]int{1,2,2,1,2,2,1,2,2},2)) // 5

    fmt.Println(maximizeSweetness2([]int{1,2,3,4,5,6,7,8,9},5)) // 6
    fmt.Println(maximizeSweetness2([]int{5,6,7,8,9,1,2,3,4},8)) // 1
    fmt.Println(maximizeSweetness2([]int{1,2,2,1,2,2,1,2,2},2)) // 5
}