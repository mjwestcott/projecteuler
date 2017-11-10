"""
problem79.py

https://projecteuler.net/problem=79

A common security method used for online banking is to ask the user for three
random characters from a passcode. For example, if the passcode was 531278, they
may ask for the 2nd, 3rd, and 5th characters; the expected reply would be: 317.

The text file, keylog.txt, contains fifty successful login attempts.

Given that the three characters are always asked for in order, analyse the file
so as to determine the shortest possible secret passcode of unknown length.
"""
from collections import defaultdict, deque
from itertools import dropwhile


def to_digits(num):
    return map(int, str(num))

def to_num(digits):
    return int(''.join(map(str, digits)))

# Use 'breadth-first tree search', inspired by Peter Norvig's version in AIMA.
def solve(codes):
    # Store all relations specified in the codes in a dict. Each digit
    # is mapped to those digits appearing after it.
    after = defaultdict(set)
    for code in codes:
        a, b, c = to_digits(code)
        after[a].add(b)
        after[a].add(c)
        after[b].add(c)

    # We will use lists to represent nodes in the tree, each of which is
    # a candidate solution. So, initialise the frontier to the possible
    # starting values.
    frontier = deque([x] for x in after)
    while frontier:
        node = frontier.popleft()
        if goal_state(node, after):
            return node
        # Use the 'after' dict to find the values, x, reachable from the end of
        # the current node. Child nodes are then node + [x].
        frontier.extend(node + [x] for x in after[node[-1]])

def goal_state(node, after):
    """Check whether, for all the relations specified in the 'after' dict,
    the node satisfies them."""
    # For each key, x, in the 'after' dict, the values, y, in after[x] must
    # exist after the first occurrence of x in the node.
    return all(y in dropwhile(lambda dgt: dgt != x, node)
               for x in after
               for y in after[x])

def problem79():
    with open("data/keylog.txt", "r") as f:
        codes = [int(x) for x in f.readlines()]
    solution = solve(codes)
    return to_num(solution)

if __name__ == "__main__":
    print(problem79())
