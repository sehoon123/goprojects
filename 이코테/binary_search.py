import sys

input = sys.stdin.readline

array = [1,2,3,5,6,7,8,9,10,11,12,13,14,15]

def binary_search(array: list, start: int, end: int, target: int):
    if start > end:
        return -1

    while start <= end:
        mid = (start + end) // 2

        if array[mid] == target:
            return mid
        elif array[mid] > target:
            end = mid - 1
        else:
            start = mid + 1

    return -1

print(binary_search(array, 0, len(array)-1, 5))
