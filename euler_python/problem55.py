"""
problem55.py

If we take 47, reverse and add, 47 + 74 = 121, which is palindromic. A number
that never forms a palindrome through the reverse and add process is called a
Lychrel number. How many Lychrel numbers are there below ten-thousand? (Only
consider fifty iterations)
"""
from toolset import iterate, quantify, take, memoize

def rev(n):
    """Return the reverse of n's digits"""
    return int(''.join(reversed(str(n))))

@memoize
def is_palindromic(n):
    return n == rev(n)

def is_lychrel(n):
    start = n + rev(n)
    iterations = iterate(lambda x: x + rev(x), start)
    return not any(is_palindromic(y) for y in take(50, iterations))

def problem55():
    return quantify(range(1, 10000), pred=is_lychrel)
