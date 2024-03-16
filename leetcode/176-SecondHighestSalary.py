# 176. Second Highest Salary
# Table: Employee
# +-------------+------+
# | Column Name | Type |
# +-------------+------+
# | id          | int  |
# | salary      | int  |
# +-------------+------+
# id is the primary key column for this table.
# Each row of this table contains information about the salary of an employee.
# Write an SQL query to report the second highest salary from the Employee table. If there is no second highest salary, the query should report null.

# The query result format is in the following example.

# Example 1:
# Input:
# Employee table:
# +----+--------+
# | id | salary |
# +----+--------+
# | 1  | 100    |
# | 2  | 200    |
# | 3  | 300    |
# +----+--------+
# Output:
# +---------------------+
# | SecondHighestSalary |
# +---------------------+
# | 200                 |
# +---------------------+

# Example 2:
# Input:
# Employee table:
# +----+--------+
# | id | salary |
# +----+--------+
# | 1  | 100    |
# +----+--------+
# Output:
# +---------------------+
# | SecondHighestSalary |
# +---------------------+
# | null                |
# +---------------------+

import pandas as pd

def second_highest_salary(employee: pd.DataFrame) -> pd.DataFrame:
    N = 2
    # 去重
    employee = employee[["salary"]].drop_duplicates()
    # 要处理 N <= 0 的情况
    # 如果去重后 数量不够 N 直接返回 None
    if len(employee) < N or N <= 0:
        return pd.DataFrame({'SecondHighestSalary': [None]})
    # salary DESC 排序(sort_values(by = 'salary', ascending=False))  & 取 第 N 个( .head(N).tail(1) )
    employee = employee.sort_values(by = 'salary', ascending=False).head(N).tail(1)
    # 重命名 column
    return employee.rename(columns={'salary': 'SecondHighestSalary'})

if __name__ == "__main__":
    data = [[1, 100], [2, 200], [3, 300]]
    employee = pd.DataFrame(data, columns=['id', 'salary']).astype({'id':'int64', 'salary':'int64'})
    print(second_highest_salary(employee))