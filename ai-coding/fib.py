# -*- coding: utf-8 -*-
import sys

def fib(n):
    """Return the nth Fibonacci number (starting from 0)"""
    if n < 0:
        raise ValueError("n must be a non-negative integer")
    if n == 0:
        return 0
    if n == 1:
        return 1
    a, b = 0, 1
    for _ in range(2, n + 1):
        a, b = b, a + b
    return b

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python fib.py <n>")
        sys.exit(1)
    try:
        n = int(sys.argv[1])
        result = fib(n)
        print(result)
    except ValueError as e:
        print("Error: {}".format(e))
        sys.exit(1)
