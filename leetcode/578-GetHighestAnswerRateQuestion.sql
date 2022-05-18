-- 578. Get Highest Answer Rate Question
-- Table: SurveyLog
--
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | id          | int  |
-- | action      | ENUM |
-- | question_id | int  |
-- | answer_id   | int  |
-- | q_num       | int  |
-- | timestamp   | int  |
-- +-------------+------+
-- There is no primary key for this table. It may contain duplicates.
-- action is an ENUM of the type: "show", "answer", or "skip".
-- Each row of this table indicates the user with ID = id has taken an action with the question question_id at time timestamp.
-- If the action taken by the user is "answer", answer_id will contain the id of that answer, otherwise, it will be null.
-- q_num is the numeral order of the question in the current session.
--  
-- The answer rate for a question is the number of times a user answered the question by the number of times a user showed the question.
-- Write an SQL query to report the question that has the highest answer rate.
-- If multiple questions have the same maximum answer rate, report the question with the smallest question_id.
-- The query result format is in the following example.
--
-- Example 1:
--
-- Input:
-- SurveyLog table:
-- +----+--------+-------------+-----------+-------+-----------+
-- | id | action | question_id | answer_id | q_num | timestamp |
-- +----+--------+-------------+-----------+-------+-----------+
-- | 5  | show   | 285         | null      | 1     | 123       |
-- | 5  | answer | 285         | 124124    | 1     | 124       |
-- | 5  | show   | 369         | null      | 2     | 125       |
-- | 5  | skip   | 369         | null      | 2     | 126       |
-- +----+--------+-------------+-----------+-------+-----------+
-- Output:
-- +------------+
-- | survey_log |
-- +------------+
-- | 285        |
-- +------------+
-- Explanation:
-- Question 285 was showed 1 time and answered 1 time. The answer rate of question 285 is 1.0
-- Question 369 was showed 1 time and was not answered. The answer rate of question 369 is 0.0
-- Question 285 has the highest answer rate.
--
# Write your MySQL query statement below
SELECT
    s.question_id AS survey_log
FROM
    (
        SELECT
            COUNT(*) AS show_num,
            question_id
        FROM
            SurveyLog
        WHERE
            action = 'show'
        GROUP BY
            question_id
    ) AS s,
    (
        SELECT
            COUNT(*) AS answer_num,
            question_id
        FROM
            SurveyLog
        WHERE
            action = 'answer'
        GROUP BY
            question_id
    ) AS a
WHERE
    s.question_id = a.question_id
ORDER BY
    a.answer_num / s.show_num DESC,
    s.question_id ASC -- 如果有多个问题具有相同的最大 回答率 ，返回 question_id 最小的那个。
LIMIT 1