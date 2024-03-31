def dfs(graph, visited, start):
    visited[start] = 1
    print(start, end=" ")

    for node in graph[start]:
        if visited[node] == 0:
            dfs(graph, visited, node)

graph=[[],[2,3],[4,5],[6,7],[],[],[],[]]
visited = [0] * 8

dfs(graph, visited, 1)