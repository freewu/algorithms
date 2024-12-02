-- 3374. First Letter Capitalization II
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

-- Convert the first letter of each word to uppercase and the remaining letters to lowercase
-- Special handling for words containing special characters:
-- For words connected with a hyphen -, both parts should be capitalized (e.g., top-rated → Top-Rated)
-- All other formatting and spacing should remain unchanged
-- Return the result table that includes both the original content_text and the modified text following the above rules.

-- The result format is in the following example.

-- Example:
-- Input:
-- user_content table:
-- +------------+---------------------------------+
-- | content_id | content_text                    |
-- +------------+---------------------------------+
-- | 1          | hello world of SQL              |
-- | 2          | the QUICK-brown fox             |
-- | 3          | modern-day DATA science         |
-- | 4          | web-based FRONT-end development |
-- +------------+---------------------------------+
-- Output:
-- +------------+---------------------------------+---------------------------------+
-- | content_id | original_text                   | converted_text                  |
-- +------------+---------------------------------+---------------------------------+
-- | 1          | hello world of SQL              | Hello World Of Sql              |
-- | 2          | the QUICK-brown fox             | The Quick-Brown Fox             |
-- | 3          | modern-day DATA science         | Modern-Day Data Science         |
-- | 4          | web-based FRONT-end development | Web-Based Front-End Development |
-- +------------+---------------------------------+---------------------------------+
-- Explanation:
-- For content_id = 1:
-- Each word's first letter is capitalized: "Hello World Of Sql"
-- For content_id = 2:
-- Contains the hyphenated word "QUICK-brown" which becomes "Quick-Brown"
-- Other words follow normal capitalization rules
-- For content_id = 3:
-- Hyphenated word "modern-day" becomes "Modern-Day"
-- "DATA" is converted to "Data"
-- For content_id = 4:
-- Contains two hyphenated words: "web-based" → "Web-Based"
-- And "FRONT-end" → "Front-End"

-- CREATE TABLE If not exists user_content (
--     content_id INT,
--     content_text VARCHAR(255)
-- )
-- Truncate table user_content
-- insert into user_content (content_id, content_text) values ('1', 'hello world of SQL')
-- insert into user_content (content_id, content_text) values ('2', 'the QUICK-brown fox')
-- insert into user_content (content_id, content_text) values ('3', 'modern-day DATA science')
-- insert into user_content (content_id, content_text) values ('4', 'web-based FRONT-end development')

-- repalce 大法好
SELECT
    content_id, 
    content_text AS original_text,
    REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(CONCAT(UPPER(SUBSTRING(content_text, 1, 1)), LOWER(SUBSTRING(content_text, 2, LENGTH(content_text) - 1))), " z", " Z"), " y", " Y"), " x", " X"), " w", " W"), " v", " V"), " u", " U"), " t", " T"), " s", " S"), " r", " R"), " q", " Q"), " p", " P"), " o", " O"), " n", " N"), " m", " M"), " l", " L"), " k", " K"), " j", " J"), " i", " I"), " h", " H"), " g", " G"), " f", " F"), " e", " E"), " d", " D"), " c", " C"), " b", " B"), " a", " A") ,"-z", "-Z"),"-y", "-Y"),"-x", "-X"),"-w", "-W"),"-v", "-V"),"-u", "-U"),"-t", "-T"),"-s", "-S"),"-r", "-R"),"-q", "-Q"),"-p", "-P"),"-o", "-O"),"-n", "-N"),"-m", "-M"),"-l", "-L"),"-k", "-K"),"-j", "-J"),"-i", "-I"),"-h", "-H"),"-g", "-G"),"-f", "-F"),"-e", "-E"),"-d", "-D"),"-c", "-C"),"-b", "-B"),"-a", "-A") AS converted_text 
FROM 
    user_content