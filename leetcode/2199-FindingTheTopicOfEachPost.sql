-- 2199. Finding the Topic of Each Post
-- Table: Keywords
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | topic_id    | int     |
-- | word        | varchar |
-- +-------------+---------+
-- (topic_id, word) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table contains the id of a topic and a word that is used to express this topic.
-- There may be more than one word to express the same topic and one word may be used to express multiple topics.
 
-- Table: Posts
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | post_id     | int     |
-- | content     | varchar |
-- +-------------+---------+
-- post_id is the primary key (column with unique values) for this table.
-- Each row of this table contains the ID of a post and its content.
-- Content will consist only of English letters and spaces.
 
-- Leetcode has collected some posts from its social media website and is interested in finding the topics of each post. Each topic can be expressed by one or more keywords. If a keyword of a certain topic exists in the content of a post (case insensitive) then the post has this topic.
-- Write a solution to find the topics of each post according to the following rules:
--     If the post does not have keywords from any topic, its topic should be "Ambiguous!".
--     If the post has at least one keyword of any topic, its topic should be a string of the IDs of its topics sorted in ascending order and separated by commas ','. The string should not contain duplicate IDs.

-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Keywords table:
-- +----------+----------+
-- | topic_id | word     |
-- +----------+----------+
-- | 1        | handball |
-- | 1        | football |
-- | 3        | WAR      |
-- | 2        | Vaccine  |
-- +----------+----------+
-- Posts table:
-- +---------+------------------------------------------------------------------------+
-- | post_id | content                                                                |
-- +---------+------------------------------------------------------------------------+
-- | 1       | We call it soccer They call it football hahaha                         |
-- | 2       | Americans prefer basketball while Europeans love handball and football |
-- | 3       | stop the war and play handball                                         |
-- | 4       | warning I planted some flowers this morning and then got vaccinated    |
-- +---------+------------------------------------------------------------------------+
-- Output: 
-- +---------+------------+
-- | post_id | topic      |
-- +---------+------------+
-- | 1       | 1          |
-- | 2       | 1          |
-- | 3       | 1,3        |
-- | 4       | Ambiguous! |
-- +---------+------------+
-- Explanation: 
-- 1: "We call it soccer They call it football hahaha"
-- "football" expresses topic 1. There is no other word that expresses any other topic.
-- 2: "Americans prefer basketball while Europeans love handball and football"
-- "handball" expresses topic 1. "football" expresses topic 1. 
-- There is no other word that expresses any other topic.
-- 3: "stop the war and play handball"
-- "war" expresses topic 3. "handball" expresses topic 1.
-- There is no other word that expresses any other topic.
-- 4: "warning I planted some flowers this morning and then got vaccinated"
-- There is no word in this sentence that expresses any topic. Note that "warning" is different from "war" although they have a common prefix. 
-- This post is ambiguous.
-- Note that it is okay to have one word that expresses more than one topic.

-- Create table If Not Exists Keywords (topic_id int, word varchar(25))
-- Create table If Not Exists Posts (post_id int, content varchar(100))
-- Truncate table Keywords
-- insert into Keywords (topic_id, word) values ('1', 'handball')
-- insert into Keywords (topic_id, word) values ('1', 'football')
-- insert into Keywords (topic_id, word) values ('3', 'WAR')
-- insert into Keywords (topic_id, word) values ('2', 'Vaccine')
-- Truncate table Posts
-- insert into Posts (post_id, content) values ('1', 'We call it soccer They call it football hahaha')
-- insert into Posts (post_id, content) values ('2', 'Americans prefer basketball while Europeans love handball and football')
-- insert into Posts (post_id, content) values ('3', 'stop the war and play handball')
-- insert into Posts (post_id, content) values ('4', 'warning I planted some flowers this morning and then got vaccinated')

-- INSTR(str，substr)是MySQL的Sring函数。此方法返回字符串str中子字符串substr的第一次出现。
-- 语法：SELECT instr(str,substr); 
-- Select instr('Javatpoint','point'); -- 6
-- Select instr('MySql String Function', 'String'); -- 7

WITH t AS ( -- 得到每个帖子对应出现了哪些关键字
    SELECT 
        distinct p.post_id,
        k.topic_id 
    FROM 
        Keywords AS k,
        Posts AS p
    WHERE 
        INSTR(CONCAT(' ', p.content, ' '), CONCAT(' ',k.word, ' ')) > 0
)
-- SELECT * FROM t;
-- | post_id | topic_id |
-- | ------- | -------- |
-- | 1       | 1        |
-- | 2       | 1        |
-- | 3       | 3        |
-- | 3       | 1        |

(-- 按文章合并所有的主题
    SELECT 
        post_id, 
        GROUP_CONCAT(topic_id ORDER BY topic_id,'') AS topic 
    FROM 
        t
    GROUP BY 
        post_id
)
UNION ALL
( -- 处理没有匹配到关键字的文章
    SELECT 
        post_id,
        'Ambiguous!' AS topic 
    FROM 
        Posts
    WHERE
        post_id NOT IN ( SELECT DISTINCT post_id FROM t )
)
