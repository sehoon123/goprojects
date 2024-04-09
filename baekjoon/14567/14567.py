import sys
from collections import deque
input = sys.stdin.readline

N, M = map(int, input().split())
indegree = [0] * (N+1)
graph = [[] for _ in range(N+1)]
result = [1] * (N+1)

for _ in range(M):
    a, b = map(int, input().split())
    indegree[b] += 1
    graph[a].append(b)


q = deque()

for i in range(1,N+1):
    if indegree[i] == 0:
        q.append(i)

while q:
    now = q.popleft()
    for node in graph[now]:
        indegree[node] -= 1
        result[node] = result[now]+1

        if indegree[node] == 0:
            q.append(node)

print(*result[1:])