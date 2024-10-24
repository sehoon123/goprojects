from collections import deque

# Graph representation
graph = [
    [],     # 0 (not used)
    [2],    # 1
    [1, 3, 4],  # 2
    [2],    # 3
    [2, 5, 6],  # 4
    [4],    # 5
    [4, 7],  # 6
    [6]     # 7
]

def custom_traversal(graph, start):
    visited = [False] * len(graph)
    queue = deque([start])
    order = []

    while queue:
        node = queue.popleft()
        if not visited[node]:
            visited[node] = True
            order.append(node)
            
            # Add neighbors in ascending order
            for neighbor in sorted(graph[node]):
                if not visited[neighbor]:
                    queue.append(neighbor)

    return order

# Perform the traversal
# traversal_order = custom_traversal(graph, 1)
# print("Traversal order:", traversal_order)

def check(num, start, end):
    if start >= end:
        return True  # 재귀 종료 조건: 범위가 올바르지 않을 때는 True 반환
    
    
    mid = (start + end) // 2
    
    # 가운데 값이 1이 아니면 False 반환
    if num[mid] != '1':
        return False
    
    # 양쪽 하위 배열에 대해서도 같은 작업을 재귀적으로 수행
    left_check = check(num, start, mid - 1)
    right_check = check(num, mid + 1, end)
    
    # 양쪽 모두 True일 때만 전체 결과가 True
    return left_check and right_check

print(check("0001111", 0, 6))