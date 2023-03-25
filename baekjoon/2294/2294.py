import sys
input = sys.stdin.readline

n, k = map(int, input().split())

dp = [float('inf')] * (k + 1)
dp[0] = 0  # base case

coins = set()

for _ in range(n):
    coins.add(int(input()))

def solve(n):
    if dp[n] != float('inf'):
        return dp[n]
    for coin in coins:
        if n - coin >= 0:
            dp[n] = min(dp[n], solve(n - coin) + 1)
    return dp[n]

ans = solve(k)

print(ans if ans != float('inf') else -1)
