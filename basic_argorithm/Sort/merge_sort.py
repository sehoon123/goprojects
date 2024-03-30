arr = [1,3,5,7,9,8,6,4,2,0]

sorted_list = [0] * len(arr)
def merge(arr, start, mid, end):
    i = start
    j = mid + 1
    k = start

    while i <= mid and j <= end:
        if arr[i] <= arr[j]:
            sorted_list[k] = arr[i]
            i += 1
        else:
            sorted_list[k] = arr[j]
            j += 1
        k += 1
    
    if i > mid:
        for x in range(j, end+1):
            sorted_list[k] = arr[x]
            k+= 1
    else:
        for x in range(i, mid+1):
            sorted_list[k] = arr[x]
            k += 1
    
    for x in range(start, end+1):
        arr[x] = sorted_list[x]
        
def merge_sort(arr, start, end):
    if start < end:
        mid = (start + end) // 2
        merge_sort(arr, start, mid)
        merge_sort(arr, mid+1, end)
        merge(arr, start, mid, end)

merge_sort(arr, 0, len(arr)-1)       

print(arr)
