import sys
from collections import deque

input = sys.stdin.readline

T = int(input())

for _ in range(T):
    N, K = map(int, input().split())
    time = [0] + list(map(int, input().split()))
    graph = [[] for _ in range(N + 1)]
    degree = [0] * (N + 1)
    dp = [0] * (N + 1)

    for _ in range(K):
        u, v = map(int, input().split())
        graph[u].append(v)
        degree[v] += 1


    q = deque()
    for i in range(1, N + 1):
        if degree[i] == 0:
            q.append(i)
            dp[i] = time[i]
    
    result = []
    
    while q:
        now = q.popleft()
        result.append(now)
        for i in graph[now]:
            degree[i] -= 1
            dp[i] = max(dp[i], dp[now] + time[i])

            if degree[i] == 0:
                q.append(i)
    
    print(dp[int(input())])