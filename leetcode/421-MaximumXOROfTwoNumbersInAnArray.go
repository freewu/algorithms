package main

// 421. Maximum XOR of Two Numbers in an Array
// Given an integer array nums, return the maximum result of nums[i] XOR nums[j], where 0 <= i <= j < n.

// Example 1:
// Input: nums = [3,10,5,25,2,8]
// Output: 28
// Explanation: The maximum result is 5 XOR 25 = 28.

// Example 2:
// Input: nums = [14,70,53,83,49,91,36,80,92,51,66,70]
// Output: 127
 
// Constraints:
//     1 <= nums.length <= 2 * 10^5
//     0 <= nums[i] <= 2^31 - 1

import "fmt"
import "math/bits"
import "slices"

func findMaximumXOR(nums []int) int {
    res, mask := 0, 0
    /*The res is a record of the largest XOR we got so far. if it's 11100 at i = 2, it means
        before we reach the last two bits, 11100 is the biggest XOR we have, and we're going to explore
        whether we can get another two '1's and put them into res

        This is a greedy part, since we're looking for the largest XOR, we start
        from the very begining, aka, the 31st postition of bits. */
    for i := 31; i >= 0; i-- {
        //The mask will grow like  100..000 , 110..000, 111..000,  then 1111...111
        //for each iteration, we only care about the left parts
        mask = mask | (1 << uint(i))
        m := make(map[int]bool)
        for _, num := range nums {
            /* num&mask: we only care about the left parts, for example, if i = 2, then we have
            {1100, 1000, 0100, 0000} from {1110, 1011, 0111, 0010}*/
            m[num & mask] = true
        }
        // if i = 1 and before this iteration, the res we have now is 1100,
        // my wish is the res will grow to 1110, so I will try to find a candidate
        // which can give me the greedyTry;
        greedyTry := res | (1 << uint(i))
        for anotherNum := range m {
            //This is the most tricky part, coming from a fact that if a ^ b = c, then a ^ c = b;
            // now we have the 'c', which is greedyTry, and we have the 'a', which is leftPartOfNum
            // If we hope the formula a ^ b = c to be valid, then we need the b,
            // and to get b, we need a ^ c, if a ^ c exisited in our set, then we're good to go
            if m[anotherNum ^ greedyTry] == true {
                res = greedyTry
                break
            }
        }
        // If unfortunately, we didn't get the greedyTry, we still have our max,
        // So after this iteration, the max will stay at 1100.
    }
    return res
}

func findMaximumXOR1(nums []int) int {
    res, mask, set := 0, 0, make(map[int]struct{})
    highBit := bits.Len(uint(slices.Max(nums))) - 1
    for i := highBit; i >= 0; i-- {
        mask |= 1 << i
        tmp := res | 1 << i
        for _, v := range nums {
            v = v & mask
            if _, ok := set[tmp ^ v]; ok {
                res = tmp
                break
            }
            set[v] = struct{}{}
        }
        clear(set)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,10,5,25,2,8]
    // Output: 28
    // Explanation: The maximum result is 5 XOR 25 = 28.
    fmt.Println(findMaximumXOR([]int{3,10,5,25,2,8})) // 28
    // Example 2:
    // Input: nums = [14,70,53,83,49,91,36,80,92,51,66,70]
    // Output: 127
    fmt.Println(findMaximumXOR([]int{14,70,53,83,49,91,36,80,92,51,66,70})) // 127

    fmt.Println(findMaximumXOR1([]int{3,10,5,25,2,8})) // 28
    fmt.Println(findMaximumXOR1([]int{14,70,53,83,49,91,36,80,92,51,66,70})) // 127
}