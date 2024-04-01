package main

// 278. First Bad Version
// You are a product manager and currently leading a team to develop a new product. 
// Unfortunately, the latest version of your product fails the quality check.
// Since each version is developed based on the previous version, all the versions after a bad version are also bad.

// Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

// You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. 
// You should minimize the number of calls to the API.

// Example 1:
// Input: n = 5, bad = 4
// Output: 4
// Explanation:
// call isBadVersion(3) -> false
// call isBadVersion(5) -> true
// call isBadVersion(4) -> true
// Then 4 is the first bad version.

// Example 2:
// Input: n = 1, bad = 1
// Output: 1
 
// Constraints:
//     1 <= bad <= n <= 2^31 - 1

import "fmt"

var bad int

func isBadVersion(version int) bool {
    return version == bad
}

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {
    left, right := 1, n
    for left + 1 < right {
        mid := (left + right) / 2
        if isBadVersion(mid) {
            right = mid
        } else {
            left = mid
        }
    }
    if isBadVersion(left) {
        return left
    }
    return right
}

func main() {
    // n = 5, bad = 4
    // call isBadVersion(3) -> false
    // call isBadVersion(5) -> true
    // call isBadVersion(4) -> true
    // Then 4 is the first bad version.
    bad = 4
    fmt.Println(firstBadVersion(5)) // 4

    // Input: n = 1, bad = 1
    // Output: 1
    bad = 1
    fmt.Println(firstBadVersion(1)) // 1

}