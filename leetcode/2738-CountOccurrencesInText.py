# 2738. Count Occurrences in Text
# Table: Files
# +-------------+---------+
# | Column Name | Type    |
# +-- ----------+---------+
# | file_name   | varchar |
# | content     | text    |
# +-------------+---------+
# file_name is the column with unique values of this table. 
# Each row contains file_name and the content of that file.
# Write a solution to find the number of files that have at least one occurrence of the words 'bull' and 'bear' as a standalone word, respectively, disregarding any instances where it appears without space on either side (e.g. 'bullet', 'bears', 'bull.', or 'bear' at the beginning or end of a sentence will not be considered) 

# Return the word 'bull' and 'bear' along with the corresponding number of occurrences in any order.
# The result format is in the following example.

# Example 1:
# Input: 
# Files table:
# +------------+----------------------------------------------------------------------------------+
# | file_name  | content                                                                         | 
# +------------+----------------------------------------------------------------------------------+
# | draft1.txt | The stock exchange predicts a bull market which would make many investors happy. | 
# | draft2.txt | The stock exchange predicts a bull market which would make many investors happy, |
# |            | but analysts warn of possibility of too much optimism and that in fact we are    |
# |            | awaiting a bear market.                                                          | 
# | draft3.txt | The stock exchange predicts a bull market which would make many investors happy, |
# |            | but analysts warn of possibility of too much optimism and that in fact we are    |
# |            | awaiting a bear market. As always predicting the future market is an uncertain   |
# |            | game and all investors should follow their instincts and best practices.         | 
# +------------+----------------------------------------------------------------------------------+
# Output: 
# +------+-------+
# | word | count |  
# +------+-------+
# | bull | 3     | 
# | bear | 2     | 
# +------+-------+
# Explanation: 
# - The word "bull" appears 1 time in "draft1.txt", 1 time in "draft2.txt", and 1 time in "draft3.txt". Therefore, the total number of occurrences for the word "bull" is 3.
# - The word "bear" appears 1 time in "draft2.txt", and 1 time in "draft3.txt". Therefore, the total number of occurrences for the word "bear" is 2.

import pandas as pd

# find + len
def count_occurrences(files: pd.DataFrame) -> pd.DataFrame:
    data = [
        ["bull", len(files[files['content'].str.find(' bull ') != -1])],
        ["bear", len(files[files['content'].str.find(' bear ') != -1])]
    ]
    return pd.DataFrame(data, columns=["word","count"])

# contains + sum
def count_occurrences1(files: pd.DataFrame) -> pd.DataFrame:
    data = [
        ["bull", files['content'].str.contains(' bull ', regex = False).sum()],
        ["bear", files['content'].str.contains(' bear ', regex = False).sum()]
    ]
    return pd.DataFrame(data, columns=["word","count"])

if __name__ == "__main__":
    data = [['draft1.txt', 'The stock exchange predicts a bull market which would make many investors happy.'], ['draft2.txt', 'The stock exchange predicts a bull market which would make many investors happy, but analysts warn of possibility of too much optimism and that in fact we are awaiting a bear market.'], ['final.txt', 'The stock exchange predicts a bull market which would make many investors happy, but analysts warn of possibility of too much optimism and that in fact we are awaiting a bear market. As always predicting the future market is an uncertain game and all investors should follow their instincts and best practices.']]
    files = pd.DataFrame(data, columns=['file_name', 'content']).astype({'file_name':'object', 'content':'object'})
    print(count_occurrences(files))
    print(count_occurrences1(files))
