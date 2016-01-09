from itertools import count
from functools import lru_cache
from math import sqrt, floor

memoize = lru_cache(maxsize=None)

@memoize
def is_prime(n):
    """Return True if n is a prime number (1 is not considered prime)."""
    # Inspired by https://github.com/tokland/pyeuler
    if n < 3:
        return (n == 2)
    elif n % 2 == 0:
        return False
    elif any(n % x == 0 for x in range(3, floor(sqrt(n))+1, 2)):
        return False
    return True

def get_primes(start=2):
    """Yield prime numbers from start."""
    return filter(is_prime, count(start))

#------------------------------------------------------------------------------
# From the Python itertools docs

def quantify(iterable, pred=bool):
    "Count how many times the predicate is true"
    return sum(map(pred, iterable))

def first_true(iterable, default=False, pred=None):
    """Returns the first true value in the iterable.

    If no true value is found, returns *default*

    If *pred* is not None, returns the first item
    for which pred(item) is true.
    """
    # first_true([a,b,c], x) --> a or b or c or x
    # first_true([a,b], x, f) --> a if f(a) else b if f(b) else x
    return next(filter(pred, iterable), default)
