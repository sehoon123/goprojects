import sys

input = sys.stdin.readline

Days, Top = map(int, input().split())

cache = [[0] * (2*Days+1) for _ in range(Days+1)]
cache[1][1] = 1
cache[1][2] = 3

for i in range(2, Days+1):
    for j in range(1, 2*i+1):
        if (cache[i-1][j-2] != 0 and j-2 > 0):
            cache[i][j] += cache[i-1][j-2]
        if (cache[i-1][j-1] != 0 and j-1 > 0):
            cache[i][j] += cache[i-1][j-1]

print(sum(cache[Days][Top:]) / sum(cache[Days]))


for v in cache:
    print(v)