package main

// 2343. Query Kth Smallest Trimmed Number
// You are given a 0-indexed array of strings nums, where each string is of equal length and consists of only digits.
// You are also given a 0-indexed 2D integer array queries where queries[i] = [ki, trimi]. For each queries[i], you need to:
//     Trim each number in nums to its rightmost trimi digits.
//     Determine the index of the kith smallest trimmed number in nums. If two trimmed numbers are equal, the number with the lower index is considered to be smaller.
//     Reset each number in nums to its original length.

// Return an array answer of the same length as queries, where answer[i] is the answer to the ith query.

// Note:
//     To trim to the rightmost x digits means to keep removing the leftmost digit, until only x digits remain.
//     Strings in nums may contain leading zeros.

// Example 1:
// Input: nums = ["102","473","251","814"], queries = [[1,1],[2,3],[4,2],[1,2]]
// Output: [2,2,1,0]
// Explanation:
// 1. After trimming to the last digit, nums = ["2","3","1","4"]. The smallest number is 1 at index 2.
// 2. Trimmed to the last 3 digits, nums is unchanged. The 2nd smallest number is 251 at index 2.
// 3. Trimmed to the last 2 digits, nums = ["02","73","51","14"]. The 4th smallest number is 73.
// 4. Trimmed to the last 2 digits, the smallest number is 2 at index 0.
//    Note that the trimmed number "02" is evaluated as 2.

// Example 2:
// Input: nums = ["24","37","96","04"], queries = [[2,1],[2,2]]
// Output: [3,0]
// Explanation:
// 1. Trimmed to the last digit, nums = ["4","7","6","4"]. The 2nd smallest number is 4 at index 3.
//    There are two occurrences of 4, but the one at index 0 is considered smaller than the one at index 3.
// 2. Trimmed to the last 2 digits, nums is unchanged. The 2nd smallest number is 24.
 
// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i].length <= 100
//     nums[i] consists of only digits.
//     All nums[i].length are equal.
//     1 <= queries.length <= 100
//     queries[i].length == 2
//     1 <= ki <= nums.length
//     1 <= trimi <= nums[i].length
    
// Follow up: Could you use the Radix Sort Algorithm to solve this problem? What will be the complexity of that solution?

import "fmt"

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
    maxTrim := -1 // we'll do radix sort until the highest trim requested in the queries
    for _, q := range queries {
        if q[1] > maxTrim {
            maxTrim = q[1]
        }
    }
    // Create a matrix where each row shows the index of the nums, sorted by trim number 
    // eg: The first row will just be 0, 1, 2 ... n as this is the original unsorted nums 
    // The row i will be sorting the previous row, using count sorting on the corresponding rightmost digit. 
    indexesOfSortedByTrim := make([][]int, maxTrim + 1)
    // Initialize with unsorted
    indexesOfSortedByTrim[0] = make([]int, len(nums))
    for i:=0; i < len(nums); i++ {
        indexesOfSortedByTrim[0][i] = i
    }
    sortIndexesForTrim := func(nums []string, indexes []int, trim int) []int {
        stringLen := len(nums[0]) // they are all the same 
        counts := make([]int, 10) // let's count the digts (0-9)
        for _, index := range indexes {
            digit := nums[index][stringLen-trim] - 48 // this converts the ascii char to its number
            counts[digit] ++
        }
        // we now overwrite our original counts with the starting index
        // of each element in the final sorted array
        startingIndex := 0
        for i := 0; i < 10; i++ {
            count := counts[i]
            counts[i] = startingIndex
            startingIndex = startingIndex + count
        }
        // Now iterate over indexes, and once again use the corresponding digit and use the starting index to place and reduce it
        sortedArray := make([]int, len(indexes))
        for _, index := range indexes {
            digit := nums[index][stringLen-trim] - 48 // this converts the ascii char to its number
            sortedArray[counts[digit]] = index
            // since we have placed an item in index counts[elem], we need to
            // increment counts[elem] index by 1 so the next duplicate element
            // is placed in appropriate index
            counts[digit] ++
        }
        return sortedArray
    }
    // Now keep adding the sorted by trim lines 
    for i := 1; i <= maxTrim; i++ {
        indexesOfSortedByTrim[i] = sortIndexesForTrim(nums, indexesOfSortedByTrim[i-1], i)
    }
    res := make([]int, len(queries))
    // We now just use all the calculations to address all queries
    for i, q := range queries {
        res[i] = indexesOfSortedByTrim[q[1]][q[0]-1]
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = ["102","473","251","814"], queries = [[1,1],[2,3],[4,2],[1,2]]
    // Output: [2,2,1,0]
    // Explanation:
    // 1. After trimming to the last digit, nums = ["2","3","1","4"]. The smallest number is 1 at index 2.
    // 2. Trimmed to the last 3 digits, nums is unchanged. The 2nd smallest number is 251 at index 2.
    // 3. Trimmed to the last 2 digits, nums = ["02","73","51","14"]. The 4th smallest number is 73.
    // 4. Trimmed to the last 2 digits, the smallest number is 2 at index 0.
    //    Note that the trimmed number "02" is evaluated as 2.
    fmt.Println(smallestTrimmedNumbers([]string{"102","473","251","814"},[][]int{{1,1},{2,3},{4,2},{1,2}})) // [2,2,1,0]
    // Example 2:
    // Input: nums = ["24","37","96","04"], queries = [[2,1],[2,2]]
    // Output: [3,0]
    // Explanation:
    // 1. Trimmed to the last digit, nums = ["4","7","6","4"]. The 2nd smallest number is 4 at index 3.
    //    There are two occurrences of 4, but the one at index 0 is considered smaller than the one at index 3.
    // 2. Trimmed to the last 2 digits, nums is unchanged. The 2nd smallest number is 24.
    fmt.Println(smallestTrimmedNumbers([]string{"24","37","96","04"},[][]int{{2,1},{2,2}})) // [3,0]
}