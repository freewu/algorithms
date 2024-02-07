// 2637. Promise Time Limit
// Given an asynchronous function fn and a time t in milliseconds, return a new time limited version of the input function. 
// fn takes arguments provided to the time limited function.
// The time limited function should follow these rules:
// If the fn completes within the time limit of t milliseconds, the time limited function should resolve with the result.
// If the execution of the fn exceeds the time limit, the time limited function should reject with the string "Time Limit Exceeded".
 
// Example 1:
// Input: 
// fn = async (n) => { 
//   await new Promise(res => setTimeout(res, 100)); 
//   return n * n; 
// }
// inputs = [5]
// t = 50
// Output: {"rejected":"Time Limit Exceeded","time":50}
// Explanation:
// const limited = timeLimit(fn, t)
// const start = performance.now()
// let result;
// try {
//    const res = await limited(...inputs)
//    result = {"resolved": res, "time": Math.floor(performance.now() - start)};
// } catch (err) {
//    result = {"rejected": err, "time": Math.floor(performance.now() - start)};
// }
// console.log(result) // Output
// The provided function is set to resolve after 100ms. However, the time limit is set to 50ms. 
// It rejects at t=50ms because the time limit was reached.

// Example 2:
// Input: 
// fn = async (n) => { 
//   await new Promise(res => setTimeout(res, 100)); 
//   return n * n; 
// }
// inputs = [5]
// t = 150
// Output: {"resolved":25,"time":100}
// Explanation:
// The function resolved 5 * 5 = 25 at t=100ms. The time limit is never reached.

// Example 3:
// Input: 
// fn = async (a, b) => { 
//   await new Promise(res => setTimeout(res, 120)); 
//   return a + b; 
// }
// inputs = [5,10]
// t = 150
// Output: {"resolved":15,"time":120}
// Explanation:
// ​​​​The function resolved 5 + 10 = 15 at t=120ms. The time limit is never reached.

// Example 4:
// Input: 
// fn = async () => { 
//   throw "Error";
// }
// inputs = []
// t = 1000
// Output: {"rejected":"Error","time":0}
// Explanation:
// The function immediately throws an error.
 
// Constraints:
//         0 <= inputs.length <= 10
//         0 <= t <= 1000
//         fn returns a promise

/**
 * @param {Function} fn
 * @param {number} t
 * @return {Function}
 */
var timeLimit = function(fn, t) {
	return async function(...args) {
        // 使用 Promise.race 实现
        // 对于超时的情况可以再构造一个Promise，里面触发一个定时器，如果超过时间t则reject('Time Limit Exceeded')。
        // 两个Promise的结果谁先返回则谁就是最终的结果，这里需要用到 Promise.race 来做判断
        return Promise.race([
            // 正常执行方法
            fn.apply(null, args), 
            // 再构造一个Promise，里面触发一个定时器，如果超过时间t则reject('Time Limit Exceeded')。
            new Promise((resolve, reject) => {
                setTimeout(() => {
                    reject('Time Limit Exceeded');
                }, t);
            })
        ]);
    }
};

/**
 * const limited = timeLimit((t) => new Promise(res => setTimeout(res, t)), 100);
 * limited(150).catch(console.log) // "Time Limit Exceeded" at t=100ms
 */

// Promise.race() 方法接收一个包含多个 Promise 对象的可迭代对象，并返回一个新的 Promise 对象，
// 该 Promise 对象在多个 Promise 中任意一个 Promise 对象状态变为 fulfilled 或 rejected 时立即返回该 Promise 对象的值或原因。
//     const promise1 = new Promise((resolve, reject) => {
//         setTimeout(() => {
//             resolve('Promise 1 resolved');
//         }, 1000);
//     });
//     const promise2 = new Promise((resolve, reject) => {
//         setTimeout(() => {
//             resolve('Promise 2 resolved');
//         }, 2000);
//     });
//     Promise.race([promise1, promise2]).then(result => {
//         console.log(result); // "Promise 1 resolved" 因为 promise1 先返回
//     });