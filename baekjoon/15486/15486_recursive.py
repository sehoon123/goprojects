import sys
sys.setrecursionlimit(100_000_000)
input = sys.stdin.readline
from functools import lru_cache

@lru_cache(maxsize=None)
def max_profit(day):
    if day > n:
        return 0
    
    # 상담을 하지 않는 경우
    not_take = max_profit(day + 1)
    
    # 상담을 하는 경우
    take = 0
    if day + t[day] - 1 <= n:
        take = p[day] + max_profit(day + t[day])
    
    return max(take, not_take)

n = int(input())
t = [0]
p = [0]

for _ in range(n):
    a, b = map(int, input().split())
    t.append(a)
    p.append(b)

print(max_profit(1))
