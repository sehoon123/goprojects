import sys
input = sys.stdin.readline

start = 100

N = input().strip()
buttons = [0,1,2,3,4,5,6,7,8,9]

_ = int(input())
tmp = list(map(int, input().split()))
for v in tmp:
    buttons.remove(v)

answer = []
for i in range(len(N)):
    tmp = 100
    digit = 0
    for b in buttons:
        if i >= 1 and int(answer[0]) < int(N[0]):
            candidate = abs(10+int(N[i])-b)
        else:
            candidate = abs(int(N[i])-b)

        if candidate < tmp:
            tmp = candidate
            digit = b
    answer.append(str(digit))

if answer[0] == 0 and len(answer) > 1:
    length = len(answer) - 1
else:
    length = len(answer)

answer = int(''.join(answer))

print(answer)

if abs(start - int(N)) <= abs(answer - int(N)):
    print(abs(start-int(N)))
else:
    print(abs(answer-int(N))+ length) 