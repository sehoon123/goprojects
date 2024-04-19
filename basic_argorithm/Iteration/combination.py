
def combination(arr:list, x: int):
    if x == 0:
        yield []
    elif x > 0:
        for i in range(len(arr)):
            for ele in combination(arr[i+1:], x-1):
                yield [arr[i]] + ele
    
arr = [1,2,3,4,5]

print(*combination(arr,0))
print(*combination(arr,1))
        

def permutation(arr:list, x: int):
    if x == 0:
        yield []
    elif x > 0:
        for i in range(len(arr)):
            for ele in permutation(arr[:i]+arr[i+1:], x-1):
                yield [arr[i]] + ele

print(*permutation(arr,0))
print(*permutation(arr,1))

import itertools

print(*itertools.permutations(arr,0))