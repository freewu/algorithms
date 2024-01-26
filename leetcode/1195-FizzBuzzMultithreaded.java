package leetcode

// 1195. Fizz Buzz Multithreaded
// You have the four functions:
//         printFizz that prints the word "fizz" to the console,
//         printBuzz that prints the word "buzz" to the console,
//         printFizzBuzz that prints the word "fizzbuzz" to the console, and
//         printNumber that prints a given integer to the console.

// You are given an instance of the class FizzBuzz that has four functions: 
//         fizz, buzz, fizzbuzz and number. 

// The same instance of FizzBuzz will be passed to four different threads:
//         Thread A: calls fizz() that should output the word "fizz".
//         Thread B: calls buzz() that should output the word "buzz".
//         Thread C: calls fizzbuzz() that should output the word "fizzbuzz".
//         Thread D: calls number() that should only output the integers.

// Modify the given class to output the series [1, 2, "fizz", 4, "buzz", ...] where the ith token (1-indexed) of the series is:
//         "fizzbuzz" if i is divisible by 3 and 5,
//         "fizz" if i is divisible by 3 and not 5,
//         "buzz" if i is divisible by 5 and not 3, or
//         i if i is not divisible by 3 or 5.

// Implement the FizzBuzz class:
//         FizzBuzz(int n) Initializes the object with the number n that represents the length of the sequence that should be printed.
//         void fizz(printFizz) Calls printFizz to output "fizz".
//         void buzz(printBuzz) Calls printBuzz to output "buzz".
//         void fizzbuzz(printFizzBuzz) Calls printFizzBuzz to output "fizzbuzz".
//         void number(printNumber) Calls printnumber to output the numbers.
 

// Example 1:
// Input: n = 15
// Output: [1,2,"fizz",4,"buzz","fizz",7,8,"fizz","buzz",11,"fizz",13,14,"fizzbuzz"]

// Example 2:
// Input: n = 5
// Output: [1,2,"fizz",4,"buzz"]

// Constraints:
//         1 <= n <= 50


// public CyclicBarrier(int parties)
// public CyclicBarrier(int parties, Runnable barrierAction)
// // -构造方法
// //parties 是参与线程的个数
// //第二个构造方法有一个 Runnable 参数，这个参数的意思是最后一个到达线程要做的任务
//   ---
// public int await() throws InterruptedException, BrokenBarrierException
// public int await(long timeout, TimeUnit unit) throws InterruptedException, BrokenBarrierException, TimeoutException
// //- 函数
// //线程调用 await() 表示自己已经到达栅栏
// //BrokenBarrierException 表示栅栏已经被破坏，破坏的原因可能是其中一个线程 await() 时被中断或者超时
// //调用await方法的线程告诉CyclicBarrier自己已经到达同步点，然后当前线程被阻塞。直到parties个参与线程调用了await方法
// CyclicBarrier 与 CountDownLatch 区别
//      CountDownLatch 是一次性的，CyclicBarrier 是可循环利用的
//      CountDownLatch 参与的线程的职责是不一样的，有的在倒计时，有的在等待倒计时结束。CyclicBarrier 参与的线程职责是一样的
// CyclicBarrier
class FizzBuzz {
    private int n;
    private CyclicBarrier cb = new CyclicBarrier(4);

    public FizzBuzz(int n) {
        this.n = n;
    }

    // printFizz.run() outputs "fizz".
    public void fizz(Runnable printFizz) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            // 如果这个数字可以被 3 整除，输出 "fizz" 
            if (i % 3 == 0 && i % 5 != 0) {
                printFizz.run();
            }
            try {
                cb.await();
            } catch (BrokenBarrierException e) {
                e.printStackTrace();
            }
        }
    }

    // printBuzz.run() outputs "buzz".
    public void buzz(Runnable printBuzz) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            // 如果这个数字可以被 5 整除，输出 "buzz"
            if (i % 3 != 0 && i % 5 == 0) {
                printBuzz.run();
            }
            try {
                cb.await();
            } catch (BrokenBarrierException e) {
                e.printStackTrace();
            }
        }
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            // 如果这个数字可以同时被 3 和 5 整除，输出 "fizzbuzz"
            if (i % 3 == 0 && i % 5 == 0) {
                printFizzBuzz.run();
            }
            try {
                cb.await();
            } catch (BrokenBarrierException e) {
                e.printStackTrace();
            }
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void number(IntConsumer printNumber) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            // 输出既不能被 3 整除也不能被 5 整除的数字
            if (i % 3 != 0 && i % 5 != 0) {
                printNumber.accept(i);
            }
            try {
                cb.await();
            } catch (BrokenBarrierException e) {
                e.printStackTrace();
            }
        }
    }
}

