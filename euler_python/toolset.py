from itertools import count, islice
from fractions import Fraction
from functools import lru_cache, reduce
from math import sqrt, floor
import operator

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

# From Alex Martelli: http://stackoverflow.com/a/2212090
def eratosthenes():
    """Yields the sequence of prime numbers via the Sieve of Eratosthenes."""
    D = {}  # map each composite integer to its first-found prime factor
    for q in count(2):  # q gets 2, 3, 4, 5, ... ad infinitum
        p = D.pop(q, None)
        if p is None:
            # q not a key in D, so q is prime, therefore, yield it
            yield q
            # mark q squared as not-prime (with q as first-found prime factor)
            D[q*q] = q
        else:
            # let x <- smallest (N*p)+q which wasn't yet known to be composite
            # we just learned x is composite, with p first-found prime factor,
            # since p is the first-found prime factor of q -- find and mark it
            x = p + q
            while x in D:
                x += p
            D[x] = p

def prime_factors(num, start=2):
    """Return all prime factors (ordered) of num in a list"""
    # From https://github.com/tokland/pyeuler
    candidates = range(start, int(sqrt(num)) + 1)
    factor = next((x for x in candidates if (num % x == 0)), None)
    return ([factor] + prime_factors(num // factor, factor) if factor else [num])

def phi(n):
    ps = list(unique(prime_factors(n)))
    return int(n * reduce(operator.mul, (1 - Fraction(1, p) for p in ps)))

def iterate(func, arg):
    # Based on Clojure's function of the same name.
    while True:
        yield arg
        arg = func(arg)

class memoize_mutable:
    """Memoize functions with mutable arguments."""
    # Attributed to Alex Martelli: http://stackoverflow.com/a/4669720
    def __init__(self, fn):
        self.fn = fn
        self.memo = {}
    def __call__(self, *args, **kwds):
        import pickle
        str = pickle.dumps(args, 1) + pickle.dumps(kwds, 1)
        if str not in self.memo:
            # print("miss") # DEBUG INFO
            self.memo[str] = self.fn(*args, **kwds)
        # else:
            # print("hit")  # DEBUG INFO
        return self.memo[str]

#------------------------------------------------------------------------------
# Adapted from the Python itertools docs

def take(n, iterable):
    """Return first n items of the iterable as a list."""
    return islice(iterable, n)

def quantify(iterable, pred=bool):
    """Count how many times the predicate is true."""
    return sum(map(pred, iterable))

def unique(iterable, key=None):
    "List unique elements, preserving order. Remember all elements ever seen."
    # unique('AAAABBBCCDAABBB') --> A B C D
    # unique('ABBCcAD', str.lower) --> A B C D
    seen = set()
    if key is None:
        for element in iterable:
            if element not in seen:
                seen.add(element)
                yield element
    else:
        for element in iterable:
            k = key(element)
            if k not in seen:
                seen.add(k)
                yield element
