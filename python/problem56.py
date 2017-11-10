"""
problem56.py

https://projecteuler.net/problem=56

Considering natural numbers of the form, a**b, where a, b < 100,
what is the maximum digital sum?
"""

def digit_sum(n):
    return sum(map(int, (str(n))))

def problem56():
    return max(digit_sum(a**b) for a in range(100) for b in range(100))

if __name__ == "__main__":
    print(problem56())