// Semaphore是用来保护一个或者多个共享资源的访问，Semaphore内部维护了一个计数器，其值为可以访问的共享资源的个数。
// 一个线程要访问共享资源，先获得信号量，如果信号量的计数器值大于1，意味着有共享资源可以访问，则使其计数器值减去1，再访问共享资源。
// 如果计数器值为0,线程进入休眠。当某个线程使用完共享资源后，释放信号量，并将信号量内部的计数器加1，之前进入休眠的线程将被唤醒并再次试图获得信号量
// Semaphore
class FizzBuzz {
    private int n;

    private Semaphore number = new Semaphore(1);
    private Semaphore fizz = new Semaphore(0);
    private Semaphore buzz = new Semaphore(0);
    private Semaphore fizzbuzz = new Semaphore(0);


    public FizzBuzz(int n) {
        this.n = n;
    }

    // printFizz.run() outputs "fizz".
    public void fizz(Runnable printFizz) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            if (i % 3 == 0 && i % 5 != 0) {
                fizz.acquire();
                printFizz.run();
                number.release();
            }
        }
    }

    // printBuzz.run() outputs "buzz".
    public void buzz(Runnable printBuzz) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            if (i % 3 != 0 && i % 5 == 0) {
                buzz.acquire();
                printBuzz.run();
                number.release();
            }
        }
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            if (i % 3 == 0 && i % 5 == 0) {
                fizzbuzz.acquire();
                printFizzBuzz.run();
                number.release();
            }
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void number(IntConsumer printNumber) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            number.acquire();
            if (i % 3 != 0 && i % 5 != 0) { // 输出既不能被 3 整除也不能被 5 整除的数字
                printNumber.accept(i);
                number.release();
            } else if (i % 3 == 0 && i % 5 != 0) { // 可以被 3 整除，输出 "fizz"
                fizz.release();
            } else if (i % 3 != 0 && i % 5 == 0) { // 可以被 5 整除，输出 "buzz"
                buzz.release();
            } else {
                fizzbuzz.release(); // 可以同时被 3 和 5 整除，输出 "fizzbuzz"
            }
        }
    }
}


// 只使用一个 Semaphore
class FizzBuzz {
    private int n;

    private Semaphore semaphore = new Semaphore(1);
    private int cur = 1;

    public FizzBuzz(int n) {
        this.n = n;
    }

    // printFizz.run() outputs "fizz".
    public void fizz(Runnable printFizz) throws InterruptedException {
        while (true) {
            semaphore.acquire(1);
            try {
                // 原因就在这里，循环过程中如果打印的字符串个数已经满足要求，那么会使用return来返回，终止该方法的执行。
                // 但是咱们已经获取了信号量，那么在方法返回前需要释放该信号量，否则会导致其它线程一直等待，整个程序一直不结束。
                // Java语言中try-finally可以做到这一点，try-finally代码块也是常用的一种释放资源（IO流、数据库连接等）的方式。
                // 不是程序死循环，而是其它线程在wait，导致无法退出。
                if (cur > n) return;
                if (cur % 3 == 0 && cur % 5 != 0) {
                    cur++;
                    printFizz.run();
                }
            } finally {
                semaphore.release(1);
            }
        }
    }

    // printBuzz.run() outputs "buzz".
    public void buzz(Runnable printBuzz) throws InterruptedException {
        while (true) {
            semaphore.acquire(1);
            try {
                if (cur > n) return;
                if (cur % 3 != 0 && cur % 5 == 0) {
                    cur++;
                    printBuzz.run();
                }
            } finally {
                semaphore.release(1);
            }
        }
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        while (true) {
            semaphore.acquire(1);
            try {
                if (cur > n) return;
                if (cur % 3 == 0 && cur % 5 == 0) {
                    cur++;
                    printFizzBuzz.run();
                }
            } finally {
                semaphore.release(1);
            }
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void number(IntConsumer printNumber) throws InterruptedException {
        while (true) {
            semaphore.acquire(1);
            try {
                if (cur > n) return;
                if (cur % 3 != 0 && cur % 5 != 0) {
                    printNumber.accept(cur);
                    cur++;
                }
            } finally {
                semaphore.release(1);
            }
        }
    }
}

// Thread.yield()
 class FizzBuzz {
    private int n;
    private volatile int state = 0; 
    public FizzBuzz(int n) {
        this.n = n;
    }

    public void fizz(Runnable printFizz) throws InterruptedException {
        for (int i = 3; i <= n; i += 3) {   //只输出3的倍数(不包含15的倍数)
            if (i % 15 == 0) continue;   //15的倍数不处理，交给fizzbuzz()方法处理
            while (state != 3)
                Thread.yield();

            printFizz.run();
            state = 0;
        }
    }

    public void buzz(Runnable printBuzz) throws InterruptedException {
        for (int i = 5; i <= n; i += 5) {   //只输出5的倍数(不包含15的倍数)
            if (i % 15 == 0)    //15的倍数不处理，交给fizzbuzz()方法处理
                continue;
            while (state != 5)
                Thread.yield();
            printBuzz.run();
            state = 0;
        }
    }

    public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        for (int i = 15; i <= n; i += 15) {   //只输出15的倍数
            while (state != 15)
                Thread.yield();
            printFizzBuzz.run();
            state = 0;
        }
    }

    public void number(IntConsumer printNumber) throws InterruptedException {
        for (int i = 1; i <= n; ++i) {
            while (state != 0)
                Thread.yield();
            if (i % 3 != 0 && i % 5 != 0)
                printNumber.accept(i);
            else {
                if (i % 15 == 0)
                    state = 15;    //交给fizzbuzz()方法处理
                else if (i % 5 == 0)
                    state = 5;    //交给buzz()方法处理
                else
                    state = 3;    //交给fizz()方法处理
            }
        }
    }
}

// ReentrantLock+Condition
class FizzBuzz {
    private int n;
    private ReentrantLock lock = new ReentrantLock();
    int state = 0;
    private Condition condition = lock.newCondition();

