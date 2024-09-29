package main

// 1386. Cinema Seat Allocation
// <img src="https://assets.leetcode.com/uploads/2020/02/14/cinema_seats_1.png" />
// A cinema has n rows of seats, numbered from 1 to n and there are ten seats in each row, labelled from 1 to 10 as shown in the figure above.

// Given the array reservedSeats containing the numbers of seats already reserved, 
// for example, reservedSeats[i] = [3,8] means the seat located in row 3 and labelled with 8 is already reserved.

// Return the maximum number of four-person groups you can assign on the cinema seats. 
// A four-person group occupies four adjacent seats in one single row. 
// Seats across an aisle (such as [3,3] and [3,4]) are not considered to be adjacent, but there is an exceptional case on which an aisle split a four-person group, 
// in that case, the aisle split a four-person group in the middle, which means to have two people on each side.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/02/14/cinema_seats_3.png" />
// Input: n = 3, reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
// Output: 4
// Explanation: The figure above shows the optimal allocation for four groups, where seats mark with blue are already reserved and contiguous seats mark with orange are for one group.

// Example 2:
// Input: n = 2, reservedSeats = [[2,1],[1,8],[2,6]]
// Output: 2

// Example 3:
// Input: n = 4, reservedSeats = [[4,3],[1,4],[4,6],[1,7]]
// Output: 4

// Constraints:
//     1 <= n <= 10^9
//     1 <= reservedSeats.length <= min(10*n, 10^4)
//     reservedSeats[i].length == 2
//     1 <= reservedSeats[i][0] <= n
//     1 <= reservedSeats[i][1] <= 10
//     All reservedSeats[i] are distinct.

import "fmt"

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
    reserved := make(map[int]int)
    for _, seat := range reservedSeats {
        reserved[seat[0]-1] |= 1 << (seat[1] - 1)
    }
    res := 0
    bms := []int{ 2 + 4 + 8 + 16, 32 + 64 + 128 + 256, 8 + 16 + 32 + 64 }
    for _, seats := range reserved {
        count := 0
        if seats & bms[0] == 0 {
            count++
        }
        if seats & bms[1] == 0 {
            count++
        }
        if (seats & bms[2] == 0) && count == 0 {
            count++
        }
        res += count
    }
    return res + (n - len(reserved)) * 2
}

func maxNumberOfFamilies1(n int, reservedSeats [][]int) int {
    mp := map[int]int{}
    for _, v := range reservedSeats {
        row, col := v[0], v[1]
        if col == 1 || col == 10 { // 每行10个位置
            continue
        }
        mp[row] += 1 << (9 - col) // 用二进制计算
    }
    res := n * 2
    for _, v := range mp {
        if v & 0xF0 == 0 || v & 0x3C == 0 || v & 0x0F == 0 { // 2-9 任何一个位置被占，只能有一个了
            res--
        } else {
            res -= 2
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/02/14/cinema_seats_3.png" />
    // Input: n = 3, reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
    // Output: 4
    // Explanation: The figure above shows the optimal allocation for four groups, where seats mark with blue are already reserved and contiguous seats mark with orange are for one group.
    fmt.Println(maxNumberOfFamilies(3, [][]int{{1,2},{1,3},{1,8},{2,6},{3,1},{3,10}})) // 4
    // Example 2:
    // Input: n = 2, reservedSeats = [[2,1],[1,8],[2,6]]
    // Output: 2
    fmt.Println(maxNumberOfFamilies(2, [][]int{{2,1},{1,8},{2,6}})) // 2
    // Example 3:
    // Input: n = 4, reservedSeats = [[4,3],[1,4],[4,6],[1,7]]
    // Output: 4
    fmt.Println(maxNumberOfFamilies(4, [][]int{{4,3},{1,4},{4,6},{1,7}})) // 4

    fmt.Println(maxNumberOfFamilies1(3, [][]int{{1,2},{1,3},{1,8},{2,6},{3,1},{3,10}})) // 4
    fmt.Println(maxNumberOfFamilies1(2, [][]int{{2,1},{1,8},{2,6}})) // 2
    fmt.Println(maxNumberOfFamilies1(4, [][]int{{4,3},{1,4},{4,6},{1,7}})) // 4
}