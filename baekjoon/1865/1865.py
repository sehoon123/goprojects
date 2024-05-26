import sys
input = sys.stdin.readline

T = int(input())

def bellman_ford() -> bool:
    for i in range(n):
        for node in range(1, n+1):
            for nxt, cost in graph[node]:
                if dist[nxt] > dist[node] + cost:
                    dist[nxt] = dist[node] + cost

                    if i == n-1:
                        return True
    
    return False



for _ in range(T):
    n, m, w = map(int, input().split())
    dist = [int(1e9)] * (n+1)
    graph = [[] for _ in range(n+1)]

    for _ in range(m):
        s, e, t = map(int, input().split())
        graph[s].append((e, t))
        graph[e].append((s, t))
    
    for _ in range(w):
        s, e, t = map(int, input().split())
        graph[s].append((e, -t))
    
    result = bellman_ford()
    if result:
        print("YES")
    else:
        print("NO")