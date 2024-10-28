-- 3338. Second Highest Salary II
-- Table: employees
-- +------------------+---------+
-- | Column Name      | Type    |
-- +------------------+---------+
-- | emp_id           | int     |
-- | salary           | int     |
-- | dept             | varchar |
-- +------------------+---------+
-- emp_id is the unique key for this table.
-- Each row of this table contains information about an employee including their ID, name, manager, salary, department, start date, and building assignment.
-- Write a solution to find the employees who earn the second-highest salary in each department. If multiple employees have the second-highest salary, include all employees with that salary.

-- Return the result table ordered by emp_id in ascending order.

-- The result format is in the following example.
-- Example:
-- Input:
-- employees table:
-- +--------+--------+-----------+
-- | emp_id | salary | dept      |
-- +--------+--------+-----------+
-- | 1      | 70000  | Sales     |
-- | 2      | 80000  | Sales     |
-- | 3      | 80000  | Sales     |
-- | 4      | 90000  | Sales     |
-- | 5      | 55000  | IT        |
-- | 6      | 65000  | IT        |
-- | 7      | 65000  | IT        |
-- | 8      | 50000  | Marketing |
-- | 9      | 55000  | Marketing |
-- | 10     | 55000  | HR        |
-- +--------+--------+-----------+
-- Output:
-- +--------+-----------+
-- | emp_id | dept      |
-- +--------+-----------+
-- | 2      | Sales     |
-- | 3      | Sales     |
-- | 5      | IT        |
-- | 8      | Marketing |
-- +--------+-----------+
-- Explanation:
-- Sales Department:
-- Highest salary is 90000 (emp_id: 4)
-- Second-highest salary is 80000 (emp_id: 2, 3)
-- Both employees with salary 80000 are included
-- IT Department:
-- Highest salary is 65000 (emp_id: 6, 7)
-- Second-highest salary is 55000 (emp_id: 5)
-- Only emp_id 5 is included as they have the second-highest salary
-- Marketing Department:
-- Highest salary is 55000 (emp_id: 9)
-- Second-highest salary is 50000 (emp_id: 8)
-- Employee 8 is included
-- HR Department:
-- Only has one employee
-- Not included in the result as it has fewer than 2 employees

-- Create table if not exists Employees (
--     emp_id int,
--     salary int,
--     dept varchar(50)
-- )
-- Truncate table employees
-- insert into employees (emp_id, salary, dept) values ('1', '70000', 'Sales')
-- insert into employees (emp_id, salary, dept) values ('2', '80000', 'Sales')
-- insert into employees (emp_id, salary, dept) values ('3', '80000', 'Sales')
-- insert into employees (emp_id, salary, dept) values ('4', '90000', 'Sales')
-- insert into employees (emp_id, salary, dept) values ('5', '55000', 'IT')
-- insert into employees (emp_id, salary, dept) values ('6', '65000', 'IT')
-- insert into employees (emp_id, salary, dept) values ('7', '65000', 'IT')
-- insert into employees (emp_id, salary, dept) values ('8', '50000', 'Marketing')
-- insert into employees (emp_id, salary, dept) values ('9', '55000', 'Marketing')
-- insert into employees (emp_id, salary, dept) values ('10', '55000', 'HR')

# Write your MySQL query statement below
-- 需要使用 DENSE_RANK() 序列是连续的
-- 使用 RANK 不连续 按人数跳排名
WITH t AS ( -- 按部门给薪水由大到小编号
    SELECT 
        *,
        DENSE_RANK() OVER (PARTITION BY dept ORDER BY salary DESC) AS `rk`
    FROM 
       employees  
)
-- SELECT * FROM t;
-- | emp_id | salary | dept      | rk |
-- | ------ | ------ | --------- | -- |
-- | 10     | 55000  | HR        | 1  |
-- | 6      | 65000  | IT        | 1  |
-- | 7      | 65000  | IT        | 1  |
-- | 5      | 55000  | IT        | 2  |
-- | 9      | 55000  | Marketing | 1  |
-- | 8      | 50000  | Marketing | 2  |
-- | 4      | 90000  | Sales     | 1  |
-- | 2      | 80000  | Sales     | 2  |
-- | 3      | 80000  | Sales     | 2  |
-- | 1      | 70000  | Sales     | 3  |

SELECT
    emp_id,
    dept
FROM 
    t
WHERE
    rk = 2
ORDER BY 
    emp_id -- the result table ordered by emp_id in ascending order.