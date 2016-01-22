"""
problem63.py

The 5-digit number, 16807=7**5, is also a fifth power. Similarly, the 9-digit
number, 134217728=8**9, is a ninth power. How many n-digit positive integers
exist which are also an nth power?
"""
from itertools import count, takewhile

def find_powers(n):
    """Return the list of powers to which one can raise n such that
    the result of exponentiation is an integer with number of digits == power"""
    # find_powers(6) --> [1, 2, 3, 4]
    return list(takewhile(lambda x: len(str(n**x)) == x, count(1)))

def problem63():
    # Take results from find_powers(i) for i in count(1) until the empty list
    # indicates no more results. Find the length of all items of every resulting list.
    results = takewhile(lambda xs: xs != [], (find_powers(i) for i in count(1)))
    return sum(len(xs) for xs in results)
