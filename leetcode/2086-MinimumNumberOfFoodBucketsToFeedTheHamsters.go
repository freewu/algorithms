package main

// 2086. Minimum Number of Food Buckets to Feed the Hamsters
// You are given a 0-indexed string hamsters where hamsters[i] is either:
//     'H' indicating that there is a hamster at index i, or
//     '.' indicating that index i is empty.

// You will add some number of food buckets at the empty indices in order to feed the hamsters. 
// A hamster can be fed if there is at least one food bucket to its left or to its right. 
// More formally, a hamster at index i can be fed if you place a food bucket at index i - 1 and/or at index i + 1.

// Return the minimum number of food buckets you should place at empty indices to feed all the hamsters or -1 if it is impossible to feed all of them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/11/01/example1.png" />
// Input: hamsters = "H..H"
// Output: 2
// Explanation: We place two food buckets at indices 1 and 2.
// It can be shown that if we place only one food bucket, one of the hamsters will not be fed.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/11/01/example2.png" />
// Input: hamsters = ".H.H."
// Output: 1
// Explanation: We place one food bucket at index 2.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/11/01/example3.png" />
// Input: hamsters = ".HHH."
// Output: -1
// Explanation: If we place a food bucket at every empty index as shown, the hamster at index 2 will not be able to eat.

// Constraints:
//     1 <= hamsters.length <= 10^5
//     hamsters[i] is either'H' or '.'.

import "fmt"

// greedy 
func minimumBuckets(hamsters string) int {
    res, i, n := 0, 0, len(hamsters)
    for i < n {
        if hamsters[i] == '.' {
            i++
            continue
        } else if i < n - 1 && hamsters[i + 1] == '.' {
            res++
            i += 3
        } else if i > 0 && hamsters[i - 1] == '.' {
            res++
            i++
        } else {
            return -1
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/11/01/example1.png" />
    // Input: hamsters = "H..H"
    // Output: 2
    // Explanation: We place two food buckets at indices 1 and 2.
    // It can be shown that if we place only one food bucket, one of the hamsters will not be fed.
    fmt.Println(minimumBuckets("H..H")) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/11/01/example2.png" />
    // Input: hamsters = ".H.H."
    // Output: 1
    // Explanation: We place one food bucket at index 2.
    fmt.Println(minimumBuckets(".H.H.")) // 1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/11/01/example3.png" />
    // Input: hamsters = ".HHH."
    // Output: -1
    // Explanation: If we place a food bucket at every empty index as shown, the hamster at index 2 will not be able to eat.
    fmt.Println(minimumBuckets(".HHH.")) // -1
}