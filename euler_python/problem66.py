"""
problem66.py

Consider quadratic Diophantine equations of the form: x**2 – Dy**2 = 1.
For example, when D=13, the minimal solution in x is 6492 – 13×1802 = 1.
It can be assumed that there are no solutions in positive integers when
D is square. By finding minimal solutions in x for D = {2, 3, 5, 6, 7},
we obtain the following:

    3**2 – 2×2**2 = 1
    2**2 – 3×1**2 = 1
    9**2 – 5×4**2 = 1
    5**2 – 6×2**2 = 1
    8**2 – 7×3**2 = 1

Hence, by considering minimal solutions in x for D ≤ 7, the largest x is
obtained when D=5.  Find the value of D ≤ 1000 in minimal solutions of x for
which the largest value of x is obtained.
"""
from collections import deque
from fractions import Fraction
from itertools import count
from math import floor, sqrt
from toolset import take

# Each iteration through the convergents of the continued fraction of sqrt(D),
# we want to check whether the numerator and denominator provide a solution to
# the Diophantine equation: https://en.wikipedia.org/wiki/Pell%27s_equation
# See the section entitled 'Fundamental solution via continued fractions'

def process_cf(D):
    """Yield the values in the continued fraction representation of sqrt(D),
    e.g. sqrt(23) = [4;(1,3,1,8)], so yield 4, 1, 3, 1, 8, 1, 3, 1, 8, ..."""
    # See problem64.py for a link explaining this algorithm. Here we use 'D'
    # in place of 'S' to be consistent with the wording of the question.
    m = 0
    d = 1
    a = floor(sqrt(D))
    while True:
        yield a
        m = (d * a) - m
        d = (D - m**2) / d
        a = floor((floor(sqrt(D)) + m) / d)

def convergent(D, n):
    """Return the nth convergent of the continued fraction for sqrt(D),
    where D is a non-square positive integer."""
    if n == 1:
        return next(process_cf(D))
    # Collect the first n partial values of D.
    values = deque(take(n, process_cf(D)))
    # Construct the continued fraction, where 'tail' is the recursive component.
    return Fraction(values.popleft() + Fraction(1, tail(values)))

def tail(values):
    "Recursively return the tail end of the continued fraction for sqrt(D)"
    next = values.popleft()
    if len(values) == 0:
        return next
    return next + Fraction(1, tail(values))

def solve_pells_equation(D):
    def is_solution(frac):
        "Check whether the convergent satisfies the Diophantine equation"
        x, y = frac.numerator, frac.denominator
        return x**2 - D*(y**2) == 1
    # Find the solution with the minimal value of x satisfying the equation.
    candidates = (convergent(D, n) for n in count(1))
    solution = next(filter(is_solution, candidates))
    # For the purpose of problem 66, we only need the value of x
    return solution.numerator

def problem66():
    solutions = [(i, solve_pells_equation(i))
                 for i in range(1, 1000+1)
                 if sqrt(i).is_integer() == False]
    # Find the solution wth the largest value of x
    answer = max(solutions, key=lambda s: s[1])
    # Return the value of D for which that value of x was obtained
    return answer[0]
