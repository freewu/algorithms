package main

// 1538. Guess the Majority in a Hidden Array
// We have an integer array nums, where all the integers in nums are 0 or 1. 
// You will not be given direct access to the array, instead, you will have an API ArrayReader which have the following functions:
//     int query(int a, int b, int c, int d): 
//         where 0 <= a < b < c < d < ArrayReader.length(). 
//         The function returns the distribution of the value of the 4 elements and returns:
//             4 : if the values of the 4 elements are the same (0 or 1).
//             2 : if three elements have a value equal to 0 and one element has value equal to 1 or vice versa.
//             0 : if two element have a value equal to 0 and two elements have a value equal to 1.
//     int length(): Returns the size of the array.

// You are allowed to call query() 2 * n times at most where n is equal to ArrayReader.length().
// Return any index of the most frequent value in nums, in case of tie, return -1.

// Example 1:
// Input: nums = [0,0,1,0,1,1,1,1]
// Output: 5
// Explanation: The following calls to the API
// reader.length() // returns 8 because there are 8 elements in the hidden array.
// reader.query(0,1,2,3) // returns 2 this is a query that compares the elements nums[0], nums[1], nums[2], nums[3]
// // Three elements have a value equal to 0 and one element has value equal to 1 or viceversa.
// reader.query(4,5,6,7) // returns 4 because nums[4], nums[5], nums[6], nums[7] have the same value.
// we can infer that the most frequent value is found in the last 4 elements.
// Index 2, 4, 6, 7 is also a correct answer.

// Example 2:
// Input: nums = [0,0,1,1,0]
// Output: 0

// Example 3:
// Input: nums = [1,0,1,0,1,0,1,0]
// Output: -1
 
// Constraints:
//     5 <= nums.length <= 10^5
//     0 <= nums[i] <= 1

// Follow up: What is the minimum number of calls needed to find the majority element?

import "fmt"

/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 * // Compares 4 different elements in the array
 * // return 4 if the values of the 4 elements are the same (0 or 1).
 * // return 2 if three elements have a value equal to 0 and one element has value equal to 1 or vice versa.
 * // return 0 : if two element have a value equal to 0 and two elements have a value equal to 1.
 * func (this *ArrayReader) query(a, b, c, d int) int {}
 * 
 * // Returns the length of the array
 * func (this *ArrayReader) length() int {}
 */

func guessMajority(reader *ArrayReader) int {
    n := reader.length() //获取长度
    num := reader.query(0,1,2,3) // 获取前四个的值
    count := 1 // 计数，这里我是计算3号位置的数字的个数
    flag := -1 // 记录不等于三号位数字的下标
    for i := 4; i < n; i++ {
        if num == reader.query(0,1,2,i) { // 判断第 4 个数字之后的值是否和第三个数字相等，相等则count + 1，
            count++
        } else { // 否则记录该元素下标到 flag 中
            flag = i
        }
    }
    if num == 0  { // 如果 num 等于 0，则证明 0，1，2 三个位置还有个和 3 号数字一样，则 count + 1
        count++
    }
    if num == 4 { // 如果 num 等于 4 ，则证明 0，1，2 三个位置都等于 3 号位的数字，则 count + 3
        count += 3
    } 
    if num == 2 { // 如果等于2，则证明0，1，2可能有两个和3号位数字相等，也有可能没有
        if flag != -1 { // 如果 flag不为 -1，则证明后面出现过和3号位数字不等的数字，
            // 查询0，1，2，flag，如果为 4，则证明和 0，1，2,flag相同,则count不加，否则count + 2
            if reader.query(0,1,2,flag) != 4 {
                count += 2
            }
        } else {
            // 如果flag为-1，证明3号及之后的数字全相同，因为至少有五个，因此判断0，1，3，4号位置，如果大于0
            // 则证明 0，1，2还有两个和3号位相等，直接返回3号位
            if reader.query(0,1,3,4) > 0 {
                return 3
            }
            // 否则证明0，1，2和3号位及后面的不相等，则直接用长度来判断
            // n = 6,说明各有3个，小于6，则说明0，1，2有3个，返回0，大于6，证明和3号位相等的超过3个，返回 3
            if n == 6 {  return -1; }
            if n < 6 { return 0; }
            return 3
        }
    }
    // 如果count * 2大于n了，则证明超过了一半，返回3，等于返回-1，小于返回flag
    if count * 2 > n {
        return 3;   
    } else if count * 2 == n {
        return -1
    }
    // 解释为啥这里flag不可能为初始值-1，因为只有当num等于0或4才能走到这里，此时这里count必定大于等于3个
    //如果后面没有出现与3不同的数，是不可能返回flag的。
    return flag
}

func main() {
    fmt.Println()
}