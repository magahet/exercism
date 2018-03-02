import collections


def detect_anagrams(word, candidates):
    count = collections.Counter(word.lower())
    return [
        c for c in candidates
        if c.lower() != word.lower()
        and collections.Counter(c.lower()) == count
    ]