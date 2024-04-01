import sys
input = sys.stdin.readline

N, M = map(int, input().split())

parent = [i for i in range(N+1)]
edge = []

for _ in range(M):
    edge.append(list(map(int, input().split())))

edge.sort(key=lambda x:x[-1])

def find_parent(parent, x):
    if parent[x] != x:
        parent[x] = find_parent(parent,parent[x])
    return parent[x]

def union_parent(parent, x, y):
    x_root = find_parent(parent,x)
    y_root = find_parent(parent,y)

    if x_root < y_root:
        parent[y_root] = x_root
    else:
        parent[x_root] = y_root

def is_same_parent(parent, x, y):
    x_root = find_parent(parent,x)
    y_root = find_parent(parent,y)
    return x_root == y_root

result = 0
largest = 0
for i in range(M):
    if not is_same_parent(parent, edge[i][0], edge[i][1]):
        result += edge[i][2]
        largest = edge[i][2]
        union_parent(parent,edge[i][0], edge[i][1])

print(result-largest)

