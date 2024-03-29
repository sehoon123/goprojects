import sys

N, M = list(map(int,input().split()))
board = [[0]*M for _ in range(N)]

for i in range(N):
    tmp = input()
    for j in range(M):
        board[i][j] = tmp[j]

