import sys
input = sys.stdin.readline

N = int(input())

arr = [0] * N

for i in range(N):
    arr[i] = int(input())

def quick_sort2(arr, start, end):
    if start >= end:
        return
    
    pivot = start
    i = start + 1
    j = end

    while i <= j :
        while i <= end and arr[i] < arr[pivot]:
            i += 1
        while j >= start and arr[j] > arr[pivot]:
            j -= 1
        
        if i > j:
            arr[j], arr[pivot] = arr[pivot], arr[j]
        else:
            arr[i], arr[j] = arr[j], arr[i]
        
    quick_sort2(arr, start, j-1)
    quick_sort2(arr, j+1, end)

    return arr


answer = quick_sort2(arr, 0, N-1)

for num in answer:
    print(num)