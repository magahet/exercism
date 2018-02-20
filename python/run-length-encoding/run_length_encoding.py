def decode(string):

    def iter_next_char(chars):
        num = ''
        for char in chars:
            if char.isdigit():
                num += char
            elif num:
                yield char * int(num)
                num = ''
            else:
                yield char

    return ''.join(iter_next_char(string))


def encode(string):

    def iter_next_char_count(chars):
        next_char = chars.pop(0)
        offset = -1
        for idx, char in enumerate(chars):
            if next_char == char:
                continue

            yield str(idx - offset), next_char

            offset = idx
            next_char = char

        if char == next_char:
            yield str(idx - offset + 1), next_char

    if not string:
        return string

    return ''.join([
        r[1] if r[0] == '1' else ''.join(r)
        for r in iter_next_char_count(list(string))
    ])
