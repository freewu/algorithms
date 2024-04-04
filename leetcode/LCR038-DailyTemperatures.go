package main 

// LCR 038. 每日温度
// 请根据每日 气温 列表 temperatures ，重新生成一个列表，要求其对应位置的输出为：
//     要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。

// 示例 1:
// 输入: temperatures = [73,74,75,71,69,72,76,73]
// 输出: [1,1,4,2,1,1,0,0]

// 示例 2:
// 输入: temperatures = [30,40,50,60]
// 输出: [1,1,1,0]

// 示例 3:
// 输入: temperatures = [30,60,90]
// 输出: [1,1,0]
 
// 提示：
//     1 <= temperatures.length <= 10^5
//     30 <= temperatures[i] <= 100

import "fmt"

// 暴力解法
func dailyTemperatures(T []int) []int {
    res, j := make([]int, len(T)), 0
    for i := 0; i < len(T); i++ {
        // 从 i + 1 开始找
        for j = i + 1; j < len(T); j++ {
            // 如果发现后面温度大于该天，计算之间相差天数
            if T[j] > T[i] {
                res[i] = j - i
                break
            }
        }
    }
    return res
}

// 使用 stack 
func dailyTemperatures1(T []int) []int {
    res := make([]int, len(T))
    stack := []int{}
    for i, t := range T {
        for len(stack) > 0 && T[stack[len(stack)-1]] < t {
            // 维护一个单调递减的单调栈
            idx := stack[len(stack)-1]
            res[idx] = i - idx
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return res
}

func main() {
	fmt.Println(dailyTemperatures([]int{73,74,75,71,69,72,76,73})) // [1,1,4,2,1,1,0,0]
	fmt.Println(dailyTemperatures([]int{30,40,50,60})) // [1,1,1,0]
	fmt.Println(dailyTemperatures([]int{30,60,90})) // [1,1,0]

	fmt.Println(dailyTemperatures1([]int{73,74,75,71,69,72,76,73})) // [1,1,4,2,1,1,0,0]
	fmt.Println(dailyTemperatures1([]int{30,40,50,60})) // [1,1,1,0]
	fmt.Println(dailyTemperatures1([]int{30,60,90})) // [1,1,0]
}