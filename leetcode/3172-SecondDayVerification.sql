-- 3172. Second Day Verification
-- Table: emails
-- +-------------+----------+
-- | Column Name | Type     | 
-- +-------------+----------+
-- | email_id    | int      |
-- | user_id     | int      |
-- | signup_date | datetime |
-- +-------------+----------+
-- (email_id, user_id) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table contains the email ID, user ID, and signup date.

-- Table: texts
-- +---------------+----------+
-- | Column Name   | Type     | 
-- +---------------+----------+
-- | text_id       | int      |
-- | email_id      | int      |
-- | signup_action | enum     |
-- | action_date   | datetime |
-- +---------------+----------+
-- (text_id, email_id) is the primary key (combination of columns with unique values) for this table. 
-- signup_action is an enum type of ('Verified', 'Not Verified'). 
-- Each row of this table contains the text ID, email ID, signup action, and action date.
-- Write a Solution to find the user IDs of those who verified their sign-up on the second day.

-- Return the result table ordered by user_id in ascending order.
-- The result format is in the following example.

-- Example:
-- Input:
-- emails table:
-- +----------+---------+---------------------+
-- | email_id | user_id | signup_date         |
-- +----------+---------+---------------------+
-- | 125      | 7771    | 2022-06-14 09:30:00|
-- | 433      | 1052    | 2022-07-09 08:15:00|
-- | 234      | 7005    | 2022-08-20 10:00:00|
-- +----------+---------+---------------------+
-- texts table:
-- +---------+----------+--------------+---------------------+
-- | text_id | email_id | signup_action| action_date         |
-- +---------+----------+--------------+---------------------+
-- | 1       | 125      | Verified     | 2022-06-15 08:30:00|
-- | 2       | 433      | Not Verified | 2022-07-10 10:45:00|
-- | 4       | 234      | Verified     | 2022-08-21 09:30:00|
-- +---------+----------+--------------+---------------------+
-- Output:
-- +---------+
-- | user_id |
-- +---------+
-- | 7005    |
-- | 7771    |
-- +---------+
-- Explanation:
-- User with email_id 7005 signed up on 2022-08-20 10:00:00 and verified on second day of the signup.
-- User with email_id 7771 signed up on 2022-06-14 09:30:00 and verified on second day of the signup.

-- CREATE TABLE If not exists emails (
--     email_id INT,
--     user_id INT,
--     signup_date DATETIME
-- );
-- CREATE TABLE If not exists texts (
--     text_id INT,
--     email_id INT,
--     signup_action ENUM('Verified', 'Not Verified'),
--     action_date DATETIME
-- );
-- Truncate table emails
-- insert into emails (email_id, user_id, signup_date) values ('125', '7771', '2022-06-14 09:30:00')
-- insert into emails (email_id, user_id, signup_date) values ('433', '1052', '2022-07-09 08:15:00')
-- insert into emails (email_id, user_id, signup_date) values ('234', '7005', '2022-08-20 10:00:00')
-- Truncate table texts
-- insert into texts (text_id, email_id, signup_action, action_date) values ('1', '125', 'Verified', '2022-06-15 08:30:00')
-- insert into texts (text_id, email_id, signup_action, action_date) values ('2', '433', 'Not Verified', '2022-07-10 10:45:00')
-- insert into texts (text_id, email_id, signup_action, action_date) values ('4', '234', 'Verified', '2022-08-21 09:30:00')

-- Write your MySQL query statement below
SELECT
    e.user_id 
FROM
    emails AS e
LEFT JOIN
    texts AS t
ON
    e.email_id = t.email_id
WHERE
    t.signup_action = "Verified" AND
    DATE_FORMAT(DATE_ADD(e.signup_date, INTERVAL 1 DAY), "%Y-%m-%d") = DATE_FORMAT(t.action_date, "%Y-%m-%d")
ORDER BY
    e.user_id 