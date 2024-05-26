import sys
sys.setrecursionlimit(1_000_000_000)
input = sys.stdin.readline

def dfs(node):
    visited[node] = True
    for next_node in graph[node]:
        if visited[next_node]:
            continue
        dfs(next_node)
        dp[node][0] += dp[next_node][1]
        dp[node][1] += min(dp[next_node])

n = int(input())
graph = [[] for _ in range(n+1)]
visited = [0]*(n+1)

dp= [[0, 1] for _ in range(n+1)]

for i in range(n-1):
    a, b = map(int, input().split())
    graph[a].append(b)
    graph[b].append(a)

dfs(1)
print(min(dp[1]))


