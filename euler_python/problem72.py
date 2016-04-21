"""
problem72.py

https://projecteuler.net/problem=72

Consider the fraction, n/d, where n and d are positive integers. If n<d and
HCF(n,d)=1, it is called a reduced proper fraction.

If we list the set of reduced proper fractions for d ≤ 8 in ascending order of
size, we get:

1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8

It can be seen that there are 21 elements in this set. How many elements would
be contained in the set of reduced proper fractions for d ≤ 1,000,000?
"""

def phi_sieve(limit):
    """Yield all values of phi(i) for i in range(2, limit+1).
    By convention phi(0)=0 and phi(1)=1, but Project Euler
    appears to disagree, so those values are ignored."""
    # Solved in the style of the sieve of Eratosthenes.
    phi = list(range(limit+1))
    marked = set()
    for i in range(2, limit+1):
        if i not in marked:
            for j in range(i, limit+1, i):
                phi[j] *= (1 - 1/i)
                marked.add(j)
        yield phi[i]

def problem72():
    # See https://en.wikipedia.org/wiki/Farey_sequence
    # for the connection to Euler's phi function.
    # In particular, note that |F n| = |F n-1| + phi(n).
    return sum(phi_sieve(1000000))
