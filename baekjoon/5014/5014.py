import sys
from collections import deque

input = sys.stdin.readline

F, S, G, U, D = map(int,input().split())

graph = [-1] * (F+1)
graph[S] = 0

dx = [U, -1*D]


q = deque()

q.append(S)

while q:
    now = q.popleft()

    if now == G:
        print(graph[now])
        break

    for i in range(2):
        if 0 < now + dx[i] and now + dx[i] <= F and graph[now+dx[i]] == -1:
            graph[now+dx[i]] = graph[now] + 1
            q.append(now+dx[i])


if graph[G] == -1:
    print("use the stairs")
        


