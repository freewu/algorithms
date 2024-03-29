package main

// LCR 001. 两数相除
// 给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。
// 注意：
// 		整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
// 		假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31, 2^31−1]。本题中，如果除法结果溢出，则返回 2^31 − 1
 
// 示例 1：
// 输入：a = 15, b = 2
// 输出：7
// 解释：15/2 = truncate(7.5) = 7

// 示例 2：
// 输入：a = 7, b = -3
// 输出：-2
// 解释：7/-3 = truncate(-2.33333..) = -2

// 示例 3：
// 输入：a = 0, b = 1
// 输出：0

// 示例 4：
// 输入：a = 1, b = 1
// 输出：1
 
// 提示:
// 		-2^31 <= a, b <= 2^31 - 1
// 		b != 0

import "fmt"
import "math"

func divide(a int, b int) int {
	if a == math.MinInt32 { // 考虑被除数为最小值的情况
        if b == 1 {
            return math.MinInt32
        }
        if b == -1 {
            return math.MaxInt32
        }
    }
    if b == math.MinInt32 { // 考虑除数为最小值的情况
        if a == math.MinInt32 {
            return 1
        }
        return 0
    }
    if a == 0 { // 考虑被除数为 0 的情况
        return 0
    }
	// 求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 
	// return a / b

	// 一般情况，使用二分查找
    // 将所有的正数取相反数，这样就只需要考虑一种情况
	// 判断正负
    rev := false
    if a > 0 {
        a = -a
        rev = !rev
    }
    if b > 0 {
        b = -b
        rev = !rev
    }

    candidates := []int{b}
    for y := b; y >= a-y; { // 注意溢出
        y += y
        candidates = append(candidates, y)
    }
    ans := 0
    for i := len(candidates) - 1; i >= 0; i-- {
        if candidates[i] >= a {
            ans |= 1 << i
            a -= candidates[i]
        }
    }
    if rev {
        return -ans
    }
    return ans
}

func main() {
	fmt.Println(divide(15,2)) // 7
	fmt.Println(divide(7,-3)) // -2
	fmt.Println(divide(0,1)) // 0
	fmt.Println(divide(1,1)) // 1
	fmt.Println(divide(-2147483648,-1)) // 2147483647
}