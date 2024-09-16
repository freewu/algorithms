package main

// 2590. Design a Todo List
// Design a Todo List Where users can add tasks, mark them as complete, or get a list of pending tasks. 
// Users can also add tags to tasks and can filter the tasks by certain tags.

// Implement the TodoList class:
//     TodoList() 
//         Initializes the object.
//     int addTask(int userId, String taskDescription, int dueDate, List<String> tags) 
//         Adds a task for the user with the ID userId with a due date equal to dueDate and a list of tags attached to the task. 
//         The return value is the ID of the task. This ID starts at 1 and is sequentially increasing. 
//         That is, the first task's id should be 1, the second task's id should be 2, and so on.
//     List<String> getAllTasks(int userId) 
//         Returns a list of all the tasks not marked as complete for the user with ID userId, ordered by the due date. 
//         You should return an empty list if the user has no uncompleted tasks.
//     List<String> getTasksForTag(int userId, String tag) 
//         Returns a list of all the tasks that are not marked as complete for the user with the ID userId and have tag as one of their tags, ordered by their due date. 
//         Return an empty list if no such task exists.
//     void completeTask(int userId, int taskId) 
//         Marks the task with the ID taskId as completed only if the task exists and the user with the ID userId has this task, and it is uncompleted.

// Example 1:
// Input
// ["TodoList", "addTask", "addTask", "getAllTasks", "getAllTasks", "addTask", "getTasksForTag", "completeTask", "completeTask", "getTasksForTag", "getAllTasks"]
// [[], [1, "Task1", 50, []], [1, "Task2", 100, ["P1"]], [1], [5], [1, "Task3", 30, ["P1"]], [1, "P1"], [5, 1], [1, 2], [1, "P1"], [1]]
// Output
// [null, 1, 2, ["Task1", "Task2"], [], 3, ["Task3", "Task2"], null, null, ["Task3"], ["Task3", "Task1"]]
// Explanation
// TodoList todoList = new TodoList();
// todoList.addTask(1, "Task1", 50, []); // return 1. This adds a new task for the user with id 1.
// todoList.addTask(1, "Task2", 100, ["P1"]); // return 2. This adds another task for the user with id 1.
// todoList.getAllTasks(1); // return ["Task1", "Task2"]. User 1 has two uncompleted tasks so far.
// todoList.getAllTasks(5); // return []. User 5 does not have any tasks so far.
// todoList.addTask(1, "Task3", 30, ["P1"]); // return 3. This adds another task for the user with id 1.
// todoList.getTasksForTag(1, "P1"); // return ["Task3", "Task2"]. This returns the uncompleted tasks that have the tag "P1" for the user with id 1.
// todoList.completeTask(5, 1); // This does nothing, since task 1 does not belong to user 5.
// todoList.completeTask(1, 2); // This marks task 2 as completed.
// todoList.getTasksForTag(1, "P1"); // return ["Task3"]. This returns the uncompleted tasks that have the tag "P1" for the user with id 1.
//                                   // Notice that we did not include "Task2" because it is completed now.
// todoList.getAllTasks(1); // return ["Task3", "Task1"]. User 1 now has 2 uncompleted tasks.

// Constraints:
//     1 <= userId, taskId, dueDate <= 100
//     0 <= tags.length <= 100
//     1 <= taskDescription.length <= 50
//     1 <= tags[i].length, tag.length <= 20
//     All dueDate values are unique.
//     All the strings consist of lowercase and uppercase English letters and digits.
//     At most 100 calls will be made for each method.

import "fmt"
import "sort"

type TodoList struct {
    No int
    Data map[int][]*TodoItem
}

type TodoItem struct {
    Id int
    Description string
    Date int
    Tags []string
    FinishLabel bool
}

func Constructor() TodoList {
    return TodoList{ No: 1, Data: map[int][]*TodoItem{} }
}

