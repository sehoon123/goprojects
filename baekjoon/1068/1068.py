import sys
import functools
input = sys.stdin.readline

@functools.lru_cache(maxsize=None)
def dfs(start: int) -> int:
    tmp = 0
    visited[start] = 1

    if graph[start] == []:
        return 1

    for next_node in graph[start]:
        if visited[next_node] == 0:
            tmp += dfs(next_node)
    
    return tmp

N = int(input())
parents = list(map(int,input().split()))
target = int(input())
graph = [[] for _ in range(N)]
parents[target] = -4
visited= [0]*N

for i in range(len(parents)):
    if parents[i] >= 0:
        graph[parents[i]].append(i)


result = 0
for i in range(len(parents)):
    if parents[i] == -1:
        result += dfs(i)

print(result)
