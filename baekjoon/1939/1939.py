import sys
from collections import deque
input = sys.stdin.readline

N, M = map(int, input().split())
graph = [[] for _ in range(N+1)]
for _ in range(M):
    a, b, c = map(int, input().split())
    graph[a].append((b,c))
    graph[b].append((a,c))
start, end = map(int, input().split())


def binary_search() -> int:
    left, right = 0, 1_000_000_000
    result = 0

    while left <= right:
        mid = (left + right) // 2

        if bfs(mid):
            result = mid
            left = mid + 1
        else:
            right = mid -1
    
    return result


def bfs(weight: int) -> bool:
    q = deque()
    visited = [0] * (N+1)

    q.append(start)
    visited[start] = 1

    while q:
        node = q.popleft()

        for nxt, w in graph[node]:
            if visited[nxt] == 0 and weight <= w:
                q.append(nxt)
                visited[nxt] = 1
    
    return visited[end]

print(binary_search())


