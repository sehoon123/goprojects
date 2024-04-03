import math
import sys
import heapq
from collections import deque
input = sys.stdin.readline

M, N = map(int, input().split())

dx = [-1,0,0,1]
dy = [0,1,-1,0]
WALL = 100_000
NO_WALL = 1
graph = []
for _ in range(N):
    graph.append(input().strip())

visit = [[float('inf')]*M for _ in range(N)]


def dijkstara(graph, visit, sx, sy):
    pq = []
    heapq.heappush(pq,[0, sx, sy])

    while pq:
        cost, x, y = heapq.heappop(pq)

        if cost > visit[x][y]:
            continue

        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]

            if 0 <= nx < N and 0 <= ny < M:
                if graph[nx][ny] == '1':
                    distance = cost + WALL
                else:
                    distance = cost + NO_WALL

                if visit[nx][ny] > distance:
                    visit[nx][ny] = distance
                    heapq.heappush(pq, [distance, nx, ny])
    
    return visit[N-1][M-1]

answer = dijkstara(graph,visit,0,0)//WALL
if math.isnan(answer):
    print(0)
else:
    print(answer)
