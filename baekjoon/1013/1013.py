import re
import sys

input = sys.stdin.readline
T = int(input())

for _ in range(T):
    sign = sys.stdin.readline().replace('\n', '')
    pattern ='(100+1+|01)+'
    m = re.fullmatch(pattern, sign)
    if m:
        print("YES")
    else:
        print("NO")
