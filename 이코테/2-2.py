import sys

input = sys.stdin.readline

N, M, K = map(int, input().split())

arr = list(map(int, input().split()))
arr.sort(reverse=True)

sum = 0
count = 0

for _ in range(M):
    if count == K:
        sum += arr[1]
        count = 0
    else:
        count += 1
        sum += arr[0]

    print(count, sum)


print(sum)