def bubble_sort(arr):
    for i in range(len(arr)):
        for j in range(len(arr)):
            if arr[i] < arr[j]:
                arr[i],arr[j] = arr[j], arr[i]
    
    return arr

def selection_sort(arr):
    for i in range(len(arr)):
        minimum = i
        for j in range(i+1, len(arr)):
            if arr[j] < arr[minimum]:
                minimum = j
        arr[i], arr[minimum] = arr[minimum], arr[i]
    
    return arr

def insertion_sort(arr):
    for i in range(1, len(arr)):
        key = arr[i]
        if key < arr[0]:
            arr = [key] + arr[:i] + arr[i+1:]
        else:
            for j in range(1, i):
                if arr[j-1] < key <= arr[j]:
                    arr = arr[:j] + [key] + arr[j:i] + arr[i+1:]
                    break
            else:
                arr = arr[:i] + [key] + arr[i+1:]
    return arr


def merge_sort(arr):
    if len(arr) == 1:
        return arr
    
    mid = len(arr)//2
    left = merge_sort(arr[:mid])
    right = merge_sort(arr[mid:])
    
    return merge(left, right)

def merge(left, right):
    result = []
    i, j = 0, 0

    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i+=1
        else:
            result.append(right[j])
            j+=1
        
    result.extend(left[i:])
    result.extend(right[j:])

    return result


arr = [3,2,3,1,2,4,5,5,6]

# print(bubble_sort(arr))
# print(insertion_sort(arr))
# print(merge_sort(arr))

def quick_sort(arr):
    if len(arr) <= 1:
        return arr

    pivot = arr[0]
    mid = [x for x in arr if x == pivot ]
    left = [x for x in arr if x < pivot]
    right = [x for x in arr if x > pivot]

    return quick_sort(left) + mid + quick_sort(right)

print(quick_sort(arr))
