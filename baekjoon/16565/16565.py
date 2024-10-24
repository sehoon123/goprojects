import sys

input = sys.stdin.readline

N = int(input())
MOD = 10007

dp = [0] * 53
dp[4] = 13
dp[52] = 1

for i in range(5,53):
    