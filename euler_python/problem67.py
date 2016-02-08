"""
problem67.py

https://projecteuler.net/problem=67

Find the maximum total from top to bottom in triangle.txt a 15K text
file containing a triangle with one-hundred rows.
"""
from toolset import memoize_mutable

@memoize_mutable
def max_route(triangle):
    """Recursively find the maximum value of the root node plus the
    largest of its children, and so on, all the way to the base."""
    # where triangle is a list of lists such as [[1], [2, 3], [4, 5, 6]]
    # representing a tree of the form:
    #   1
    #  2 3
    # 4 5 6
    root = triangle[0][0]
    if len(triangle) == 1:
        return root
    a, b = children(triangle)
    return root + max(max_route(a), max_route(b))

def children(triangle):
    "Split the triangle in two below the root node"
    # [[1], [2, 3], [4, 5, 6]] --> [[2], [4, 5]], [[3], [5, 6]]
    # the two children triangles of the root node.
    a = [row[:-1] for row in triangle[1:]]
    b = [row[1:] for row in triangle[1:]]
    return a, b

def problem67():
    with open("data/triangle.txt", "r") as f:
        triangle = [list(map(int, row.split())) for row in f]
    return max_route(triangle)
