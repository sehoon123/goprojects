import sys

input = sys.stdin.readline

C = int(input())
MOD = 1000000007

def tiles(n: int) -> int:
    if n <= 1:
        return 1
    if n == 2:
        return 2
    if cache[n] != 0:
        return cache[n]

    cache[n] = (tiles(n-1) + tiles(n-2)) % MOD
    return cache[n]

def asymmetric(n: int) -> int:
    if n % 2 == 1:
        return (tiles(n) - tiles(n//2) + MOD) % MOD
    ret = tiles(n)
    ret = (ret - tiles(n//2) + MOD) % MOD
    ret = (ret - tiles(n//2-1) + MOD) % MOD
    return ret

for _ in range(C):
    n = int(input())

    cache = [0] * (n+1)
    print(tiles(n))
