package main 

// LCR 003. 比特位计数
// 给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。

// 示例 1:
// 输入: n = 2
// 输出: [0,1,1]
// 解释: 
// 0 --> 0
// 1 --> 1
// 2 --> 10

// 示例 2:
// 输入: n = 5
// 输出: [0,1,1,2,1,2]
// 解释:
// 0 --> 0
// 1 --> 1
// 2 --> 10
// 3 --> 11
// 4 --> 100
// 5 --> 101

// 说明 :
//      0 <= n <= 10^5
 
// 进阶:
//      给出时间复杂度为 O(n*sizeof(integer)) 的解答非常容易。但你可以在线性时间 O(n) 内用一趟扫描做到吗？
//      要求算法的空间复杂度为 O(n) 。
//      你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount ）来执行此操作。

import "fmt"

//   X & 1 ==1 or ==0，可以用 X & 1 判断奇偶性，X & 1 > 0 即奇数  X & 1 == 0 即偶数
//   X = X & (X-1) 清零最低位的1
//   X & -X => 得到最低位的1 
//   X &~X=>0
func countBits(n int) []int {
	bits := make([]int, n + 1)
	for i := 1; i <= n; i++ {
        fmt.Printf("i: %v i & (i-1): %b, %v\n", i,i & (i-1),i & (i-1))
        // X & (X-1) 清零最低位的 1
		bits[i] += bits[ i & (i-1) ] + 1
	}
	return bits
}

func countBits1(n int) []int {
	res := []int{0}
	k := 1
	for i := 1; i <= n; i++ {
		if k*2 == i {
			k = i
		}
		res = append(res, res[i-k]+1)
	}
	return res
}

func main() {
    fmt.Println(countBits(2)) // [0,1,1]
    fmt.Println(countBits(5)) // [0,1,1,2,1,2]
    fmt.Println(countBits(10)) // [0 1 1 2 1 2 2 3 1 2 2]
    fmt.Println(countBits(16)) // [0 1 1 2 1 2 2 3 1 2 2 3 2 3 3 4 1]
    //fmt.Println(countBits(1024)) // 

    fmt.Println(countBits1(2)) // [0,1,1]
    fmt.Println(countBits1(5)) // [0,1,1,2,1,2]
    fmt.Println(countBits1(10)) // [0 1 1 2 1 2 2 3 1 2 2]
    fmt.Println(countBits1(16)) // [0 1 1 2 1 2 2 3 1 2 2 3 2 3 3 4 1]
    //fmt.Println(countBits1(1024)) // 
}