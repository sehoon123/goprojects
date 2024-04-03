import sys
import heapq

input = sys.stdin.readline

N = int(input())
M = int(input())

graph = [[] for _ in range(N+1)]
visit = [float('inf')] * (N+1)

for _ in range(M):
    start, end, cost = map(int, input().split())
    graph[start].append([cost, end])

start, end = map(int, input().split())

def dijkstra(graph, start):
    pq = []
    heapq.heappush(pq, [0, start])

    while pq:
        current_distance, current_node = heapq.heappop(pq)
        
        if current_distance > visit[current_node]:
            continue

        for cost, dest in graph[current_node]:
            distance = current_distance + cost

            if distance < visit[dest]:
                visit[dest] = distance

                heapq.heappush(pq, [distance, dest])
    
    return visit


dijkstra(graph, start)
print(visit[end])

