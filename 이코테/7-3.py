import sys

input = sys.stdin.readline

n, m = map(int, input().split())
height = list(map(int, input().split()))


def solve(x: int) -> int:
    res = 0
    for i in height:
        if i > x:
            res += i - x

    return res


def binary_search(array: list, start: int, end: int, target: int) -> int:
    result = 0
    while start <= end:
        mid = (start+end) // 2

        if solve(mid) >= target:
            start = mid + 1
            result = mid  # 이부분에서 결과를 빼내는것이 포인트
        else:
            end = mid - 1

    return result


res = binary_search(height, 0, 2_000_000_000, m)

print(res)

