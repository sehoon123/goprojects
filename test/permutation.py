def permutation(num, answer):
    if len(answer) == len(num):
        print(" ".join(map(str, answer)))
        return
    
    for i in range(1, len(num)+1):
        if i not in answer:
            answer.append(i)
            permutation(num, answer)
            answer.pop()

N = 3
num = list(range(1,6))
# permutation(num, [])

def combination(num, r, start, answer):
    if r == 0:
        print(answer)
        return

    for i in range(start, len(num)):
        if i not in answer:
            answer.append(i)
            combination(num, r-1, i + 1,answer)
            answer.pop()

combination(num, 2, 0, [])
    