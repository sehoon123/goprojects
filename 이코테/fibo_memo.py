import sys

sys.setrecursionlimit(100_000_000)

input = sys.stdin.readline

cache = [0] * 101


def fibo(x: int) -> int:
    if x <= 2:
        return 1

    if cache[x] != 0:
        return cache[x]

    cache[x] = fibo(x-1) + fibo(x-2)

    return cache[x]


print(fibo(100))


def pivot_pythonic(array: list):
    if len(array) <= 1:
        return array

    pivot = array[0]
    tails = array[1:]

    left = [x for x in tails if x <= pivot]
    right = [x for x in tails if x > pivot]

    return pivot_pythonic(left) + [pivot] + pivot_pythonic(right)


array = [3,5,1,3,6,8,10,15,2,4,5]
res = pivot_pythonic(array)

print(res)



