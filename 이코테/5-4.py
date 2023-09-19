import sys
from collections import deque

input = sys.stdin.readline

N, M = map(int, input().split())

graph = []
visited = [[0]*M for _ in range(N)]

for _ in range(N):
    graph.append(list(map(int, input().rstrip())))

dx = [-1,0,1,0]
dy = [0,1,0,-1]

def bfs(x: int, y: int):
    queue = deque()
    queue.append((x, y))
    visited[x][y] = 1

    while queue:
        x, y = queue.popleft()

        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]

            if nx < 0 or nx >= N or ny < 0 or ny >= M:
                continue

            if graph[nx][ny] == 0:
                continue

            if visited[nx][ny] == 0:
                queue.append((nx, ny))
                visited[nx][ny] = visited[x][y] + 1

bfs(0, 0)

print(visited)