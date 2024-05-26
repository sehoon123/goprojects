import sys
import heapq
input = sys.stdin.readline

def print_log(end: int):
    tmp = []
    tmp.append(end)

    while 1:
        now = tmp[-1]
        if before[now] == 0:
            break
        tmp.append(before[now])
    
    print(len(tmp))
    print(*tmp[::-1])


def dijkstra(start: int):
    dist[start] = 0
    pq = []
    heapq.heappush(pq, (0, start))

    while pq:
        cost, now = heapq.heappop(pq)

        if dist[now] < cost:
            continue
        
        for next, next_cost in graph[now]:
            if dist[next] > cost + next_cost:
                dist[next] = cost + next_cost
                before[next] = now
                heapq.heappush(pq, (dist[next], next))


n = int(input())
m = int(input())

graph = [[] for _ in range(n+1)]
dist = [float('inf')] * (n+1)
before = [0] * (n+1)

for _ in range(m):
    s, e, c = map(int, input().split())
    graph[s].append((e, c))

start, end = map(int, input().split())

dijkstra(start)
print(dist[end])
print_log(end)    