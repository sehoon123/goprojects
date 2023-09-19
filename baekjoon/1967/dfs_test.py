import sys

input = sys.stdin.readline

n = int(input())

tree = [[] for _ in range(n+1)]
visited = [False] * (n+1)

for _ in range(n-1):
    p, c, w = map(int, input().split())
    tree[p].append((c, w))

def dfs(start):
    visited[start] = True
    print(start)
    for node, weight in tree[start]:
        if not visited[node]:
            dfs(node)

dfs(1)