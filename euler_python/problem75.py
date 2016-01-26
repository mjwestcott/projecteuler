"""
problem75.py

It turns out that 12 cm is the smallest length of wire that can be bent to form
an integer sided right angle triangle in exactly one way, but there are many
more examples.

    12 cm: (3,4,5)
    24 cm: (6,8,10)
    30 cm: (5,12,13)
    36 cm: (9,12,15)
    40 cm: (8,15,17)
    48 cm: (12,16,20)

In contrast, some lengths of wire, like 20 cm, cannot be bent to form an integer
sided right angle triangle, and other lengths allow more than one solution to be
found; for example, using 120 cm it is possible to form exactly three different
integer sided right angle triangles.

    120 cm: (30,40,50), (20,48,52), (24,45,51)

Given that L is the length of the wire, for how many values of L ≤ 1,500,000 can
exactly one integer sided right angle triangle be formed?
"""
from collections import Counter
from itertools import count, takewhile

def children(triple):
    """Given a pythagorean triple, return its three children triples."""
    # See Berggren's ternary tree, which will produce all infinitely many
    # primitive triples without duplication.
    a, b, c = triple
    a1, b1, c1 = (-a + 2*b + 2*c), (-2*a + b + 2*c), (-2*a + 2*b + 3*c)
    a2, b2, c2 = (+a + 2*b + 2*c), (+2*a + b + 2*c), (+2*a + 2*b + 3*c)
    a3, b3, c3 = (+a - 2*b + 2*c), (+2*a - b + 2*c), (+2*a - 2*b + 3*c)
    return (a1, b1, c1), (a2, b2, c2), (a3, b3, c3)

def problem75():
    limit = 1500000
    # A mapping from values of L to the number of right-angled triangles with
    # the perimeter L
    triangles = Counter()
    # Use a depth-first search to exhaust the search space, starting with the
    # first pythagorean triple.
    frontier = [(3, 4, 5)]
    while frontier:
        triple = frontier.pop()
        L = sum(triple)
        if L > limit:
            continue
        triangles[L] += 1
        a, b, c = triple
        # We're not only interested in 'primitive triples', but multiples too.
        multiples = takewhile(lambda m: sum(m) < limit, ((i*a, i*b, i*c) for i in count(2)))
        for m in multiples:
            triangles[sum(m)] += 1
        for child in children(triple):
            frontier.append(child)
    return sum(triangles[L] == 1 for L in triangles)
