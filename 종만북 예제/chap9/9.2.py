import sys

input = sys.stdin.readline

C = int(input())

def solve(cap: int, item: int) -> int:
    if item == N:
        return 0

    ret = cache[cap][item]
    if ret != -1:
        return ret

    ret = solve(cap, item+1)
    if cap >= stuff[item][1]:
        ret = max(ret, solve(cap-stuff[item][1], item+1) + stuff[item][2])

    cache[cap][item] = ret
    return ret


for _ in range(C):
    N, W = map(int, input().split())
    cache = [[-1]*1001 for _ in range(100)]

    stuff = []

    for _ in range(N):
        a, b, c = input().split()
        stuff.append((a, int(b), int(c)))

    print(solve(W, 0))