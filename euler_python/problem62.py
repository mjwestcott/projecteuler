"""
problem62.py

The cube, 41063625 (345**3), can be permuted to produce two other cubes:
56623104 (384**3) and 66430125 (405**3). In fact, 41063625 is the smallest cube
which has exactly three permutations of its digits which are also cube. Find the
smallest cube for which exactly five permutations of its digits are cube.
"""
from collections import defaultdict
from itertools import count

def problem62():
    cubes = defaultdict(list)
    for i in count():
        # Add i**3 to a dict, the keys of which are an arbitrarily specified
        # canonical permutation of its digits; the values of which are lists of
        # all cube numbers seen so far which are permutable to the key. Every
        # iteration we check whether the list under the current key contains
        # five members, and if so return the smallest member.
        canonical = ''.join(sorted(str(i**3)))
        seen = cubes[canonical]
        seen.append(i**3)
        if len(seen) == 5:
            return min(seen)
