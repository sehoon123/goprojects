import sys

sys.setrecursionlimit(100_000_000)

input = sys.stdin.readline


def top_down(n, array, cache):
    if n <= 0:
        return 0
    if cache[n] != 0:
        return cache[n]
    cache[n] = max(top_down(n-1, array, cache),
                   top_down(n-2, array, cache) + array[n-1])
    return cache[n]


n = int(input())
array = list(map(int, input().split()))

cache = [0] * (n+1)
result = top_down(n, array, cache)

print(result)
