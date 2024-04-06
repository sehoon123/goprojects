import sys
from collections import deque
input = sys.stdin.readline

T = int(input())

for _ in range(T):
    n = int(input())
    indegree = [0] * (n+1)
    graph = [[] for _ in range(n+1)]
    tmp = list(map(int,input().split()))

    for i in range(n-1):
        graph[tmp[i]].append(tmp[i+1])
        indegree[tmp[i+1]] += 1

    m = int(input())
    for _ in range(m):
        a, b = map(int, input().split())
    
    print(graph)
    print(indegree)