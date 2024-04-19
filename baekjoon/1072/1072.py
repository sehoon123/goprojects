import sys

input = sys.stdin.readline

X, Y = map(int, input().split())

MAX = 1_000_000_000

current = int(Y*100/X)

left, right = 0, MAX
while left < right:
    mid = (left+right)//2
    changed = int((Y+mid)*100/(X+mid))
    
    if changed < current:
        left = mid + 1
    else:
        right = mid

    
if right != MAX:
    print(left)
else:
    print(-1)

