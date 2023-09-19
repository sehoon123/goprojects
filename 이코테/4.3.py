import sys

input = sys.stdin.readline

dx = [-1,0,1,0]
dy = [0,1,0,-1]

N, M = map(int, input().split())

x, y, direction = map(int, input().split())

graph = []
for _ in range(N):
    graph.append(list(map(int, input().split())))

count = 1
graph[x][y] = 2
flag = 0

while 1:
    direction = (direction + 3) % 4
    nx = x + dx[direction]
    ny = y + dy[direction]

    for v in graph:
        print(v)
    print(direction)
    print()

    if nx < 0 or nx >= N or ny < 0 or ny >= M:
        flag += 1
        continue

    if graph[nx][ny] == 0:
        graph[nx][ny] = 2
        x = nx
        y = ny
        count += 1
        flag = 0
        continue

    flag += 1

    if flag == 4:
        nx = x - dx[direction]
        ny = y - dy[direction]

        if graph[nx][ny] == 1:
            break
        else:
            x = nx
            y = ny
            flag = 0


for v in graph:
    print(v)
print()

print(count)