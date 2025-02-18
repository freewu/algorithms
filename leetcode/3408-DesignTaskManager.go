package main

// 3408. Design Task Manager
// There is a task management system that allows users to manage their tasks, each associated with a priority. 
// The system should efficiently handle adding, modifying, executing, and removing tasks.

// Implement the TaskManager class:
//     TaskManager(vector<vector<int>>& tasks) 
//         initializes the task manager with a list of user-task-priority triples. 
//         Each element in the input list is of the form [userId, taskId, priority], which adds a task to the specified user with the given priority.
//     void add(int userId, int taskId, int priority) 
//         adds a task with the specified taskId and priority to the user with userId. 
//         It is guaranteed that taskId does not exist in the system.
//     void edit(int taskId, int newPriority) 
//         updates the priority of the existing taskId to newPriority. 
//         It is guaranteed that taskId exists in the system.
//     void rmv(int taskId) 
//         removes the task identified by taskId from the system. 
//         It is guaranteed that taskId exists in the system.
//     int execTop() 
//         executes the task with the highest priority across all users. 
//         If there are multiple tasks with the same highest priority, execute the one with the highest taskId. 
//         After executing, the taskId is removed from the system. 
//         Return the userId associated with the executed task. 
//         If no tasks are available, return -1.

// Note that a user may be assigned multiple tasks.

// Example 1:
// Input:
// ["TaskManager", "add", "edit", "execTop", "rmv", "add", "execTop"]
// [[[[1, 101, 10], [2, 102, 20], [3, 103, 15]]], [4, 104, 5], [102, 8], [], [101], [5, 105, 15], []]
// Output:
// [null, null, null, 3, null, null, 5]
// Explanation
// TaskManager taskManager = new TaskManager([[1, 101, 10], [2, 102, 20], [3, 103, 15]]); // Initializes with three tasks for Users 1, 2, and 3.
// taskManager.add(4, 104, 5); // Adds task 104 with priority 5 for User 4.
// taskManager.edit(102, 8); // Updates priority of task 102 to 8.
// taskManager.execTop(); // return 3. Executes task 103 for User 3.
// taskManager.rmv(101); // Removes task 101 from the system.
// taskManager.add(5, 105, 15); // Adds task 105 with priority 15 for User 5.
// taskManager.execTop(); // return 5. Executes task 105 for User 5.

// Constraints:
//     1 <= tasks.length <= 10^5
//     0 <= userId <= 10^5
//     0 <= taskId <= 10^5
//     0 <= priority <= 10^9
//     0 <= newPriority <= 10^9
//     At most 2 * 10^5 calls will be made in total to add, edit, rmv, and execTop methods.
//     The input is generated such that taskId will be valid.

import "fmt"
import "container/heap"

type Task struct {
    UserId int
    TaskId int
    Priority int
}

type TaskHeap []Task

func (th TaskHeap) Len() int {
    return len(th)
}
func (th TaskHeap) Less(i, j int) bool {
    if th[i].Priority == th[j].Priority{
        if th[i].TaskId==th[j].TaskId{
            return th[i].UserId>th[j].UserId
        }
        return th[i].TaskId>th[j].TaskId
    }
    return th[i].Priority > th[j].Priority
}

func (th TaskHeap) Swap(i, j int) {
    th[i], th[j] = th[j], th[i]
}

func (th *TaskHeap) Push(val interface{}) {
    *th = append(*th, val.(Task))
}

func (th *TaskHeap) Pop() any {
    old := *th
    n := len(old)
    x := old[n-1]
    *th = old[0 : n-1]
    return x
}

type TaskManager struct {
    Heap TaskHeap
    Index map[int]Task
}

func Constructor(tasks [][]int) TaskManager {
    tm := TaskManager{
        Heap : make(TaskHeap,0,len(tasks)),
        Index: make(map[int]Task,len(tasks)),
    }
    for _,task:= range tasks{
        t := Task{task[0],task[1],task[2]}
        tm.Heap.Push(t)
        tm.Index[t.TaskId] = t
    }
    heap.Init(&tm.Heap)
    return tm
}

func (this *TaskManager) Add(userId int, taskId int, priority int)  {
    t := Task{userId, taskId, priority}
    heap.Push(&this.Heap, t)
    this.Index[t.TaskId]=t 
}

func (this *TaskManager) Edit(taskId int, newPriority int)  {
    t,_ := this.Index[taskId]
    t.Priority = newPriority
    this.Index[taskId]=t
    heap.Push(&this.Heap,t)
}

func (this *TaskManager) Rmv(taskId int)  {
    delete(this.Index,taskId)
}

func (this *TaskManager) ExecTop() int {
    if this.Heap.Len()>0{
        pop := heap.Pop(&this.Heap).(Task)
        t,ok := this.Index[pop.TaskId]
        for(!ok||t.Priority!=pop.Priority)&&this.Heap.Len()>0{
            pop = heap.Pop(&this.Heap).(Task)
            t,ok = this.Index[pop.TaskId]
        }
        if ok && t.Priority==pop.Priority{
            delete(this.Index,pop.TaskId)
            return pop.UserId
        }
    }
    return -1
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */

func main() {
    // TaskManager taskManager = new TaskManager([[1, 101, 10], [2, 102, 20], [3, 103, 15]]); // Initializes with three tasks for Users 1, 2, and 3.
    obj := Constructor([][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}})
    fmt.Println(obj)
    // taskManager.add(4, 104, 5); // Adds task 104 with priority 5 for User 4.
    obj.Add(4, 104, 5)
    fmt.Println(obj)
    // taskManager.edit(102, 8); // Updates priority of task 102 to 8.
    obj.Edit(102, 8)
    fmt.Println(obj)
    // taskManager.execTop(); // return 3. Executes task 103 for User 3.
    fmt.Println(obj.ExecTop()) // 3
    // taskManager.rmv(101); // Removes task 101 from the system.
    obj.Rmv(101)
    fmt.Println(obj)
    // taskManager.add(5, 105, 15); // Adds task 105 with priority 15 for User 5.
    obj.Add(5, 105, 15)
    fmt.Println(obj)
    // taskManager.execTop(); // return 5. Executes task 105 for User 5.
    fmt.Println(obj.ExecTop()) // 5
}