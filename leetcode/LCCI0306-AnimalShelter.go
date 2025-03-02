package main

// 面试题 03.06. Animal Shelter LCCI
// An animal shelter, which holds only dogs and cats, operates on a strictly"first in, first out" basis. 
// People must adopt either the"oldest" (based on arrival time) of all animals at the shelter, or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of that type). 
// They cannot select which specific animal they would like. 
// Create the data structures to maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog, and dequeueCat. 
// You may use the built-in Linked list data structure.

// enqueue method has a animal parameter, animal[0] represents the number of the animal, animal[1] represents the type of the animal, 0 for cat and 1 for dog.

// dequeue* method returns [animal number, animal type], if there's no animal that can be adopted, return [-1, -1].

// Example1:
// Input: 
// ["AnimalShelf", "enqueue", "enqueue", "dequeueCat", "dequeueDog", "dequeueAny"]
// [[], [[0, 0]], [[1, 0]], [], [], []]
// Output: 
// [null,null,null,[0,0],[-1,-1],[1,0]]

// Example2:
// Input: 
// ["AnimalShelf", "enqueue", "enqueue", "enqueue", "dequeueDog", "dequeueCat", "dequeueAny"]
// [[], [[0, 0]], [[1, 0]], [[2, 1]], [], [], []]
// Output: 
// [null,null,null,null,[2,1],[0,0],[1,0]]

// Note:
//     The number of animals in the shelter will not exceed 20000.

import "fmt"

type AnimalShelf struct {
    queue [][]int
}

func Constructor() AnimalShelf {
    return AnimalShelf{ queue: [][]int{} }
}

func (this *AnimalShelf) pick(t int) []int {
    for i, v := range this.queue {
        if v[1] == t {
            res := this.queue[i]
            this.queue = append(this.queue[0:i], this.queue[i + 1:]...)
            return res
        }
    }
    return []int{-1,-1}
}

func (this *AnimalShelf) empty() bool {
    return len(this.queue) == 0
}

func (this *AnimalShelf) Enqueue(animal []int)  {
    this.queue = append(this.queue, animal)
}

func (this *AnimalShelf) DequeueAny() []int {
    if this.empty() { return []int{-1,-1} }
    v := this.queue[0]
    this.queue = this.queue[1:]
    return v
}

func (this *AnimalShelf) DequeueDog() []int {
    if this.empty() { return []int{-1,-1} }
    return this.pick(1)
}

func (this *AnimalShelf) DequeueCat() []int {
    if this.empty() { return []int{-1,-1} }
    return this.pick(0)
}


/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */

func main() {
    // Example1:
    // Input: 
    // ["AnimalShelf", "enqueue", "enqueue", "dequeueCat", "dequeueDog", "dequeueAny"]
    // [[], [[0, 0]], [[1, 0]], [], [], []]
    // Output: 
    // [null,null,null,[0,0],[-1,-1],[1,0]]
    obj1 := Constructor()
    fmt.Println(obj1)
    obj1.Enqueue([]int{0, 0})
    fmt.Println(obj1)
    obj1.Enqueue([]int{1, 0})
    fmt.Println(obj1)
    fmt.Println(obj1.DequeueCat()) // [0,0]
    fmt.Println(obj1.DequeueDog()) // [-1,-1]
    fmt.Println(obj1.DequeueAny()) // [1,0]

    // Example2:
    // Input: 
    // ["AnimalShelf", "enqueue", "enqueue", "enqueue", "dequeueDog", "dequeueCat", "dequeueAny"]
    // [[], [[0, 0]], [[1, 0]], [[2, 1]], [], [], []]
    // Output: 
    // [null,null,null,null,[2,1],[0,0],[1,0]]
    obj2 := Constructor()
    fmt.Println(obj2)
    obj2.Enqueue([]int{0, 0})
    fmt.Println(obj2)
    obj2.Enqueue([]int{1, 0})
    fmt.Println(obj2)
    obj2.Enqueue([]int{2, 1})
    fmt.Println(obj2)
    fmt.Println(obj2.DequeueDog()) // [2,1]
    fmt.Println(obj2.DequeueCat()) // [0,0]
    fmt.Println(obj2.DequeueAny()) // [1,0]
}