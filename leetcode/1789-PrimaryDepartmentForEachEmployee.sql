-- 1789. Primary Department for Each Employee
-- Table: Employee
-- +---------------+---------+
-- | Column Name   |  Type   |
-- +---------------+---------+
-- | employee_id   | int     |
-- | department_id | int     |
-- | primary_flag  | varchar |
-- +---------------+---------+
-- (employee_id, department_id) is the primary key (combination of columns with unique values) for this table.
-- employee_id is the id of the employee.
-- department_id is the id of the department to which the employee belongs.
-- primary_flag is an ENUM (category) of type ('Y', 'N'). If the flag is 'Y', the department is the primary department for the employee. If the flag is 'N', the department is not the primary.
 
-- Employees can belong to multiple departments. When the employee joins other departments, they need to decide which department is their primary department. Note that when an employee belongs to only one department, their primary column is 'N'.
-- Write a solution to report all the employees with their primary department. For employees who belong to one department, report their only department.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Employee table:
-- +-------------+---------------+--------------+
-- | employee_id | department_id | primary_flag |
-- +-------------+---------------+--------------+
-- | 1           | 1             | N            |
-- | 2           | 1             | Y            |
-- | 2           | 2             | N            |
-- | 3           | 3             | N            |
-- | 4           | 2             | N            |
-- | 4           | 3             | Y            |
-- | 4           | 4             | N            |
-- +-------------+---------------+--------------+
-- Output: 
-- +-------------+---------------+
-- | employee_id | department_id |
-- +-------------+---------------+
-- | 1           | 1             |
-- | 2           | 1             |
-- | 3           | 3             |
-- | 4           | 3             |
-- +-------------+---------------+
-- Explanation: 
-- - The Primary department for employee 1 is 1.
-- - The Primary department for employee 2 is 1.
-- - The Primary department for employee 3 is 3.
-- - The Primary department for employee 4 is 3.

-- Create table If Not Exists Employee (employee_id int, department_id int, primary_flag ENUM('Y','N'))
-- Truncate table Employee
-- insert into Employee (employee_id, department_id, primary_flag) values ('1', '1', 'N')
-- insert into Employee (employee_id, department_id, primary_flag) values ('2', '1', 'Y')
-- insert into Employee (employee_id, department_id, primary_flag) values ('2', '2', 'N')
-- insert into Employee (employee_id, department_id, primary_flag) values ('3', '3', 'N')
-- insert into Employee (employee_id, department_id, primary_flag) values ('4', '2', 'N')
-- insert into Employee (employee_id, department_id, primary_flag) values ('4', '3', 'Y')
-- insert into Employee (employee_id, department_id, primary_flag) values ('4', '4', 'N')

-- Write your MySQL query statement below
SELECT 
    employee_id,
    department_id
FROM
    Employee
WHERE
    primary_flag = 'Y' OR -- 多个部门通过标志标取
    employee_id IN ( -- 只取一个部门的
        SELECT
            employee_id
        FROM 
            Employee
        GROUP By
            employee_id
        HAVING(COUNT(*)) = 1
    )

-- UNION
(-- 只有一个部门的员工
    SELECT
        employee_id,
        department_id
    FROM
        Employee 
    GROUP By
        employee_id
    HAVING
        COUNT(department_id) = 1
)
UNION
(-- 多个部门取 primary_flag = 'Y' 的部门编号
    SELECT
        employee_id,
        department_id
    FROM
        Employee 
    WHERE 
        primary_flag = 'Y'
)