from collections import deque

def topology_sort(graph):
    n = len(graph)

    indegree = [0] * n

    for nodes in graph:
        for node in nodes:
            indegree[node] += 1

    q = deque()
    for i in range(n):
        if indegree[i] == 0:
            q.append(i)
    
    result = []

    while q:
        node = q.popleft()
        result.append(node)

        for next_node in graph[node]:
            indegree[next_node] -= 1

            if indegree[next_node] == 0:
                q.append(next_node)
    
    return result