"""
problem57.py

https://projecteuler.net/problem=57

It is possible to show that the square root of two can be expressed as an
infinite continued fraction.

    √ 2 = 1 + 1/(2 + 1/(2 + 1/(2 + ... ))) = 1.414213...

By expanding this for the first four iterations, we get:

    1 + 1/2 = 3/2 = 1.5
    1 + 1/(2 + 1/2) = 7/5 = 1.4
    1 + 1/(2 + 1/(2 + 1/2)) = 17/12 = 1.41666...
    1 + 1/(2 + 1/(2 + 1/(2 + 1/2))) = 41/29 = 1.41379...

The next three expansions are 99/70, 239/169, and 577/408, but the eighth
expansion, 1393/985, is the first example where the number of digits in the
numerator exceeds the number of digits in the denominator. In the first
one-thousand expansions, how many fractions contain a numerator with more
digits than denominator?
"""
from fractions import Fraction
from toolset import iterate, quantify, take

def num_digits(n):
    """How many digits in integer n?"""
    return len(str(n))

def check_numerator(frac):
    """Does the numerator have more digits than the denominator?"""
    return num_digits(frac.numerator) > num_digits(frac.denominator)

def tail(n):
    """The repeating pattern at the end of the expansions"""
    return 2 + Fraction(1, n)

def problem57():
    generate_tails = iterate(tail, 2) # yields 2, tail(2), tail(tail(2)), ...
    expansions = (1 + Fraction(1, t) for t in generate_tails)
    return quantify(take(1000, expansions), pred=check_numerator)
