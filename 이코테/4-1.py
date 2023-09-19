import sys

input = sys.stdin.readline

N = int(input())
start = (1, 1)
R = (0, 1)
L = (0, -1)
U = (-1, 0)
D = (1, 0)

commands = list(input().split())

for v in commands:
    if v == 'R':
        new = tuple(map(sum, zip(start, R)))
        if new[1] > N or new[1] < 1 or new[0] > N or new[0] < 1:
            continue
        else:
            start = new
    elif v == 'L':
        new = tuple(map(sum, zip(start, L)))
        if new[1] > N or new[1] < 1 or new[0] > N or new[0] < 1:
            continue
        else:
            start = new
    elif v == 'U':
        new = tuple(map(sum, zip(start, U)))
        if new[1] > N or new[1] < 1 or new[0] > N or new[0] < 1:
            continue
        else:
            start = new
    elif v == 'D':
        new = tuple(map(sum, zip(start, D)))
        if new[1] > N or new[1] < 1 or new[0] > N or new[0] < 1:
            continue
        else:
            start = new

print(*start)