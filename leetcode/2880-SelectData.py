# 2880. Select Data
# DataFrame students
# +-------------+--------+
# | Column Name | Type   |
# +-------------+--------+
# | student_id  | int    |
# | name        | object |
# | age         | int    |
# +-------------+--------+

# Write a solution to select the name and age of the student with student_id = 101.
# The result format is in the following example.

# Example 1:
# Input:
# +------------+---------+-----+
# | student_id | name    | age |
# +------------+---------+-----+
# | 101        | Ulysses | 13  |
# | 53         | William | 10  |
# | 128        | Henry   | 6   |
# | 3          | Henry   | 11  |
# +------------+---------+-----+
# Output:
# +---------+-----+
# | name    | age | 
# +---------+-----+
# | Ulysses | 13  |
# +---------+-----+
# Explanation:
# Student Ulysses has student_id = 101, we select the name and age.
import pandas as pd

def selectData(students: pd.DataFrame) -> pd.DataFrame:
    # return students.loc[students['student_id']==101,['name','age']]
    # 指定条件的行 student_id = 101 
    s = students[students["student_id"] == 101]
    # 指定列
    return s[["name","age"]]

if __name__ == "__main__":
    l = [
        [101, 1, 15],
        [102, 2, 11],
        [103, 3, 11],
        [104, 4, 20]
    ]
    print(selectData(pd.DataFrame(l,columns=["student_id","name","age"])))