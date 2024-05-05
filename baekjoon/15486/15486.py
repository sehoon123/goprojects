import sys
input = sys.stdin.readline

n = int(input())

t = [0]
p = [0]
dp = [0] * (n+1)

for _ in range(n):
    a, b = map(int, input().split())
    t.append(a)
    p.append(b)

for i in range(1,n+1):
    dp[i] = max(dp[i], dp[i-1])
    finish = i + t[i] - 1
    if finish <= n:
        dp[finish] = max(dp[finish], dp[i-1] + p[i])

print(dp[n])


