# 184. Department Highest Salary
# Table: Employee
# +--------------+---------+
# | Column Name  | Type    |
# +--------------+---------+
# | id           | int     |
# | name         | varchar |
# | salary       | int     |
# | departmentId | int     |
# +--------------+---------+
# id is the primary key column for this table.
# departmentId is a foreign key of the ID from the Department table.
# Each row of this table indicates the ID, name, and salary of an employee. It also contains the ID of their department.

# Table: Department
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | id          | int     |
# | name        | varchar |
# +-------------+---------+
# id is the primary key column for this table.
# Each row of this table indicates the ID of a department and its name.
#  
# Write an SQL query to find employees who have the highest salary in each of the departments.
# Return the result table in any order.
# The query result format is in the following example.

# Example 1:
# Input:
# Employee table:
# +----+-------+--------+--------------+
# | id | name  | salary | departmentId |
# +----+-------+--------+--------------+
# | 1  | Joe   | 70000  | 1            |
# | 2  | Jim   | 90000  | 1            |
# | 3  | Henry | 80000  | 2            |
# | 4  | Sam   | 60000  | 2            |
# | 5  | Max   | 90000  | 1            |
# +----+-------+--------+--------------+
# Department table:
# +----+-------+
# | id | name  |
# +----+-------+
# | 1  | IT    |
# | 2  | Sales |
# +----+-------+
# Output:
# +------------+----------+--------+
# | Department | Employee | Salary |
# +------------+----------+--------+
# | IT         | Jim      | 90000  |
# | Sales      | Henry    | 80000  |
# | IT         | Max      | 90000  |
# +------------+----------+--------+
# Explanation: Max and Jim both have the highest salary in the IT department and Henry has the highest salary in the Sales department.

import pandas as pd

def department_highest_salary(employee: pd.DataFrame, department: pd.DataFrame) -> pd.DataFrame:
    # 连接两张表
    merged_data = pd.merge(employee, department, left_on='departmentId', right_on='id')
    # print(merged_data) # id_x name_x  salary  departmentId  id_y name_y
    # 按 name_y( department.name ) group by
    result = merged_data.groupby('name_y').apply(
        # 只取每个部门最大的薪水值
        lambda x: x[x['salary'] == x['salary'].max()]
    ).reset_index(drop=True)[['name_y', 'name_x', 'salary']]
    # 重命名 column
    result.columns = ['Department', 'Employee', 'Salary']
    return result

def department_highest_salary1(employee: pd.DataFrame, department: pd.DataFrame) -> pd.DataFrame:
    # 重命名 部门名称
    department.rename(columns = {'name': 'Department'}, inplace = True)
    # join employee & department 数据 employee,departmentId = department.id
    df = employee.merge(department, how = 'left', left_on = 'departmentId', right_on = 'id')
    # 按部门 给每个人排名
    df['rank'] = df.groupby('Department')['salary'].rank(method = 'dense', ascending = False)
    # 重命名 name => Employee  salary => Salary
    df.rename(columns = {'name': 'Employee', 'salary': 'Salary'}, inplace = True)
    # 取排行为1的数据 rank == 1
    return df[df['rank'] == 1].loc[:,['Department', 'Employee', 'Salary']]

if __name__ == "__main__":
    data = [[1, 'Joe', 70000, 1], [2, 'Jim', 90000, 1], [3, 'Henry', 80000, 2], [4, 'Sam', 60000, 2], [5, 'Max', 90000, 1]]
    employee = pd.DataFrame(data, columns=['id', 'name', 'salary', 'departmentId']).astype({'id':'Int64', 'name':'object', 'salary':'Int64', 'departmentId':'Int64'})
    data = [[1, 'IT'], [2, 'Sales']]
    department = pd.DataFrame(data, columns=['id', 'name']).astype({'id':'Int64', 'name':'object'})
    print(department_highest_salary(employee, department))
    print(department_highest_salary1(employee, department))