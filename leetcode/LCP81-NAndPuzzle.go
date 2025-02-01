package main

// LCP 81. 与非的谜题
// 在永恒之森中，封存着有关万灵之树线索的卷轴，只要探险队通过最后的考验，便可以获取前往万灵之树的线索。

// 探险队需要从一段不断变化的谜题数组中找到最终的密码，初始的谜题为长度为 n 的数组 arr（下标从 0 开始），数组中的数字代表了 k 位二进制数。 
// 破解谜题的过程中，需要使用 与非（NAND） 运算方式，operations[i] = [type,x,y] 表示第 i 次进行的谜题操作信息：
//     1. 若 type = 0，表示修改操作，将谜题数组中下标 x 的数字变化为 y；
//     2. 若 type = 1，表示运算操作，将数字 y 进行 x*n 次「与非」操作，第 i 次与非操作为 y = y NAND arr[i%n]；

// 运算操作结果即：y NAND arr[0%n] NAND arr[1%n] NAND arr[2%n] ... NAND arr[(x*n-1)%n]

// 最后，将所有运算操作的结果按顺序逐一进行 异或（XOR）运算，从而得到最终解开封印的密码。请返回最终解开封印的密码。

// 注意:
//     「与非」（NAND）的操作为：先进行 与 操作，后进行 非 操作。
//     例如：两个三位二进制数2和3，其与非结果为 NOT ((010) AND (011)) = (101) = 5

// 示例 1：
// 输入: k = 3 arr = [1,2] operations = [[1,2,3],[0,0,3],[1,2,2]]
// 输出: 2
// 解释： 初始的谜题数组为 [1,2]，二进制位数为 3， 第 0 次进行运算操作，将数字 3(011) 进行 2*2 次「与非」运算， 运算操作结果为 3 NAND 1 NAND 2 NAND 1 NAND 2 = 5 第 1 次进行修改操作，谜题数组的第 0 个数字变化为 3，谜题变成 [3,2] 第 2 次进行运算操作，将数字 2(010) 进行 2*2 次「与非」运算， 运算操作结果为 2 NAND 3 NAND 2 NAND 3 NAND 2 = 7 所有运算操作结果进行「异或」运算为 5 XOR 7 = 2 因此得到的最终密码为 2。

// 示例 2：
// 输入: k = 4 arr = [4,6,4,7,10,9,11] operations = [[1,5,7],[1,7,14],[0,6,7],[1,6,5]] 
// 输出: 9 
// 解释: 初始的谜题数组为 [4,6,4,7,10,9,11], 第 0 次进行运算操作，运算操作结果为 5； 第 1 次进行运算操作，运算操作结果为 5； 第 2 次进行修改操作，修改后谜题数组为 [4, 6, 4, 7, 10, 9, 7]； 第 3 次进行运算操作，运算操作结果为 9； 所有运算操作结果进行「异或」运算为 5 XOR 5 XOR 9 = 9； 因此得到的最终密码为 9。

// 提示:
//     1 <= arr.length, operations.length <= 10^4
//     1 <= k <= 30
//     0 <= arr[i] < 2^k
//     若 type = 0，0 <= x < arr.length 且 0 <= y < 2^k
//     若 type = 1，1 <= x < 10^9 且 0 <= y < 2^k
//     保证存在 type = 1 的操作

import "fmt"

type Segment []struct {
    l, r int
    to   [2]int
}

func (t Segment) set(o, val int) {
	t[o].to[1] = t[o].to[0] ^ val
}

func (t Segment) maintain(o, k int) {
    a, b, c := t[o << 1].to, t[o << 1|1].to, [2]int{}
    for i := 0; i < k; i++ {
        c[0] |= b[a[0]>>i&1] >> i & 1 << i
        c[1] |= b[a[1]>>i&1] >> i & 1 << i
    }
    t[o].to = c
}

func (t Segment) build(a []int, k, o, l, r int) {
    t[o].l, t[o].r = l, r
    t[o].to[0] = 1 << k - 1
    if l == r {
        t.set(o, a[l-1])
        return
    }
    m := (l + r) >> 1
    t.build(a, k, o<<1, l, m)
    t.build(a, k, o<<1|1, m + 1, r)
    t.maintain(o, k)
}

func (t Segment) update(o, i, val, k int) {
    if t[o].l == t[o].r {
        t.set(o, val)
        return
    }
    m := (t[o].l + t[o].r) >> 1
    if i <= m {
        t.update(o << 1, i, val, k)
    } else {
        t.update(o << 1 | 1, i, val, k)
    }
    t.maintain(o, k)
}

func getNandResult(k int, arr []int, operations [][]int) int {
    res := 0
    t := make(Segment, len(arr) * 4)
    t.build(arr, k, 1, 1, len(arr))
    for _, op := range operations {
        if op[0] == 0 {
            t.update(1, op[1]+1, op[2], k)
            continue
        }
        x, y, to := op[1], op[2], t[1].to
        for i := 0; i < k; i++ {
            v := 0
            y := y >> i & 1
            y1 := to[y] >> i & 1 // 穿过 arr 一次
            if y1 == y { // 不变
                v = y
            } else if x == 1 || to[y1]>>i&1 == y1 {
                // 只穿过一次，或者穿过两次和穿过一次相同
                v = y1
            } else {
                v = y ^ x%2 // 奇变偶不变
            }
            res ^= v << i
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入: k = 3 arr = [1,2] operations = [[1,2,3],[0,0,3],[1,2,2]]
    // 输出: 2
    // 解释： 初始的谜题数组为 [1,2]，二进制位数为 3， 第 0 次进行运算操作，将数字 3(011) 进行 2*2 次「与非」运算， 运算操作结果为 3 NAND 1 NAND 2 NAND 1 NAND 2 = 5 第 1 次进行修改操作，谜题数组的第 0 个数字变化为 3，谜题变成 [3,2] 第 2 次进行运算操作，将数字 2(010) 进行 2*2 次「与非」运算， 运算操作结果为 2 NAND 3 NAND 2 NAND 3 NAND 2 = 7 所有运算操作结果进行「异或」运算为 5 XOR 7 = 2 因此得到的最终密码为 2。
    fmt.Println(getNandResult(3, []int{1,2}, [][]int{{1,2,3},{0,0,3},{1,2,2}})) // 2
    // 示例 2：
    // 输入: k = 4 arr = [4,6,4,7,10,9,11] operations = [[1,5,7],[1,7,14],[0,6,7],[1,6,5]] 
    // 输出: 9 
    // 解释: 初始的谜题数组为 [4,6,4,7,10,9,11], 第 0 次进行运算操作，运算操作结果为 5； 第 1 次进行运算操作，运算操作结果为 5； 第 2 次进行修改操作，修改后谜题数组为 [4, 6, 4, 7, 10, 9, 7]； 第 3 次进行运算操作，运算操作结果为 9； 所有运算操作结果进行「异或」运算为 5 XOR 5 XOR 9 = 9； 因此得到的最终密码为 9。
    fmt.Println(getNandResult(4, []int{4,6,4,7,10,9,11}, [][]int{{1,5,7},{1,7,14},{0,6,7},{1,6,5}} )) // 9
}