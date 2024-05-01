# import sys
# input = sys.stdin.readline

# T = int(input())
# MOD = 1_000_000_009

# dp = [[0] * 3 for _ in range(100_001)]
# dp[1] = [1,0,0]
# dp[2] = [0,1,0]
# dp[3] = [1,1,1]

# for i in range(4, 100_001):
#     dp[i][0] = (dp[i-1][1] + dp[i-1][2])%MOD
#     dp[i][1] = (dp[i-2][0] + dp[i-2][2])%MOD
#     dp[i][2] = (dp[i-3][0] + dp[i-3][1])%MOD

# for _ in range(T):
#     n = int(input())
#     print(sum(dp[n])%MOD)

import sys
input = sys.stdin.readline

T = int(input())
MOD = 1_000_000_009

def solve(n):
    if n == 1:
        return 1
    elif n == 2:
        return 2
    elif n == 3:
        return 4
    else:
        dp = [0] * (n+1)
        dp[1] = 1
        dp[2] = 2
        dp[3] = 4
        for i in range(4, n+1):
            dp[i] = (dp[i-1] + dp[i-2] + dp[i-3] - dp[i-1]) % MOD
        return dp[n]

for _ in range(T):
    n = int(input())
    print(solve(n))
