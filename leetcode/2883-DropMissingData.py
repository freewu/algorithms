# 2883. Drop Missing Data
# DataFrame students
# +-------------+--------+
# | Column Name | Type   |
# +-------------+--------+
# | student_id  | int    |
# | name        | object |
# | age         | int    |
# +-------------+--------+
# There are some rows having missing values in the name column.
# Write a solution to remove the rows with missing values.
# The result format is in the following example.

# Example 1:

# Input:
# +------------+---------+-----+
# | student_id | name    | age |
# +------------+---------+-----+
# | 32         | Piper   | 5   |
# | 217        | None    | 19  |
# | 779        | Georgia | 20  |
# | 849        | Willow  | 14  |
# +------------+---------+-----+
# Output:
# +------------+---------+-----+
# | student_id | name    | age |
# +------------+---------+-----+
# | 32         | Piper   | 5   |
# | 779        | Georgia | 20  | 
# | 849        | Willow  | 14  | 
# +------------+---------+-----+
# Explanation: 
# Student with id 217 havs empty value in the name column, so it will be removed.

import pandas as pd

def dropMissingData(students: pd.DataFrame) -> pd.DataFrame:
    #return students[students["name"] != None]
    #return students.dropna(subset='name',how='any',axis=0)
    students.dropna(subset=['name'], inplace=True)
    return students
    

if __name__ == "__main__":
    l = [
        [101, 1, 15],
        [101, 2, 11],
        [103, None, 11],
        [104, 4, 20]
    ]
    print(dropMissingData(pd.DataFrame(l,columns=["email","name","age"])))