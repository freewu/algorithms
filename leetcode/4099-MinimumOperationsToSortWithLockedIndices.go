package main

import "fmt"

// 4099. Minimum Operations to Sort With Locked Indices
// You are given an integer array nums of length n, where nums is a permutation of the integers in the range [1, n].

// You are also given a binary array locked of length n and an integer k.
//     1. locked[i] = 1 means index i is locked.
//     2. locked[i] = 0 means index i is unlocked.

// You may perform the following operations in any order:
//     1. Swap: Choose an index i such that both indices i and i + 1 are currently unlocked, and swap nums[i] and nums[i + 1].
//     2. Unlock: Choose an index i such that locked[i] = 1 and permanently unlock it by setting locked[i] = 0.

// You may perform any number of swap operations but at most k unlock operations.

// Return the minimum total number of operations required to sort nums in strictly increasing order.
// If it is impossible, return -1.

// Example 1:
// Input: nums = [2,1,3,4], locked = [0,1,0,1], k = 1
// Output: 2
// Explanation:
// - Unlock index 1. The array locked becomes [0,0,0,1].
// - Swap nums[0] and nums[1]. The array nums becomes [1,2,3,4].
// The total number of operations is 1 + 1 = 2.

// Example 2:
// Input: nums = [3,1,2], locked = [1,0,1], k = 2
// Output: 4
// Explanation:
// - Unlock indices 0 and 2. All indices are now unlocked.
// - Swap indices 0 and 1. The array nums becomes [1,3,2].
// - Swap indices 1 and 2. The array nums becomes [1,2,3].
// The total number of operations is 2 + 2 = 4.

// Example 3:
// Input: nums = [2,3,1], locked = [1,0,0], k = 0
// Output: -1
// Explanation:
// Index 0 is locked, so its value cannot be moved. Since k = 0, no indices can be unlocked. Therefore, sorting the array is impossible.

// Constraints:
//     - 1 <= n == nums.length == locked.length <= 2000
//     - 1 <= nums[i] <= n
//     - 0 <= locked[i] <= 1
//     - nums is a permutation of the integers in the range [1, n].
//     - 0 <= k <= n

type Fenwick []int

// a[i] 增加 val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f Fenwick) update(i int, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f Fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func minOperations(nums, locked []int, k int) int {
	res, n := 0, len(nums)
	t := make(Fenwick, n+1)
	for i := 0; i < n; {
		if nums[i] == i+1 { // nums[i] 已在正确位置上
			i++
			continue
		}
		mx := -1
		for ; i < n && mx != i; i++ {
			if locked[i] > 0 {
				if k == 0 { // 无法解锁
					return -1
				}
				k--
				res++
			}
			x := nums[i]
			v := n + 1 - x
			res += t.pre(v) // x 与其左侧元素的逆序对个数
			t.update(v, 1)
			mx = max(mx, x)
		}
	}
	return res
}

func minOperations1(nums []int, locked []int, k int) int {
	var dnc func([]int, []int, int, int) int64
	dnc = func(a []int, tmp []int, l int, r int) int64 {
		if r-l <= 1 {
			return 0
		}
		m := (l + r) / 2
		res := dnc(a, tmp, l, m) + dnc(a, tmp, m, r)
		i, j, p := l, m, l
		for i < m && j < r {
			if a[i] <= a[j] {
				tmp[p] = a[i]
				p++
				i++
			} else {
				tmp[p] = a[j]
				p++
				j++
				res += int64(m - i)
			}
		}
		for i < m {
			tmp[p] = a[i]
			p++
			i++
		}
		for j < r {
			tmp[p] = a[j]
			p++
			j++
		}
		for x := l; x < r; x++ {
			a[x] = tmp[x]
		}
		return res
	}
	n := len(nums)
	var need int64
	var mx int64
	for i := 0; i < n; i++ {
		mx = max(mx, int64(nums[i]))
		if locked[i] != 0 && !(nums[i] == i+1 && mx == int64(i+1)) {
			need++
		}
	}
	if need > int64(k) {
		return -1
	}
	tmp := make([]int, n)
	return int(dnc(nums, tmp, 0, n) + need)
}

func main() {
	// Example 1:
	// Input: nums = [2,1,3,4], locked = [0,1,0,1], k = 1
	// Output: 2
	// Explanation:
	// - Unlock index 1. The array locked becomes [0,0,0,1].
	// - Swap nums[0] and nums[1]. The array nums becomes [1,2,3,4].
	// The total number of operations is 1 + 1 = 2.
	fmt.Println(minOperations([]int{2, 1, 3, 4}, []int{0, 1, 0, 1}, 1)) // 2
	// Example 2:
	// Input: nums = [3,1,2], locked = [1,0,1], k = 2
	// Output: 4
	// Explanation:
	// - Unlock indices 0 and 2. All indices are now unlocked.
	// - Swap indices 0 and 1. The array nums becomes [1,3,2].
	// - Swap indices 1 and 2. The array nums becomes [1,2,3].
	// The total number of operations is 2 + 2 = 4.
	fmt.Println(minOperations([]int{3, 1, 2}, []int{1, 0, 1}, 2)) // 2
	// Example 3:
	// Input: nums = [2,3,1], locked = [1,0,0], k = 0
	// Output: -1
	// Explanation:
	// Index 0 is locked, so its value cannot be moved. Since k = 0, no indices can be unlocked. Therefore, sorting the array is impossible.
	fmt.Println(minOperations([]int{2, 3, 1}, []int{1, 0, 0}, 0)) // -1

	fmt.Println(minOperations1([]int{2, 1, 3, 4}, []int{0, 1, 0, 1}, 1)) // 2
	fmt.Println(minOperations1([]int{3, 1, 2}, []int{1, 0, 1}, 2))       // 2
	fmt.Println(minOperations1([]int{2, 3, 1}, []int{1, 0, 0}, 0))       // -1
}
