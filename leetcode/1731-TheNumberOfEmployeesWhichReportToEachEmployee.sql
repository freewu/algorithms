-- 1731. The Number of Employees Which Report to Each Employee
-- Table: Employees
-- +-------------+----------+
-- | Column Name | Type     |
-- +-------------+----------+
-- | employee_id | int      |
-- | name        | varchar  |
-- | reports_to  | int      |
-- | age         | int      |
-- +-------------+----------+
-- employee_id is the column with unique values for this table.
-- This table contains information about the employees and the id of the manager they report to. Some employees do not report to anyone (reports_to is null). 
 
-- For this problem, we will consider a manager an employee who has at least 1 other employee reporting to them.
-- Write a solution to report the ids and the names of all managers, the number of employees who report directly to them, and the average age of the reports rounded to the nearest integer.
-- Return the result table ordered by employee_id.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Employees table:
-- +-------------+---------+------------+-----+
-- | employee_id | name    | reports_to | age |
-- +-------------+---------+------------+-----+
-- | 9           | Hercy   | null       | 43  |
-- | 6           | Alice   | 9          | 41  |
-- | 4           | Bob     | 9          | 36  |
-- | 2           | Winston | null       | 37  |
-- +-------------+---------+------------+-----+
-- Output: 
-- +-------------+-------+---------------+-------------+
-- | employee_id | name  | reports_count | average_age |
-- +-------------+-------+---------------+-------------+
-- | 9           | Hercy | 2             | 39          |
-- +-------------+-------+---------------+-------------+
-- Explanation: Hercy has 2 people report directly to him, Alice and Bob. Their average age is (41+36)/2 = 38.5, which is 39 after rounding it to the nearest integer.

-- Create table If Not Exists Employees(employee_id int, name varchar(20), reports_to int, age int)
-- Truncate table Employees
-- insert into Employees (employee_id, name, reports_to, age) values ('9', 'Hercy', 'None', '43')
-- insert into Employees (employee_id, name, reports_to, age) values ('6', 'Alice', '9', '41')
-- insert into Employees (employee_id, name, reports_to, age) values ('4', 'Bob', '9', '36')
-- insert into Employees (employee_id, name, reports_to, age) values ('2', 'Winston', 'None', '37')

-- Write your MySQL query statement below
SELECT
    a.employee_id,
    a.name,
    count(b.employee_id) AS reports_count,
    ROUND(AVG(b.age),0) AS average_age 
FROM 
    Employees a,
    Employees b
WHERE
    a.employee_id = b.reports_to 
GROUP By
    a.employee_id
ORDER BY 
    a.employee_id