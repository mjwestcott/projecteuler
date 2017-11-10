"""
problem83.py

https://projecteuler.net/problem=83

In the 5 by 5 matrix below, the minimal path sum from the top left to the bottom
right, by moving left, right, up, and down, is indicated in brackets and is
equal to 2297.

    [131]  673  [234] [103]  [18]
    [201]  [96] [342]  965  [150]
     630   803   746  [422] [111]
     537   699   497  [121]  956
     805   732   524   [37] [331]

Find the minimal path sum, in matrix.txt (right click and "Save Link/Target
As..."), a 31K text file containing a 80 by 80 matrix, from the top left to the
bottom right by moving left, right, up, and down.
"""

# The only differences from problems 81 and 82 are the successor function,
# the goal test, and the initial frontier.

with open("data/matrix.txt", "r") as f:
    # Represent the matrix as a two-dimensional array such that matrix[r][c]
    # retrieves the value at row r and column c.
    matrix = [[int(x) for x in line.split(",")] for line in f.readlines()]

def value(node):
    r, c = node
    return matrix[r][c]

def successors(node):
    r, c = node
    if r > 0: yield (r-1, c) # up
    if c > 0: yield (r, c-1) # left
    if r < len(matrix)-1: yield (r+1, c) # down
    if c < len(matrix)-1: yield (r, c+1) # right

def goal_state(node):
    r, c = node
    return r == c == len(matrix)-1

def uniform_cost_search():
    root = (0, 0)
    frontier = [root]
    explored = set()
    path_sum = {root: value(root)}
    while frontier:
        frontier.sort(key=lambda x: path_sum[x], reverse=True) # Use heapq for bigger tasks.
        node = frontier.pop()
        if goal_state(node):
            # Since we are always popping off the node with the minimal path_sum, we can be
            # sure that if we've arrived at the goal, we've found the shortest path.
            return path_sum[node]
        explored.add(node)
        for child in successors(node):
            new_path_sum = path_sum[node] + value(child)
            if child not in explored and child not in frontier:
                path_sum[child] = new_path_sum
                frontier.append(child)
            elif child in frontier:
                # Before the child may be popped off the frontier, update its
                # value in the path_sum dict, if we found a better path.
                path_sum[child] = min(path_sum[child], new_path_sum)

def problem83():
    return uniform_cost_search()

if __name__ == "__main__":
    print(problem83())
