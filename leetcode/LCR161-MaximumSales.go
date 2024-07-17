package main

// LCR 161. 连续天数的最高销售额
// 某公司每日销售额记于整数数组 sales，请返回所有 连续 一或多天销售额总和的最大值。
// 要求实现时间复杂度为 O(n) 的算法。

// 示例 1:
// 输入：sales = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：[4,-1,2,1] 此连续四天的销售总额最高，为 6。

// 示例 2:
// 输入：sales = [5,4,-1,7,8]
// 输出：23
// 解释：[5,4,-1,7,8] 此连续五天的销售总额最高，为 23。 

// 提示：
//     1 <= arr.length <= 10^5
//     -100 <= arr[i] <= 100

import "fmt"

func maxSales(sales []int) int {
    mx, sum := -1 >> 32 - 1, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := range sales {
        sum = max(sum + sales[i], sales[i]) // 如果累加值都小于当前值，从当前值开始
        mx = max(mx, sum)
    }
    return mx
}

func maxSales1(sales []int) int {
    n, mx := len(sales), 0
    array := make([][]int, n)
    for i := 0; i < n; i++ {
        subArray := make([]int, n)
        subArray[0] = sales[i]
        for j := i + 1; j < n; j++ {
            subArray[i] += sales[j]
            if subArray[i] > mx {
                mx = subArray[i]
            }
        }
        array[i] = subArray
    }
    return mx
}

// DP
func maxSales2(sales []int) int {
    if len(sales) == 0 { return 0 }
    if len(sales) == 1 { return sales[0] }
    max := func (a int, b int) int { if a > b { return a; }; return b; }
    dp, res := make([]int, len(sales)), sales[0]
    dp[0] = sales[0]
    for i := 1; i < len(sales); i++ {
        if dp[i-1] > 0 { // 如果值大于 0 则累加
            dp[i] = sales[i] + dp[i-1]
        } else {
            dp[i] = sales[i]
        }
        res = max(res, dp[i])
    }
    return res
}

// 模拟
func maxSales3(sales []int) int {
    if len(sales) == 1 { return sales[0] }
    maxSum, res, p := sales[0], 0, 0
    for p < len(sales) {
        res += sales[p]
        if res > maxSum {
            maxSum = res
        }
        if res < 0 {
            res = 0
        }
        p++
    }
    return maxSum
}

//  快慢指针
func maxSales4(sales []int) int {
    n, fast, slow := len(sales), 0, 0
    res, sum :=  -1 >> 32 - 1, 0
    for fast < n && slow < n {
        sum = sum + sales[fast]
        if sum > res {
            res = sum
        }
        for sum< 0 {
            sum = sum - sales[slow]
            slow++
        }
        fast++
    }
    return res
}

func maxSales5(sales []int) int {
    n := len(sales)
    res, last_day_max := sales[0], make([]int, n)
    last_day_max[0] = sales[0]
    for i := 1; i < n; i++ {
        if last_day_max[i-1]+sales[i] > sales[i] {
            last_day_max[i] = last_day_max[i-1] + sales[i]
        } else {
            last_day_max[i] = sales[i]
        }
        if last_day_max[i] > res {
            res = last_day_max[i]
        }
    }
    return res
}

func main() {
    fmt.Printf("maxSales([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSales([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
    fmt.Printf("maxSales([]int{1}) = %v\n",maxSales([]int{1})) // 1
    fmt.Printf("maxSales([]int{5,4,-1,7,8}) = %v\n",maxSales([]int{5,4,-1,7,8})) // 23
    fmt.Printf("maxSales([]int{-3,-3,-3}) = %v\n",maxSales([]int{-3,-3,-3})) // -3

    fmt.Printf("maxSales1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSales1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
    fmt.Printf("maxSales1([]int{1}) = %v\n",maxSales1([]int{1})) // 1
    fmt.Printf("maxSales1([]int{5,4,-1,7,8}) = %v\n",maxSales1([]int{5,4,-1,7,8})) // 23
    fmt.Printf("maxSales1([]int{-3,-3,-3}) = %v\n",maxSales1([]int{-3,-3,-3})) // -3

    fmt.Printf("maxSales2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSales2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
    fmt.Printf("maxSales2([]int{1}) = %v\n",maxSales2([]int{1})) // 1
    fmt.Printf("maxSales2([]int{5,4,-1,7,8}) = %v\n",maxSales2([]int{5,4,-1,7,8})) // 23
    fmt.Printf("maxSales2([]int{-3,-3,-3}) = %v\n",maxSales2([]int{-3,-3,-3})) // -3

    fmt.Printf("maxSales3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSales3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
    fmt.Printf("maxSales3([]int{1}) = %v\n",maxSales3([]int{1})) // 1
    fmt.Printf("maxSales3([]int{5,4,-1,7,8}) = %v\n",maxSales3([]int{5,4,-1,7,8})) // 23
    fmt.Printf("maxSales3([]int{-3,-3,-3}) = %v\n",maxSales3([]int{-3,-3,-3})) // -3

    fmt.Printf("maxSales4([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSales4([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
    fmt.Printf("maxSales4([]int{1}) = %v\n",maxSales4([]int{1})) // 1
    fmt.Printf("maxSales4([]int{5,4,-1,7,8}) = %v\n",maxSales4([]int{5,4,-1,7,8})) // 23
    fmt.Printf("maxSales4([]int{-3,-3,-3}) = %v\n",maxSales4([]int{-3,-3,-3})) // -3 -2?

    fmt.Printf("maxSales5([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSales5([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
    fmt.Printf("maxSales5([]int{1}) = %v\n",maxSales5([]int{1})) // 1
    fmt.Printf("maxSales5([]int{5,4,-1,7,8}) = %v\n",maxSales5([]int{5,4,-1,7,8})) // 23
    fmt.Printf("maxSales5([]int{-3,-3,-3}) = %v\n",maxSales5([]int{-3,-3,-3})) // -3
}