array = [7,5,9,0,3,1,6,2,4,8]

def quick_sort(array:list, start: int, end:int):
    if start >= end: # 원소가 1개인 경우
        return
    pivot = start
    left = start + 1
    right = end

    while left <= right:
        # 피벗보다 큰 데이터를 찾을 때까지 반복
        while left <= end and array[left] <= array[pivot]:
            left += 1
        # 피벗보다 작은 데이터를 찾을 때까지 반복
        while right > start and array[right] >= array[pivot]:
            right -= 1

        if left > right: # 엇갈렸다면 작은 데이터와 피벗을 교체
            array[right], array[pivot] = array[pivot], array[right]
        else: # 엇갈리지 않았다면 작은 데이터와 큰 데이터를 교체
            array[left], array[right] = array[right], array[left]


    quick_sort(array, start, right-1)
    quick_sort(array, right+1, end)


    def pythonic_quick_sort(array: list):
        if len(array) <= 1:
            return array

        pivot = array[0]
        tails = array[1:]

        left = [x for x in tails if x <= pivot]
        right = [x for x in tails if x > pivot]

        return pythonic_quick_sort(left) + [pivot] + pythonic_quick_sort(right)
