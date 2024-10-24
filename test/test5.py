import math

def solution(numbers):
    answer = []
    binary_nums = [bin(x)[2:] for x in numbers]
    
    for binary in binary_nums:
        length = len(binary)
        x = 2** (math.ceil(math.log(length+1, 2)))-1
        num = str("0"*(x-length) + binary)
        
        if check(num, 0, x-1):
            answer.append(1)
        else:
            answer.append(0)
        
    return answer

    
def check(num, start, end):
    if start >end:
        return True  # 재귀 종료 조건: 범위가 올바르지 않을 때는 True 반환
    
    mid = (start + end) // 2
    
    # 가운데 값이 1이 아니면 False 반환
    if num[mid] != '1':
        if num[mid-1] == '1' or num[mid+1] == '1':
            return False
    
    # 양쪽 하위 배열에 대해서도 같은 작업을 재귀적으로 수행
    left_check = check(num, start, mid - 1)
    right_check = check(num, mid + 1, end)
    
    # 양쪽 모두 True일 때만 전체 결과가 True
    return left_check and right_check

print(check("000010000", 0, 8))