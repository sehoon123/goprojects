import sys
from collections import deque
from functools import cache
input = sys.stdin.readline

N = int(input())
graph = [list(map(int,input().split())) for _ in range(N)]

dx = [0, 1, 1]
dy = [1, 0, 1]


# 0: 가로, 1: 세로, 2: 대각선
@cache
def dfs(x, y, state):
    if x == N-1 and y == N-1:
        return 1

    count = 0
    for i in range(3):
        if state == 0 and i == 1:
            continue
        elif state == 1 and i == 0:
            continue

        nx, ny = x + dx[i], y + dy[i]

        if 0 <= nx < N and 0 <= ny < N and graph[nx][ny] != 1:
            if i == 2:
                if graph[nx-1][ny] != 1 and graph[nx][ny-1] != 1:
                    count += dfs(nx, ny, i)
            else:
                count += dfs(nx, ny, i)

    return count
        
ans = dfs(0,1,0)

print(ans)



    