func (this *TodoList) AddTask(userId int, taskDescription string, dueDate int, tags []string) int {
    taskid := this.No
    this.Data[userId] = append( this.Data[userId], &TodoItem{
        Id: taskid,
        Description: taskDescription,
        Date: dueDate,
        Tags: tags,
        FinishLabel: false,
    })
    this.No++
    return taskid
}

func (this *TodoList) GetAllTasks(userId int) []string {
    t := this.Data[userId]
    arr := []*TodoItem{}
    for _, v := range t {
        if v.FinishLabel == false { // 只取完成的
            arr = append(arr, v)
        }
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].Date < arr[j].Date
    })
    res:= []string{}
    for _, v := range arr {
        res = append(res, v.Description)
    }
    return res
}

func (this *TodoList) GetTasksForTag(userId int, tag string) []string {
    t := this.Data[userId]
    arr := []*TodoItem{}
    for _, v := range t {
        if v.FinishLabel == false {
            for _, a := range v.Tags {
                if a == tag {
                    arr = append(arr, v)
                    break // 匹配到 tag 就跳出 
                }
            }
        }
    }
    sort.Slice(arr, func(i,j int) bool {
        return arr[i].Date < arr[j].Date
    })
    res := []string{}
    for _, v := range arr {
        res = append(res, v.Description)
    }
    return res
}

func (this *TodoList) CompleteTask(userId int, taskId int)  {
    t := this.Data[userId]
    for _, v := range t {
        if v.Id == taskId {
            v.FinishLabel = true
        }
    }
}

/**
 * Your TodoList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.AddTask(userId,taskDescription,dueDate,tags);
 * param_2 := obj.GetAllTasks(userId);
 * param_3 := obj.GetTasksForTag(userId,tag);
 * obj.CompleteTask(userId,taskId);
 */

func main() {
    // TodoList todoList = new TodoList();
    obj := Constructor()
    fmt.Println(obj)
    // todoList.addTask(1, "Task1", 50, []); // return 1. This adds a new task for the user with id 1.
    obj.AddTask(1, "Task1", 50, []string{})
    fmt.Println(obj)
    // todoList.addTask(1, "Task2", 100, ["P1"]); // return 2. This adds another task for the user with id 1.
    obj.AddTask(1, "Task2", 100, []string{"P1"})
    fmt.Println(obj)
    // todoList.getAllTasks(1); // return ["Task1", "Task2"]. User 1 has two uncompleted tasks so far.
    fmt.Println(obj.GetAllTasks(1)) // ["Task1", "Task2"]
    // todoList.getAllTasks(5); // return []. User 5 does not have any tasks so far.
    fmt.Println(obj.GetAllTasks(5)) // []
    // todoList.addTask(1, "Task3", 30, ["P1"]); // return 3. This adds another task for the user with id 1.
    obj.AddTask(1, "Task3", 30, []string{"P1"})
    fmt.Println(obj)
    // todoList.getTasksForTag(1, "P1"); // return ["Task3", "Task2"]. This returns the uncompleted tasks that have the tag "P1" for the user with id 1.
    fmt.Println(obj.GetTasksForTag(1, "P1")) // ["Task3", "Task2"]
    // todoList.completeTask(5, 1); // This does nothing, since task 1 does not belong to user 5.
    obj.CompleteTask(5, 1) // does nothing, since task 1 does not belong to user 5.
    fmt.Println(obj)
    // todoList.completeTask(1, 2); // This marks task 2 as completed.
    obj.CompleteTask(1, 2) // does nothing, since task 1 does not belong to user 5.
    fmt.Println(obj)
    // todoList.getTasksForTag(1, "P1"); // return ["Task3"]. This returns the uncompleted tasks that have the tag "P1" for the user with id 1.
    //                                   // Notice that we did not include "Task2" because it is completed now.
    fmt.Println(obj.GetTasksForTag(1, "P1")) // ["Task3"]
    // todoList.getAllTasks(1); // return ["Task3", "Task1"]. User 1 now has 2 uncompleted tasks.
    fmt.Println(obj.GetAllTasks(1)) // ["Task3", "Task1"]
}