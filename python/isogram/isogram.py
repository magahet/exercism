def is_isogram(string):
    alpha_chars = [c.lower() for c in string if c.isalpha()]
    return len(alpha_chars) == len(set(alpha_chars))
