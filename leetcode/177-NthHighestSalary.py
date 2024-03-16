# 177. Nth Highest Salary
# Table: Employee
# +-------------+------+
# | Column Name | Type |
# +-------------+------+
# | id          | int  |
# | salary      | int  |
# +-------------+------+
# id is the primary key column for this table.
# Each row of this table contains information about the salary of an employee.
#  
# Write an SQL query to report the nth highest salary from the Employee table.
# If there is no nth highest salary, the query should report null.
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
# n = 2
# Output:
# +------------------------+
# | getNthHighestSalary(2) |
# +------------------------+
# | 200                    |
# +------------------------+

# Example 2:
# Input:
# Employee table:
# +----+--------+
# | id | salary |
# +----+--------+
# | 1  | 100    |
# +----+--------+
# n = 2
# Output:
# +------------------------+
# | getNthHighestSalary(2) |
# +------------------------+
# | null                   |
# +------------------------+

import pandas as pd

def nth_highest_salary(employee: pd.DataFrame, N: int) -> pd.DataFrame:
    # 排序
    employee = employee.sort_values('Salary', ascending=[False])
    # 取 Salary 列并改名 & 取排名
    return employee[["Salary"]].rename(columns = {'Salary': 'getNthHighestSalary(' + str(N) + ')' }).loc[N - 1]

import pandas as pd

def nth_highest_salary1(employee: pd.DataFrame, N: int) -> pd.DataFrame:
    # 去重
    employee = employee[["Salary"]].drop_duplicates()
    # 如果数据不够 N 直接 返回 None
    if len(employee) < N:
        return pd.DataFrame({'getNthHighestSalary('+ str(N) +')': [None]})
    # 排序后取值
    employee = employee.sort_values(by = 'Salary', ascending=False).head(N).tail(1)
    # 修改名称
    return employee.rename(columns={'Salary':'getNthHighestSalary('+ str(N) +')'})


def nth_highest_salary2(employee: pd.DataFrame, N: int) -> pd.DataFrame:
    # 去重
    employee = employee[["Salary"]].drop_duplicates()
    # 如果数据不够 N 直接 返回 None
    if len(employee) < N:
        return pd.DataFrame({'getNthHighestSalary('+ str(N) +')': [None]})
    # 排序后取值
    employee = employee.sort_values(by = 'Salary', ascending=False).loc(N - 1)
    # 修改名称
    return employee.rename(columns={'Salary':'getNthHighestSalary('+ str(N) +')'})



if __name__ == "__main__":
    data = [[1, 100], [2, 200], [3, 300]]
    employee = pd.DataFrame(data, columns=['Id', 'Salary']).astype({'Id':'Int64', 'Salary':'Int64'})
    print(nth_highest_salary(employee,2))
    print(nth_highest_salary1(employee,2))
    #print(nth_highest_salary2(employee,2))