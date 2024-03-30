arr = [9, 7, 5, 3, 1, 2, 4, 8, 6, 10]

def quick_sort(arr, start, end):
    if start >= end:
        return

    pivot = arr[start]
    i = start + 1
    j = end

    while i <= j:
        while i <= end and arr[i] <= pivot:
            i += 1
        while j > start and arr[j] >= pivot:
            j -= 1

        if i > j:
            arr[start], arr[j] = arr[j], arr[start]
        else:
            arr[i], arr[j] = arr[j], arr[i]

    quick_sort(arr, start, j - 1)
    quick_sort(arr, j + 1, end)

    return arr

print(quick_sort(arr, 0, len(arr) - 1))


def quick_sort2(arr):
    if len(arr) <= 1:
        return arr
    
    pivot = arr[0]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]

    return quick_sort2(left) + middle + quick_sort2(right)

print(quick_sort2(arr))