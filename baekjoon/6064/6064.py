import sys
import math
input = sys.stdin.readline

T = int(input())

for _ in range(T):
    M, N, x, y = map(int, input().split())
    max_year = math.lcm(M, N)
    answer = x
    while answer <= max_year:
        if (answer - 1) % N + 1 == y:
            break
        answer += M
    if answer > max_year:
        print(-1)
    else:
        print(answer)
        
