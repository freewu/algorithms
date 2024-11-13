package main

// 1850. Minimum Adjacent Swaps to Reach the Kth Smallest Number
// You are given a string num, representing a large integer, and an integer k.

// We call some integer wonderful if it is a permutation of the digits in num and is greater in value than num. 
// There can be many wonderful integers. 
// However, we only care about the smallest-valued ones.
//     For example, when num = "5489355142":
//         The 1st smallest wonderful integer is "5489355214".
//         The 2nd smallest wonderful integer is "5489355241".
//         The 3rd smallest wonderful integer is "5489355412".
//         The 4th smallest wonderful integer is "5489355421".

// Return the minimum number of adjacent digit swaps that needs to be applied to num to reach the kth smallest wonderful integer.

// The tests are generated in such a way that kth smallest wonderful integer exists.

// Example 1:
// Input: num = "5489355142", k = 4
// Output: 2
// Explanation: The 4th smallest wonderful number is "5489355421". To get this number:
// - Swap index 7 with index 8: "5489355142" -> "5489355412"
// - Swap index 8 with index 9: "5489355412" -> "5489355421"

// Example 2:
// Input: num = "11112", k = 4
// Output: 4
// Explanation: The 4th smallest wonderful number is "21111". To get this number:
// - Swap index 3 with index 4: "11112" -> "11121"
// - Swap index 2 with index 3: "11121" -> "11211"
// - Swap index 1 with index 2: "11211" -> "12111"
// - Swap index 0 with index 1: "12111" -> "21111"

// Example 3:
// Input: num = "00123", k = 1
// Output: 1
// Explanation: The 1st smallest wonderful number is "00132". To get this number:
// - Swap index 3 with index 4: "00123" -> "00132"

// Constraints:
//     2 <= num.length <= 1000
//     1 <= k <= 1000
//     num only consists of digits.

import "fmt"

func getMinSwaps(num string, k int) int {
    res, j, arr := 0, 0, []byte(num)
    nextPermutation := func (arr []byte) {
        n := len(arr)
        i := n - 2
        for ; i >= 0 && arr[i] >= arr[i + 1]; i-- {}
        if i >= 0 {
            j := n - 1
            for ; j >= 0 && arr[j] <= arr[i]; j-- {}
            arr[i], arr[j] = arr[j], arr[i]
        }
        for k, m := i + 1, n - 1; k < m; k, m = k + 1, m - 1 {
            arr[k], arr[m] = arr[m], arr[k]
        }
    }
    for i := 0; i < k; i++ {
        nextPermutation(arr)
    }
    for i := 0; i < len(num); i++ {
        if num[i] != arr[i] {
            for j = i + 1; num[i] != arr[j]; j++ { }
            for ; j != i; j-- {
                arr[j], arr[j-1] = arr[j-1], arr[j] // swap
                res++
            }
        }
    }
    return res
}

func getMinSwaps1(num string, k int) int {
    res, after, before := 0, []byte(num), []byte(num)
    nextPermutation := func(nums []byte) {
        i, n := 0, len(nums)
        for i = n - 2; i >= 0 && nums[i] >= nums[i + 1]; i-- {}
        if i >= 0 {
            j := n - 1
            for ; nums[i] >= nums[j]; j-- {}
            nums[i], nums[j] = nums[j], nums[i]
        }
        for n, i = n - 1, i + 1; i < n; {
            nums[i], nums[n] = nums[n], nums[i]
            n, i = n - 1, i + 1
        }
    }
    for ; k > 0; k-- {
        nextPermutation(after)
    }
    for i := range before {
        if after[i] == before[i] { continue }
        for j := i + 1; j < len(after); j++ {
            if after[i] == before[j] {
                res += (j - i)
                copy(before[i+1:], before[i:j])
                break
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: num = "5489355142", k = 4
    // Output: 2
    // Explanation: The 4th smallest wonderful number is "5489355421". To get this number:
    // - Swap index 7 with index 8: "5489355142" -> "5489355412"
    // - Swap index 8 with index 9: "5489355412" -> "5489355421"
    fmt.Println(getMinSwaps("5489355142", 4)) // 2
    // Example 2:
    // Input: num = "11112", k = 4
    // Output: 4
    // Explanation: The 4th smallest wonderful number is "21111". To get this number:
    // - Swap index 3 with index 4: "11112" -> "11121"
    // - Swap index 2 with index 3: "11121" -> "11211"
    // - Swap index 1 with index 2: "11211" -> "12111"
    // - Swap index 0 with index 1: "12111" -> "21111"
    fmt.Println(getMinSwaps("11112", 4)) // 4
    // Example 3:
    // Input: num = "00123", k = 1
    // Output: 1
    // Explanation: The 1st smallest wonderful number is "00132". To get this number:
    // - Swap index 3 with index 4: "00123" -> "00132"
    fmt.Println(getMinSwaps("00123", 1)) // 1

    fmt.Println(getMinSwaps1("5489355142", 4)) // 2
    fmt.Println(getMinSwaps1("11112", 4)) // 4
    fmt.Println(getMinSwaps1("00123", 1)) // 1
}