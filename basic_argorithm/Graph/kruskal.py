import sys
input = sys.stdin.readline

def get_parent(parent, x):
    if parent[x] != x:
        parent[x] = get_parent(parent,parent[x])
    
    return parent[x]

def union(parent, x, y):
    x_root = get_parent(parent, x)
    y_root = get_parent(parent, y)

    if x_root < y_root:
        parent[y_root] = x_root
    else:
        parent[x_root] = y_root
    
def same_parent(parent, x, y):
    x_root = get_parent(parent, x)
    y_root = get_parent(parent, y)

    return x_root == y_root


parent = [i for i in range(8)]
edge = []

edge.append((1,7,12))
edge.append((1,4,28))
edge.append((1,2,67))
edge.append((1,5,17))
edge.append((2,4,24))
edge.append((2,5,62))
edge.append((3,5,20))
edge.append((3,6,37))
edge.append((4,7,13))
edge.append((5,6,45))
edge.append((5,7,73))

edge.sort(key=lambda x : x[-1])
print(edge)

print(parent)

total = 0
for i in range(len(edge)):
    if not same_parent(parent, edge[i][0], edge[i][1]):
        total += edge[i][-1]
        union(parent, edge[i][0], edge[i][1])

print(total)
