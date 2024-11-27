-- 3368. First Letter Capitalization
-- Table: user_content
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | content_id  | int     |
-- | content_text| varchar |
-- +-------------+---------+
-- content_id is the unique key for this table.
-- Each row contains a unique ID and the corresponding text content.
-- Write a solution to transform the text in the content_text column by applying the following rules:

-- Convert the first letter of each word to uppercase
-- Keep all other letters in lowercase
-- Preserve all existing spaces
-- Note: There will be no special character in content_text.

-- Return the result table that includes both the original content_text and the modified text where each word starts with a capital letter.

-- The result format is in the following example.

-- Example:
-- Input:
-- user_content table:
-- +------------+-----------------------------------+
-- | content_id | content_text                      |
-- +------------+-----------------------------------+
-- | 1          | hello world of SQL                |
-- | 2          | the QUICK brown fox               |
-- | 3          | data science AND machine learning |
-- | 4          | TOP rated programming BOOKS       |
-- +------------+-----------------------------------+
-- Output:
-- +------------+-----------------------------------+-----------------------------------+
-- | content_id | original_text                     | converted_text                    |
-- +------------+-----------------------------------+-----------------------------------+
-- | 1          | hello world of SQL                | Hello World Of SQL                |
-- | 2          | the QUICK brown fox               | The Quick Brown Fox               |
-- | 3          | data science AND machine learning | Data Science And Machine Learning |
-- | 4          | TOP rated programming BOOKS       | Top Rated Programming Books       |
-- +------------+-----------------------------------+-----------------------------------+
-- Explanation:
-- For content_id = 1:
-- Each word's first letter is capitalized: Hello World Of SQL
-- For content_id = 2:
-- Original mixed-case text is transformed to title case: The Quick Brown Fox
-- For content_id = 3:
-- The word AND is converted to "And": "Data Science And Machine Learning"
-- For content_id = 4:
-- Handles word TOP rated correctly: Top Rated
-- Converts BOOKS from all caps to title case: Books

-- CREATE TABLE If not exists user_content (
--     content_id INT,
--     content_text VARCHAR(255)
-- )
-- Truncate table user_content
-- insert into user_content (content_id, content_text) values ('1', 'hello world of SQL')
-- insert into user_content (content_id, content_text) values ('2', 'the QUICK brown fox')
-- insert into user_content (content_id, content_text) values ('3', 'data science AND machine learning')
-- insert into user_content (content_id, content_text) values ('4', 'TOP rated programming BOOKS')

-- # Write your MySQL query statement below
WITH RECURSIVE data AS (
    (
        SELECT 
            content_id,
            RIGHT(content_text, LENGTH(content_text) - LENGTH(SUBSTRING_INDEX(content_text,' ', 1)) - 1) AS content_text,
            CONCAT( 
                UPPER(LEFT(SUBSTRING_INDEX(content_text,' ', 1), 1)),
                LOWER(RIGHT(SUBSTRING_INDEX(content_text,' ', 1), 
                LENGTH(SUBSTRING_INDEX(content_text,' ', 1)) - 1))
            ) AS converted_text
        FROM 
            user_content
    )
    UNION ALL
    (
        SELECT 
            content_id, 
            RIGHT(content_text, LENGTH(content_text) - LENGTH(SUBSTRING_INDEX(content_text,' ', 1)) - 1) AS content_text,
            CONCAT(
                converted_text,
                ' ',
                UPPER(LEFT(SUBSTRING_INDEX(content_text,' ', 1), 1)),
                LOWER(RIGHT(SUBSTRING_INDEX(content_text,' ', 1), LENGTH(SUBSTRING_INDEX(content_text,' ', 1)) - 1))
            ) AS converted_text
        FROM 
            data
        WHERE 
            LENGTH(content_text) != 0
    )
)

SELECT 
    d.content_id,
    u.content_text AS original_text,
    d.converted_text AS converted_text
FROM 
    data AS d
LEFT JOIN 
    user_content AS u 
ON 
    d.content_id = u.content_id
WHERE 
    LENGTH(d.content_text) = 0
ORDER BY 
    d.content_id
