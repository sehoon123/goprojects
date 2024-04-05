import sys
import heapq
input = sys.stdin.readline

dx = [-1,0,0,1]
dy = [0,1,-1,0]

def dijkstra(graph, visited, sx, sy):
    pq = []
    heapq.heappush(pq, (graph[sx][sy], sx, sy))

    while pq:
        cost, x, y = heapq.heappop(pq)

        if cost > visited[x][y]:
            continue

        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]

            if 0 <= nx < n and 0 <= ny < n:
                next_cost = cost + graph[nx][ny]

                if next_cost < visited[nx][ny]:
                    visited[nx][ny] = next_cost
                    heapq.heappush(pq,(next_cost,nx,ny))
        
    
    return visited

i = 1
while 1:
    n = int(input())
    if n == 0:
        break
    
    graph = []
    for _ in range(n):
        graph.append(list(map(int,input().split())))
    
    visited = [[float('inf')]* n for _ in range(n)]


    answer = dijkstra(graph, visited, 0, 0)
    print(answer)
    print(f"Problem {i}: {answer[n-1][n-1]}")
    i += 1
    