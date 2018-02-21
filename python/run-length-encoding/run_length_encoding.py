def decode(string):
    message = ''
    num = ''
    for char in string:
        if char.isdigit():
            num += char
        else:
            message += char * int(num) if num else char
            num = ''

    return message


def encode(string, count=1):
    if not string:
        return ''

    prefix = str(count) if count > 1 else ''

    if len(string) == 1:
        return ''.join([prefix, string[0]])
    elif string[0] == string[1]:
        return encode(string[1:], count + 1)
    else:
        return ''.join([prefix, string[0], encode(string[1:])])
