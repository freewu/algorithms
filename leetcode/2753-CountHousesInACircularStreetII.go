package main

// 2753. Count Houses in a Circular Street II
// You are given an object street of class Street that represents a circular street 
// and a positive integer k which represents a maximum bound for the number of houses in that street 
// (in other words, the number of houses is less than or equal to k). 
// Houses' doors could be open or closed initially (at least one is open).

// Initially, you are standing in front of a door to a house on this street. 
// Your task is to count the number of houses in the street.

// The class Street contains the following functions which may help you:
//     void closeDoor(): Close the door of the house you are in front of.
//     boolean isDoorOpen(): Returns true if the door of the current house is open and false otherwise.
//     void moveRight(): Move to the right house.

// Note that by circular street, we mean if you number the houses from 1 to n, 
// then the right house of housei is housei+1 for i < n, and the right house of housen is house1.

// Return ans which represents the number of houses on this street.

// Example 1:
// Input: street = [1,1,1,1], k = 10
// Output: 4
// Explanation: There are 4 houses, and all their doors are open. 
// The number of houses is less than k, which is 10.

// Example 2:
// Input: street = [1,0,1,1,0], k = 5
// Output: 5
// Explanation: There are 5 houses, and the doors of the 1st, 3rd, and 4th house (moving in the right direction) are open, and the rest are closed.
// The number of houses is equal to k, which is 5.

// Constraints:
//     n == number of houses
//     1 <= n <= k <= 10^5
//     street is circular by definition provided in the statement.
//     The input is generated such that at least one of the doors is open.

import "fmt"

type Street interface {
    OpenDoor()
    CloseDoor()
    IsDoorOpen() bool
    MoveRight()
    MoveLeft()
}

type StreetService struct {
    data []int
}

func NewStreet(arr []int) Street {
    return Street{ data: arr}
}

/**
 * Definition for a street.
 * type Street interface {
 *     CloseDoor()
 *     IsDoorOpen() bool
 *     MoveRight()
 * }
 */
func houseCount(street Street, k int) int {
    res := 0
    for !street.IsDoorOpen() {
        street.MoveRight()
    }
    for i := 1; i <= k; i++ {
        street.MoveRight()
        if street.IsDoorOpen() {
            res = i
            street.CloseDoor()
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: street = [1,1,1,1], k = 10
    // Output: 4
    // Explanation: There are 4 houses, and all their doors are open. 
    // The number of houses is less than k, which is 10.

    // Example 2:
    // Input: street = [1,0,1,1,0], k = 5
    // Output: 5
    // Explanation: There are 5 houses, and the doors of the 1st, 3rd, and 4th house (moving in the right direction) are open, and the rest are closed.
    // The number of houses is equal to k, which is 5.
    fmt.Println()
}