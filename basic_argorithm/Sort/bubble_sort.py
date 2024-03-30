arr = [1,3,5,7,9,8,6,4,2]

def bubble_sort(arr):
    n = len(arr)

    for i in range(n):
        for j in range(n-i-1):
            if arr[j+1] < arr[j]:
                arr[j+1], arr[j] = arr[j], arr[j+1]
    
    return arr


print(bubble_sort(arr))