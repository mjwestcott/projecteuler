"""
problem85.py

By counting carefully it can be seen that a rectangular grid measuring 3 by 2
contains eighteen rectangles.

Although there exists no rectangular grid that contains exactly two million
rectangles, find the area of the grid with the nearest solution.
"""
from collections import namedtuple
from math import factorial

# To create a rectangle one needs to choose two horizontal lines and two
# vertical lines. For an x by y grid the formula for number of rectangles is
# thus: (x+1 choose 2) * (y+1 choose 2)
def ncombinations(n, r):
    return factorial(n) / (factorial(r) * factorial(n - r))

def nrectangles(x, y):
    return ncombinations(x+1, 2) * ncombinations(y+1, 2)

def problem85():
    target = 2000000
    candidate = namedtuple("candidate", ["area", "nrectangles"])
    cs = (candidate(x*y, nrectangles(x, y)) for x in range(2, 100) for y in range(2, 100))
    solution = min(cs, key=lambda c: abs(c.nrectangles - target))
    return solution.area
