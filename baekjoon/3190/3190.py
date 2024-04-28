import sys
from collections import deque
input = sys.stdin.readline

N = int(input())
K = int(input())
apple = []
command = deque()
graph = [[0]*N for _ in range(N)]
shadow = [[0]*N for _ in range(N)]

for _ in range(K):
    a, b = map(int, input().split())
    graph[a-1][b-1] = 1

L = int(input())

for _ in range(L):
    a, b = input().split()
    command.append((int(a), b))


position = [ 0, 0 ]
direction = 0
length = 1

dx = [0,1,0,-1]
dy = [1,0,-1,0]

shadow[0][0] = 1

c, d = command.popleft()


time = 0
while 1:
    time += 1
    nx, ny = position[0] + dx[direction], position[1] + dy[direction]
    position[0], position[1] = nx, ny
    if 0 <= nx < N and 0 <= ny < N and shadow[nx][ny] < time:
        if graph[nx][ny] == 1:
            graph[nx][ny] = 0
            length += 1
            for i in range(N):
                for j in range(N):
                    shadow[i][j] += 1
        
        shadow[nx][ny] = time + length
    else:
        break

    if time == c:
        if d == 'L':
            direction = (direction -1 + 4) % 4
        else:
            direction = (direction +1 + 4) % 4
        
        try:
            c, d = command.popleft()
        except:
            pass
    
print(time)
