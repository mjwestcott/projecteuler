"""
problem55.py

If we take 47, reverse and add, 47 + 74 = 121, which is palindromic. A number
that never forms a palindrome through the reverse and add process is called a
Lychrel number. How many Lychrel numbers are there below ten-thousand? (Only
consider fifty iterations)
"""
from toolset import iterate, quantify, take

def to_digits(num):
    # to_digits(1234) --> [1, 2, 3, 4]
    return list(map(int, str(num)))

def to_num(digits):
    # to_num([1, 2, 3, 4]) --> 1234
    return int(''.join(map(str, digits)))

def is_palindromic(num):
    return to_digits(num) == list(reversed(to_digits(num)))

def is_lychrel(num):
    rev = lambda x: to_num(reversed(to_digits(x)))
    start = num + rev(num)
    iterations = iterate(lambda x: x + rev(x), start)
    return not any(is_palindromic(n) for n in take(50, iterations))

def problem55():
    return quantify(range(1, 10000), pred=is_lychrel)
