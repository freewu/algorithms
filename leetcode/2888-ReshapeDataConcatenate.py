# 2888. Reshape Data: Concatenate
# DataFrame df1
# +-------------+--------+
# | Column Name | Type   |
# +-------------+--------+
# | student_id  | int    |
# | name        | object |
# | age         | int    |
# +-------------+--------+

# DataFrame df2
# +-------------+--------+
# | Column Name | Type   |
# +-------------+--------+
# | student_id  | int    |
# | name        | object |
# | age         | int    |
# +-------------+--------+

# Write a solution to concatenate these two DataFrames vertically into one DataFrame.
# The result format is in the following example.

# Example 1:
# Input:
# df1
# +------------+---------+-----+
# | student_id | name    | age |
# +------------+---------+-----+
# | 1          | Mason   | 8   |
# | 2          | Ava     | 6   |
# | 3          | Taylor  | 15  |
# | 4          | Georgia | 17  |
# +------------+---------+-----+
# df2
# +------------+------+-----+
# | student_id | name | age |
# +------------+------+-----+
# | 5          | Leo  | 7   |
# | 6          | Alex | 7   |
# +------------+------+-----+
# Output:
# +------------+---------+-----+
# | student_id | name    | age |
# +------------+---------+-----+
# | 1          | Mason   | 8   |
# | 2          | Ava     | 6   |
# | 3          | Taylor  | 15  |
# | 4          | Georgia | 17  |
# | 5          | Leo     | 7   |
# | 6          | Alex    | 7   |
# +------------+---------+-----+
# Explanation:
# The two DataFramess are stacked vertically, and their rows are combined.
import pandas as pd

def concatenateTables(df1: pd.DataFrame, df2: pd.DataFrame) -> pd.DataFrame:
    return pd.concat([df1, df2], axis=0)


if __name__ == "__main__":
    l1 = [
        [101, None, 15,1.0],
        [102, None, 11,2.0],
        [103, 3, 11,3.0],
        [104, 4, 20,4.0]
    ]
    l2 = [
        [105, 3, 11,3.0],
        [106, 4, 20,4.0]
    ]
    print(concatenateTables(
        pd.DataFrame(l1,columns=["id","quantity","last","grade"]),
        pd.DataFrame(l2,columns=["id","quantity","last","grade"]),
    ))