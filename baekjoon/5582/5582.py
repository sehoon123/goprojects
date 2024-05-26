import sys
input = sys.stdin.readline

s = input().strip()
t = input().strip()
answer = 0

dp = [[0]*(len(s)+1) for _ in range((len(t)+1))]

for i in range(1, len(t)+1):
    for j in range(1, len(s) + 1):
        if t[i-1] == s[j-1]:
            dp[i][j] = dp[i-1][j-1] + 1
            answer = max(answer, dp[i][j])

print(answer)