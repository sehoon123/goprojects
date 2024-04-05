import sys
from collections import deque

input = sys.stdin.readline

N = int(input())

indegree = [0] * (N+1)
time = [0] * (N+1)
graph = [[] for _ in range(N+1)]
result = [0] * (N+1)
q = deque()


for i in range(1,N+1):
    tmp = list(map(int,input().split()))

    a, b = tmp[0], tmp[1]
    time[i] = a
    indegree[i] = b

    if b == 0:
        q.append(i)

    for node in tmp[2:]:
        graph[node].append(i)

result = time.copy()

while q:
    now = q.popleft()

    for nxt in graph[now]:
        indegree[nxt] -= 1

        if result[nxt] < result[now] + time[nxt]:
            result[nxt] = result[now] + time[nxt]

        if indegree[nxt] == 0:
            q.append(nxt)


print(max(result))