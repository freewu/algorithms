// 2805. Custom Interval
// Function customInterval
//     Given a function fn, a number delay and a number period, return a number id. 
//     customInterval is a function that should execute the provided function fn at intervals based on a linear pattern defined by the formula delay + period * count. 
//     The count in the formula represents the number of times the interval has been executed starting from an initial value of 0.

// Function customClearInterval 
//     Given the id. id is the returned value from the function customInterval. 
//     customClearInterval should stop executing provided function fn at intervals.

// Note: The setTimeout and setInterval functions in Node.js return an object, not a number.

// Example 1:
// Input: delay = 50, period = 20, cancelTime = 225
// Output: [50,120,210]
// Explanation: 
//     const t = performance.now()  
//     const result = []
//     const fn = () => {
//         result.push(Math.floor(performance.now() - t))
//     }
//     const id = customInterval(fn, delay, period)
        
//     setTimeout(() => {
//         customClearInterval(id)
//     }, 225)

//     50 + 20 * 0 = 50 // 50ms - 1st function call
//     50 + 20 * 1 = 70 // 50ms + 70ms = 120ms - 2nd function call
//     50 + 20 * 2 = 90 // 50ms + 70ms + 90ms = 210ms - 3rd function call

// Example 2:
// Input: delay = 20, period = 20, cancelTime = 150
// Output: [20,60,120]
// Explanation: 
//     20 + 20 * 0 = 20 // 20ms - 1st function call
//     20 + 20 * 1 = 40 // 20ms + 40ms = 60ms - 2nd function call
//     20 + 20 * 2 = 60 // 20ms + 40ms + 60ms = 120ms - 3rd function call

// Example 3:
//     Input: delay = 100, period = 200, cancelTime = 500
//     Output: [100,400]
//     Explanation: 
//     100 + 200 * 0 = 100 // 100ms - 1st function call
//     100 + 200 * 1 = 300 // 100ms + 300ms = 400ms - 2nd function call
 
// Constraints:
//     20 <= delay, period <= 250
//     20 <= cancelTime <= 1000

// 使用 Object.create(null) 创建一个没有原型链的纯净对象
let timers = Object.create(null);

/**
 * 递归函数，用于安排函数 fn 的下一次执行
 * @param {Function} fn - 需要定期执行的函数
 * @param {number} delay - 初始延迟时间
 * @param {number} period - 每次执行的间隔时间
 * @param {number} id - 计时器的唯一标识符
 * @param {number} count - 当前已执行的次数
 */
function interval(fn, delay, period, id, count = 0) {
    timers[id] = setTimeout(() => {
        fn(); // 执行函数
        interval(fn, delay, period, id, count + 1); // 安排下一次执行
    }, delay + period * count); // 计算下次执行的时间
  }

/**
 * @param {Function} fn
 * @param {number} delay
 * @param {number} period
 * @return {number} id
 */
function customInterval(fn, delay, period) {
    const id = Date.now(); // 使用当前时间戳作为计时器 ID
    interval(fn, delay, period, id); // 安排第一次执行
    return id; // 返回计时器 ID
}

/**
 * @param {number} id
 * @return {void}
 */
function customClearInterval(id) {
    if (timers[id]) {
        clearTimeout(timers[id]); // 清除计时器
        delete timers[id]; // 从 timers 对象中删除引用
    }
}

// example 1:
// const t = performance.now()  
const t = process.hrtime();
const result = []
const fn = () => {
    console.log(process.hrtime())
    console.log(new Date())
    //result.push(Math.floor(performance.now() - t))
    result.push(Math.floor(process.hrtime() - t))
}
const id = customInterval(fn, 50, 20)
    
setTimeout(() => {
    customClearInterval(id)
}, 225)