-- 1873. Calculate Special Bonus
-- Table: Employees
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | employee_id | int     |
-- | name        | varchar |
-- | salary      | int     |
-- +-------------+---------+
-- employee_id is the primary key for this table.
-- Each row of this table indicates the employee ID, employee name, and salary.

-- Write an SQL query to calculate the bonus of each employee.
-- The bonus of an employee is 100% of their salary if the ID of the employee is an odd number and the employee name does not start with the character 'M'. The bonus of an employee is 0 otherwise.
-- Return the result table ordered by employee_id.
-- The query result format is in the following example.

-- Example 1:
-- Input:
-- Employees table:
-- +-------------+---------+--------+
-- | employee_id | name    | salary |
-- +-------------+---------+--------+
-- | 2           | Meir    | 3000   |
-- | 3           | Michael | 3800   |
-- | 7           | Addilyn | 7400   |
-- | 8           | Juan    | 6100   |
-- | 9           | Kannon  | 7700   |
-- +-------------+---------+--------+
-- Output:
-- +-------------+-------+
-- | employee_id | bonus |
-- +-------------+-------+
-- | 2           | 0     |
-- | 3           | 0     |
-- | 7           | 7400  |
-- | 8           | 0     |
-- | 9           | 7700  |
-- +-------------+-------+
-- Explanation:
-- The employees with IDs 2 and 8 get 0 bonus because they have an even employee_id.
-- The employee with ID 3 gets 0 bonus because their name starts with 'M'.
-- The rest of the employees get a 100% bonus.

-- Create table If Not Exists Employees (employee_id int, name varchar(30), salary int)
-- Truncate table Employees
-- insert into Employees (employee_id, name, salary) values ('2', 'Meir', '3000')
-- insert into Employees (employee_id, name, salary) values ('3', 'Michael', '3800')
-- insert into Employees (employee_id, name, salary) values ('7', 'Addilyn', '7400')
-- insert into Employees (employee_id, name, salary) values ('8', 'Juan', '6100')
-- insert into Employees (employee_id, name, salary) values ('9', 'Kannon', '7700')

SELECT
    employee_id,
    IF((employee_id % 2 = 1) AND SUBSTR(name,1,1) != 'M', salary, 0) AS bonus
FROM
    Employees
ORDER BY
    employee_id 

-- ID of the employee is an odd number
-- (employee_id % 2 = 1)

-- the employee name does not start with the character 'M'
-- SUBSTR(name,1,1) != 'M'