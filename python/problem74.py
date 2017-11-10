"""
problem74.py

https://projecteuler.net/problem=74

The number 145 is well known for the property that the sum of the factorial of
its digits is equal to 145:

    1! + 4! + 5! = 1 + 24 + 120 = 145

Perhaps less well known is 169, in that it produces the longest chain of numbers
that link back to 169; it turns out that there are only three such loops that
exist:

    169 → 363601 → 1454 → 169
    871 → 45361 → 871
    872 → 45362 → 872

It is not difficult to prove that EVERY starting number will eventually get
stuck in a loop. For example,

    69 → 363600 → 1454 → 169 → 363601 (→ 1454)
    78 → 45360 → 871 → 45361 (→ 871)
    540 → 145 (→ 145)

Starting with 69 produces a chain of five non-repeating terms, but the longest
non-repeating chain with a starting number below one million is sixty terms. How
many chains, with a starting number below one million, contain exactly sixty
non-repeating terms?
"""
from math import factorial

from toolset import memoize, quantify


@memoize
def sum_factorial_digits(n):
    digits = map(int, str(n))
    return sum(factorial(x) for x in digits)

def problem74():
    # Known chain loop lengths given in problem description. We will use this
    # dictionary to cache all further results as we calculate them.
    known_loops = {145: 1, 169: 3, 1454: 3, 871: 2, 872: 2, 69: 5, 78: 4, 540: 2}
    lengths = (chain_length(n) for n in range(1, 1000000+1))
    def chain_length(n):
        chain = [n]
        next = sum_factorial_digits(n)
        while True:
            if next in chain: # We have found a new loop, add to the cache.
                result = known_loops[n] = len(chain)
                return result
            if next in known_loops: # We have found a known loop, add its length to current chain.
                result = known_loops[n] = len(chain) + known_loops[next]
                return result
            # We haven't found a loop, continue to investigate the chain.
            chain.append(next)
            next = sum_factorial_digits(next)
    return quantify(lengths, pred=lambda x: x == 60)

if __name__ == "__main__":
    print(problem74())
