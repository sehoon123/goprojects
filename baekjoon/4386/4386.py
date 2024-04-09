import sys
import math
import itertools
input = sys.stdin.readline

def calc_dist(ax, ay, bx, by) -> float:
    return math.sqrt((ax-bx)**2 + (ay-by)**2)

def find_parent(parents, x):
    if parents[x] != x:
        parents[x] = find_parent(parents, parents[x])
    return parents[x]

def union_parent(parents, x, y):
    x_root = find_parent(parents, x)
    y_root = find_parent(parents, y)

    if x_root < y_root:
        parents[y_root] = x_root
    else:
        parents[x_root] = y_root

def same_parent(parents, x, y):
    x_root = find_parent(parents, x)
    y_root = find_parent(parents, y)
    return x_root == y_root


n = int(input())
parents = [0] * n
graph = [[] for _ in range(n)]
dists = []
q = []

for i in range(n):
    parents[i] = i
    dists.append([i] + list(map(float, input().split())))


nCr = itertools.combinations(dists, 2)
nCr_list = list(nCr)

for i in range(len(nCr_list)):
    edge = nCr_list[i]
    q.append((edge[0][0], edge[1][0], calc_dist(edge[0][1], edge[0][2], edge[1][1], edge[1][2])))

q.sort(key=lambda x: x[2])


result = 0
for i in range(len(q)):
    if not same_parent(parents, q[i][0], q[i][1]):
        union_parent(parents, q[i][0], q[i][1])
        result += q[i][2]

print(result)