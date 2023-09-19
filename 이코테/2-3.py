import sys

input = sys.stdin.readline

N, M = map(int, input().split())

candidate = []
for _ in range(N):
     candidate.append(min(list(map(int, input().split()))))

print(max(candidate))