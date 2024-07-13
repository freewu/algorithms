package main

// LCR 148. 验证图书取出顺序
// 现在图书馆有一堆图书需要放入书架，并且图书馆的书架是一种特殊的数据结构，只能按照 一定 的顺序 放入 和 拿取 书籍。
// 给定一个表示图书放入顺序的整数序列 putIn，请判断序列 takeOut 是否为按照正确的顺序拿取书籍的操作序列。你可以假设放入书架的所有书籍编号都不相同。

// 示例 1：
// 输入：putIn = [6,7,8,9,10,11], takeOut = [9,11,10,8,7,6]
// 输出：true
// 解释：我们可以按以下操作放入并拿取书籍：
// push(6), push(7), push(8), push(9), pop() -> 9,
// push(10), push(11),pop() -> 11,pop() -> 10, pop() -> 8, pop() -> 7, pop() -> 6

// 示例 2：
// 输入：putIn = [6,7,8,9,10,11], takeOut = [11,9,8,10,6,7]
// 输出：false
// 解释：6 不能在 7 之前取出。
 
// 提示：
//     0 <= putIn.length == takeOut.length <= 1000
//     0 <= putIn[i], takeOut < 1000
//     putIn 是 takeOut 的排列。

import "fmt"

func validateBookSequences(putIn []int, takeOut []int) bool {
    stack, x := []int{}, 0
    for _, val := range putIn {
        stack = append(stack, val)
        for len(stack) > 0 && stack[ len(stack) - 1] == takeOut[x] {
            x++
            stack = stack[:len(stack)-1] // pop operation
        }
    }
    return len(stack) == 0
}

func validateBookSequences1(putIn []int, takeOut []int) bool {
    // 合法的栈操作顺序
    // 在拿出一本书的时候，需要判断栈里有没有该书
    // 如果没有，需要继续入栈
    // 如果有，需要立刻出栈
    // 如果栈顶元素不等于需要出栈的元素，则无效
    // 先暂时用map来查找是否存在
    t, stack := make(map[int]bool), make([]int, 0)
    index := 0 // 标记下一个需要入栈的元素
    for _, x := range takeOut {
        for !t[x] { // 没有这本书，需要入栈
            t[putIn[index]] = true
            stack = append(stack, putIn[index])
            index++
        }
        if stack[len(stack)-1] != x {
            return false
        }
        stack = stack[:len(stack)-1]
    }
    return true
}

func main() {
    // Example 1:
    // Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
    // Output: true
    // Explanation: We might do the following sequence:
    // push(1), push(2), push(3), push(4),
    // pop() -> 4,
    // push(5),
    // pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
    fmt.Println(validateBookSequences([]int{1,2,3,4,5}, []int{4,5,3,2,1})) // true
    // Example 2:
    // Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
    // Output: false
    // Explanation: 1 cannot be popped before 2.
    fmt.Println(validateBookSequences([]int{1,2,3,4,5}, []int{4,3,5,1,2})) // false

    fmt.Println(validateBookSequences([]int{6,7,8,9,10,11}, []int{9,11,10,8,7,6})) // true
    fmt.Println(validateBookSequences([]int{6,7,8,9,10,11}, []int{11,9,8,10,6,7})) // false


    fmt.Println(validateBookSequences1([]int{1,2,3,4,5}, []int{4,5,3,2,1})) // true
    fmt.Println(validateBookSequences1([]int{1,2,3,4,5}, []int{4,3,5,1,2})) // false
    fmt.Println(validateBookSequences1([]int{6,7,8,9,10,11}, []int{9,11,10,8,7,6})) // true
    fmt.Println(validateBookSequences1([]int{6,7,8,9,10,11}, []int{11,9,8,10,6,7})) // false
}