import sys
input = sys.stdin.readline

n = int(input())
answer = list(input().strip())

for _ in range(1,n):
    temp = input().strip()
    for i in range(len(answer)):
        if answer[i] != temp[i]:
            answer[i] = '?'

print(''.join(answer))

