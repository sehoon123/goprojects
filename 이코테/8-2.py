import sys

sys.setrecursionlimit(100_000_000)
input = sys.stdin.readline

x = int(input())

dp = [1e9] * (5*(x+1))
dp[1] = 0


for i in range(1, x+1):
    dp[i+1] = min(dp[i+1], dp[i]+1)
    dp[5*i] = min(dp[5*i], dp[i]+1)
    dp[3*i] = min(dp[3*i], dp[i]+1)
    dp[2*i] = min(dp[2*i], dp[i]+1)


print(dp[x])
