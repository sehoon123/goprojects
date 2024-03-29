import sys
from collections import deque

N = int(input())
time = [0]
indegree = [0] * (N+1)
dp = [0] * (N+1)
graph = [[] for _ in range(N+1)]
q = deque()

for i in range(N):
    tmp = list(map(int, input().split()))

    time.append(tmp[0])

    for node in tmp[1:]:
        if node != -1:
            graph[node].append(i+1)
            indegree[i+1] += 1

for i in range(1, N+1):
    dp[i] = time[i]
    if indegree[i] == 0:
        q.append(i)

while q:
    now = q.popleft()

    for node in graph[now]:
        indegree[node] -= 1
        dp[node] = max(dp[node], dp[now] + time[node])

        if indegree[node] == 0:
            q.append(node)
    

for res in dp[1:]:
    print(res)

