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

#------------------------------------------------------------------------------
# From the Python itertools docs

def quantify(iterable, pred=bool):
    "Count how many times the predicate is true"
    return sum(map(pred, iterable))
