from collections import Counter
import string


def word_count(phrase):
    valid_chars = string.letters + string.digits + "'"
    split_table = string.maketrans(',_', '  ')
    count = Counter([
        ''.join([c for c in word.lower() if c in valid_chars]).strip("'")
        for word in phrase.translate(split_table).split()
    ])
    del count['']
    return count