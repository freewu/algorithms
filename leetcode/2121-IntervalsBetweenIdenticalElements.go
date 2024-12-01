package main

// 2121. Intervals Between Identical Elements
// You are given a 0-indexed array of n integers arr.

// The interval between two elements in arr is defined as the absolute difference between their indices. 
// More formally, the interval between arr[i] and arr[j] is |i - j|.

// Return an array intervals of length n where intervals[i] is the sum of intervals between arr[i] 
// and each element in arr with the same value as arr[i].

// Note: |x| is the absolute value of x.

// Example 1:
// Input: arr = [2,1,3,1,2,3,3]
// Output: [4,2,7,2,4,4,5]
// Explanation:
// - Index 0: Another 2 is found at index 4. |0 - 4| = 4
// - Index 1: Another 1 is found at index 3. |1 - 3| = 2
// - Index 2: Two more 3s are found at indices 5 and 6. |2 - 5| + |2 - 6| = 7
// - Index 3: Another 1 is found at index 1. |3 - 1| = 2
// - Index 4: Another 2 is found at index 0. |4 - 0| = 4
// - Index 5: Two more 3s are found at indices 2 and 6. |5 - 2| + |5 - 6| = 4
// - Index 6: Two more 3s are found at indices 2 and 5. |6 - 2| + |6 - 5| = 5

// Example 2:
// Input: arr = [10,5,10,10]
// Output: [5,0,3,4]
// Explanation:
// - Index 0: Two more 10s are found at indices 2 and 3. |0 - 2| + |0 - 3| = 5
// - Index 1: There is only one 5 in the array, so its sum of intervals to identical elements is 0.
// - Index 2: Two more 10s are found at indices 0 and 3. |2 - 0| + |2 - 3| = 3
// - Index 3: Two more 10s are found at indices 0 and 2. |3 - 0| + |3 - 2| = 4

// Constraints:
//     n == arr.length
//     1 <= n <= 10^5
//     1 <= arr[i] <= 10^5

// Note: This question is the same as 2615: Sum of Distances.

import "fmt"

func getDistances(arr []int) []int64 {
    data, sum := make(map[int][]int), make(map[int]int64)
    for i, v := range arr {
        data[v] = append(data[v], i)
        sum[v] += int64(i)
    }
    res := make([]int64, len(arr))
    for i, rows := range data {
        left, right := int64(0), int64(0)
        for t, index := range rows {
            right = sum[i] - left - int64(index)
            res[index] = right - left - int64(len(rows) - t - 1 - t) * int64(index)
            left += int64(index)
        }
    }
    return res
}

func getDistances1(arr []int) []int64 {
    mp := make(map[int][]int) 
    for i,v := range arr { // 使用一个map[v][]int来存储相同元素的所有下标
        mp[v] = append(mp[v],i)
    }
    res := make([]int64, len(arr))
    for _, row := range mp { // 遍历 map，计算每个v的数组p的所以元素，先计算第一个与别的元素的差值sum
        sum  := int64(0)
        for _, v  := range row {  // 先统计第一个元素的总和
            sum += int64(v - row[0])
        }
        res[row[0]] = sum // 赋值第一个数字的差值
        for i := 1; i < len(row); i++ { // 依次赋值剩下的差值
            sum += int64((2 * i - len(row))) * int64((row[i] -row[i-1]))
            res[row[i]] = sum
        }
    }
    return res
}

func getDistances2(arr []int) []int64 {
    res, buckets := make([]int64, len(arr)), make(map[int][]int)
    for i, v := range arr {
        buckets[v] = append(buckets[v], i)
    }
    for _, row := range buckets {
        n := len(row)
        if n == 1 { continue }
        prefix := 0
        for i := 1; i < n; i++ {
            prefix += row[i] - row[0]
        }
        res[row[0]] = int64(prefix)
        for i := 1; i < n; i++ {
            prefix += (row[i] - row[i-1]) * (i - (n - i))
            res[row[i]] = int64(prefix)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [2,1,3,1,2,3,3]
    // Output: [4,2,7,2,4,4,5]
    // Explanation:
    // - Index 0: Another 2 is found at index 4. |0 - 4| = 4
    // - Index 1: Another 1 is found at index 3. |1 - 3| = 2
    // - Index 2: Two more 3s are found at indices 5 and 6. |2 - 5| + |2 - 6| = 7
    // - Index 3: Another 1 is found at index 1. |3 - 1| = 2
    // - Index 4: Another 2 is found at index 0. |4 - 0| = 4
    // - Index 5: Two more 3s are found at indices 2 and 6. |5 - 2| + |5 - 6| = 4
    // - Index 6: Two more 3s are found at indices 2 and 5. |6 - 2| + |6 - 5| = 5
    fmt.Println(getDistances([]int{2,1,3,1,2,3,3})) // [4,2,7,2,4,4,5]
    // Example 2:
    // Input: arr = [10,5,10,10]
    // Output: [5,0,3,4]
    // Explanation:
    // - Index 0: Two more 10s are found at indices 2 and 3. |0 - 2| + |0 - 3| = 5
    // - Index 1: There is only one 5 in the array, so its sum of intervals to identical elements is 0.
    // - Index 2: Two more 10s are found at indices 0 and 3. |2 - 0| + |2 - 3| = 3
    // - Index 3: Two more 10s are found at indices 0 and 2. |3 - 0| + |3 - 2| = 4
    fmt.Println(getDistances([]int{10,5,10,10})) // [5,0,3,4]

    fmt.Println(getDistances1([]int{2,1,3,1,2,3,3})) // [4,2,7,2,4,4,5]
    fmt.Println(getDistances1([]int{10,5,10,10})) // [5,0,3,4]

    fmt.Println(getDistances2([]int{2,1,3,1,2,3,3})) // [4,2,7,2,4,4,5]
    fmt.Println(getDistances2([]int{10,5,10,10})) // [5,0,3,4]
}