-- 1454. Active Users
-- Table: Accounts
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | name          | varchar |
-- +---------------+---------+
-- id is the primary key (column with unique values) for this table.
-- This table contains the account id and the user name of each account.
 
-- Table: Logins
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | login_date    | date    |
-- +---------------+---------+
-- This table may contain duplicate rows.
-- This table contains the account id of the user who logged in and the login date. 
-- A user may log in multiple times in the day.

-- Active users are those who logged in to their accounts for five or more consecutive days.
-- Write a solution to find the id and the name of active users.
-- Return the result table ordered by id.

-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Accounts table:
-- +----+----------+
-- | id | name     |
-- +----+----------+
-- | 1  | Winston  |
-- | 7  | Jonathan |
-- +----+----------+
-- Logins table:
-- +----+------------+
-- | id | login_date |
-- +----+------------+
-- | 7  | 2020-05-30 |
-- | 1  | 2020-05-30 |
-- | 7  | 2020-05-31 |
-- | 7  | 2020-06-01 |
-- | 7  | 2020-06-02 |
-- | 7  | 2020-06-02 |
-- | 7  | 2020-06-03 |
-- | 1  | 2020-06-07 |
-- | 7  | 2020-06-10 |
-- +----+------------+
-- Output: 
-- +----+----------+
-- | id | name     |
-- +----+----------+
-- | 7  | Jonathan |
-- +----+----------+
-- Explanation: 
-- User Winston with id = 1 logged in 2 times only in 2 different days, so, Winston is not an active user.
-- User Jonathan with id = 7 logged in 7 times in 6 different days, five of them were consecutive days, so, Jonathan is an active user.
 
-- Follow up: Could you write a general solution if the active users are those who logged in to their accounts for n or more consecutive days?

-- Create table If Not Exists Accounts (id int, name varchar(10))
-- Create table If Not Exists Logins (id int, login_date date)
-- Truncate table Accounts
-- insert into Accounts (id, name) values ('1', 'Winston')
-- insert into Accounts (id, name) values ('7', 'Jonathan')
-- Truncate table Logins
-- insert into Logins (id, login_date) values ('7', '2020-05-30')
-- insert into Logins (id, login_date) values ('1', '2020-05-30')
-- insert into Logins (id, login_date) values ('7', '2020-05-31')
-- insert into Logins (id, login_date) values ('7', '2020-06-01')
-- insert into Logins (id, login_date) values ('7', '2020-06-02')
-- insert into Logins (id, login_date) values ('7', '2020-06-02')
-- insert into Logins (id, login_date) values ('7', '2020-06-03')
-- insert into Logins (id, login_date) values ('1', '2020-06-07')
-- insert into Logins (id, login_date) values ('7', '2020-06-10')


-- Write your MySQL query statement below
-- LAG
WITH tt AS (
    SELECT
        id, 
        login_date,
        LAG(login_date,1) OVER(PARTITION BY id ORDER BY login_date DESC) as lag1,
        LAG(login_date,2) OVER(PARTITION BY id ORDER BY login_date DESC) as lag2,
        LAG(login_date,3) OVER(PARTITION BY id ORDER BY login_date DESC) as lag3,
        LAG(login_date,4) OVER(PARTITION BY id ORDER BY login_date DESC) as lag4
    FROM 
        (
            SELECT DISTINCT * FROM Logins
        ) as t1
)
SELECT 
    DISTINCT a.id AS id,
    a.name AS name 
FROM 
    tt AS b
JOIN 
    Accounts a 
ON 
    b.id = a.id
WHERE 
    DATEDIFF(login_date,lag1) = -1 AND 
    DATEDIFF(lag1, lag2) = -1 AND 
    DATEDIFF(lag2, lag3) = -1 AND 
    DATEDIFF(lag3, lag4) = -1
ORDER BY 
    id

-- WITH tt AS (
--     SELECT
--         id, 
--         login_date,
--         LAG(login_date,1) OVER(PARTITION BY id ORDER BY login_date DESC) as lag1,
--         LAG(login_date,2) OVER(PARTITION BY id ORDER BY login_date DESC) as lag2,
--         LAG(login_date,3) OVER(PARTITION BY id ORDER BY login_date DESC) as lag3,
--         LAG(login_date,4) OVER(PARTITION BY id ORDER BY login_date DESC) as lag4
--     FROM 
--         (
--             SELECT DISTINCT * FROM Logins
--         ) as t1
-- )
-- SELECT * FROM tt;
-- | id | login_date | lag1       | lag2       | lag3       | lag4       |
-- | -- | ---------- | ---------- | ---------- | ---------- | ---------- |
-- | 1  | 2020-06-07 | null       | null       | null       | null       |
-- | 1  | 2020-05-30 | 2020-06-07 | null       | null       | null       |
-- | 7  | 2020-06-10 | null       | null       | null       | null       |
-- | 7  | 2020-06-03 | 2020-06-10 | null       | null       | null       |
-- | 7  | 2020-06-02 | 2020-06-03 | 2020-06-10 | null       | null       |
-- | 7  | 2020-06-01 | 2020-06-02 | 2020-06-03 | 2020-06-10 | null       |
-- | 7  | 2020-05-31 | 2020-06-01 | 2020-06-02 | 2020-06-03 | 2020-06-10 |
-- | 7  | 2020-05-30 | 2020-05-31 | 2020-06-01 | 2020-06-02 | 2020-06-03 |

-- date_sub + DENSE_RANK
SELECT  
    DISTINCT a.id,
    b.name
FROM 
    (   
        SELECT  
            id
        FROM 
            (
                SELECT  
                    id,
                    login_date,
                    DATE_SUB(login_date,interval DENSE_RANK() OVER (PARTITION BY id ORDER BY login_date )  DAY) AS tag
                FROM    
                    Logins   
            ) AS t
        GROUP BY 
            id ,tag
        HAVING 
            COUNT(distinct login_date) >= 5
    )  AS a 
LEFT JOIN 
    Accounts  AS b 
ON 
    a.id = b.id