import sys
from collections import deque

input = sys.stdin.readline

V = int(input())
graph = [[] for _ in range(V+1)]

for i in range(V):
    tmp = list(map(int, input().split()))
    tmp = tmp[:-1]
    
    for j in range(1, len(tmp), 2):
        graph[tmp[0]].append((tmp[j],tmp[j+1]))

visited = [-1] * (V+1)
visited[1] = 0

def dfs(node, dist):
    for v, d in graph[node]:
        res = dist + d
        if visited[v] == -1:
            visited[v] = res
            dfs(v, res)
    return

dfs(1,0)
tmp = [0, 0]
for i in range(1, V+1):
    if visited[i] > tmp[1]:
        tmp[1] = visited[i]
        tmp[0] = i

visited = [-1] * (V+1)
visited[tmp[0]] = 0
dfs(tmp[0], 0)

print(max(visited))
        

    