-- 1270. All People Report to the Given Manager
-- Table: Employees
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | employee_id   | int     |
-- | employee_name | varchar |
-- | manager_id    | int     |
-- +---------------+---------+
-- employee_id is the column of unique values for this table.
-- Each row of this table indicates that the employee with ID employee_id and name employee_name reports his work to his/her direct manager with manager_id
-- The head of the company is the employee with employee_id = 1.
 
-- Write a solution to find employee_id of all employees that directly or indirectly report their work to the head of the company.
-- The indirect relation between managers will not exceed three managers as the company is small.
-- Return the result table in any order.
-- The result format is in the following example.
 
-- Example 1:
-- Input: 
-- Employees table:
-- +-------------+---------------+------------+
-- | employee_id | employee_name | manager_id |
-- +-------------+---------------+------------+
-- | 1           | Boss          | 1          |
-- | 3           | Alice         | 3          |
-- | 2           | Bob           | 1          |
-- | 4           | Daniel        | 2          |
-- | 7           | Luis          | 4          |
-- | 8           | Jhon          | 3          |
-- | 9           | Angela        | 8          |
-- | 77          | Robert        | 1          |
-- +-------------+---------------+------------+
-- Output: 
-- +-------------+
-- | employee_id |
-- +-------------+
-- | 2           |
-- | 77          |
-- | 4           |
-- | 7           |
-- +-------------+
-- Explanation: 
-- The head of the company is the employee with employee_id 1.
-- The employees with employee_id 2 and 77 report their work directly to the head of the company.
-- The employee with employee_id 4 reports their work indirectly to the head of the company 4 --> 2 --> 1. 
-- The employee with employee_id 7 reports their work indirectly to the head of the company 7 --> 4 --> 2 --> 1.
-- The employees with employee_id 3, 8, and 9 do not report their work to the head of the company directly or indirectly. 

-- Create table If Not Exists Employees (employee_id int, employee_name varchar(30), manager_id int)
-- Truncate table Employees
-- insert into Employees (employee_id, employee_name, manager_id) values ('1', 'Boss', '1')
-- insert into Employees (employee_id, employee_name, manager_id) values ('3', 'Alice', '3')
-- insert into Employees (employee_id, employee_name, manager_id) values ('2', 'Bob', '1')
-- insert into Employees (employee_id, employee_name, manager_id) values ('4', 'Daniel', '2')
-- insert into Employees (employee_id, employee_name, manager_id) values ('7', 'Luis', '4')
-- insert into Employees (employee_id, employee_name, manager_id) values ('8', 'John', '3')
-- insert into Employees (employee_id, employee_name, manager_id) values ('9', 'Angela', '8')
-- insert into Employees (employee_id, employee_name, manager_id) values ('77', 'Robert', '1')

-- 使用 RECURSIVE mysql 8.0
WITH RECURSIVE cte  AS
(
    (
        SELECT 
            employee_id 
        FROM 
            Employees a 
        WHERE 
            employee_id !=1 AND -- 不为 BOSS
            manager_id = 1 -- 经理不为 BOSS
    )
    UNION ALL
    (
        SELECT 
            a.employee_id 
        FROM 
            Employees AS a 
        JOIN 
            cte AS b 
        ON a.manager_id = b.employee_id  
    )
    
)
SELECT * FROM cte;

-- sub select
SELECT 
    employee_id 
FROM 
    Employees 
WHERE 
    employee_id <> 1 AND -- 不为 BOSS
    -- 由于公司规模较小，经理之间的间接关系 不超过 3 个经理 
    -- 7 --> 4 --> 2 --> 1
    manager_id IN
    ( -- 第二层 
        SELECT 
            employee_id 
        FROM 
            Employees
        WHERE 
            manager_id IN
            ( -- 第一层
                SELECT employee_id FROM Employees WHERE manager_id = 1 
            ) 
    );

-- join
SELECT 
    e1.employee_id 
FROM 
    employees e1, 
    employees e2, 
    employees e3, 
    employees e4
WHERE 
    -- 由于公司规模较小，经理之间的间接关系 不超过 3 个经理 
    e1.manager_id = e2.employee_id AND -- 三层
    e2.manager_id = e3.employee_id AND -- 两层
    e3.manager_id = e4.employee_id AND -- 一层
    e4.manager_id = 1 AND -- 直接给 BOSS 汇报的经理 第一层
    e1.employee_id != 1 -- 不为 BOSS