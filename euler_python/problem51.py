"""
problem51.py

https://projecteuler.net/problem=51

By replacing the 1st digit of the 2-digit number *3, it turns out that six of
the nine possible values: 13, 23, 43, 53, 73, and 83, are all prime.

By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit
number is the first example having seven primes among the ten generated
numbers, yielding the family: 56003, 56113, 56333, 56443, 56663, 56773, and
56993. Consequently 56003, being the first member of this family, is the
smallest prime with this property.

Find the smallest prime which, by replacing part of the number (not necessarily
adjacent digits) with the same digit, is part of an eight prime value family.
"""
from itertools import count, product
from toolset import quantify, is_prime, get_primes

# Our strategy is as follows. Since we are seeking an eight prime family, it
# must be the case that the pattern of digits which are replaced contains
# either 0, 1, or 2 in the smallest family member. Therefore, we can search
# through primes and replace digits in patterns specified by the locations 0,
# 1, and 2. If the family of numbers that results contains eight primes, we
# have found the solution.
#
# In the example given, 56003 is the smallest member of an eight prime family.
# We would find the pattern of 0s at indices (2, 3) to produce the
# corresponding family from 56**3.
def find_indices(num):
    """Yield three tuples, where each tuple contains the indices
    in the num of the digits 0, 1, and 2 respectively."""
    # find_indices(18209912) --> (3), (0, 6), (2, 7)
    # find_indices(56003) --> (2, 3), (), ()
    digits = str(num)
    for target in "012":
        yield tuple(i for i, x in enumerate(digits) if x == target)

def family(num, indices):
    """Yield the family of numbers resulting from replacing
    digits at the specific indices with the digits 0 to 9."""
    # family(56003, (2, 3)) --> 56003, 56113, 56223, 56333, 56443, ...
    template = str(num)
    for i in "0123456789":
        member = list(template)
        for idx in indices:
            member[idx] = i
        # yield sentinel value (-1) in case of leading zero
        yield int(''.join(member)) if member[0] != "0" else -1

def is_smallest_member(num):
    """Does the number satisfy the problem specification?"""
    return any(quantify(family(num, indices), pred=is_prime) == 8
               for indices in find_indices(num))

def problem51():
    return next(filter(is_smallest_member, get_primes(start=56995)))
