"""
problem69.py

https://projecteuler.net/problem=69

Euler's Totient function, φ(n) [sometimes called the phi function], is used to
determine the number of numbers less than n which are relatively prime to n. For
example, as 1, 2, 4, 5, 7, and 8, are all less than nine and relatively prime to
nine, φ(9)=6.

It can be seen that n=6 produces a maximum n/φ(n) for n ≤ 10. Find the value of
n ≤ 1,000,000 for which n/φ(n) is a maximum.
"""
from itertools import takewhile

from toolset import get_primes

# from fractions import Fraction
# from toolset import prime_factors
# def phi(n):
#     ps = list(unique(prime_factors(n)))
#     return n * reduce(operator.mul, (1 - Fraction(1, p) for p in ps))
# return max((n for n in range(2, 1000000+1)), key=lambda n: n/phi(n))
#
# The commented-out solution above is correct and true to the problem
# description, but slightly slower than 1 minute.
#
# So, note that the phi function multiplies n by (1 - (1/p)) for every p in
# its unique prime factors. Therefore, phi(n) will diminish as n has a
# greater number of small unique prime factors. Since we are seeking the
# largest value for n/phi(n), we want to minimize phi(n). We are therefore
# looking for the largest number <= 1e6 which is the product of the smallest
# unique prime factors, i.e successive prime numbers starting from 2.

def candidates():
    primes = get_primes()
    x = next(primes)
    while True:
        yield x
        x *= next(primes)

def problem69():
    return max(takewhile(lambda x: x < 1e6, candidates()))

if __name__ == "__main__":
    print(problem69())
