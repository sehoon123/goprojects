import sys
input = sys.stdin.readline

def find_parent(parent, x):
    if parent[x] != x:
        parent[x] = find_parent(parent, parent[x])
    return parent[x]

def union_parent(parent, x, y):
    x_root = find_parent(parent, x)
    y_root = find_parent(parent, y)

    if x_root < y_root:
        parent[y_root] = x_root
    else:
        parent[x_root] = y_root

def is_same_parent(parent, x, y):
    x_root = find_parent(parent, x)
    y_root = find_parent(parent, y)
    return x_root == y_root

while True:
    m, n = map(int,input().split())
    if m == 0 and n ==0:
        break
    parent = [i for i in range(m)]
    edge = []

    total = 0
    for _ in range(n):
        edge.append(list(map(int,input().split())))

    edge.sort(key=lambda x:x[-1])

    result = 0
    for i in range(n):
        total += edge[i][-1]
        if not is_same_parent(parent, edge[i][0], edge[i][1]):
            result += edge[i][-1]
            union_parent(parent, edge[i][0], edge[i][1])

    print(total-result)