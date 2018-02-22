def is_armstrong(number):
    digits = [int(d) for d in str(number)]
    return sum(map(lambda d: d ** len(digits), digits)) == number