// 2649. Nested Array Generator
// Given a multi-dimensional array of integers, 
// return a generator object which yields integers in the same order as inorder traversal.

// A multi-dimensional array is a recursive data structure that contains both integers and other multi-dimensional arrays.

// inorder traversal iterates over each array from left to right, 
// yielding any integers it encounters or applying inorder traversal to any arrays it encounters.

// Example 1:
// Input: arr = [[[6]],[1,3],[]]
// Output: [6,1,3]
// Explanation:
//         const generator = inorderTraversal(arr);
//         generator.next().value; // 6
//         generator.next().value; // 1
//         generator.next().value; // 3
//         generator.next().done; // true

// Example 2:
// Input: arr = []
// Output: []
// Explanation: 
//         There are no integers so the generator doesn't yield anything.

// Constraints:
//         0 <= arr.flat().length <= 10^5
//         0 <= arr.flat()[i] <= 10^5
//         maxNestingDepth <= 10^5

/**
 * @param {Array} arr
 * @return {Generator}
 */
var inorderTraversal = function*(arr) {
    // 递归把多维数据平铺成一维数组
    const arrayFlatten = (array, ans = new Array()) => {
        for(const item of array) {
            if(Array.isArray(item)) {
                arrayFlatten(item, ans);
            } else {
                ans.push(item);
            }
        }
        return ans;
    }
    // 返回一个 Generator
    for(const item of arrayFlatten(arr)) {
        yield item;
    }
};

// best solution
var inorderTraversal1 = function*(arr) {
    const queue = []
    for(let i = 0; i < arr.length; i++) {
        if (Array.isArray(arr[i])) {
            queue.unshift(...arr[i])
        } else {
            yield arr[i]
        }
        while(queue.length) {
            const item = queue.shift()
            if (Array.isArray(item)) {
                queue.unshift(...item)
            } else {
                yield item
            }
        }
    }
};

/**
 * const gen = inorderTraversal([1, [2, 3]]);
 * gen.next().value; // 1
 * gen.next().value; // 2
 * gen.next().value; // 3
 */
let gen = inorderTraversal([1, [2, 3]]);
console.log(gen.next().value); // 1
console.log(gen.next().value); // 2
console.log(gen.next().value); // 3

let generator = inorderTraversal([[[6]],[1,3],[]]);
console.log(generator.next().value); // 6
console.log(generator.next().value); // 1
console.log(generator.next().value); // 3
console.log(generator.next().done); // true

gen = inorderTraversal1([1, [2, 3]]);
console.log(gen.next().value); // 1
console.log(gen.next().value); // 2
console.log(gen.next().value); // 3

generator = inorderTraversal1([[[6]],[1,3],[]]);
console.log(generator.next().value); // 6
console.log(generator.next().value); // 1
console.log(generator.next().value); // 3
console.log(generator.next().done); // true