import sys

input = sys.stdin.readline

N = int(input())

basic = 45*15 + 15*60

result = 0
for i in range(N+1):
    if i == 3 or i == 13 or i == 23:
        result += 60*60
    else:
        result += basic

print(result)
