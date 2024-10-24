import sys

input = sys.stdin.readline

printed = []

N = int(input())
depth = 0
graph = [[] for _ in range(N)]
for j in range(N):
    tmp = list(input().split())
    if depth > int(tmp[0]):
        depth = int(tmp[0])
    for i, v in enumerate(tmp[1:]):
        graph[j].append((i, v))

graph.sort()

for v in graph:
    for w in v:
        if w not in printed:
            printed.append(w)
            a, b = w
            printed = [x for x in printed if x[0] <= a]
            print("--"*a+b)