    public FizzBuzz(int n) {
        this.n = n;
    }

    // printFizz.run() outputs "fizz".
    public void fizz(Runnable printFizz) throws InterruptedException {
        for (int i = 3; i <= n; i += 3) {
            try {
                if (i % 3 == 0 && i % 5 == 0) continue;
                lock.lock();
                while (state != 3) {
                    condition.await();
                }
                printFizz.run();
                state = 0;
                condition.signalAll();
            } finally {
                lock.unlock();
            }
        }
    }

    // printBuzz.run() outputs "buzz".
    public void buzz(Runnable printBuzz) throws InterruptedException {
        for (int i = 5; i <= n; i += 5) {
            try {
                if (i % 3 == 0 && i % 5 == 0) continue;
                lock.lock();
                while (state != 5) {
                    condition.await();
                }
                printBuzz.run();
                state = 0;
                condition.signalAll();
            } finally {
                lock.unlock();
            }
        }
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        for (int i = 15; i <= n; i += 15) {
            try {
                lock.lock();
                while (state != 15) {
                    condition.await();
                }
                printFizzBuzz.run();
                state = 0;
                condition.signalAll();
            } finally {
                lock.unlock();
            }
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void number(IntConsumer printNumber) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            try {
                lock.lock();
                while (state != 0) {
                    condition.await();
                }
                if (i % 3 != 0 && i % 5 != 0) {
                    printNumber.accept(i);
                } else {
                    if (i % 3 == 0 && i % 5 == 0) state = 15;
                    else if (i % 3 == 0) state = 3;
                    else if (i % 5 == 0) state = 5;
                    condition.signalAll();
                }
            } finally {
                lock.unlock();
            }
        }
    }
}

// synchronized + wait + notifyAll
class FizzBuzz {
    
    private int n;
    private int currentNumber = 1;
    
    public FizzBuzz(int n) {
        this.n = n;
    }

    // printFizz.run() outputs "fizz".
    public synchronized void fizz(Runnable printFizz) throws InterruptedException {
        while (currentNumber <= n) {
            if (currentNumber % 3 != 0 || currentNumber % 5 == 0) {
                wait();
                continue;
            }
            printFizz.run();
            currentNumber += 1;
            notifyAll();
        }
    }

    // printBuzz.run() outputs "buzz".
    public synchronized void buzz(Runnable printBuzz) throws InterruptedException {
        while (currentNumber <= n) {
            if (currentNumber % 5 != 0 || currentNumber % 3 == 0) {
                wait();
                continue;
            }
            printBuzz.run();
            currentNumber += 1;
            notifyAll();
        }
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public synchronized void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        while (currentNumber <= n) {
            if (currentNumber % 15 != 0) {
                wait();
                continue;
            }
            printFizzBuzz.run();
            currentNumber += 1;
            notifyAll();
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public synchronized void number(IntConsumer printNumber) throws InterruptedException {
        while (currentNumber <= n) {
            if (currentNumber % 3 == 0 || currentNumber % 5 == 0) {
                wait();
                continue;
            }
            printNumber.accept(currentNumber);
            currentNumber += 1;
            notifyAll();
        }
    }
}



class FizzBuzzMultithreaded {
    public static void main(String[] args) {
        Runnable printFizz = () -> {
            System.out.printf("%s", "fizz");
        };
        Runnable printBuzz = () -> {
            System.out.printf("%s", "buzz");
        };
        Runnable printFizzBuzz = () -> {
            System.out.printf("%s", "fizzbuzz");
        };
        IntConsumer intConsumer = new IntConsumer();
        FizzBuzz fb = new FizzBuzz(15);
        new Thread(() -> {
            try {
                fb.fizz(printFizz);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();
        new Thread(() -> {
            try {
                fb.buzz(printBuzz);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();
        new Thread(() -> {
            try {
                fb.fizzbuzz(printFizzBuzz);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();
        new Thread(() -> {
            try {
                fb.number(intConsumer);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();

    }


    public class IntConsumer {
        public void accept(int i) {
            System.out.printf("%d", i);
        }
    }

}