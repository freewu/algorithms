# 2199 · Remove empty lines from files
# Description
# Write python code that reads the contents of the file from the given path input_path by line, 
# removes the empty lines, and writes them to the given path output_path.
# Write the python code to realize these functions in solution.py,
# and we will import your module in main.py to check if your code achieve the above.

# Pay attention to the indentation between classes and methods
# Pay attention to Chinese and English punctuation issues
# Example
# The evaluator will execute your code by executing the command python main.py <input_path> <output_path>

# and passing the input_path which read the file, output_path which write to the file as arguments.

# When you read the following via input_path
# When You Are Old

# Yeats

# When you are old and grey and full of sleep,

# And nodding by the fire, take down this book,

# And slowly read, and dream of the soft look,

# Your eyes had once, and of their shadows deep;

# How many loved your moments of glad grace,

# And loved your beauty with love false or true,

# But one man loved the pilgrim soul in you,

# And loved the sorrows of your changing face;

# And bending down beside the glowing bars,

# Murmur, a little sadly, how love fled,

# And paced upon the mountains overhead.

# And hid his face amid a crowd of stars.
# You need to write the following to output_path

# When You Are Old
# Yeats
# When you are old and grey and full of sleep,
# And nodding by the fire, take down this book,
# And slowly read, and dream of the soft look,
# Your eyes had once, and of their shadows deep;
# How many loved your moments of glad grace,
# And loved your beauty with love false or true,
# But one man loved the pilgrim soul in you,
# And loved the sorrows of your changing face;
# And bending down beside the glowing bars,
# Murmur, a little sadly, how love fled,
# And paced upon the mountains overhead.
# And hid his face amid a crowd of stars.

def remove_empty_lines(input_path, output_path):
    '''
    read the file in the input_path path
    remove blank lines
    write the data with the blank lines removed to the file
    '''
    # -- write your code here --
    lines = ""
    with open(input_path, 'r', encoding="utf-8") as f:
        lines = f.readlines()

    new_lines = []
    for line in lines:
        if line == '\n': # 空行跳过
            continue
        new_lines.append(line)

    with open(output_path, 'w', encoding="utf-8") as f:
        f.write(''.join(new_lines))