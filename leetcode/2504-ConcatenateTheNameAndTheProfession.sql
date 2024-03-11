-- 2504. Concatenate the Name and the Profession
-- Table: Person
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | person_id   | int     |
-- | name        | varchar |
-- | profession  | ENUM    |
-- +-------------+---------+
-- person_id is the primary key (column with a unique value) for this table.
-- Each row in this table contains a person's ID, name, and profession.
-- The profession column in an enum of the type ('Doctor', 'Singer', 'Actor', 'Player', 'Engineer', or 'Lawyer')
-- Write a solution to report each person's name followed by the first letter of their profession enclosed in parentheses.
-- Return the result table ordered by person_id in descending order.
-- The result format is shown in the following example.

-- Example 1:
-- Input: 
-- Person table:
-- +-----------+-------+------------+
-- | person_id | name  | profession |
-- +-----------+-------+------------+
-- | 1         | Alex  | Singer     |
-- | 3         | Alice | Actor      |
-- | 2         | Bob   | Player     |
-- | 4         | Messi | Doctor     |
-- | 6         | Tyson | Engineer   |
-- | 5         | Meir  | Lawyer     |
-- +-----------+-------+------------+
-- Output: 
-- +-----------+----------+
-- | person_id | name     |
-- +-----------+----------+
-- | 6         | Tyson(E) |
-- | 5         | Meir(L)  |
-- | 4         | Messi(D) |
-- | 3         | Alice(A) |
-- | 2         | Bob(P)   |
-- | 1         | Alex(S)  |
-- +-----------+----------+
-- Explanation: Note that there should not be any white space between the name and the first letter of the profession.

-- Create table If Not Exists Person (person_id int, name varchar(30), profession ENUM('Doctor','Singer','Actor','Player','Engineer','Lawyer'))
-- Truncate table Person
-- insert into Person (person_id, name, profession) values ('1', 'Alex', 'Singer')
-- insert into Person (person_id, name, profession) values ('3', 'Alice', 'Actor')
-- insert into Person (person_id, name, profession) values ('2', 'Bob', 'Player')
-- insert into Person (person_id, name, profession) values ('4', 'Messi', 'Doctor')
-- insert into Person (person_id, name, profession) values ('6', 'Tyson', 'Engineer')
-- insert into Person (person_id, name, profession) values ('5', 'Meir', 'Lawyer')

-- Write your MySQL query statement below
SELECT 
    person_id,
    CONCAT(name,"(",SUBSTRING(profession,1,1),")") AS name
FROM
    Person
ORDER BY
    person_id DESC -- the result table ordered by person_id in descending order.

-- SUBSTRING函数是文本处理函数，可以截取字符串
-- 格式: SUBSTRING(s, start, length)
-- 从字符串s的start位置截取长度为length的子字符串
-- 如果SUBSTRING()函数接收2个参数:SUBSTRING(s,start)，
--     则第一个参数为待截取的字符串，
--     第二个参数为截取的起始位置。如果第二个参数为负整数，则为倒数的起始位置

-- SUBSTRING()函数接收两个参数
-- 操作示例	示例结果
-- SELECT SUBSTRING('abc123',2)	bc123
-- SELECT SUBSTRING('abc123',-3)	123
-- 如果SUBSTRING()函数接收3个参数:SUBSTRING(s,start,length)，
--     则第一个参数为待截取的字符串，
--     第二个参数为截取的起始位置，
--     第三个参数为截取的长度。如果第二个参数为负整数，则为倒数的起始位置

-- SUBSTRING()函数接收3个参数
-- 操作示例	示例结果
-- SELECT SUBSTRING('abc123',2,3)	bc1
-- SELECT SUBSTRING('abc123',-3,2)	12

-- 使用方法：CONCAT(str1,str2,…) 
-- 返回结果为连接参数产生的字符串。如有任何一个参数为NULL ，则返回值为 NULL。
-- 如果所有参数均为非二进制字符串，则结果为非二进制字符串。 
-- 如果自变量中含有任一二进制字符串，则结果为一个二进制字符串。
-- 一个数字参数被转化为与之相等的二进制字符串格式；若要避免这种情况，可使用显式类型 cast, 例如：
-- SELECT CONCAT(CAST(int_xxx AS CHAR), char_col)
-- MySQL的concat函数可以连接一个或者多个字符串,如

-- mysql> SELECT CONCAT('my', 's', 'ql');
-- -> 'mysql'
-- mysql> SELECT CONCAT('my', NULL, 'ql');
-- -> NULL
-- mysql> SELECT CONCAT(14.3);
-- -> '14.3'

SELECT 
    person_id,
    CONCAT(name, '(', LEFT(profession, 1), ')') AS name
FROM 
    Person 
ORDER BY 
    person_id DESC