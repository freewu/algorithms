#/usr/bin/env python
# -*- coding:utf-8 -*- 

# Count the number of k's between 0 and n. k can be 0 - 9.
# Example
# Example 1:

# Input:
# k = 1, n = 1
# Output:
# 1
# Explanation:
# In [0, 1], we found that 1 appeared once (1).
# Example 2:

# Input:
# k = 1, n = 12
# Output:
# 5
# Explanation:
# In [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12], we found that one appeared five times (1, 10, 11, 12)(Note that there are two 1

# 0 1  2  3  4  5  6  7  8  9
# 10 11 12 13 14 15 15 17 18 19
# 20 21 22 23 24 25 26 27 28 29

class Solution:
    """
    @param k: An integer
    @param n: An integer
    @return: An integer denote the count of digit k in 1..n
    """
    def digitCounts(self, k, n):
        # write your code here
        c = 0
        for i in range(1,n + 1):
            if i < 10:
                if i == k: c = c + 1
            else:
                while (i / 10) >= 1:
                    if i < 10:
                        if i == k: c = c + 1
                        break

                    t = i % 10
                    # print t
                    if t == k: c = c + 1
                    i = i / 10
        return c

if __name__ == "__main__":
    s = Solution()
    print s.digitCounts(1,11)
    print s.digitCounts(8,11)