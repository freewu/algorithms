package leetcode

// 1226. The Dining Philosophers
// Five silent philosophers sit at a round table with bowls of spaghetti. 
// Forks are placed between each pair of adjacent philosophers.
// Each philosopher must alternately think and eat. 
// However, a philosopher can only eat spaghetti when they have both left and right forks. 
// Each fork can be held by only one philosopher and so a philosopher can use the fork only if it is not being used by another philosopher.
// After an individual philosopher finishes eating, they need to put down both forks so that the forks become available to others. A philosopher can take the fork on their right or the one on their left as they become available, but cannot start eating before getting both forks.
// Eating is not limited by the remaining amounts of spaghetti or stomach space; 
// an infinite supply and an infinite demand are assumed.

// Design a discipline of behaviour (a concurrent algorithm) such that no philosopher will starve; i.e., each can forever continue to alternate between eating and thinking, assuming that no philosopher can know when others may want to eat or think.

// The problem statement and the image above are taken from wikipedia.org

// The philosophers' ids are numbered from 0 to 4 in a clockwise order. Implement the function void wantsToEat(philosopher, pickLeftFork, pickRightFork, eat, putLeftFork, putRightFork) where:

// philosopher is the id of the philosopher who wants to eat.
// pickLeftFork and pickRightFork are functions you can call to pick the corresponding forks of that philosopher.
// eat is a function you can call to let the philosopher eat once he has picked both forks.
// putLeftFork and putRightFork are functions you can call to put down the corresponding forks of that philosopher.
// The philosophers are assumed to be thinking as long as they are not asking to eat (the function is not being called with their number).
// Five threads, each representing a philosopher, will simultaneously use one object of your class to simulate the process. The function may be called for the same philosopher more than once, even before the last call ends.

// Example 1:
// Input: n = 1
// Output: [[4,2,1],[4,1,1],[0,1,1],[2,2,1],[2,1,1],[2,0,3],[2,1,2],[2,2,2],[4,0,3],[4,1,2],[0,2,1],[4,2,2],[3,2,1],[3,1,1],[0,0,3],[0,1,2],[0,2,2],[1,2,1],[1,1,1],[3,0,3],[3,1,2],[3,2,2],[1,0,3],[1,1,2],[1,2,2]]
// Explanation:
// n is the number of times each philosopher will call the function.
// The output array describes the calls you made to the functions controlling the forks and the eat function, its format is:
// output[i] = [a, b, c] (three integers)
// - a is the id of a philosopher.
// - b specifies the fork: {1 : left, 2 : right}.
// - c specifies the operation: {1 : pick, 2 : put, 3 : eat}.
 
// Constraints:

//         1 <= n <= 60

// n 表示每个哲学家需要进餐的次数。
// 输出数组描述了叉子的控制和进餐的调用，它的格式如下：
// output[i] = [a, b, c] (3个整数)
// - a 哲学家编号。
// - b 指定叉子：{1 : 左边, 2 : 右边}.
// - c 指定行为：{1 : 拿起, 2 : 放下, 3 : 吃面}。
// 如 [4,2,1] 表示 4 号哲学家拿起了右边的叉子。

// ReentrantLock + Semaphore
class DiningPhilosophers {
    // 1个Fork视为1个ReentrantLock，5个叉子即5个ReentrantLock，将其都放入数组中
    private final ReentrantLock[] lockList = {
        new ReentrantLock(),
        new ReentrantLock(),
        new ReentrantLock(),
        new ReentrantLock(),
        new ReentrantLock()
    };

    // 限制 最多只有4个哲学家去持有叉子
    private Semaphore eatLimit = new Semaphore(4);

    public DiningPhilosophers() {

    }

    // call the run() method of any runnable to execute its code
    public void wantsToEat(int philosopher,
                           Runnable pickLeftFork,
                           Runnable pickRightFork,
                           Runnable eat,
                           Runnable putLeftFork,
                           Runnable putRightFork) throws InterruptedException {

        int leftFork = (philosopher + 1) % 5;    // 左边的叉子 的编号
        int rightFork = philosopher;    // 右边的叉子 的编号

        eatLimit.acquire();    // 限制的人数 -1

        lockList[leftFork].lock();    // 拿起左边的叉子
        lockList[rightFork].lock();    // 拿起右边的叉子

        pickLeftFork.run();    // 拿起左边的叉子 的具体执行
        pickRightFork.run();    //拿起右边的叉子 的具体执行

        eat.run();    // 吃意大利面 的具体执行

        putLeftFork.run();    // 放下左边的叉子 的具体执行
        putRightFork.run();    // 放下右边的叉子 的具体执行

        lockList[leftFork].unlock();    // 放下左边的叉子
        lockList[rightFork].unlock();    // 放下右边的叉子

        eatLimit.release(); // 限制的人数 +1
    }
}

// 只允许1个哲学家就餐 one ReentrantLock
class DiningPhilosophers {
    // 只允许1个哲学家就餐
    private ReentrantLock pickBothForks = new ReentrantLock();

    public DiningPhilosophers() {

    }

