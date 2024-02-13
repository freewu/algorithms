// 2636. Promise Pool
// Given an array of asynchronous functions functions and a pool limit n, return an asynchronous function promisePool. 
// It should return a promise that resolves when all the input functions resolve.
// Pool limit is defined as the maximum number promises that can be pending at once. 
// promisePool should begin execution of as many functions as possible and continue executing new functions when old promises resolve. 
// promisePool should execute functions[i] then functions[i + 1] then functions[i + 2], etc. 
// When the last promise resolves, promisePool should also resolve.

// For example:
//     if n = 1, promisePool will execute one function at a time in series. However, 
//     if n = 2, it first executes two functions. When either of the two functions resolve, 
//     a 3rd function should be executed (if available), and so on until there are no functions left to execute.

// You can assume all functions never reject. 
// It is acceptable for promisePool to return a promise that resolves any value.

// Example 1:
// Input: 
// functions = [
//   () => new Promise(res => setTimeout(res, 300)),
//   () => new Promise(res => setTimeout(res, 400)),
//   () => new Promise(res => setTimeout(res, 200))
// ]
// n = 2
// Output: [[300,400,500],500]
// Explanation:
//         Three functions are passed in. They sleep for 300ms, 400ms, and 200ms respectively.
//         They resolve at 300ms, 400ms, and 500ms respectively. The returned promise resolves at 500ms.
//         At t=0, the first 2 functions are executed. The pool size limit of 2 is reached.
//         At t=300, the 1st function resolves, and the 3rd function is executed. Pool size is 2.
//         At t=400, the 2nd function resolves. There is nothing left to execute. Pool size is 1.
//         At t=500, the 3rd function resolves. Pool size is zero so the returned promise also resolves.

// Example 2:
// Input:
// functions = [
//   () => new Promise(res => setTimeout(res, 300)),
//   () => new Promise(res => setTimeout(res, 400)),
//   () => new Promise(res => setTimeout(res, 200))
// ]
// n = 5
// Output: [[300,400,200],400]
// Explanation:
//         The three input promises resolve at 300ms, 400ms, and 200ms respectively.
//         The returned promise resolves at 400ms.
//         At t=0, all 3 functions are executed. The pool limit of 5 is never met.
//         At t=200, the 3rd function resolves. Pool size is 2.
//         At t=300, the 1st function resolved. Pool size is 1.
//         At t=400, the 2nd function resolves. Pool size is 0, so the returned promise also resolves.

// Example 3:
// Input:
// functions = [
//   () => new Promise(res => setTimeout(res, 300)),
//   () => new Promise(res => setTimeout(res, 400)),
//   () => new Promise(res => setTimeout(res, 200))
// ]
// n = 1
// Output: [[300,700,900],900]
// Explanation:
//         The three input promises resolve at 300ms, 700ms, and 900ms respectively.
//         The returned promise resolves at 900ms.
//         At t=0, the 1st function is executed. Pool size is 1.
//         At t=300, the 1st function resolves and the 2nd function is executed. Pool size is 1.
//         At t=700, the 2nd function resolves and the 3rd function is executed. Pool size is 1.
//         At t=900, the 3rd function resolves. Pool size is 0 so the returned promise resolves.

// Constraints:
//         0 <= functions.length <= 10
//         1 <= n <= 10

/**
 * @param {Function[]} functions
 * @param {number} n
 * @return {Promise<any>}
 */
// 异步/等待 + Promise.all() + Array.shift()
var promisePool = async function(functions, n) {
    // 创建一个 n 长度的 promise 数组
    await Promise.all([...new Array(n)].map(async () => {
        // 每个promise的作用都是一致的，那么每个promise都可以作为一个 "线程"，
        // 不停的从functions中取出还未完成的任务进行执行，执行结束了继续执行下一个
        while (functions.length) {
            await functions.shift()()
        }
    }))
};

// 递归辅助函数
var promisePool1 = async function(functions, n) {
    return new Promise((resolve) => {
        // 每次我们执行一个新的函数，我们都会增加 functionIndex，并且我们会增加 inProgressCount
        let inProgressCount = 0;
        let functionIndex = 0;
        function helper() {
            // 任务执行完成,退出 functionIndex >= functions.length
            if (functionIndex >= functions.length) {
                if (inProgressCount === 0) resolve();
                return;
            }
            // inProgressCount < n 控制每次并行的任务数量
            // functionIndex < functions.length 用来保证执行完所有的任务
            while (inProgressCount < n && functionIndex < functions.length) {
                inProgressCount++;
                const promise = functions[functionIndex]();
                functionIndex++;
                promise.then(() => {
                    // 每次一个 promise 解决，我们都会减少 inProgressCount,并递归调用之前的
                    inProgressCount--;
                    helper();
                });
            }
        }
        helper();
    });
};


/**
 * const sleep = (t) => new Promise(res => setTimeout(res, t));
 * promisePool([() => sleep(500), () => sleep(400)], 1)
 *   .then(console.log) // After 900ms
 */

const sleep = (t) => new Promise(res => setTimeout(res, t));
promisePool([() => sleep(500), () => sleep(400)], 1)
.then(console.log) // After 900ms

promisePool1([() => sleep(500), () => sleep(400)], 1)
.then(console.log) // After 900ms