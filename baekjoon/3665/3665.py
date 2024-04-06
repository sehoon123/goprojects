import sys
input = sys.stdin.readline

def read_input():
    n = int(input())
    graph = [[] for _ in range(n + 1)]
    in_degree = [0 for _ in range(n + 1)]
    team = list(map(int, input().split()))

    for i in range(n - 1):
        for j in range(i + 1, n):
            graph[team[i]].append(team[j])
            in_degree[team[j]] += 1

    m = int(input())
    for _ in range(m):
        first, second = map(int, input().split())
        if second in graph[first]:
            graph[first].remove(second)
            in_degree[second] -= 1
            graph[second].append(first)
            in_degree[first] += 1
        else:
            graph[second].remove(first)
            in_degree[first] -= 1
            graph[first].append(second)
            in_degree[second] += 1

    return n, graph, in_degree

def topological_sort(n, graph, in_degree):
    queue = []
    answer = []

    for i in range(1, n + 1):
        if in_degree[i] == 0:
            queue.append(i)

    if not queue:
        return "IMPOSSIBLE"

    while queue:
        if len(queue) > 1:
            return "IMPOSSIBLE"

        node = queue.pop()
        answer.append(node)
        for neighbor in graph[node]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)
            elif in_degree[neighbor] < 0:
                return "IMPOSSIBLE"

    if len(answer) < n:
        return "IMPOSSIBLE"
    else:
        return " ".join(map(str, answer))

t = int(input())

for _ in range(t):
    n, graph, in_degree = read_input()
    result = topological_sort(n, graph, in_degree)
    print(result)
