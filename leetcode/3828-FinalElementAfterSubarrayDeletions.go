package main

// 3828. Final Element After Subarray Deletions
// You are given an integer array nums.

// Two players, Alice and Bob, play a game in turns, with Alice playing first.
//     1. In each turn, the current player chooses any subarray nums[l..r] such that r - l + 1 < m, where m is the current length of the array.
//     2. The selected subarray is removed, and the remaining elements are concatenated to form the new array.
//     3. The game continues until only one element remains.

// Alice aims to maximize the final element, while Bob aims to minimize it. 
// Assuming both play optimally, return the value of the final remaining element.

// Example 1:
// Input: nums = [1,5,2]
// Output: 2
// Explanation:
// One valid optimal strategy:
// Alice removes [1], array becomes [5, 2].
// Bob removes [5], array becomes [2]​​​​​​​. Thus, the answer is 2.

// Example 2:
// Input: nums = [3,7]
// Output: 7
// Explanation:
// Alice removes [3], leaving the array [7]. Since Bob cannot play a turn now, the answer is 7.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func finalElement(nums []int) int {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(nums[0], nums[len(nums)-1])
}

func main() {
    // Example 1:
    // Input: nums = [1,5,2]
    // Output: 2
    // Explanation:
    // One valid optimal strategy:
    // Alice removes [1], array becomes [5, 2].
    // Bob removes [5], array becomes [2]​​​​​​​. Thus, the answer is 2.
    fmt.Println(finalElement([]int{1,5,2})) // 2
    // Example 2:
    // Input: nums = [3,7]
    // Output: 7
    // Explanation:
    // Alice removes [3], leaving the array [7]. Since Bob cannot play a turn now, the answer is 7.
    fmt.Println(finalElement([]int{3,7})) // 7

    fmt.Println(finalElement([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(finalElement([]int{9,8,7,6,5,4,3,2,1})) // 9
}