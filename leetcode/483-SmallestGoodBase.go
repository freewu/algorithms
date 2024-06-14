package main

// 483. Smallest Good Base
// Given an integer n represented as a string, return the smallest good base of n.
// We call k >= 2 a good base of n, if all digits of n base k are 1's.

// Example 1:
// Input: n = "13"
// Output: "3"
// Explanation: 13 base 3 is 111.

// Example 2:
// Input: n = "4681"
// Output: "8"
// Explanation: 4681 base 8 is 11111.

// Example 3:
// Input: n = "1000000000000000000"
// Output: "999999999999999999"
// Explanation: 1000000000000000000 base 999999999999999999 is 11.
 
// Constraints:
//     n is an integer in the range [3, 10^18].
//     n does not contain any leading zeros.

import "fmt"
import "strconv"
import "math"

func smallestGoodBase(n string) string {
    num, _ := strconv.ParseInt(n, 10, 64)
    binarySearch := func(num int64, length int) int64 {
        l, r := int64(2), int64(math.Pow(float64(num), 1.0/ float64(length))) + 1
        // note, we need to convert length to float64, otherwise, 1.0/int64(n)=0 if n>1
        for l < r {
            sum, cur, mid := int64(1), int64(1), l + (r-l) / 2
            for i := 1; i <= length; i++ {
                cur *= mid
                sum += cur
            }
            if sum == num {
                return mid
            }
            if sum > num {
                r = mid
            } else {
                l = mid + 1
            }
        }
        return -1
    }
    //the target is to get the longest 1111111....
    //and the longest 1... length is 64, the shortest should be 2, since the range of n is [3,10^18]
    x := int64(1)
    for i := 62; i >= 1; i-- {
        if (x << uint(i)) < num {
            //make sure that the 2 base number is smaller than num
            //now e.g. 10000<num, then we check if there is a number 11111 (base from 2 to num^(1/5)) is equal to num
            if val := binarySearch(num, i); val != -1 {
                return strconv.FormatInt(val, 10)
            }
        }
    }
    return strconv.FormatInt(num - 1, 10)
}

func main() {
    // Example 1:
    // Input: n = "13"
    // Output: "3"
    // Explanation: 13 base 3 is 111.
    fmt.Println(smallestGoodBase("13")) // "3"
    // Example 2:
    // Input: n = "4681"
    // Output: "8"
    // Explanation: 4681 base 8 is 11111.
    fmt.Println(smallestGoodBase("4681")) // "8"
    // Example 3:
    // Input: n = "1000000000000000000"
    // Output: "999999999999999999"
    // Explanation: 1000000000000000000 base 999999999999999999 is 11.
    fmt.Println(smallestGoodBase("1000000000000000000")) // "999999999999999999"
}