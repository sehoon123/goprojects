import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())

tree = [[] for _ in range(n+1)]
dist = [-1] * (n+1)
dist[1] = 0


for _ in range(n-1):
    p, c, w = map(int, input().split())
    tree[p].append([c, w])
    tree[c].append([p, w])


def dfs(start, weight):
    for c, w in tree[start]:
        if dist[c] == -1:
            dist[c] = weight + w
            dfs(c, weight + w)


dfs(1, 0)

start = dist.index(max(dist))
dist = [-1] * (n+1)
dist[start] = 0
dfs(start, 0)

print(max(dist))