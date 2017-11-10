"""
problem80.py

https://projecteuler.net/problem=80

It is well known that if the square root of a natural number is not an integer,
then it is irrational. The decimal expansion of such square roots is infinite
without any repeating pattern at all.

The square root of two is 1.41421356237309504880..., and the digital sum of the
first one hundred decimal digits is 475.

For the first one hundred natural numbers, find the total of the digital sums of
the first one hundred decimal digits for all the irrational square roots.
"""
from decimal import Decimal, getcontext
from math import sqrt

# Set the precision of decimal arithmetic operations to slightly above 100 to
# ensure that the first 100 digits are correct.
getcontext().prec = 102

def square_root_digits(n):
    """Find the first one hundred decimal digits of the square root of n."""
    return Decimal(n).sqrt().as_tuple().digits[:-2] # Cut off the last 2 of the 102 digits.

def has_irrational_sqrt(n):
    return sqrt(n).is_integer() == False

def problem80():
    return sum(sum(square_root_digits(n))
               for n in range(1, 100)
               if has_irrational_sqrt(n))

if __name__ == "__main__":
    print(problem80())
