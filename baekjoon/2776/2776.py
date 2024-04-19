import sys
input = sys.stdin.readline

T = int(input())

def binary_search(arr: list, x: int) -> bool:
    left, right = 0, len(arr)-1

    while left <= right:
        mid = (left + right) // 2

        if arr[mid] == x:
            return True
        
        if arr[mid] < x:
            left = mid + 1
        else:
            right = mid -1
    
    return False



for _ in range(T):
    N = int(input())
    arr1 = list(map(int,input().split()))
    M = int(input())
    arr2 = list(map(int,input().split()))

    arr1.sort()

    for w in arr2:
        if binary_search(arr1,w):
            print(1)
        else:
            print(0)