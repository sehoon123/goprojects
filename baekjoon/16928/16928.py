import sys
from collections import deque
input = sys.stdin.readline

N, M = map(int, input().split())

ladder = []
snake = []
for _ in range(N):
    a, b = map(int, input().split())
    ladder.append((a, b))

for _ in range(M):
    a, b = map(int, input().split())
    snake.append((a, b))

q = deque()
visited = [0] * 101

# 위치, 주사위 횟수
q.append((1, 0))
visited[1] = 1

while q:
    location, count = q.popleft()

    if location == 100:
        print(count)
        break

    for i in range(location+1, location+7):
        if i > 100:
            break

        if visited[i] == 0:
            if any(k == i for k, _ in ladder):
                for k, v in ladder:
                    if i == k:
                        visited[v] = 1
                        q.append((v, count + 1))
            elif any(k == i for k, _ in snake):
                for k, v in snake:
                    if i == k:
                        visited[v] = 1
                        q.append((v, count + 1))
            else:
                visited[i] = 1
                q.append((i, count + 1))
