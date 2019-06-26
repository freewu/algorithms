#/usr/bin/env python
# -*- coding:utf-8 -*- 

# Given number n. Print number from 1 to n. But:

# when number is divided by 3, print "fizz".
# when number is divided by 5, print "buzz".
# when number is divided by both 3 and 5, print "fizz buzz".
# when number can't be divided by either 3 or 5, print the number itself.
# Example
# If n = 15, you should return:
# [
#   "1", "2", "fizz",
#   "4", "buzz", "fizz",
#   "7", "8", "fizz",
#   "buzz", "11", "fizz",
#   "13", "14", "fizz buzz"
# ]

# If n = 10, you should return:
# [
#   "1", "2", "fizz",
#   "4", "buzz", "fizz",
#   "7", "8", "fizz",
#   "buzz"
# ]

class Solution:
    """
    @param n: An integer
    @return: A list of strings.
    """
    def fizzBuzz(self, n):
        # write your code here
        arr = []
        for i in range(0,n):
            l = i + 1
            t3 = l % 3
            t5 = l % 5
            if 0 == t3 and 0 == t5:
                arr.append("fizz buzz")
            elif 0 == t3:
                arr.append("fizz")    
            elif 0 == t5:
                arr.append("buzz")
            else:
                arr.append(str(l))

        return arr

if __name__ == "__main__":
    s = Solution()
    print s.fizzBuzz(15)