import sys
input = sys.stdin.readline

for _ in range(int(input())):
    N = int(input())
    coin = tuple(map(int, input().split()))
    target = int(input())
    dp = [0] * (target+1)
    dp[0] = 1

    for c in coin:
        for i in range(1, target+1):
            if i >= c:
                dp[i] += dp[i - c]
    
    print(dp)

