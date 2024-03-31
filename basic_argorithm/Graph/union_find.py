def find_parent(parent, x):
    if parent[x] != x:
        parent[x] = find_parent(parent,parent[x])
    
    return parent[x]

def union(parent,x, y):
    x_root = find_parent(parent, x)
    y_root = find_parent(parent, y)

    if x_root < y_root:
        parent[y_root] = x_root
    else:
        parent[x_root] = y_root
    

n = 5 # 예시로 5개의 요소를 사용합니다.
parent = list(range(n))

# 집합 합치기
union(parent, 0, 1) # 0과 1을 합칩니다.
union(parent, 1, 2) # 1과 2를 합칩니다. (이미 0과 1이 합쳐져 있으므로, 0과 2가 합쳐집니다.)
union(parent, 3, 4) # 3과 4를 합칩니다.

# 결과 확인
print("Parent of 0:", find_parent(parent, 0)) # 0의 부모는 0입니다.
print("Parent of 1:", find_parent(parent, 1)) # 1의 부모는 0입니다.
print("Parent of 2:", find_parent(parent, 2)) # 2의 부모는 0입니다.
print("Parent of 3:", find_parent(parent, 3)) # 3의 부모는 3입니다.
print("Parent of 4:", find_parent(parent, 4)) # 4의 부모는 3입니다.