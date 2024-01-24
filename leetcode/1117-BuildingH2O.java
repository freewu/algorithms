package leetcode 

// 1117. Building H2O
// There are two kinds of threads: oxygen and hydrogen. 
// Your goal is to group these threads to form water molecules.
// There is a barrier where each thread has to wait until a complete molecule can be formed. 
// Hydrogen and oxygen threads will be given releaseHydrogen and releaseOxygen methods respectively, 
// which will allow them to pass the barrier. These threads should pass the barrier in groups of three, 
// and they must immediately bond with each other to form a water molecule. 
// You must guarantee that all the threads from one molecule bond before any other threads from the next molecule do.
// In other words:

//         If an oxygen thread arrives at the barrier when no hydrogen threads are present, it must wait for two hydrogen threads.
//         If a hydrogen thread arrives at the barrier when no other threads are present, it must wait for an oxygen thread and another hydrogen thread.

// We do not have to worry about matching the threads up explicitly; 
// the threads do not necessarily know which other threads they are paired up with. 
// The key is that threads pass the barriers in complete sets; thus, if we examine the sequence of threads that bind and divide them into groups of three, each group should contain one oxygen and two hydrogen threads.

// Write synchronization code for oxygen and hydrogen molecules that enforces these constraints.

// Example 1:
// Input: water = "HOH"
// Output: "HHO"
// Explanation: "HOH" and "OHH" are also valid answers.

// Example 2:
// Input: water = "OOHHHH"
// Output: "HHOHHO"
// Explanation: "HOHHHO", "OHHHHO", "HHOHOH", "HOHHOH", "OHHHOH", "HHOOHH", "HOHOHH" and "OHHOHH" are also valid answers.
 

// Constraints:
//         3 * n == water.length
//         1 <= n <= 20
//         water[i] is either 'H' or 'O'.
//         There will be exactly 2 * n 'H' in water.
//         There will be exactly n 'O' in water.

// synchronized + notifyAll
class H2O {
    public H2O() {

    }
    private int flag = 0;
    public void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
		synchronized (this) {
            // 有两个H时等待 O的生成
            while (flag == 2) {
                this.wait();
            }
            // releaseHydrogen.run() outputs "H". Do not change or remove this line.
            releaseHydrogen.run();
            flag++;
            this.notifyAll();
        }
    }

    public void oxygen(Runnable releaseOxygen) throws InterruptedException {
        synchronized (this) {
            // 需要两个 H 才 生成一个 O
            while (flag != 2) {
                this.wait();
            }
            // releaseOxygen.run() outputs "O". Do not change or remove this line.
            releaseOxygen.run();
            flag = 0; // 开始的新的生产
            this.notifyAll();
        }
    }
}

// 使用 Semaphore
class H2O {
    private Semaphore h;
    private Semaphore o;
    private int number;
    public H2O() {
        h = new Semaphore(2);
        o = new Semaphore(0);
        number = 0;
    }

    public void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
		h.acquire();
        // releaseHydrogen.run() outputs "H". Do not change or remove this line.
        releaseHydrogen.run();
        number++;
        if (number % 2 == 0) { // 每两个 H 就需要一个 O
            o.release();
        }
    }

    public void oxygen(Runnable releaseOxygen) throws InterruptedException {
        o.acquire();
        // releaseOxygen.run() outputs "O". Do not change or remove this line.
		releaseOxygen.run();
        // release 两次
        h.release();
        h.release();
    }
}

// AtomicInteger
class H2O {
    private static final int []states={0,0,1}; 
    private int stateIndex;
    private final AtomicInteger state = new AtomicInteger(states[stateIndex]);
    public void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
        while(state.get() != 0)
            Thread.yield();
        releaseHydrogen.run();
        state.set(states[stateIndex=(stateIndex+1)%3]); // 0,1
    }
    public void oxygen(Runnable releaseOxygen) throws InterruptedException {
        while(state.get() != 1)
            Thread.yield();
        releaseOxygen.run();
        state.set(states[stateIndex=(stateIndex+1)%3]); // 2
    }
}