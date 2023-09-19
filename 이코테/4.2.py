import sys

input = sys.stdin.readline

dx = [2, 2, -2, -2, 1, 1, -1, -1]
dy = [1, -1, 1, -1, 2, -2, 2, -2]

loc = input().rstrip()

x = ord(loc[0]) - ord('a') + 1
y = int(loc[1])

count = 0
for i in range(8):
    nx = x + dx[i]
    ny = y + dy[i]

    if nx < 1 or nx > 8 or ny < 1 or ny > 8:
        continue
    count += 1

print(count)