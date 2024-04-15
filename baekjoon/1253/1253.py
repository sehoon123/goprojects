import sys
import itertools
input = sys.stdin.readline

N = int(input())
arr = list(map(int, input().split()))
arr.sort()

count = 0

for i in range(N):
    target = arr[i]
    left, right = 0, N-1

    while left < right:
        if left == i:
            left += 1
            continue

        if right == i:
            right -= 1
            continue

        tmp = arr[left] + arr[right]

        if tmp == target:
            count += 1
            break

        if tmp < target:
            left += 1
        
        else:
            right -= 1

print(count)