package main

// 2231. Largest Number After Digit Swaps by Parity
// You are given a positive integer num. 
// You may swap any two digits of num that have the same parity (i.e. both odd digits or both even digits).

// Return the largest possible value of num after any number of swaps.

// Example 1:
// Input: num = 1234
// Output: 3412
// Explanation: Swap the digit 3 with the digit 1, this results in the number 3214.
// Swap the digit 2 with the digit 4, this results in the number 3412.
// Note that there may be other sequences of swaps but it can be shown that 3412 is the largest possible number.
// Also note that we may not swap the digit 4 with the digit 1 since they are of different parities.

// Example 2:
// Input: num = 65875
// Output: 87655
// Explanation: Swap the digit 8 with the digit 6, this results in the number 85675.
// Swap the first digit 5 with the digit 7, this results in the number 87655.
// Note that there may be other sequences of swaps but it can be shown that 87655 is the largest possible number.

// Constraints:
//     1 <= num <= 10^9

import "fmt"
import "strconv"
import "strings"
import "sort"
import "math"

func largestInteger(num int) int {
    sums := strconv.Itoa(num)
    largestNum := num
    swapNums := func (arr string, a,b int) int {
        curr := strings.Split(arr, "") // turn into a []string
        curr[a],curr[b]=curr[b],curr[a] // peform swap
        curr_num,_ := strconv.Atoi(strings.Join(curr,""))
        return curr_num
    }
    for i := 0; i < len(sums); i++ {
        in, _ := strconv.Atoi(string(sums[i]))
        for j := i + 1; j < len(sums); j++ {
            jn, _ := strconv.Atoi(string(sums[j]))
            // if both even or both odd
            if (in % 2 == 0 && jn % 2 == 0) ||  (in % 2 == 1 && jn % 2 == 1){
                curr := swapNums(sums,i,j)
                if curr > largestNum {
                    largestNum = curr
                    sums = strconv.Itoa(largestNum)
                }
            }
        }
    }
    return largestNum
}

func largestInteger1(num int) int {
    // 奇偶都进行排序，倒序
    str := strconv.Itoa(num)
    queue, odd, even := []int{}, []int{}, []int{}
    for i := 0; i < len(str); i++ {
        v, _ := strconv.Atoi(string(str[i]))
        if v % 2 == 0 {
            even = append(even, i)
        } else {
            odd = append(odd, i)
        }
        queue = append(queue, v)
    }
    sort.Slice(odd, func(i, j int) bool {
        return queue[odd[i]] > queue[odd[j]]
    })
    sort.Slice(even, func(i, j int) bool {
        return queue[even[i]] > queue[even[j]]
    })
    res, oddLoc, evenLoc := 0, 0, 0
    for i := 0; i < len(queue); i++ {
        advance := int(math.Pow(float64(10), float64(len(queue) - i - 1)))
        if queue[i] % 2 == 0 { // 偶数
            res += queue[even[evenLoc]] * advance
            evenLoc++
        } else { // 奇数
            res += queue[odd[oddLoc]] * advance
            oddLoc++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: num = 1234
    // Output: 3412
    // Explanation: Swap the digit 3 with the digit 1, this results in the number 3214.
    // Swap the digit 2 with the digit 4, this results in the number 3412.
    // Note that there may be other sequences of swaps but it can be shown that 3412 is the largest possible number.
    // Also note that we may not swap the digit 4 with the digit 1 since they are of different parities.
    fmt.Println(largestInteger(1234)) // 3412
    // Example 2:
    // Input: num = 65875
    // Output: 87655
    // Explanation: Swap the digit 8 with the digit 6, this results in the number 85675.
    // Swap the first digit 5 with the digit 7, this results in the number 87655.
    // Note that there may be other sequences of swaps but it can be shown that 87655 is the largest possible number.
    fmt.Println(largestInteger(65875)) // 87655

    fmt.Println(largestInteger1(1234)) // 3412
    fmt.Println(largestInteger1(65875)) // 87655
}