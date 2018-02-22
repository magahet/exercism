import string


def rotate(text, key):
    low = string.ascii_lowercase
    up = string.ascii_uppercase
    table = string.maketrans(
        low + up,
        low[key:] + low[:key] + up[key:] + up[:key])
    return text.translate(table)