package main

// 4116. Assign Umbrella on Rainy Days
// You are given a binary array rain of length n, 
// where rain[i] == 1 means it rains on day i, and rain[i] == 0 means it does not.

// Alice and Bob share a single umbrella over these n days. On each day i:
//     1. Alice needs the umbrella on day i if rain[i] == rain[i + 1] == 1 for 0 <= i < n - 1.
//     2. Bob needs the umbrella if rain[i] == 1.
//     3. At most one of them can receive the umbrella.
//     4. If only one person needs it, that person receives it.
//     5. If both need it, it goes to whoever has received it fewer times so far. If they are tied, it goes to Alice.
//     6. If neither needs it, no one receives it.

// A person's count increases only on a day they receive the umbrella. 
// Initially, both counts are 0.

// Return a string array answer of length n, where answer[i] is "Alice", "Bob", or "None", denoting who receives the umbrella on day i.
 
// A binary array is an array that contains only 0 and 1.

// Example 1:
// Input: rain = [1,1,0,1,1]
// Output: ["Alice","Bob","None","Alice","Bob"]
// Explanation:
//     Day 0: rain[0] == 1 and rain[1] == 1, so both need it. Counts are tied, so Alice receives it.
//     Day 1: rain[1] == 1, so Bob needs it, but rain[2] == 0, so Alice does not need it. Therefore Bob receives it.
//     Day 2: No rain, so no one receives it.
//     Day 3: rain[3] == 1 and rain[4] == 1, so both need it. Counts are tied at 1, so Alice receives it.
//     Day 4: Last day, so only Bob needs it, and he receives it.
//     Therefore, answer = ["Alice", "Bob", "None", "Alice", "Bob"].

// Example 2:
// Input: rain = [0,1,0,1]
// Output: ["None","Bob","None","Bob"]
// Explanation:
//     It never rains on two consecutive days, so Alice never needs the umbrella.
//     Bob receives it on days 1 and 3, and no one receives it on days 0 and 2.
//     Therefore, answer = ["None", "Bob", "None", "Bob"].

// Constraints:
//     1 <= n == rain.length <= 100
//     rain[i] is either 0 or 1.

import "fmt"

func assignUmbrella(rain []int) []string {
    n, alice, bob := len(rain), 0, 0
    res := make([]string, n)
    for i := 0; i < n; i++ {
        aliceneeds := (i < n-1 && rain[i] == 1 && rain[i+1] == 1)
        bobneeds := (rain[i] == 1)
        if aliceneeds && bobneeds {
            if alice <= bob {
                res[i] = "Alice"
                alice++
            } else {
                res[i] = "Bob"
                bob++
            }
        } else if bobneeds {
            res[i] = "Bob"
            bob++
        } else {
            res[i] = "None"
        }
    }
    return res
}

func main() {
// Example 1:
// Input: rain = [1,1,0,1,1]
// Output: ["Alice","Bob","None","Alice","Bob"]
// Explanation:
//     Day 0: rain[0] == 1 and rain[1] == 1, so both need it. Counts are tied, so Alice receives it.
//     Day 1: rain[1] == 1, so Bob needs it, but rain[2] == 0, so Alice does not need it. Therefore Bob receives it.
//     Day 2: No rain, so no one receives it.
//     Day 3: rain[3] == 1 and rain[4] == 1, so both need it. Counts are tied at 1, so Alice receives it.
//     Day 4: Last day, so only Bob needs it, and he receives it.
//     Therefore, answer = ["Alice", "Bob", "None", "Alice", "Bob"].
fmt.Println(assignUmbrella([]int{1,1,0,1,1})) // ["Alice","Bob","None","Alice","Bob"]
// Example 2:
// Input: rain = [0,1,0,1]
// Output: ["None","Bob","None","Bob"]
// Explanation:
//     It never rains on two consecutive days, so Alice never needs the umbrella.
//     Bob receives it on days 1 and 3, and no one receives it on days 0 and 2.
//     Therefore, answer = ["None", "Bob", "None", "Bob"].
fmt.Println(assignUmbrella([]int{0,1,0,1})) // ["None","Bob","None","Bob"]
}