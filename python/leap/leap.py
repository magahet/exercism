def is_leap_year(year):
    '''Return whether a given year is a leap year.'''

    def is_divisible(dividend, divisor):
        '''Return whether a dividend can be evenly divided by a divisor.'''
        return dividend % divisor == 0

    return is_divisible(year, 4) and (
        not is_divisible(year, 100) or is_divisible(year, 400))