    // call the run() method of any runnable to execute its code
    public void wantsToEat(int philosopher,
                           Runnable pickLeftFork,
                           Runnable pickRightFork,
                           Runnable eat,
                           Runnable putLeftFork,
                           Runnable putRightFork) throws InterruptedException {

        int leftFork = (philosopher + 1) % 5;    //左边的叉子 的编号
        int rightFork = philosopher;    //右边的叉子 的编号

        // 拿起左右叉子 + 吃意面 + 放下左右叉子 一套流程走完之后才退出临界区
        pickBothForks.lock();    // 进入临界区

        pickLeftFork.run();    //拿起左边的叉子 的具体执行
        pickRightFork.run();    //拿起右边的叉子 的具体执行

        eat.run();    //吃意大利面 的具体执行

        putLeftFork.run();    //放下左边的叉子 的具体执行
        putRightFork.run();    //放下右边的叉子 的具体执行

        pickBothForks.unlock();    // 退出临界区
    }
}

// 如何避免死锁。
// 而当5个哲学家都左手持有其左边的叉子 或 当5个哲学家都右手持有其右边的叉子时，会发生死锁。
// 故只需设计1个避免发生上述情况发生的策略即可。
// 即可以让一部分哲学家优先去获取其左边的叉子，再去获取其右边的叉子；
// 再让剩余哲学家优先去获取其右边的叉子，再去获取其左边的叉子
class DiningPhilosophers {
    //1个Fork视为1个ReentrantLock，5个叉子即5个ReentrantLock，将其都放入数组中
    private final ReentrantLock[] lockList = {
        new ReentrantLock(),
        new ReentrantLock(),
        new ReentrantLock(),
        new ReentrantLock(),
        new ReentrantLock()
    };

    public DiningPhilosophers() {

    }

    // call the run() method of any runnable to execute its code
    public void wantsToEat(int philosopher,
                           Runnable pickLeftFork,
                           Runnable pickRightFork,
                           Runnable eat,
                           Runnable putLeftFork,
                           Runnable putRightFork) throws InterruptedException {

        int leftFork = (philosopher + 1) % 5;    //左边的叉子 的编号
        int rightFork = philosopher;    //右边的叉子 的编号

        // 编号为偶数的哲学家，优先拿起左边的叉子，再拿起右边的叉子
        if (philosopher % 2 == 0) {
            lockList[leftFork].lock();    //拿起左边的叉子
            lockList[rightFork].lock();    //拿起右边的叉子
        } else { // 编号为奇数的哲学家，优先拿起右边的叉子，再拿起左边的叉子
            lockList[rightFork].lock();    //拿起右边的叉子
            lockList[leftFork].lock();    //拿起左边的叉子
        }

        pickLeftFork.run();    //拿起左边的叉子 的具体执行
        pickRightFork.run();    //拿起右边的叉子 的具体执行

        eat.run();    //吃意大利面 的具体执行

        putLeftFork.run();    //放下左边的叉子 的具体执行
        putRightFork.run();    //放下右边的叉子 的具体执行

        lockList[leftFork].unlock();    //放下左边的叉子
        lockList[rightFork].unlock();    //放下右边的叉子
    }
}

// 位运算就可以表示5个叉子的使用状态，只需用1个volatile修饰的int变量即可 + CAS操作即可 AtomicInteger
class DiningPhilosophers {
    // 初始化为0, 二进制表示则为00000, 说明当前所有叉子都未被使用
    private AtomicInteger fork = new AtomicInteger(0);
    // 每个叉子的int值(即二进制的00001, 00010, 00100, 01000, 10000)
    private final int[] forkMask = new int[]{1, 2, 4, 8, 16};
    // 限制 最多只有4个哲学家去持有叉子
    private Semaphore eatLimit = new Semaphore(4);

    public DiningPhilosophers() {

    }

    // call the run() method of any runnable to execute its code
    public void wantsToEat(int philosopher,
                           Runnable pickLeftFork,
                           Runnable pickRightFork,
                           Runnable eat,
                           Runnable putLeftFork,
                           Runnable putRightFork) throws InterruptedException {

        int leftMask = forkMask[(philosopher + 1) % 5], rightMask = forkMask[philosopher];
        eatLimit.acquire();    //限制的人数 -1

        while (!pickFork(leftMask)) Thread.sleep(1);    //拿起左边的叉子
        while (!pickFork(rightMask)) Thread.sleep(1);   //拿起右边的叉子

        pickLeftFork.run();    //拿起左边的叉子 的具体执行
        pickRightFork.run();    //拿起右边的叉子 的具体执行

        eat.run();    //吃意大利面 的具体执行

        putLeftFork.run();    //放下左边的叉子 的具体执行
        putRightFork.run();    //放下右边的叉子 的具体执行

        while (!putFork(leftMask)) Thread.sleep(1);     //放下左边的叉子
        while (!putFork(rightMask)) Thread.sleep(1);    //放下右边的叉子

        eatLimit.release(); //限制的人数 +1
    }

    private boolean pickFork(int mask) {
        int expect = fork.get();
        return (expect & mask) > 0 ? false : fork.compareAndSet(expect, expect ^ mask);
    }

    private boolean putFork(int mask) {
        int expect = fork.get();
        return fork.compareAndSet(expect, expect ^ mask);
    }
}