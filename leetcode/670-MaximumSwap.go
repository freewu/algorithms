package main

// 670. Maximum Swap
// You are given an integer num. You can swap two digits at most once to get the maximum valued number.
// Return the maximum valued number you can get.

// Example 1:
// Input: num = 2736
// Output: 7236
// Explanation: Swap the number 2 and the number 7.

// Example 2:
// Input: num = 9973
// Output: 9973
// Explanation: No swap.
 
// Constraints:
// 	0 <= num <= 10^8

import "fmt"

func maximumSwap(num int) int {
	if num == 0 {
        return 0
    }
	// 先把数字拆解成数组 如 2736 => [2,7,3,6]
    digs := []int{}
    for num > 0 {
        digs = append(digs, num%10)
        num /= 10
    }
    max := make([]int, len(digs))
    idx := make([]int, len(digs))
    for i, d := range digs {
        if i == 0 || max[i-1] < d {
            max[i] = d
            idx[i] = i
        } else {
            max[i] = max[i-1]
            idx[i] = idx[i-1]
        }
		fmt.Println("max: ", max)
		fmt.Println("idx: ", idx)
    }
    for i := len(digs)-1; i >= 0; i-- {
        d := digs[i]
		// 最大值在后面  idx[i] 记录了 最大值的位置
		// 把最大值替换到前面
        if d < max[i] {
            digs[i], digs[idx[i]] = digs[idx[i]], digs[i]
            break
        }
    }
    for i := len(digs)-1; i >= 0; i-- {
        num *= 10
        num += digs[i]
    }
    return num
}

func main() {
	fmt.Println(maximumSwap(2736)) // 7236
	fmt.Println(maximumSwap(9973)) // 9973
}

