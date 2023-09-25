import sys

sys.setrecursionlimit(100_000_000)
input = sys.stdin.readline

n = int(input())

cache = [0] * 1001

cache[1] = 1
cache[2] = 3


def solve(x: int) -> int:
    if x <= 0:
        return 0

    if cache[x] != 0:
        return cache[x]

    cache[x] = solve(x-1) + solve(x-2) * 2 

    return cache[x]


print(solve(n))

