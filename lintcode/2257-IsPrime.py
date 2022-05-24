# 2257 · Determining whether a number is prime or not
# # Description
# Write Python code that implements a function named prime. 
# This function will determine if the integer num is a prime number, and return True if it is, or False if it is not.
# Write the code for the prime function in solution.py,
#  and we will run your code in main.py by importing it to check if it does the above correctly..

# The range of values of the integer num is 0 <= num <= 10000 
# A prime number is a number that is not divisible by any other natural number except 1 and itself
# 0 and 1 are neither prime nor composite
# Example
# The evaluator will execute your code by executing python main.py {num} and main.py will print out the result statement of the judgment.

# Example 1
# If the input data is:
#   11
# then the output is:
#   11 is a prime number

# Example 2
# If the input data is:
#   0
# then the output is:
#   0 isn't a prime number

def prime(num: int) -> int:
    """
    :param num: a random integer
    :return: determine if the result is a prime number and return 1 otherwise other values
    """
	# write your code here
    if num == 0 or num == 1:
        return 0
    else:
        for i in range(2, int(num / 2)): # 取一半数据就行了
            if num % i == 0:
                return 0
        return 1

def test(num: int):
    n = prime(num)
    if n:
        print("%d is a prime number" % num)
    else:
        print("%d isn't a prime number" % num)

if __name__ == "__main__": 
    test(0)
    test(11)