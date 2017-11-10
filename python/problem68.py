"""
problem68.py

https://projecteuler.net/problem=68

What is the maximum 16-digit string for a 'magic' 5-gon ring?
"""
from itertools import chain

flatten = chain.from_iterable

def five_gon_rings(n):
    """Return list of solutions to the 'magic' 5-gon ring problem, each line
    summing to n. The empty list will be returned if there are no solutions."""
    rings = [([a, b, c], [d, c, e], [f, e, g], [h, g, i], [j, i, b])
             for a in range(1, 10+1)
             for b in range(1, 10+1) if b != a
             for c in range(1, 10+1) if c not in [a, b]
             if a + b + c == n
             for d in range(1, 10+1) if d not in [a, b, c]
             for e in range(1, 10+1) if e not in [a, b, c, d]
             if d + c + e == n
             for f in range(1, 10+1) if f not in [a, b, c, d, e]
             for g in range(1, 10+1) if g not in [a, b, c, d, e, f]
             if f + e + g == n
             for h in range(1, 10+1) if h not in [a, b, c, d, e, f, g]
             for i in range(1, 10+1) if i not in [a, b, c, d, e, f, g, h]
             if h + g + i == n
             for j in range(1, 10+1) if j not in [a, b, c, d, e, f, g, h, i]
             if j + i + b == n
             if a < min(d, f, h, j)]
    # Each solution can be described uniquely starting from the group of three
    # with the numerically lowest external node and working clockwise.
    # So we specified at the end that a < min(d, f, h, j)
    return rings

def problem68():
    START = 6   # each line cannot sum to less than 6 (1+2+3)
    END = 27+1  # or greater than 27 (8+9+10)

    # Collect solution candidates, flattening into one array of solutions.
    rings = flatten(five_gon_rings(n) for n in range(START, END))
    # Filter out the empty lists
    rings = filter(bool, rings)
    # Transform each solution tuple into a string of digits.
    rings = [''.join(str(x) for x in flatten(solution)) for solution in rings]
    # Find the max 16-digit string.
    return int(max(solution for solution in rings if len(solution) == 16))

if __name__ == "__main__":
    print(problem68())
