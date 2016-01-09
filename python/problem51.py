"""
problem51.py

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
from toolset import quantify, first_true, is_prime, get_primes

def to_digits(num):
    # to_digits(1234) --> [1, 2, 3, 4]
    return map(int, str(num))

def to_num(digits):
    # to_num([1, 2, 3, 4]) --> 1234
    return int(''.join(map(str, digits)))

def digitlen(num):
    # digitlen(9999) --> 4
    return len(str(abs(num)))

# We will use tuples of 1s and 0s to represent all the ways
# to replace digits in a number. For example, (0, 1, 0) means
# replace the second digit of a three digit number. For this
# problem, we do not want masks ending in a set bit because
# the final digit of a prime over 5 must end in 1, 3, 7, or 9.
# Thus it would be impossible to make an eight prime value family
# which replaces the final digit.
def binary_masks(n):
    """Return iterator over all n-digit binary masks except those ending
    in a set bit."""
    # binary_masks(3) --> (0, 0, 0), (0, 1, 0), (1, 0, 0), (1, 1, 0)
    return filter(lambda x: x[-1] == 0, product([0, 1], repeat=n))

def replace_digits(num, mask, val):
    """Replace digits in num with val at indicies specified by the mask.
    If result leads with a zero, return a sentinel value (-1)."""
    # replace_digits(3537, [1, 1, 0, 0], 9) --> 9937
    replace_if_set = lambda bit, x: val if bit else x
    digits = list(map(replace_if_set, mask, to_digits(num)))
    return to_num(digits) if digits[0] != 0 else -1

def problem51():
    # For each prime above 56993, for all possible binary masks
    # representing ways to replace digits in that number, yield
    # the corresponding family of ten digits
    families = ([replace_digits(n, mask, val) for val in range(10)]
                for n in get_primes(start=56995)
                for mask in binary_masks(digitlen(n)))
    # Find the next solution, where solution is the first prime
    # member of the family for which eight members are prime.
    return next(first_true(family, pred=is_prime)
                for family in families
                if quantify(family, pred=is_prime) == 8)
