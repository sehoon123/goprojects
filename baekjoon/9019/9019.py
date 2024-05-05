import sys
from collections import deque
input = sys.stdin.readline

T = int(input())

for _ in range(T):
    a, b = map(int, input().split())

    visited = [0] * 10_001
    q = deque()
    q.append([a, ''])
    visited[a] = 1

    while q:
        n, c = q.popleft()

        if n == b:
            print(c)
            break
            
        d = n * 2 % 10000
        if not visited[d]:
            visited[d] = 1
            q.append([d, c+'D'])
        
        s = (n - 1) % 10_000
        if not visited[s]:
            visited[s] = 1
            q.append([s, c+'S'])

        
        l = n // 1000 + (n % 1000) * 10
        if not visited[l]:
            visited[l] = 1
            q.append([l, c+'L'])

        r = n // 10 + (n % 10) * 1000
        if not visited[r]:
            visited[r] = 1
            q.append([r, c+'R'])
        


