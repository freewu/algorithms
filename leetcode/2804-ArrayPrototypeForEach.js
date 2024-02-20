// 2804. Array Prototype ForEach
// Write your version of method forEach that enhances all arrays such that you can call the array.forEach(callback, context) method on any array and it will execute callback on each element of the array. 
// Method forEach should not return anything.

// callback accepts the following arguments:
//     currentValue - represents the current element being processed in the array. It is the value of the element in the current iteration.
//     index - represents the index of the current element being processed in the array.
//     array - represents the array itself, allowing access to the entire array within the callback function.

// The context is the object that should be passed as the function context parameter to the callback function, 
// ensuring that the this keyword within the callback function refers to this context object.

// Try to implement it without using the built-in array methods.

// Example 1:
// Input: 
// arr = [1,2,3], 
// callback = (val, i, arr) => arr[i] = val * 2, 
// context = {"context":true}
// Output: [2,4,6]
// Explanation: 
// arr.forEach(callback, context)  
// console.log(arr) // [2,4,6]
// The callback is executed on each element of the array.

// Example 2:
// Input: 
// arr = [true, true, false, false], 
// callback = (val, i, arr) => arr[i] = this, 
// context = {"context": false}
// Output: [{"context":false},{"context":false},{"context":false},{"context":false}]
// Explanation: 
// arr.forEach(callback, context) 
// console.log(arr) // [{"context":false},{"context":false},{"context":false},{"context":false}]
// The callback is executed on each element of the array with the right context.

// Example 3:
// Input: 
// arr = [true, true, false, false], 
// callback = (val, i, arr) => arr[i] = !val, 
// context = {"context": 5}
// Output: [false,false,true,true]
 
// Constraints:
//     arr is a valid JSON array
//     context is a valid JSON object
//     fn is a function
//     0 <= arr.length <= 10^5

/**
 * @param {Function} callback
 * @param {Object} context
 * @return {void}
 */
Array.prototype.forEach = function(callback, context) {
    for (let i = 0; i < this.length; i++) {
        // 使用 call 方法调用回调函数，并将当前元素的值、索引和数组本身作为参数传递给回调函数
        // 将上下文对象作为 this 绑定到回调函数
        // context 的传入
        callback.call(context, this[i], i, this)
    }
}

// # call  调用一个对象的一个方法，以另一个对象替换当前对象。
//      call([thisObj[,arg1[, arg2[, [,.argN]]]]])
// 参数:
//      thisObj
//      可选项。将被用作当前对象的对象。
//      arg1, arg2, , argN

// Example 1:
let arr = [1,2,3];
let callback = (val, i, arr) => arr[i] = val * 2;
let context = {"context":true};
arr.forEach(callback, context)  
console.log(arr) // [2,4,6]

// Example 2:
arr = [true, true, false, false];
callback = (val, i, arr) => arr[i] = this;
context = {"context": false}
arr.forEach(callback, context) 
console.log(arr) // [{"context":false},{"context":false},{"context":false},{"context":false}]


// Example 3:
arr = [true, true, false, false], 
callback = (val, i, arr) => arr[i] = !val, 
context = {"context": 5}
arr.forEach(callback, context) 
console.log(arr) // Output: [false,false,true,true]
 