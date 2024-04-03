import sys
input = sys.stdin.readline

dp = [0] * 251

def solve(x):
    if x <= 1:
        return 1
    
    if dp[x] != 0:
        return dp[x]

    dp[x] = 2*solve(x-2) + solve(x-1)

    return dp[x]

while True:
    try:
        n = int(input())
        print(solve(n))
    except:
        break

