package main

import "container/heap"
import "fmt"
import "sort"

// LCP 30. 魔塔游戏
// 小扣当前位于魔塔游戏第一层，共有 N 个房间，编号为 0 ~ N-1。
// 每个房间的补血道具/怪物对于血量影响记于数组 nums，其中正数表示道具补血数值，即血量增加对应数值；负数表示怪物造成伤害值，即血量减少对应数值；0 表示房间对血量无影响。

// 小扣初始血量为 1，且无上限。
// 假定小扣原计划按房间编号升序访问所有房间补血/打怪，为保证血量始终为正值，小扣需对房间访问顺序进行调整，每次仅能将一个怪物房间（负数的房间）调整至访问顺序末尾。
// 请返回小扣最少需要调整几次，才能顺利访问所有房间。
// 若调整顺序也无法访问完全部房间，请返回 -1。

// 示例 1：
// 输入：nums = [100,100,100,-250,-60,-140,-50,-50,100,150]
// 输出：1
// 解释：初始血量为 1。至少需要将 nums[3] 调整至访问顺序末尾以满足要求。

// 示例 2：
// 输入：nums = [-200,-300,400,0]
// 输出：-1
// 解释：调整访问顺序也无法完成全部房间的访问。

// 提示：
// 		1 <= nums.length <= 10^5
// 		-10^5 <= nums[i] <= 10^5

type minHeap struct{ sort.IntSlice } // 继承 Len, Less, Swap
func (h *minHeap) Push(v any) { 
	h.IntSlice = append(h.IntSlice, v.(int)) 
}
func (h *minHeap) Pop() any   { 
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1];
	return v 
}

func magicTower(nums []int) int {
	// 先计算整体值，如果总和都为负 直接返回 -1
	sum := 0
	for i := 0; i < len(nums); i = i + 1 {
		sum = sum + nums[i]
	}
	if sum < 0 {
		return -1
	}

	room := &minHeap{} // 当前扣血房间
	now := 1  // 当前血量
	ans := 0  // 要调整的步数
	for i := 0; i < len(nums); i = i + 1 {
		if nums[i] >= 0 { // 加血房间或不影响血量的房间
			now = now + nums[i]
		} else { // 扣血房间
			now = now + nums[i]
			heap.Push(room, nums[i])
			// 如果当前血量不是正值，则需要将之前扣血最多的房间移动到末尾（贪心）
			for now <= 0 && room.Len() > 0 {
				now = now -  heap.Pop(room).(int)
				ans += 1
			}
			// 如果当前血量不是正值，且已经没有可以移动的房间，则说明无法通过
			if now <= 0 &&  room.Len() > 0 {
				return -1
			}
		}
	}
	return ans
}

func magicTower1(nums []int) (ans int) {
    sum := 0
    for _, x := range nums {
        sum += x
    }
    if sum < 0 {
        return -1
    }

    hp := 1
    h := &minHeap{}
    for _, x := range nums {
        if x < 0 {
            heap.Push(h, x)
        }
        hp += x
        if hp < 1 {
            // 这意味着 x < 0，所以前面必然会把 x 入堆
            // 所以堆必然不是空的，并且堆顶 <= x
            hp -= heap.Pop(h).(int) // 反悔
            ans++
        }
    }
    return
}


func main() {
	fmt.Println(magicTower([]int{100,100,100,-250,-60,-140,-50,-50,100,150})) // 1
	fmt.Println(magicTower([]int{-200,-300,400,0})) // -1

	fmt.Println(magicTower1([]int{100,100,100,-250,-60,-140,-50,-50,100,150})) // 1
	fmt.Println(magicTower1([]int{-200,-300,400,0})) // -1
}