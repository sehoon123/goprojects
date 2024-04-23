import sys

input = sys.stdin.readline

for _ in range(int(input())):
    command = input().rstrip()
    n = int(input())
    arr = input().strip()
    arr = eval(arr)

    left = 0
    right = 0
    reverse = False

    for c in command:
        if c == 'R':
            reverse = not reverse
        elif c == 'D':
            if reverse:
                right += 1
            else:
                left += 1

    if left + right > n:
        print("error")
    else:
        if reverse:
            arr = arr[left:n-right][::-1]
        else:
            arr = arr[left:n-right]
        print(str(arr).replace(' ', ''))
