import sys

input = sys.stdin.readline

N, M = map(int,input().split())

d = [[float('inf')]*N for _ in range(N)]

for _ in range(M):
    a, b = map(int,input().split())
    d[a-1][b-1] = 1

for i in range(N):
    d[i][i] = 0



for k in range(N):
    for i in range(N):
        for j in range(N):
            if d[i][k] + d[k][j] < d[i][j]:
                d[i][j] = d[i][k] + d[k][j]



for i in range(N):
    for j in range(N):
        if d[i][j] != float('inf'):
            d[j][i] = d[i][j]


count = 0
for w in d:
    if sum(w) != float('inf'):
        count +=1


print(count)