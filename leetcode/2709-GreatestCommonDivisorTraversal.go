package main

// 2709. Greatest Common Divisor Traversal
// You are given a 0-indexed integer array nums, and you are allowed to traverse between its indices.
// You can traverse between index i and index j, i != j, if and only if gcd(nums[i], nums[j]) > 1, where gcd is the greatest common divisor.
// Your task is to determine if for every pair of indices i and j in nums, where i < j, there exists a sequence of traversals that can take us from i to j.
// Return true if it is possible to traverse between all such pairs of indices, or false otherwise.

// Example 1:
// Input: nums = [2,3,6]
// Output: true
// Explanation: In this example, there are 3 possible pairs of indices: (0, 1), (0, 2), and (1, 2).
// To go from index 0 to index 1, we can use the sequence of traversals 0 -> 2 -> 1, where we move from index 0 to index 2 because gcd(nums[0], nums[2]) = gcd(2, 6) = 2 > 1, and then move from index 2 to index 1 because gcd(nums[2], nums[1]) = gcd(6, 3) = 3 > 1.
// To go from index 0 to index 2, we can just go directly because gcd(nums[0], nums[2]) = gcd(2, 6) = 2 > 1. Likewise, to go from index 1 to index 2, we can just go directly because gcd(nums[1], nums[2]) = gcd(3, 6) = 3 > 1.

// Example 2:
// Input: nums = [3,9,5]
// Output: false
// Explanation: No sequence of traversals can take us from index 0 to index 2 in this example. So, we return false.

// Example 3:
// Input: nums = [4,3,12,8]
// Output: true
// Explanation: There are 6 possible pairs of indices to traverse between: (0, 1), (0, 2), (0, 3), (1, 2), (1, 3), and (2, 3). A valid sequence of traversals exists for each pair, so we return true.
 
// Constraints:
//         1 <= nums.length <= 10^5
//         1 <= nums[i] <= 10^5

import "fmt"
import "math"

func canTraverseAllPairs(nums []int) bool {
    if len(nums) == 1{
        return true
    }
    
    //处理有很多相同数字的情况，只需要保留其中一个即可
    ma := make(map[int]interface{})
    for _, num := range nums{
        ma[num] = struct{}{}
    }
    nums = make([]int,0)
    for k, _ := range ma{
        nums = append(nums, k)
    }
    
    n := len(nums)
    // 全部是相同的数字，则判断这个数字本身有没有大于1的最大公约数
    if n == 1{
        return judge(nums[0], nums[0])
    }
    
    // 用邻接表存图
    graph := make(map[int][]int)
    for i := 0; i<n-1; i++{
        for j := i+1; j<n; j++{
            can := judge(nums[i], nums[j])
            if can{
                graph[i] = append(graph[i], j)
                graph[j] = append(graph[j], i)
            }
        }
        // 剪枝！！！因为已经发现了一个根本不可能到达的孤点
        if len(graph[i]) == 0{
            return false
        }
    }
    
    // 遍历图，看是否连通
    visit := make([]bool, n)
    Graph(visit, graph, 0)
    for _, v := range visit{
        if !v {
            return false
        }
    }
    return true
}

func Graph(visit []bool, graph map[int][]int, source int){
    visit[source] = true
    for _, to := range graph[source]{
        if !visit[to]{
            Graph(visit,graph,to)
        }
    }
}

func judge(i,j int) bool{
   return gcb(i,j) > 1
}

func gcb(i, j int) int{
    if j > 0 {
        return gcb(j, i%j)
    } else {
        return i
    }
}


func canTraverseAllPairs1(nums []int) bool {
    var max int
    var oneFound bool
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            oneFound = true
        }
        if nums[i] > max {
            max = nums[i]
        }
    }
    if oneFound && len(nums) > 1 {
        // We can't have 1 anywhere because gcd(1, x) = 1
        return false
    }
    root := make([]int, max+1)
    rank := make([]int, max+1)
    for i := 0; i < len(root); i++ {
        root[i] = i
    }
    for _, num := range nums {
        for _, factor := range factorize(num) {
            union(root, rank, num, factor)
        }
    }
    r := find(root, nums[0])
    for i := 1; i < len(nums); i++ {
        if find(root, nums[i]) != r {
            return false
        }
    }
    return true
}

func find(root []int, x int) int {
    if root[x] == x {
        return x
    }
    root[x] = find(root, root[x])
    return root[x]
}

func factorize(num int) []int {
    var factors []int
    factor := 2
    end := int(math.Sqrt(float64(num)))
    for num > 1 && factor <= end {
        appended := false
        for num%factor == 0 {
            if !appended {
                factors = append(factors, factor)
                appended = true
            }
            num /= factor
        }
        factor++
    }
    if num > 1 { // num is a really big prime number
        factors = append(factors, num)
    }
    return factors
}

func union(root, rank []int, x, y int) {
    rootX := find(root, x)
    rootY := find(root, y)
    if rootX != rootY {
        if rank[rootX] > rank[rootY] {
            root[rootY] = rootX
        } else if rank[rootX] < rank[rootY] {
            root[rootX] = rootY
        } else {
            root[rootY] = rootX
            rank[rootX]++
        }
    }
}

func main() {
    
    fmt.Println(canTraverseAllPairs([]int{2,3,6})) // true
    fmt.Println(canTraverseAllPairs([]int{3,9,5})) // false
    fmt.Println(canTraverseAllPairs([]int{4,3,12,8})) // true

    fmt.Println(canTraverseAllPairs1([]int{2,3,6})) // true
    fmt.Println(canTraverseAllPairs1([]int{3,9,5})) // false
    fmt.Println(canTraverseAllPairs1([]int{4,3,12,8})) // true

}
