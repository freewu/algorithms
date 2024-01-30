# 2886. Change Data Type
# DataFrame students
# +-------------+--------+
# | Column Name | Type   |
# +-------------+--------+
# | student_id  | int    |
# | name        | object |
# | age         | int    |
# | grade       | float  |
# +-------------+--------+
# Write a solution to correct the errors:
# The grade column is stored as floats, convert it to integers.
# The result format is in the following example.

# Example 1:
# Input:
# DataFrame students:
# +------------+------+-----+-------+
# | student_id | name | age | grade |
# +------------+------+-----+-------+
# | 1          | Ava  | 6   | 73.0  |
# | 2          | Kate | 15  | 87.0  |
# +------------+------+-----+-------+
# Output:
# +------------+------+-----+-------+
# | student_id | name | age | grade |
# +------------+------+-----+-------+
# | 1          | Ava  | 6   | 73    |
# | 2          | Kate | 15  | 87    |
# +------------+------+-----+-------+
# Explanation: 
# The data types of the column grade is converted to int.

import pandas as pd

def changeDatatype(students: pd.DataFrame) -> pd.DataFrame:
    students['grade'] = students['grade'].astype(int)
    return students

if __name__ == "__main__":
    l = [
        [101, 1, 15,1.0],
        [101, 2, 11,2.0],
        [103, 3, 11,3.0],
        [104, 4, 20,4.0]
    ]
    print(changeDatatype(pd.DataFrame(l,columns=["id","first","last","grade"])))