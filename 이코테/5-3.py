import sys

input = sys.stdin.readline

N, M = map(int, input().split())

dx = [-1,0,1,0]
dy = [0,1,0,-1]

graph = []


for i in range(N):
    graph.append(list(map(int, input().rstrip())))


def dfs(x: int, y: int):
    graph[x][y] = 1

    for i in range(4):
        nx = x + dx[i]
        ny = y + dy[i]

        if nx < 0 or nx >= N or ny < 0 or ny >= M:
            continue

        if graph[nx][ny] == 0:
            dfs(nx, ny)


count = 0
for i in range(N):
    for j in range(M):
        if graph[i][j] == 0:
            dfs(i, j)
            count += 1

print(count)