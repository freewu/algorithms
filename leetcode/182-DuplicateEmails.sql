-- 182. Duplicate Emails
-- Table: Person
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | email       | varchar |
-- +-------------+---------+
-- id is the primary key column for this table.
-- Each row of this table contains an email. The emails will not contain uppercase letters.
--  
-- Write an SQL query to report all the duplicate emails.
-- Return the result table in any order.
-- The query result format is in the following example.
--
-- Example 1:
-- Input:
-- Person table:
-- +----+---------+
-- | id | email   |
-- +----+---------+
-- | 1  | a@b.com |
-- | 2  | c@d.com |
-- | 3  | a@b.com |
-- +----+---------+
-- Output:
-- +---------+
-- | Email   |
-- +---------+
-- | a@b.com |
-- +---------+
-- Explanation: a@b.com is repeated two times.

-- Create table If Not Exists Person (id int, email varchar(255))
-- Truncate table Person
-- insert into Person (id, email) values ('1', 'a@b.com')
-- insert into Person (id, email) values ('2', 'c@d.com')
-- insert into Person (id, email) values ('3', 'a@b.com')

-- Write your MySQL query statement below
SELECT
    email AS Email
FROM
    Person
GROUP BY
    email
HAVING COUNT(*) >= 2