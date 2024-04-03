n = int(input())
dp = [0] * 31
dp[2] = 3

def solve(n):
    if n <= 1:
        return 1
    
    if dp[n] != 0:
        return dp[n]
    
    dp[n] = solve(n-2) * 2 + solve(n-1)
    return dp[n]


if n % 2 == 1:
    result = (solve(n) + solve(n // 2)) // 2
else:
    result = (solve(n) + solve(n // 2) + 2 * solve(n//2 -1)) // 2

print(result)
