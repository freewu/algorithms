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

def nth_highest_salary1(employee: pd.DataFrame, N: int) -> pd.DataFrame:
    employee = employee.sort_values(by='Salary',ascending=False)
    employee = employee.drop_duplicates(['Salary'])
    if len(employee['Salary']) < N or N <= 0:
        return pd.DataFrame({'getNthHighestSalary('+str(N)+')':[None]})
    else:
        return pd.DataFrame({'getNthHighestSalary('+str(N)+')':[ employee.iloc[N-1]['Salary']]} )


def nth_highest_salary2(employee: pd.DataFrame, N: int) -> pd.DataFrame:
    # 去重
    df = employee[["Salary"]].drop_duplicates()
    # 要处理 N <= 0 的情况
    # 如果去重后 数量不够 N 直接返回 None
    if len(df) < N or N <= 0:
        return pd.DataFrame({'getNthHighestSalary('+ str(N) +')': [None]})
    # salary DESC 排序(sort_values(by = 'salary', ascending=False))  & 取 第 N 个( .head(N).tail(1) )
    df = df.sort_values(by = 'Salary', ascending=False).head(N).tail(1)
    # 重命名 column
    return df.rename(columns={'Salary':'getNthHighestSalary('+ str(N) +')'})


if __name__ == "__main__":
    data = [[1, 100], [2, 200], [3, 300]]
    employee = pd.DataFrame(data, columns=['Id', 'Salary']).astype({'Id':'Int64', 'Salary':'Int64'})
    print(nth_highest_salary(employee,2))
    print(nth_highest_salary1(employee,2))
    print(nth_highest_salary2(employee,2))