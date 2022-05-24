# 2252 · Judgment Results
# Description
# The problem has a variable marks and we need you to give a grade for that grade based on marks, refine the code in solution.py to achieve this.

# if marks is less than 40, print Maybe you should work hard and assign grade the value No grade.
# if marks is greater than or equal to 90, 80, 70, 60, 40, assign grade to A B C D E.
# finally return the grade result grade.
# We will import your code from solution.py in main.py and run it, if your code logic is correct,
# the program will output the grade corresponding to the grade marks.

# mask is a positive integer greater than zero and less than one hundred
# Example
# The evaluator executes your code by executing the command python main.py {marks}, 
# passing marks as a command line argument, and you can see how the code is run in main.py.

# Example 1 
# When marks = 55, the result you get is
#   E

# Example 2
# When marks = 0, the result obtained is
#   Maybe you should work hard
#   No grade

def mark_jud(marks: int) -> str:
    """
    :param marks: The source int
    :return: A str of grade
    """
    # --write your code here--
    if (marks >= 90): return "A"
    if (marks >= 80): return "B"
    if (marks >= 70): return "C"
    if (marks >= 60): return "D"
    if (marks >= 40): return "E"
    return "Maybe you should work hard\nNo grade"

if __name__ == "__main__":
    print(mark_jud(55)) # E
    print(mark_jud(0)) # No grade