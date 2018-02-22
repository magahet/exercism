import calendar
import datetime


def meetup_day(year, month, day_of_the_week, which):
    # Get day of the week: Mon = 0, Tue = 1, etc.
    dow = [calendar.day_name[i] for i in xrange(7)].index(day_of_the_week)
    mon = calendar.monthcalendar(year, month)
    # Get list of days that land on a given day of the week
    days = [w[dow] for w in mon if w[dow] != 0]
    # lookup the specific day using which
    if which[0].isdigit():
        day = days[int(which[0]) - 1]
    elif which == 'last':
        day = days[-1]
    elif which == 'teenth':
        day = filter(lambda d: d in range(13, 20), days)[0]
    else:
        raise Exception('Could not parse which')
    return datetime.date(year, month, day)