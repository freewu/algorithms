package main

import (
	"fmt"
	"time"
)

/**
 * 问题: 假设你正在爬楼梯，需要n步你才能到达顶部。但每次你只能爬一步或者两步，你能有多少种不同的方法爬到楼顶部？
 * 解:
 * 1 阶时 1种 解法(1)
 * 2 阶时 2种 解法(1,1 / 2)
 * 3 阶时 3种 解法(1,1,1 / 1,2 / 2,1)
 * 4 阶时 5种 解法(1,1,1,1 / 1,1,2 / 1,2,1 / 2,1,1 / 2,2)
 * 5 阶时 8种 解法(1,1,1,1,1 / 1,1,1,2 / 1,1,2,1 / 1,2,1,1 / 1.2.2 / 2,1,1,1 / 2.1,2 / 2,2,1 )
 *
 * 1  2  3  5  8 ...
 * 3 = 2 + 1
 * 5 = 3 + 2
 * 8 = 5 + 3
 * 得到 公式  f(n) = f(n-1) + f(n-2)
 */

// 递归调用 (最简答的处理方式，性能最差)
func fib1(n int64) int64 {
	if n == 1 || n == 2 {
		return n
	}
	return fib1(n - 1) + fib1(n - 2)
}

// 备忘录法
func fib2(n int64) int64  {
	if n == 1 || n == 2 {
		return n
	} else {
		arr := make([]int64, n + 1)
		arr[1] = 1
		arr[2] = 2
		return dfs(n,arr)
	}
}

func dfs(n int64,arr []int64) int64 {
	if arr[n] != 0 {
		return arr[n]
	} else {
		arr[n] = dfs(n -1,arr) + dfs(n - 2,arr)
		return arr[n]
	}
}

// 动态规划法 (利用数组来存储)
func fib3(n int64) int64  {
	if n == 1 || n == 2 {
		return n
	}
	arr := make([]int64, n + 1)
	arr[1] = 1
	arr[2] = 2
	for i := int64(3) ; i <= n; i++ {
		arr[i] = arr[i - 1] + arr[i - 2]
	}
	return arr[n]
}

// 状态压缩法(又称滚动数组、滑动窗口，用于优化动态规划法的空间复杂度)
func fib4(n int64) int64 {
	if n == 1 || n == 2 {
		return n
	}
	var one,two,result int64
	one = 1
	two = 2
	result = 0
	for {
		if n <= 2 {
			break
		}
		result = one + two
		one = two
		two = result
		n--
	}
	return result
}

// 斐波那契数列的通项公式
func fib5(n int64) int64 {
	if n == 1 || n == 2 {
		return n
	}
	//math.Floor( 1 / math.Sqrt(5) * (math.Pow(1 )))
	return n
}
/*
    public int fib06(int n) {
        if (n == 0)
            return 1;
        if (n == 1 || n == 2)
            return n;
        int result = (int) Math.floor(
                1 / Math.sqrt(5) * (Math.pow((1 + Math.sqrt(5)) / 2, n + 1) - Math.pow((1 - Math.sqrt(5)) / 2, n + 1)));
        return result;
    }
 */

func main() {
	var start,end,total int64

	start = time.Now().UnixNano()
	total = fib1(30)
	end = time.Now().UnixNano()
	fmt.Printf("fib1(30) = %v  used: %d ns \n",total,end - start)

	start = time.Now().UnixNano()
	total = fib2(1500)
	time.Sleep(time.Second * 1)
	end = time.Now().UnixNano()
	fmt.Printf("fib2(1500) = %v  used: %d ns \n",total,end - start)

	start = time.Now().UnixNano()
	total = fib3(1500)
	time.Sleep(time.Second * 1)
	end = time.Now().UnixNano()
	fmt.Printf("fib3(1500) = %v  used: %d ns \n",total,end - start)

	start = time.Now().UnixNano()
	total = fib4(1500)
	time.Sleep(time.Second * 1)
	end = time.Now().UnixNano()
	fmt.Printf("fib4(1500) = %v  used: %d ns \n",total,end - start)

}