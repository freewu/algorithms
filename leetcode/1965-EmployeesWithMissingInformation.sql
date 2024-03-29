-- 1965. Employees With Missing Information
-- Table: Employees
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | employee_id | int     |
-- | name        | varchar |
-- +-------------+---------+
-- employee_id is the primary key for this table.
-- Each row of this table indicates the name of the employee whose ID is employee_id.
--
-- Table: Salaries
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | employee_id | int     |
-- | salary      | int     |
-- +-------------+---------+
-- employee_id is the primary key for this table.
-- Each row of this table indicates the salary of the employee whose ID is employee_id.
--
-- Write an SQL query to report the IDs of all the employees with missing information. The information of an employee is missing if:
--      The employee's name is missing, or
--      The employee's salary is missing.

-- Return the result table ordered by employee_id in ascending order.
-- The query result format is in the following example.
--
-- Example 1:
-- Input:
-- Employees table:
-- +-------------+----------+
-- | employee_id | name     |
-- +-------------+----------+
-- | 2           | Crew     |
-- | 4           | Haven    |
-- | 5           | Kristian |
-- +-------------+----------+
-- Salaries table:
-- +-------------+--------+
-- | employee_id | salary |
-- +-------------+--------+
-- | 5           | 76071  |
-- | 1           | 22517  |
-- | 4           | 63539  |
-- +-------------+--------+
-- Output:
-- +-------------+
-- | employee_id |
-- +-------------+
-- | 1           |
-- | 2           |
-- +-------------+
-- Explanation:
-- Employees 1, 2, 4, and 5 are working at this company.
-- The name of employee 1 is missing.
-- The salary of employee 2 is missing.

-- Create table If Not Exists Employees (employee_id int, name varchar(30))
-- Create table If Not Exists Salaries (employee_id int, salary int)
-- Truncate table Employees
-- insert into Employees (employee_id, name) values ('2', 'Crew')
-- insert into Employees (employee_id, name) values ('4', 'Haven')
-- insert into Employees (employee_id, name) values ('5', 'Kristian')
-- Truncate table Salaries
-- insert into Salaries (employee_id, salary) values ('5', '76071')
-- insert into Salaries (employee_id, salary) values ('1', '22517')
-- insert into Salaries (employee_id, salary) values ('4', '63539')

-- union + not in
SELECT
    employee_id
FROM
    ( -- 取 Employees 表中 Salaries 没有的员工编号
        (
            SELECT
                e.employee_id AS employee_id
            FROM
                Employees AS e
            WHERE
                e.employee_id NOT IN (
                    SELECT
                        employee_id
                    FROM
                        Salaries
                )
        )
        UNION
        ( -- 取 Salaries 表中 Employees 没有的员工编号
            SELECT
                s.employee_id AS employee_id
            FROM
                Salaries AS s
            WHERE
                s.employee_id NOT IN (
                    SELECT
                        employee_id
                    FROM
                        Employees
                )
        )
    ) AS t
ORDER BY
    employee_id ASC

-- union all + having
SELECT
    a.employee_id AS employee_id
FROM
    (
        SELECT employee_id FROM employees
        UNION ALL
        SELECT employee_id FROM salaries
    ) AS a
GROUP BY a.employee_id
HAVING COUNT (a.employee_id) = 1 -- 两张表合并只出一次的 肯定只存在一张表中，另一张表一定不会存在
ORDER BY a.employee_id ASC