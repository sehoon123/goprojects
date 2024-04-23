n = int(input())

dp = [0] * 16
dp[0] = 2

tmp = 1
for i in range(1, 16):
    dp[i] = dp[i-1] + pow(2, i-1)

print(dp[n]**2)

