from collections import deque

def bfs(graph, visited, start):
    q = deque()
    q.append(start)

    while q:
        now = q.popleft()
        print(now, end=" ")
        visited[now] = 1
        for node in graph[now]:
            if visited[node] == 0:
                q.append(node)


graph=[[],[2,3],[4,5],[6,7],[],[],[],[]]
visited = [0] * 8

bfs(graph, visited, 1)