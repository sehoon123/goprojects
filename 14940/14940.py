import sys
from collections import deque
input = sys.stdin.readline

n, m = map(int, input().split())

graph = [[0]*m for _ in range(n)]
visited = [[0]*m for _ in range(n)]

sx, sy = -1, -1

for i in range(n):
    tmp = list(map(int, input().split()))
    for j in range(m):
        graph[i][j] = tmp[j]
        if tmp[j] == 2:
            sx, sy = i, j


dx = [-1, 0, 0, 1]
dy = [0, -1, 1, 0]

def bfs(sx: int, sy: int) -> None:
    q = deque()
    q.append([sx, sy])

    visited[sx][sy] = 0

    while q:
        x, y = q.popleft()

        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]
            if 0 <= nx < n and 0 <= ny < m and graph[nx][ny] == 1:
                if visited[nx][ny] == 0:
                    visited[nx][ny] =  visited[x][y] + 1
                    q.append([nx, ny])

bfs(sx, sy)

for i in range(n):
    for j in range(m):
        if graph[i][j] == 1 and visited[i][j] == 0:
            print(-1, end=" ")
        else:
            print(visited[i][j], end=" ")
    print()




