from collections import deque

def bfs(graph: list, v: int, visited: list):
    queue = deque([v])
    visited[v] = True

    while queue:
        node = queue.popleft()
        print(node, end=' ')

        for v in graph[node]:
            if not visited[v]:
                queue.append(v)
                visited[v] = True


graph = [
    [],
    [2, 3, 8],
    [1, 7],
    [1, 4, 5],
    [3, 5],
    [3, 4],
    [7],
    [2, 6, 8],
    [1, 7]
]

visited = [False] * 9

bfs(graph, 1, visited)
