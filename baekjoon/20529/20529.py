import sys
from itertools import combinations
input = sys.stdin.readline

def combination(arr, n):
    if n == 0:
        yield []
    else:
        for i in range(len(arr)):
            for ele in combination(arr[i+1:], n-1):
                yield [arr[i]] + ele

for _ in range(int(input())):
    n = int(input())
    arr = list(input().split())

    if n > 32:
        print(0)
    else:
        comb = combinations(arr, 3)
        comb = list(comb)
        # print(comb)

        answer = float('inf')

        for v in comb:
            temp = 0
            three = combinations(v, 2)
            three = list(three)
            for a, b in three:
                for i in range(4):
                    if a[i] != b[i]:
                        temp += 1
        
            if temp <= answer:
                answer = temp

        print(answer)