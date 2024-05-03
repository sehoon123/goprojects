import sys
from collections import deque
input = sys.stdin.readline

# s = int(input())
s = 5
q = deque()
q.append((1, 0))

dp = [[float('inf')] * 1001 for _ in range(1001)]
dp[1][0] = 0

while q:
    screen, clipboard = q.popleft()
    if dp[screen][screen] == float('inf'): #클립보드 저장
        dp[screen][screen] = dp[screen][clipboard] + 1
        q.append((screen, screen))
    if screen + clipboard <= 1000 and dp[screen+clipboard][clipboard] == float('inf'): # 붙여넣기
        dp[screen+clipboard][clipboard] = dp[screen][clipboard] + 1
        q.append((screen+clipboard, clipboard))
    if screen - 1 >= 0 and dp[screen-1][clipboard] == float('inf'):
        dp[screen-1][clipboard] = dp[screen][clipboard] + 1
        q.append((screen-1, clipboard))


print(min(dp[s]))

INF = sys.maxsize

print(INF)