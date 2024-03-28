package main

// 29. Divide Two Integers
// Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
// The integer division should truncate toward zero, which means losing its fractional part. 
// For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

// Return the quotient after dividing dividend by divisor.

// Note: 
//     Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−2^31, 2^31 − 1]. 
//     For this problem, if the quotient is strictly greater than 2^31 - 1, then return 2^31 - 1, and if the quotient is strictly less than -231, then return -231.

// Example 1:
// Input: dividend = 10, divisor = 3
// Output: 3
// Explanation: 10/3 = 3.33333.. which is truncated to 3.

// Exam
// Input: dividend = 7, divisor = -3
// Output: -2
// Explanation: 7/-3 = -2.33333.. which is truncated to -2.
 
// Constraints:
//     -2^31 <= dividend, divisor <= 2^31 - 1
//     divisor != 0

import "fmt"
import "math"

func divide(dividend int, divisor int) int {
    if (dividend / divisor) > math.MaxInt32 {
        return math.MaxInt32
    } else {
        return dividend / divisor
    }
}

// 二分法+位用算乘法
// 注意的是我们把除数和被除数都转成了负数，当 y*mid >= x 时，mid 需要往 right 方向取
func divide1(dividend, divisor int) int {
    if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
        if divisor == 1 {
            return math.MinInt32
        }
        if divisor == -1 {
            return math.MaxInt32
        }
    }
    if divisor == math.MinInt32 { // 考虑除数为最小值的情况
        if dividend == math.MinInt32 {
            return 1
        }
        return 0
    }
    if dividend == 0 { // 考虑被除数为 0 的情况
        return 0
    }

    // 一般情况，使用二分查找
    // 将所有的正数取相反数，这样就只需要考虑一种情况
    flag := false // 正负符号位
    if dividend > 0 {
        dividend = -dividend
        flag = !flag
    }
    if divisor > 0 {
        divisor = -divisor
        flag = !flag
    }

    multiply := func (x, y int) int {
        res := 0
        for y > 0 {
            if y & 1 == 1 { // 判断最后一位[低位是否为 1]
                res += x
            }
            x <<= 1
            y >>= 1
        }
        return res
    }
    // 快速乘
    // x 和 y 是负数，z 是正数
    // 判断 z * y >= x 是否成立
    quickAdd := func (y, z, x int) bool { if multiply(y, z) < x { return false; }; return true; }

    res := 0
    left, right := 1, math.MaxInt32
    for left <= right {
        mid := left + (right-left) >> 1 // 注意溢出，并且不能使用除法
        if quickAdd(divisor, mid, dividend) {
            res = mid
            if mid == math.MaxInt32 { // 注意溢出
                break
            }
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    if flag {
        return -res
    }
    return res
}

func divide2(dividend int, divisor int) int {
	sign := 1
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	pair := [][2]int{}
	{
		i := divisor
		j := 1
		for i <= dividend {
			pair = append(pair, [2]int{i, j})
			i += i
			j += j
		}
	}
	res := 0
	for i := len(pair) - 1; i >= 0; i-- {
		if pair[i][0] <= dividend {
			dividend -= pair[i][0]
			res += pair[i][1] * sign

			if res >= math.MaxInt32 {
				return math.MaxInt32
			}
			if res <= math.MinInt32 {
				return math.MinInt32
			}
		}
	}
	return res
}

func main() {
    // Explanation: 10/3 = 3.33333.. which is truncated to 3.
    fmt.Println(divide(10,3)) // 3

    // Explanation: 7/-3 = -2.33333.. which is truncated to -2.
    fmt.Println(divide(7,-3)) // -2

    fmt.Println(divide1(10,3)) // 3
    fmt.Println(divide1(7,-3)) // -2

    fmt.Println(divide2(10,3)) // 3
    fmt.Println(divide2(7,-3)) // -2

    fmt.Println(math.MinInt32)
    fmt.Println(math.MaxInt32)
}