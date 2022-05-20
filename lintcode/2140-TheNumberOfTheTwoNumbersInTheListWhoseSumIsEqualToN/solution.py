# 2140 · The number of the two numbers in the list whose sum is equal to n
# Description
# Python's for loop can traverse items of any sequence , such as a list or a string, 
# and can traverse the elements in the list through a loop. For the list type variable lst, there are:
#   for i in range(len(lst)):
#       pass
# Here len(lst) is the length of the operation list, 
# range is an iterable list that can be produced, and pass is a placeholder

# This question has a set of ordered array lst whose input type is list (the elements are arranged in ascending order)

# lst = [1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 18, 19, 20, 21, 29, 34, 54, 65]
# And a target value n whose input type is integer

# You are required to find the position of two numbers in the list lst whose sum is equal to the variable n by using a double for loop in this task, 
# and you only need to print out a set of data, and require the first number to be as small as possible.
# You need to complete the code at the specific location in solution.py, 
# we will check in main.py to determine whether your code find the data that meets the requirements correctly, 
# if found, the program will output the position number of this group of data, if If all data does not meet the condition, output "not found".

# The output is the position of the data in the list and not the data itself
# The program will return not found if no data is found.

# Example
# When n = 9, your code should output:
#    (0, 5)
# When n = 4, your code should output:
#   not found

# When n = 5, your code should output:
#   (0, 1)

def twonums_sum(n, lst):
# -- write your code here --
    for i in range(len(lst)):
        for j in range(len(lst)):
            if (lst[i] + lst[j] ) == n:
                return print("({i}, {j})".format(i = i,j = j))
    print("not found")