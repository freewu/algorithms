package main

/**
147 · Narcissistic Number
# Description
Narcissistic Number is a number that is the sum of its own digits each raised to the power of the number of digits. 
See wiki： https://en.wikipedia.org/wiki/Narcissistic_number
For example the 3-digit decimal number 15^3 is a narcissistic number because 

	153 = 1^3 + 5^3 + 3^3.

And the 4-digit decimal number 1634 is a narcissistic number because 

	1634 = 1^4 + 6^4 + 3^4 + 4^4.

Given n, return all narcissistic numbers with n digits.
You may assume n is smaller than 8.

Example 1:

	Input: 1
	Output: [0,1,2,3,4,5,6,7,8,9]

Example 2:

	Input:  2
	Output: []
	Explanation: There is no Narcissistic Number with 2 digits.

在数论中，水仙花数（Narcissistic number），
也被称为超完全数字不变数（pluperfect digital invariant, PPDI）、自恋数、自幂数、阿姆斯壮数或阿姆斯特朗数（Armstrong number），
用来描述一个N位非负整数，其各位数字的N次方和等于该数本身。

*/

/*
python 的解法
def getNarcissisticNumbers(self,n):
      l = []  #存放满足要求的水仙花数
      if n == 1:
          l.append(0) #如果n等于1，下面循环从range（1,10），少了0，所以先补上
      store = [] 
      for item in range(10):
          store.append(pow(item,n)) #存放0-9的N次幂。这里先把N次幂算出来，下面要用的话就直接取，不会再每一个都算一遍
                               
      for j in range(pow(10,n-1),pow(10,n)):   #j表示几位数，例如n == 4，就是所有的四位数
          j_str = str(j)  #先将整数j转换为字符串，方便依次取下标
          sums = 0
          for k in range(len(j_str)):
              sums += store[int(j_str[k])]  #这里int（j_str）这里循环依次取出整数j的每一个数字
          if j == sums:
              l.append(j)
      return l
*/

import "fmt"

/**
 * @param n: The number of digits
 * @return: All narcissistic numbers with n digits
 */
 func GetNarcissisticNumbers(n int) []int {
    // write your code here
    res := make([]int) // 用于保存满足要求的水仙花数的结果
    if n == 1 { // 如果n等于 1，下面循环从range（1,10），少了0，所以先补上
        res = append(res,0)
    }
    start := Pow(10,n - 1)
    for i = start; i < start * 10; i++ {
        if isNarcissistic(i) {
            res = append(res,i)
        }
    }
}

// m^n 的 m 的 n 次方
func Pow(m int ,n int) int {
    res := m
    for i := 1; i < n; i++ {
        res = res * m
    }
    return res
}

func isNarcissistic(n int) bool {
    s := 0
    tmp := n
    m := len(fmt.Sprint(n))
    for tmp >= 1 {
        s += Pow(m,tmp % 10)
        tmp /= 10
    }
    return s == n
}