arr = [1,3,5,7,9,8,6,4,2]

def insertion_sort(arr):
    n = len(arr)

    for i in range(1, n):
        key = arr[i]

        for j in range(i-1,-1,-1):
            if arr[j] > key:
                arr[j+1] = arr[j]
            else:
                break
        arr[j+1] = key

    return arr