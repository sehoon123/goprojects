import sys

input = sys.stdin.readline

N = int(input())
stock = list(map(int, input().split()))

M = int(input())
orders = list(map(int, input().split()))

stock.sort()


def binary_search(array: list, start: int, end: int, target: int) -> int:
    while start <= end:
        mid = (start + end) // 2

        if array[mid] == target:
            return mid
        elif array[mid] > target:
            end = mid - 1
        else:
            start = mid + 1

    return -1


for i in orders:
    result = binary_search(stock, 0, len(stock)-1, i)
    if result == -1:
        print("no", end=" ")
    else:
        print("yes", end=" ")
