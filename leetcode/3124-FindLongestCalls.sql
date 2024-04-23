-- -- 3124. Find Longest Calls
-- Table: Contacts
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | first_name  | varchar |
-- | last_name   | varchar |
-- +-------------+---------+
-- id is the primary key (column with unique values) of this table.
-- id is a foreign key (reference column) to Calls table.
-- Each row of this table contains id, first_name, and last_name.

-- Table: Calls
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | contact_id  | int  |
-- | type        | enum |
-- | duration    | int  |
-- +-------------+------+
-- (contact_id, type, duration) is the primary key (column with unique values) of this table.
-- type is an ENUM (category) type of ('incoming', 'outgoing').
-- Each row of this table contains information about calls, comprising of contact_id, type, and duration in seconds.
-- Write a solution to find the three longest incoming and outgoing calls.

-- Return the result table ordered by type, duration, and first_name in descending order and duration must be formatted as HH:MM:SS.
-- The result format is in the following example.

-- Example 1:
-- Input:
-- Contacts table:
-- +----+------------+-----------+
-- | id | first_name | last_name |
-- +----+------------+-----------+
-- | 1  | John       | Doe       |
-- | 2  | Jane       | Smith     |
-- | 3  | Alice      | Johnson   |
-- | 4  | Michael    | Brown     |
-- | 5  | Emily      | Davis     |
-- +----+------------+-----------+        
-- Calls table:
-- +------------+----------+----------+
-- | contact_id | type     | duration |
-- +------------+----------+----------+
-- | 1          | incoming | 120      |
-- | 1          | outgoing | 180      |
-- | 2          | incoming | 300      |
-- | 2          | outgoing | 240      |
-- | 3          | incoming | 150      |
-- | 3          | outgoing | 360      |
-- | 4          | incoming | 420      |
-- | 4          | outgoing | 200      |
-- | 5          | incoming | 180      |
-- | 5          | outgoing | 280      |
-- +------------+----------+----------+
-- Output:
-- +-----------+----------+-------------------+
-- | first_name| type     | duration_formatted|
-- +-----------+----------+-------------------+
-- | Michael   | incoming | 00:07:00          |
-- | Jane      | incoming | 00:05:00          |
-- | Emily     | incoming | 00:03:00          |
-- | Alice     | outgoing | 00:06:00          |
-- | Emily     | outgoing | 00:04:40          |
-- | Jane      | outgoing | 00:04:00          |
-- +-----------+----------+-------------------+
-- Explanation:
-- Michael had an incoming call lasting 7 minutes.
-- Jane had an incoming call lasting 5 minutes.
-- Emily had an incoming call lasting 3 minutes.
-- Alice had an outgoing call lasting 6 minutes.
-- Emily had an outgoing call lasting 4 minutes and 40 seconds.
-- Jane had an outgoing call lasting 4 minutes.
-- Note: Output table is sorted by type, duration, and first_name in descending order.

-- Create table if Not Exists Contacts(id int, first_name varchar(20), last_name varchar(20))
-- Create table if Not Exists Calls(contact_id int, type ENUM('incoming', 'outgoing'), duration int)
-- Truncate table Contacts
-- insert into Contacts (id, first_name, last_name) values ('1', 'John', 'Doe')
-- insert into Contacts (id, first_name, last_name) values ('2', 'Jane', 'Smith')
-- insert into Contacts (id, first_name, last_name) values ('3', 'Alice', 'Johnson')
-- insert into Contacts (id, first_name, last_name) values ('4', 'Michael', 'Brown')
-- insert into Contacts (id, first_name, last_name) values ('5', 'Emily', 'Davis')
-- Truncate table Calls
-- insert into Calls (contact_id, type, duration) values ('1', 'incoming', '120')
-- insert into Calls (contact_id, type, duration) values ('1', 'outgoing', '180')
-- insert into Calls (contact_id, type, duration) values ('2', 'incoming', '300')
-- insert into Calls (contact_id, type, duration) values ('2', 'outgoing', '240')
-- insert into Calls (contact_id, type, duration) values ('3', 'incoming', '150')
-- insert into Calls (contact_id, type, duration) values ('3', 'outgoing', '360')
-- insert into Calls (contact_id, type, duration) values ('4', 'incoming', '420')
-- insert into Calls (contact_id, type, duration) values ('4', 'outgoing', '200')
-- insert into Calls (contact_id, type, duration) values ('5', 'incoming', '180')
-- insert into Calls (contact_id, type, duration) values ('5', 'outgoing', '280')

-- Write your MySQL query statement below
-- SELECT 
--     *,
--     RANK() OVER(PARTITION BY type ORDER BY duration DESC) AS rk
-- FROM
--     Calls 

SELECT  
    u.first_name,
    c.type,
    -- DATE_ADD("2024-01-01", INTERVAL duration SECOND),
    DATE_FORMAT(DATE_ADD("2024-01-01", INTERVAL duration SECOND), "%H:%i:%s") AS duration_formatted -- duration must be formatted as HH:MM:SS
FROM
    Contacts AS u,
    (
        SELECT 
            *,
            RANK() OVER(PARTITION BY type ORDER BY duration DESC) AS rk
        FROM
            Calls 
    ) AS c 
WHERE
    u.id = c.contact_id AND
    c.rk <= 3
ORDER  BY 
    c.type, c.duration DESC -- Return the result table ordered by type, duration, and first_name in descending orde