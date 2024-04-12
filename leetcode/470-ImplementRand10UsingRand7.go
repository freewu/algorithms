package main

// 470. Implement Rand10() Using Rand7()
// Given the API rand7() that generates a uniform random integer in the range [1, 7], write a function rand10() that generates a uniform random integer in the range [1, 10]. You can only call the API rand7(), and you shouldn't call any other API. Please do not use a language's built-in random API.
// Each test case will have one internal argument n, the number of times that your implemented function rand10() will be called while testing. Note that this is not an argument passed to rand10().

// Example 1:
// Input: n = 1
// Output: [2]

// Example 2:
// Input: n = 2
// Output: [2,8]

// Example 3:
// Input: n = 3
// Output: [3,8,10]
 
// Constraints:
//     1 <= n <= 10^5
 
// Follow up:
//     What is the expected value for the number of calls to rand7() function?
//     Could you minimize the number of calls to rand7()?

// # 解题思路
// # 给出 rand7() 要求实现 rand10()
//     rand7() 等概率地产生1，2，3，4，5，6，7。要想得到 rand10() 即等概率的生成 1-10 。
//     解题思路是先构造一个 randN()，这个 N 必须是 10 的整数倍，然后 randN % 10 就可以得到 rand10() 了。
//     所以可以从 rand7() 先构造出 rand49()，再把 rand49() 中大于等于 40 的都过滤掉，这样就得到了 rand40()，在对 10 取余即可。

//     具体构造步骤，rand7() --> rand49() --> rand40() --> rand10()：
//         rand7() 等概率地产生 1,2,3,4,5,6,7.
//         rand7() - 1 等概率地产生 [0,6].
//         (rand7() - 1) * 7 等概率地产生 0, 7, 14, 21, 28, 35, 42
//         (rand7() - 1) * 7 + (rand7() - 1) 等概率地产生 [0, 48] 这 49 个数字

//     如果步骤 4 的结果大于等于 40，那么就重复步骤 4，直到产生的数小于 40
//     把步骤 5 的结果 mod 10 再加 1， 就会等概率的随机生成 [1, 10]

// 这道题可以推广到生成任意数的随机数问题。用 randN() 实现 randM()，M > N。步骤如下：
//     1 用 randN() 先实现 randX()，其中 X ≥ M，X 是 M 的整数倍。如这题中的 49 > 10；
//     2 再用 randX() 生成 randM()，如这题中的 49 —> 40 —> 10。

//     举个例子，
//         用 rand3() 生成 rand11()，可以先生成 rand27()，然后变成以 22 为界限，因为 22 是 11 的倍数。生成 rand27() 的方式：3 * 3 * (rand3() - 1) + 3 * (rand3() - 1) + (rand3() - 1)，最后生成了 rand11()；
//         用 rand7() 生成 rand9()，可以先生成 rand49()，然后变成以 45 为界限，因为 45 是 9 的倍数。生成 rand49() 的方式：(rand7() - 1) * 7 + (rand7() - 1)，最后生成了 rand9()；
//         用 rand6() 生成 rand13()，可以先生成 rand36()，然后变成以 26 为界限，因为 26 是 13 的倍数。生成 rand36() 的方式：(rand6() - 1) * 6 + (rand6() - 1)，最后生成了 rand13()；

import "fmt"
import "math/rand"

func rand7() int {
    return rand.Intn(7)
}

func rand10() int {
    res := 49
    for res > 40 {
        res = (rand7() - 1) * 7 + rand7()
    }
    return res % 10 + 1
}

func rand101() int {
    res := 10
    for res >= 10 {
        res = (rand7() - 1) + rand7()
    }
    return res % 10 + 1
}

func rand102() int {
    return ((rand7()+rand7())%7 + (rand7()+rand7())%7 + (rand7()+rand7())%7 + (rand7()+rand7())%7 + (rand7()+rand7())%7)%10+1
}


func main() {
    fmt.Println(rand10())
    fmt.Println(rand101())
    fmt.Println(rand102())
}