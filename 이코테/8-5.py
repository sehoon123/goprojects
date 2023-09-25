import sys

input = sys.stdin.readline

n, m = map(int, input().split())

coins = [0] * n
cache = [1e9] * (m + 1)

for i in range(n):
    coins[i] = int(input())

coins.sort()

for i in coins:
    cache[i] = 1

for i in range(1, m+1):
    for coin in coins:
        if cache[i] != 1e9 and (i+coin) < m+1:
            cache[i + coin] = min(cache[i] + 1, cache[i+coin])


if cache[m] != 1e9:
    print(cache[m])
else:
    print(-1)
