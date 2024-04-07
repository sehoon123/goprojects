import sys
import heapq
input = sys.stdin.readline

N, M, K, X = map(int, input().split())

dist = [float('inf')] * (N + 1)
graph = [[] for _ in range(N+1)]

for _ in range(M):
    s, e = map(int, input().split())
    graph[s].append([1, e])

def dijkstra(graph, dist, start):
    pq = []
    heapq.heappush(pq, [0, start])

    while pq:
        current_cost, current_node = heapq.heappop(pq)

        if current_cost > dist[current_node]:
            continue
    
        for cost, next_node in graph[current_node]:
            distance = current_cost + cost

            if distance < dist[next_node]:
                dist[next_node] = distance
                heapq.heappush(pq, [distance, next_node])
    
    return dist

dist[X] = 0
dijkstra(graph, dist, X)

flag = False
for i in range(len(dist)):
    if dist[i] == K:
        print(i)
        flag = True

if not flag:
    print(-1)