#/usr/bin/env python
# -*- coding:utf-8 -*- 

import random
import time

# 0-5
def rand5():
    return (int)((random.random() * 100) % 5)

# 0-7
def rand7():
    while(1):
        # 构造等概率的0, 1, 2, 3, 4, 5, 6, 6, 7, 8, ..., 20, 21, 22, 23, 24, 25
        x = rand5() * 5 + rand5()
        if x >= 21: # 剔除21, 22, 23, 24,25  只留 0-20
            continue
        else:
            return x % 7

if __name__ == "__main__":
    start = time.time()
    len = 100000
    a = [0]* 5
    for i in range(0,len):
        val = rand5()
        a[val] = a[val] +1

    for i in range(0,5):
        print "%s : %f : %d" % (i, float(a[i]) / float(len), a[i])
    
    print "耗时: %f " % (time.time() - start)

    start = time.time()
    b = [0] * 7
    for i in range(0,len):
        val = rand7()
        b[val] = b[val] + 1

    for i in range(0,7):
        print "%s : %f : %d" % (i, float(b[i]) / float(len) , b[i])
    
    print "耗时: %f " % (time.time() - start)

    # for i in range(0,21):
    #     print "%d : %d" % (i, i % 7)

