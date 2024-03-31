import heapq

def heap_sort(arr):
    heap = []
    for value in arr:
        heapq.heappush(heap, value)
    
    sorted_list = []
    while heap:
        sorted_list.append(heapq.heappop(heap))
    
    return sorted_list

arr = [12,11,13,5,6,7]

sorted_arr = heap_sort(arr)
print(sorted_arr)