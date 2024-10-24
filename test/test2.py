from collections import deque
import sys

input = sys.stdin.readline
INF = sys.maxsize

# graph = [
#     [INF, -1, 0, INF],
#     [INF, INF, INF, -1],
#     [INF, -1, INF, -1],
#     [0, -1, INF, INF]
# ]

graph = [
    [INF, -1, 0 ,INF],
    [-1, INF, INF, -1],
    [INF, -1, INF, -1],
    [0, -1, INF, INF]
]

dx = [-1, 0, 0, 1]
dy = [0, -1, 1, 0]

n = len(graph)
m = len(graph[0])

q = deque()

for i in range(n):
    for j in range(m):
        if graph[i][j] == 0:
            q.append((i, j))


while q:
    x, y = q.popleft()

    for i in range(4):
        nx = x + dx[i]
        ny = y + dy[i]
    
        if 0 <= nx < n and 0 <= ny < m and graph[nx][ny] != -1 and graph[nx][ny] >= graph[x][y] + 1:
            q.append((nx, ny))
            graph[nx][ny] = graph[x][y] + 1

for w in graph:
    print(w)

print(graph[0][0] == INF)