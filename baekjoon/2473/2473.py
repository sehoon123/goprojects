import sys
input = sys.stdin.readline

N = int(input())
arr = list(map(int, input().split()))
arr.sort()

answer = (float('inf'))
ans_arr = []
exit_loop = False

for i in range(N):
    left = i+1
    right = N-1

    while left < right:
        left_val = arr[left]
        right_val = arr[right]
        total = arr[i] + left_val + right_val

        if abs(total) < answer:
            answer = abs(total)
            ans_arr = [arr[i], left_val, right_val]
            if answer == 0:
                exit_loop = True
                break
        
        if total < 0:
            left += 1
        else:
            right -= 1
    if exit_loop:
        break

print(*ans_arr)


